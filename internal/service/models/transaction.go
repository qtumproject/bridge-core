package models

import (
	"encoding/json"
	"gitlab.com/tokend/bridge/core/resources"
	"math/rand"
)

func NewTransactionResponse(tx json.RawMessage) resources.TransactionResponse {
	return resources.TransactionResponse{
		Data: resources.Transaction{
			Key: resources.NewKeyInt64(rand.Int63(), resources.TRANSACTION),
			Attributes: resources.TransactionAttributes{
				EncodedTx: tx,
			},
		},
	}
}
