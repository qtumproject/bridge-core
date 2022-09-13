package evm

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"math/big"
)

func skipSig(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return transaction, nil
}

func encodeTx(tx *types.Transaction, from string, chainID *big.Int) (json.RawMessage, error) {
	txObj := jsonTx{
		From:  from,
		To:    tx.To().Hex(),
		Value: 0,
		Data:  hexutil.Encode(tx.Data()),
		Chain: fmt.Sprintf("0x%x", chainID.Int64()),
	}

	rawTx, err := json.Marshal(txObj)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode tx")
	}
	return rawTx, nil
}

type jsonTx struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value int64  `json:"value"`
	Data  string `json:"data"`
	Chain string `json:"chain_id"`
}
