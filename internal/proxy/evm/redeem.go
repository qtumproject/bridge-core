package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/signature"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"math/big"
)

func (p *evmProxy) RedeemFungible(params types.FungibleRedeemParams) (interface{}, error) {
	if err := p.containsHash(params.TxHash, params.EventIndex); err != nil {
		return nil, err
	}

	sender := common.HexToAddress(params.Sender)
	if params.TokenChain.AutoSend {
		sender = p.signer.Address()
	}

	var tx *ethTypes.Transaction
	var err error
	switch params.TokenChain.TokenType {
	case tokenTypeNative:
		tx, err = p.redeemNative(params, sender)
	case tokenTypeErc20:
		tx, err = p.redeemErc20(params, sender)
	default:
		return nil, errors.Errorf("unknown token type: %s, token: %s", params.TokenChain.TokenType,
			params.TokenChain.TokenID)
	}
	if err != nil {
		return nil, err
	}

	if params.TokenChain.AutoSend {
		return p.sendTx(tx, params.TokenChain.ChainID)
	}

	return encodeTx(tx, common.HexToAddress(params.Sender), p.chainID, params.TokenChain.ChainID)
}

func (p *evmProxy) RedeemNonFungible(params types.NonFungibleRedeemParams) (interface{}, error) {
	if err := p.containsHash(params.TxHash, params.EventIndex); err != nil {
		return nil, err
	}

	sender := common.HexToAddress(params.Sender)
	if params.TokenChain.AutoSend {
		sender = p.signer.Address()
	}

	var tx *ethTypes.Transaction
	var err error
	switch params.TokenChain.TokenType {
	case tokenTypeErc721:
		tx, err = p.redeemErc721(params, sender)
	case tokenTypeErc1155:
		tx, err = p.redeemErc1155(params, sender)
	default:
		return nil, errors.Errorf("unknown token type: %s, token: %s", params.TokenChain.TokenType,
			params.TokenChain.TokenID)
	}
	if err != nil {
		return nil, err
	}

	if params.TokenChain.AutoSend {
		return p.sendTx(tx, params.TokenChain.ChainID)
	}

	return encodeTx(tx, common.HexToAddress(params.Sender), p.chainID, params.TokenChain.ChainID)
}

func (p *evmProxy) containsHash(txHash string, eventIndex int) error {
	//containsHash, err := p.bridge.ContainsHash(&bind.CallOpts{}, common.HexToHash(txHash), big.NewInt(int64(eventIndex)))
	//if err != nil {
	//	return errors.Wrap(err, "failed to check contains hash")
	//}
	//if containsHash {
	//	return types.ErrAlreadyRedeemed
	//}

	return nil
}

func (p *evmProxy) redeemNative(params types.FungibleRedeemParams, sender common.Address) (*ethTypes.Transaction, error) {
	txHash := common.HexToHash(params.TxHash)
	amount := params.Amount.Int()

	log := signature.NativeLog{
		Amount:     amount,
		Receiver:   params.Receiver,
		TxHash:     txHash,
		EventIndex: params.EventIndex,
		ChainID:    p.chainID,
	}
	sign, err := p.signer.Sign(&log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign log")
	}

	return p.bridge.WithdrawNative(buildTransactOpts(sender),
		amount,
		common.HexToAddress(params.Receiver),
		txHash,
		big.NewInt(int64(params.EventIndex)),
		[][]byte{sign})
}

func (p *evmProxy) redeemErc20(params types.FungibleRedeemParams, sender common.Address) (*ethTypes.Transaction, error) {
	decimals, err := p.getDecimals(*params.TokenChain.ContractAddress)
	if err != nil {
		return nil, err
	}
	txHash := common.HexToHash(params.TxHash)
	amount := params.Amount.IntWithPrecision(decimals)

	log := signature.Erc20Log{
		TokenAddress: *params.TokenChain.ContractAddress,
		Amount:       amount,
		Receiver:     params.Receiver,
		TxHash:       txHash,
		EventIndex:   params.EventIndex,
		ChainID:      p.chainID,
		IsWrapped:    isWrappedToken(params.TokenChain.BridgingType),
	}
	sign, err := p.signer.Sign(&log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign log")
	}

	return p.bridge.WithdrawERC20(buildTransactOpts(sender),
		common.HexToAddress(*params.TokenChain.ContractAddress),
		amount,
		common.HexToAddress(params.Receiver),
		txHash,
		big.NewInt(int64(params.EventIndex)),
		log.IsWrapped,
		[][]byte{sign},
	)
}

func (p *evmProxy) redeemErc721(params types.NonFungibleRedeemParams, sender common.Address) (*ethTypes.Transaction, error) {
	txHash := common.HexToHash(params.TxHash)
	nftId, ok := big.NewInt(0).SetString(params.NftId, 10)
	if !ok {
		return nil, errors.New("failed to parse nft id")
	}

	log := signature.Erc721Log{
		TokenAddress: *params.TokenChain.ContractAddress,
		TokenID:      nftId,
		Receiver:     params.Receiver,
		TxHash:       txHash,
		EventIndex:   params.EventIndex,
		ChainID:      p.chainID,
		TokenUri:     params.NftUri,
		IsWrapped:    isWrappedToken(params.TokenChain.BridgingType),
	}
	sign, err := p.signer.Sign(&log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign log")
	}

	return p.bridge.WithdrawERC721(buildTransactOpts(sender),
		common.HexToAddress(*params.TokenChain.ContractAddress),
		nftId,
		common.HexToAddress(params.Receiver),
		txHash,
		big.NewInt(int64(params.EventIndex)),
		params.NftUri,
		log.IsWrapped,
		[][]byte{sign},
	)
}

func (p *evmProxy) redeemErc1155(params types.NonFungibleRedeemParams, sender common.Address) (*ethTypes.Transaction, error) {
	txHash := common.HexToHash(params.TxHash)
	nftId, ok := big.NewInt(0).SetString(params.NftId, 10)
	if !ok {
		return nil, errors.New("failed to parse nft id")
	}
	amount := big.NewInt(1)

	log := signature.Erc1155Log{
		TokenAddress: *params.TokenChain.ContractAddress,
		TokenID:      nftId,
		Amount:       amount,
		Receiver:     params.Receiver,
		TxHash:       txHash,
		EventIndex:   params.EventIndex,
		ChainID:      p.chainID,
		TokenUri:     params.NftUri,
		IsWrapped:    isWrappedToken(params.TokenChain.BridgingType),
	}
	sign, err := p.signer.Sign(&log)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign log")
	}

	return p.bridge.WithdrawERC1155(buildTransactOpts(sender),
		common.HexToAddress(*params.TokenChain.ContractAddress),
		nftId,
		amount,
		common.HexToAddress(params.Receiver),
		txHash,
		big.NewInt(int64(params.EventIndex)),
		params.NftUri,
		log.IsWrapped,
		[][]byte{sign},
	)
}

func (p *evmProxy) sendTx(tx *ethTypes.Transaction, chain string) (interface{}, error) {
	tx, err := p.signer.SignTx(tx, p.chainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign tx")
	}

	err = p.client.SendTransaction(context.TODO(), tx)

	return encodeProcessedTx(tx.Hash(), chain), errors.Wrap(err, "failed to send tx")
}
