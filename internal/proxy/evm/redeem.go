package evm

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/enums"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/bridge"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/signature"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"math/big"
	"reflect"
)

func (p *evmProxy) RedeemFungible(params types.FungibleRedeemParams) (interface{}, error) {
	if err := p.containsHash(params.TxHash, params.EventIndex); err != nil {
		return nil, err
	}

	threshold, err := p.getThreshold()
	if err != nil {
		return nil, err
	}

	sender := common.HexToAddress(params.Sender)
	if params.TokenChain.AutoSend {
		sender = p.signer.Address()
	}

	var tx *ethTypes.Transaction
	switch params.TokenChain.TokenType {
	case TokenTypeNative:
		tx, err = p.redeemNative(params, sender)
	case TokenTypeErc20:
		tx, err = p.redeemErc20(params, sender)
	default:
		return nil, errors.Errorf("unknown token type: %s, token: %s", params.TokenChain.TokenType,
			params.TokenChain.TokenID)
	}
	if err != nil {
		return nil, err
	}

	signNumber := int64(1)

	// if tx provided check it and sign; otherwise use created tx
	if params.RawTxData != nil {
		tx, signNumber, err = p.checkTxDataAndSign(buildTransactOpts(sender), tx, *params.RawTxData)
		if err != nil {
			return nil, err
		}
	}

	confirmed := signNumber >= threshold

	if params.TokenChain.AutoSend && confirmed {
		return p.sendTx(tx, params.TokenChain.ChainID)
	}

	return encodeTx(tx, common.HexToAddress(params.Sender), p.chainID, params.TokenChain.ChainID, &confirmed)
}

func (p *evmProxy) RedeemNonFungible(params types.NonFungibleRedeemParams) (interface{}, error) {
	if err := p.containsHash(params.TxHash, params.EventIndex); err != nil {
		return nil, err
	}

	threshold, err := p.getThreshold()
	if err != nil {
		return nil, err
	}

	sender := common.HexToAddress(params.Sender)
	if params.TokenChain.AutoSend {
		sender = p.signer.Address()
	}

	var tx *ethTypes.Transaction
	switch params.TokenChain.TokenType {
	case TokenTypeErc721:
		tx, err = p.redeemErc721(params, sender)
	case TokenTypeErc1155:
		tx, err = p.redeemErc1155(params, sender)
	default:
		return nil, errors.Errorf("unknown token type: %s, token: %s", params.TokenChain.TokenType,
			params.TokenChain.TokenID)
	}
	if err != nil {
		return nil, err
	}

	signNumber := int64(1)

	// if tx provided check it and sign; otherwise use created tx
	if params.RawTxData != nil {
		tx, signNumber, err = p.checkTxDataAndSign(buildTransactOpts(sender), tx, *params.RawTxData)
		if err != nil {
			return nil, err
		}
	}

	confirmed := signNumber >= threshold

	if params.TokenChain.AutoSend && confirmed {
		return p.sendTx(tx, params.TokenChain.ChainID)
	}

	return encodeTx(tx, common.HexToAddress(params.Sender), p.chainID, params.TokenChain.ChainID, &confirmed)
}

func (p *evmProxy) containsHash(txHash string, eventIndex int) error {
	containsHash, err := p.bridge.ContainsHash(&bind.CallOpts{}, common.HexToHash(txHash), big.NewInt(int64(eventIndex)))
	if err != nil {
		return errors.Wrap(err, "failed to check contains hash")
	}
	if containsHash {
		return types.ErrAlreadyRedeemed
	}

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
		BridgingType: enums.ToErc20BridgingType(params.TokenChain.BridgingType),
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
		uint8(log.BridgingType),
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
		BridgingType: enums.ToErc721BridgingType(params.TokenChain.BridgingType),
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
		uint8(log.BridgingType),
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
		BridgingType: enums.ToErc1155BridgingType(params.TokenChain.BridgingType),
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
		uint8(log.BridgingType),
		[][]byte{sign},
	)
}

func (p *evmProxy) sendTx(tx *ethTypes.Transaction, chain string) (interface{}, error) {
	hash, err := p.relayer.SendTx(tx, p.chainID)
	return encodeProcessedTx(hash, chain), errors.Wrap(err, "failed to send tx")
}

func (p *evmProxy) getThreshold() (int64, error) {
	threshold, err := p.bridge.SignaturesThreshold(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}

	return threshold.Int64(), nil
}

func (p *evmProxy) checkTxDataAndSign(opts *bind.TransactOpts, tx *ethTypes.Transaction, rawData []byte) (*ethTypes.Transaction, int64, error) {
	abi, err := bridge.BridgeMetaData.GetAbi()
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to parse bridge ABI")
	}

	newParams, newMethod, err := decodeTxParams(*abi, tx.Data())
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to decode tx params")
	}

	oldParams, oldMethod, err := decodeTxParams(*abi, rawData)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to decode tx params")
	}

	if len(oldParams) != len(newParams) {
		return nil, 0, types.ErrWrongSignedTx
	}

	if newMethod.Name != oldMethod.Name {
		return nil, 0, types.ErrWrongSignedTx
	}

	// Check if all params except signature is equal
	if !reflect.DeepEqual(oldParams[:len(oldParams)-1], newParams[:len(newParams)-1]) {
		return nil, 0, types.ErrWrongSignedTx
	}

	signatures := oldParams[len(oldParams)-1].([][]byte)
	newSig := newParams[len(newParams)-1].([][]byte)[0]

	for _, sig := range signatures {
		if reflect.DeepEqual(sig, newSig) {
			return nil, int64(len(signatures)), errors.New("double signature")
		}
	}

	// Add signature to params and encode new transaction
	newParams[len(newParams)-1] = append(signatures, newSig)
	if signs, ok := newParams[len(newParams)-1].([][]byte); ok {
		contract := bind.NewBoundContract(p.bridgeContract, *abi, nil, p.client, nil)
		newTx, err := contract.Transact(opts, newMethod.Name, newParams...)
		return newTx, int64(len(signs)), err
	}

	return nil, 0, errors.New("failed to cast signatures param")
}
