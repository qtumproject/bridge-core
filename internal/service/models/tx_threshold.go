package models

import (
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/resources"
)

type TxResponseWithThreshold struct {
	Data     DataWithThreshold  `json:"data"`
	Included resources.Included `json:"included"`
}

type DataWithThreshold struct {
	Tx               interface{} `json:"tx"`
	ThresholdReached bool        `json:"threshold_reached"`
}

func NewTxResponseWithThreshold(tx interface{}, chain data.Chain, thresholdReached bool) TxResponseWithThreshold {
	response := TxResponseWithThreshold{
		Data: DataWithThreshold{
			Tx:               tx,
			ThresholdReached: thresholdReached,
		},
		Included: resources.Included{},
	}

	response.Included.Add(newChainModel(chain))

	return response
}
