package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/amount"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/bridge"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc1155"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc721"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
)

const (
	depositedNativeEventName  = "DepositedNative"
	depositedERC20EventName   = "DepositedERC20"
	depositedERC721EventName  = "DepositedERC721"
	depositedERC1155EventName = "DepositedERC1155"
)

func (p *evmProxy) CheckFungibleLockEvent(txHash string, eventIndex int, tokenChain data.TokenChain) (*types.FungibleLockEvent, error) {
	receipt, err := p.getTxReceipt(txHash)
	if err != nil {
		return nil, err
	}

	switch tokenChain.TokenType {
	case tokenTypeNative:
		return p.checkNativeLockEvent(receipt, eventIndex)
	case tokenTypeErc20:
		return p.checkErc20LockEvent(receipt, eventIndex, tokenChain)
	default:
		return nil, errors.New("unsupported token type")
	}
}

func (p *evmProxy) CheckNonFungibleLockEvent(txHash string, eventIndex int, tokenChain data.TokenChain) (*types.NonFungibleLockEvent, error) {
	receipt, err := p.getTxReceipt(txHash)
	if err != nil {
		return nil, err
	}

	switch tokenChain.TokenType {
	case tokenTypeErc721:
		return p.checkErc721LockEvent(receipt, eventIndex, tokenChain)
	case tokenTypeErc1155:
		return p.checkErc1155LockEvent(receipt, eventIndex, tokenChain)
	default:
		return nil, errors.New("unsupported token type")
	}
}

func (p *evmProxy) getTxReceipt(txHash string) (*ethTypes.Receipt, error) {
	receipt, err := p.client.TransactionReceipt(context.TODO(), common.HexToHash(txHash))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tx by hash")
	}

	if receipt.Status != ethTypes.ReceiptStatusSuccessful {
		return nil, types.ErrTxFailed
	}

	height, err := p.client.BlockNumber(context.TODO())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current blockchain height")
	}
	// -1 because including in block mean already one confirmation
	if receipt.BlockNumber.Uint64()+uint64(p.confirmations-1) > height {
		return nil, types.ErrTxNotConfirmed
	}

	return receipt, nil
}

func (p *evmProxy) checkNativeLockEvent(receipt *ethTypes.Receipt, eventIndex int) (*types.FungibleLockEvent, error) {
	var log bridge.BridgeDepositedNative
	err := getBridgeEvent(&log, depositedNativeEventName, eventIndex, receipt)
	if err != nil {
		return nil, err
	}

	return &types.FungibleLockEvent{
		Receiver:         log.Receiver,
		DestinationChain: log.Network,
		Amount:           amount.NewFromBigInt(log.Amount),
	}, nil
}

func (p *evmProxy) checkErc20LockEvent(receipt *ethTypes.Receipt, eventIndex int, tokenChain data.TokenChain) (*types.FungibleLockEvent, error) {
	var log bridge.BridgeDepositedERC20
	err := getBridgeEvent(&log, depositedERC20EventName, eventIndex, receipt)
	if err != nil {
		return nil, err
	}

	if !compareAddresses(log.Token, common.HexToAddress(*tokenChain.ContractAddress)) {
		return nil, types.ErrWrongLockEvent
	}
	if log.IsWrapped && tokenChain.BridgingType != data.BridgingTypeWrapped {
		return nil, types.ErrWrongLockEvent
	}

	decimals, err := p.getDecimals(*tokenChain.ContractAddress)
	if err != nil {
		return nil, err
	}

	return &types.FungibleLockEvent{
		Receiver:         log.Receiver,
		DestinationChain: log.Network,
		Amount:           amount.NewFromIntWithPrecision(log.Amount, decimals),
	}, nil
}

func (p *evmProxy) checkErc721LockEvent(receipt *ethTypes.Receipt, eventIndex int, tokenChain data.TokenChain) (*types.NonFungibleLockEvent, error) {
	var log bridge.BridgeDepositedERC721
	err := getBridgeEvent(&log, depositedERC721EventName, eventIndex, receipt)
	if err != nil {
		return nil, err
	}

	tokenAddress := common.HexToAddress(*tokenChain.ContractAddress)
	if !compareAddresses(log.Token, tokenAddress) {
		return nil, types.ErrWrongLockEvent
	}
	if log.IsWrapped && tokenChain.BridgingType != data.BridgingTypeWrapped {
		return nil, types.ErrWrongLockEvent
	}

	token, err := erc721.NewErc721(tokenAddress, p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create erc721 token instance")
	}
	uri, err := token.TokenURI(&bind.CallOpts{}, log.TokenId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get token uri")
	}

	return &types.NonFungibleLockEvent{
		Receiver:         log.Receiver,
		DestinationChain: log.Network,
		NftId:            log.TokenId.String(),
		NftUri:           uri,
	}, nil
}

func (p *evmProxy) checkErc1155LockEvent(receipt *ethTypes.Receipt, eventIndex int, tokenChain data.TokenChain) (*types.NonFungibleLockEvent, error) {
	var log bridge.BridgeDepositedERC1155
	err := getBridgeEvent(&log, depositedERC1155EventName, eventIndex, receipt)
	if err != nil {
		return nil, err
	}

	tokenAddress := common.HexToAddress(*tokenChain.ContractAddress)
	if !compareAddresses(log.Token, tokenAddress) {
		return nil, types.ErrWrongLockEvent
	}
	if log.IsWrapped && tokenChain.BridgingType != data.BridgingTypeWrapped {
		return nil, types.ErrWrongLockEvent
	}
	if log.Amount.Uint64() != 1 {
		return nil, types.ErrWrongLockEvent
	}

	token, err := erc1155.NewErc1155(tokenAddress, p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create erc1155 token instance")
	}
	uri, err := token.Uri(&bind.CallOpts{}, log.TokenId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get token uri")
	}

	return &types.NonFungibleLockEvent{
		Receiver:         log.Receiver,
		DestinationChain: log.Network,
		NftId:            log.TokenId.String(),
		NftUri:           uri,
	}, nil
}

func getBridgeEvent(dest interface{}, logName string, eventIndex int, receipt *ethTypes.Receipt) error {
	abi, err := bridge.BridgeMetaData.GetAbi()
	if err != nil {
		return errors.Wrap(err, "failed to parse bridge ABI")
	}
	contract := bind.NewBoundContract(common.Address{}, *abi, nil, nil, nil)

	index := 0
	for _, l := range receipt.Logs {
		if l == nil {
			continue
		}
		err := contract.UnpackLog(dest, logName, *l)
		if err == nil {
			if index == eventIndex {
				return nil
			}
			index++
		}
	}

	return types.ErrEventNotFound
}
