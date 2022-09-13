package evm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/bridge"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/tokenerc20"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/tokenerc721"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"gitlab.com/tokend/bridge/core/internal/signature"
	"io"
	"math/big"
	"net/http"
)

func NewProxy(rpc string, signer signature.Signer) (types.Proxy, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to rpc")
	}

	chainID, err := client.ChainID(context.TODO())
	if err != nil {
		return nil, errors.Wrap(err, "failed to load chain id")
	}

	return &evmProxy{
		client:  client,
		signer:  signer,
		chainID: chainID,
	}, nil
}

type evmProxy struct {
	client  *ethclient.Client
	signer  signature.Signer
	chainID *big.Int
}

func (p *evmProxy) GetNftMetadata(tokenAddress string, nftID string) (*data.NFTMetadata, error) {
	contract, err := tokenerc721.NewTokenerc721Caller(common.HexToAddress(tokenAddress), p.client)
	if err != nil {
		return nil, err
	}

	tokenId := big.NewInt(0)
	tokenId.SetString(nftID, 10)
	url, err := contract.TokenURI(&bind.CallOpts{}, tokenId)
	if err != nil {
		return nil, err
	}

	return getDetails(url)
}

func getDetails(url string) (*data.NFTMetadata, error) {
	metadata := data.NFTMetadata{
		MetadataURL: url,
	}

	url = wrapIpfsUrl(url)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, errors.New(fmt.Sprintf("request to get details failed, status - %d, body - %s", response.StatusCode, string(body)))
	}

	var model tokenDetails
	err = json.Unmarshal(body, &model)
	if err != nil {
		return nil, err
	}

	var iconUrl string
	if model.Image != "" {
		iconUrl = wrapIpfsUrl(model.Image)
	} else {
		iconUrl = wrapIpfsUrl(model.ImageOpenSea)
	}

	metadata.Name = model.Name
	metadata.IconURL = iconUrl
	return &metadata, nil
}

type tokenDetails struct {
	Name         string `json:"name"`
	Image        string `json:"image"`
	ImageOpenSea string `json:"image_url"`
}

func wrapIpfsUrl(url string) string {
	return url
}

func (p *evmProxy) CreateApprovalTx(tokenId *big.Int, tokenAddress, ownerAddress, receiverAddress string) (json.RawMessage, error) {
	token, err := tokenerc721.NewTokenerc721(common.HexToAddress(tokenAddress), p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transactor")
	}

	status, err := token.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(ownerAddress), common.HexToAddress(receiverAddress))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get approval")
	}
	if status {
		// Token is already approved
		return nil, nil
	}

	tx, err := token.Approve(&bind.TransactOpts{
		From:     common.HexToAddress(ownerAddress),
		Signer:   skipSig,
		NoSend:   true,
		GasLimit: 300000,
	}, common.HexToAddress(receiverAddress), tokenId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build token approve tx")
	}

	return encodeTx(tx, ownerAddress, p.chainID)
}

func (p *evmProxy) CreateNonFungibleWithdrawTx(params types.NonFungibleWithdrawParams) (json.RawMessage, error) {
	transactor, err := tokenerc721.NewTokenerc721Transactor(common.HexToAddress(params.TokenAddress), p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transactor")
	}

	id := big.NewInt(0)
	id.SetString(params.NftID, 10)
	tx, err := transactor.SafeTransferFrom(&bind.TransactOpts{
		From:     common.HexToAddress(params.Sender),
		Signer:   skipSig,
		NoSend:   true,
		GasLimit: 300000,
	}, common.HexToAddress(params.Sender), common.HexToAddress(params.TokenAddress), id)

	return encodeTx(tx, params.Sender, p.chainID)
}

func (p *evmProxy) CreateNonFungibleDepositTx(params types.NonFungibleDepositParams) (json.RawMessage, error) {
	transactor, err := tokenerc721.NewTokenerc721Transactor(common.HexToAddress(params.TokenAddress), p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transactor")
	}

	id := big.NewInt(0)
	id.SetString(params.NftID, 10)
	log := signature.DepositNonFungibleLog{
		ChainID:      p.chainID,
		TxHash:       params.WithdrawTxHash,
		TokenAddress: params.TokenAddress,
		Receiver:     params.Receiver,
		TokenID:      id,
	}
	sigParams, err := p.signer.Sign(log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign log")
	}
	r, s, v, err := sigParams.ToGeth()
	if err != nil {
		return nil, errors.Wrap(err, "failed to map sign params")
	}
	tx, err := transactor.Mint(&bind.TransactOpts{
		From:     common.HexToAddress(params.Receiver),
		Signer:   skipSig,
		NoSend:   true,
		GasLimit: 300000,
	}, log.TxHash, params.NftMetadata.MetadataURL, common.HexToAddress(params.Receiver), id, [][32]byte{r}, [][32]byte{s}, []uint8{v})

	return encodeTx(tx, params.Receiver, p.chainID)
}

func (p *evmProxy) CreateFungibleWithdrawTx(params types.FungibleWithdrawParams) (json.RawMessage, error) {
	//transactor, err := bridge.NewBridgeTransactor(common.HexToAddress(params.BridgeAddress), p.client)
	//if err != nil {
	//	return nil, errors.Wrap(err, "failed to create transactor")
	//}
	//
	//amount := big.Int(*params.Amount)
	//tx, err := transactor.DepositERC20(&bind.TransactOpts{
	//	From:     common.HexToAddress(params.Sender),
	//	Signer:   skipSig,
	//	NoSend:   true,
	//	GasLimit: 300000,
	//}, common.HexToAddress(params.TokenAddress), params.Receiver, &amount)
	transactor, err := tokenerc20.NewTokenerc20Transactor(common.HexToAddress(params.TokenAddress), p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transactor")
	}

	amount := big.Int(*params.Amount)
	tx, err := transactor.Transfer(&bind.TransactOpts{
		From:     common.HexToAddress(params.Sender),
		Signer:   skipSig,
		NoSend:   true,
		GasLimit: 300000,
	}, common.HexToAddress(params.BridgeAddress), &amount)

	return encodeTx(tx, params.Sender, p.chainID)
}

func (p *evmProxy) CreateFungibleDepositTx(params types.FungibleDepositParams) (json.RawMessage, error) {
	transactor, err := bridge.NewBridgeTransactor(common.HexToAddress(params.BridgeAddress), p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transactor")
	}

	amount := big.Int(*params.Amount)
	log := signature.DepositFungibleLog{
		ChainID:      p.chainID,
		TxHash:       params.WithdrawTxHash,
		TokenAddress: params.TokenAddress,
		Receiver:     params.Receiver,
		Amount:       &amount,
	}
	sigParams, err := p.signer.Sign(log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign log")
	}
	r, s, v, err := sigParams.ToGeth()
	if err != nil {
		return nil, errors.Wrap(err, "failed to map sign params")
	}
	tx, err := transactor.WithdrawERC200(&bind.TransactOpts{
		From:     common.HexToAddress(params.Receiver),
		Signer:   skipSig,
		NoSend:   true,
		GasLimit: 300000,
	}, common.HexToAddress(params.TokenAddress), params.WithdrawTxHash, &amount, [][32]byte{r}, [][32]byte{s}, []uint8{v})

	return encodeTx(tx, params.Sender, p.chainID)
}
