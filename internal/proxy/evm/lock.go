package evm

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc20"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"math/big"
)

func (p *evmProxy) LockFungible(params types.FungibleLockParams) (interface{}, error) {
	switch params.TokenChain.TokenType {
	case tokenTypeNative:
		return p.lockNative(params)
	case tokenTypeErc20:
		return p.lockErc20(params)
	default:
		return nil, errors.New("unsupported token type")
	}
}

func (p *evmProxy) LockNonFungible(params types.NonFungibleLockParams) (interface{}, error) {
	switch params.TokenChain.TokenType {
	case tokenTypeErc721:
		return p.lockErc721(params)
	case tokenTypeErc1155:
		return p.lockErc1155(params)
	default:
		return nil, errors.New("unsupported token type")
	}
}

func (p *evmProxy) lockNative(params types.FungibleLockParams) (interface{}, error) {
	senderAddr := common.HexToAddress(params.Sender)
	tx, err := p.bridge.DepositNative(buildTransactOptsWithValue(senderAddr, params.Amount.Int()),
		params.Receiver, params.DestinationChain)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build deposit native transaction")
	}

	return encodeTx(tx, senderAddr, p.chainID)
}

func (p *evmProxy) lockErc20(params types.FungibleLockParams) (interface{}, error) {
	token, err := erc20.NewErc20(common.HexToAddress(*params.TokenChain.ContractAddress), p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create erc20 contract")
	}

	decimals, err := token.Decimals(&bind.CallOpts{})
	if err != nil {
		return &big.Int{}, errors.Wrap(err, "failed to get decimals")
	}

	senderAddr := common.HexToAddress(params.Sender)
	tx, err := p.bridge.DepositERC20(
		buildTransactOpts(senderAddr),
		common.HexToAddress(*params.TokenChain.ContractAddress),
		params.Amount.IntWithPrecision(int(decimals)),
		params.Receiver,
		params.DestinationChain,
		isWrappedToken(params.TokenChain.BridgingType),
	)

	return encodeTx(tx, senderAddr, p.chainID)
}

func (p *evmProxy) lockErc721(params types.NonFungibleLockParams) (interface{}, error) {
	senderAddr := common.HexToAddress(params.Sender)
	tokenId, ok := big.NewInt(0).SetString(params.NftId, 10)
	if !ok {
		return nil, errors.New("failed to parse token id")
	}
	tx, err := p.bridge.DepositERC721(
		buildTransactOpts(senderAddr),
		common.HexToAddress(*params.TokenChain.ContractAddress),
		tokenId,
		params.Receiver,
		params.DestinationChain,
		isWrappedToken(params.TokenChain.BridgingType),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build deposit erc721 transaction")
	}

	return encodeTx(tx, senderAddr, p.chainID)
}

func (p *evmProxy) lockErc1155(params types.NonFungibleLockParams) (interface{}, error) {
	senderAddr := common.HexToAddress(params.Sender)
	tokenId, ok := big.NewInt(0).SetString(params.NftId, 10)
	if !ok {
		return nil, errors.New("failed to parse token id")
	}
	tx, err := p.bridge.DepositERC1155(
		buildTransactOpts(senderAddr),
		common.HexToAddress(*params.TokenChain.ContractAddress),
		tokenId,
		// TODO: Allow to specify amount for fungible tokens
		big.NewInt(1),
		params.Receiver,
		params.DestinationChain,
		isWrappedToken(params.TokenChain.BridgingType),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build deposit erc1155 transaction")
	}

	return encodeTx(tx, senderAddr, p.chainID)
}
