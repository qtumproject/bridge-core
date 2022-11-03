package evm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc20"
	"gitlab.com/tokend/bridge/core/resources"
	"math/big"
	"strings"
)

const gasLimit = 300000

func (p *evmProxy) getDecimals(tokenAddress string) (int, error) {
	token, err := erc20.NewErc20(common.HexToAddress(tokenAddress), p.client)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create erc20 contract")
	}

	decimals, err := token.Decimals(&bind.CallOpts{})
	if err != nil {
		return 0, errors.Wrap(err, "failed to get decimals")
	}

	return int(decimals), nil
}

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

func compareAddresses(a, b common.Address) bool {
	return strings.ToLower(a.String()) == strings.ToLower(b.String())
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

func encodeTx(tx *types.Transaction, from common.Address, chainID *big.Int, chain string) (interface{}, error) {
	return resources.EvmTransaction{
		Key: resources.Key{
			ID:   tx.Hash().String(),
			Type: resources.EVM_TRANSACTION,
		},
		Attributes: resources.EvmTransactionAttributes{
			TxBody: resources.EvmTransactionTxBody{
				From:    from.String(),
				To:      tx.To().String(),
				Value:   tx.Value().String(),
				Data:    hexutil.Encode(tx.Data()),
				ChainId: fmt.Sprintf("0x%x", chainID.Int64()),
			},
		},
		Relationships: resources.EvmTransactionRelationships{
			Chain: resources.Key{
				ID:   chain,
				Type: resources.CHAIN,
			}.AsRelation(),
		},
	}, nil
}

func encodeProcessedTx(txHash common.Hash, chain string) interface{} {
	return resources.ProcessedTransaction{
		Key: resources.Key{
			ID:   txHash.String(),
			Type: resources.PROCESSED_TRANSACTION,
		},
		Relationships: resources.ProcessedTransactionRelationships{
			Chain: resources.Key{
				ID:   chain,
				Type: resources.CHAIN,
			}.AsRelation(),
		},
	}
}
