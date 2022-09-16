package models

import (
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/resources"
)

type TxResponse struct {
	Data     interface{}        `json:"data"`
	Included resources.Included `json:"included"`
}

func NewTxResponse(tx interface{}, chain data.Chain) TxResponse {
	response := TxResponse{
		Data:     tx,
		Included: resources.Included{},
	}

	response.Included.Add(newChainModel(chain))

	return response
}
