package evm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/resources"
	"math/big"
)

const gasLimit = 300000

func buildTransactOpts(from common.Address) *bind.TransactOpts {
	return buildTransactOptsWithValue(from, nil)
}

func buildTransactOptsWithValue(from common.Address, value *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     from,
		Signer:   skipSig,
		NoSend:   true,
		GasLimit: gasLimit,
		Value:    value,
	}
}

func isWrappedToken(bridgingType data.BridgingType) bool {
	switch bridgingType {
	case data.BridgingTypeWrapped:
		return true
	case data.BridgingTypeLP:
		return false
	default:
		panic("unsupported bridging type")
	}
}

func skipSig(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return transaction, nil
}

func encodeTx(tx *types.Transaction, from common.Address, chainID *big.Int) (interface{}, error) {
	return resources.EvmTransaction{
		Key: resources.Key{
			ID:   tx.Hash().String(),
			Type: resources.EVM_TRANSACTION,
		},
		Attributes: resources.EvmTransactionAttributes{
			From:  from.String(),
			To:    tx.To().String(),
			Value: tx.Value().String(),
			Data:  hexutil.Encode(tx.Data()),
			Chain: fmt.Sprintf("0x%x", chainID.Int64()),
		},
	}, nil
}
