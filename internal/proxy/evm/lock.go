package evm

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/amount"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/enums"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"math/big"
)

func (p *evmProxy) LockFungible(params types.FungibleLockParams) (interface{}, error) {
	switch params.TokenChain.TokenType {
	case TokenTypeNative:
		return p.lockNative(params)
	case TokenTypeErc20:
		return p.lockErc20(params)
	default:
		return nil, errors.New("unsupported token type")
	}
}

func (p *evmProxy) LockNonFungible(params types.NonFungibleLockParams) (interface{}, error) {
	switch params.TokenChain.TokenType {
	case TokenTypeErc721:
		return p.lockErc721(params)
	case TokenTypeErc1155:
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

	return encodeTx(tx, senderAddr, p.chainID, params.TokenChain.ChainID, nil)
}

func (p *evmProxy) lockErc20(params types.FungibleLockParams) (interface{}, error) {
	decimals, err := p.getDecimals(*params.TokenChain.ContractAddress)
	if err != nil {
		return nil, err
	}

	senderAddr := common.HexToAddress(params.Sender)
	tx, err := p.bridge.DepositERC20(
		buildTransactOpts(senderAddr),
		common.HexToAddress(*params.TokenChain.ContractAddress),
		params.Amount.IntWithPrecision(decimals),
		params.Receiver,
		params.DestinationChain,
		uint8(enums.ToErc20BridgingType(params.TokenChain.BridgingType)),
	)

	return encodeTx(tx, senderAddr, p.chainID, params.TokenChain.ChainID, nil)
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
		uint8(enums.ToErc721BridgingType(params.TokenChain.BridgingType)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build deposit erc721 transaction")
	}

	return encodeTx(tx, senderAddr, p.chainID, params.TokenChain.ChainID, nil)
}

func (p *evmProxy) lockErc1155(params types.NonFungibleLockParams) (interface{}, error) {
	senderAddr := common.HexToAddress(params.Sender)
	tokenId, ok := big.NewInt(0).SetString(params.NftId, 10)
	am := amount.MustNewFromString("1")
	if params.Amount != nil {
		am = *params.Amount
	}
	if !ok {
		return nil, errors.New("failed to parse token id")
	}
	tx, err := p.bridge.DepositERC1155(
		buildTransactOpts(senderAddr),
		common.HexToAddress(*params.TokenChain.ContractAddress),
		tokenId,
		am.IntWithPrecision(0),
		params.Receiver,
		params.DestinationChain,
		uint8(enums.ToErc1155BridgingType(params.TokenChain.BridgingType)),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build deposit erc1155 transaction")
	}

	return encodeTx(tx, senderAddr, p.chainID, params.TokenChain.ChainID, nil)
}
