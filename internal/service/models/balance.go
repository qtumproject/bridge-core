package models

import (
	"fmt"
	"gitlab.com/tokend/bridge/core/internal/amount"
	"gitlab.com/tokend/bridge/core/resources"
)

func NewBalanceResponse(tokenId string, address string, nft *string, balance amount.Amount) resources.BalanceResponse {
	var id string
	if nft == nil {
		id = fmt.Sprintf("%s-%s", tokenId, address)
	} else {
		id = fmt.Sprintf("%s-%s-%s", tokenId, address, *nft)
	}

	return resources.BalanceResponse{
		Data: resources.Balance{
			Key: resources.Key{
				ID:   id,
				Type: resources.BALANCE,
			},
			Attributes: resources.BalanceAttributes{
				Amount: balance,
			},
		},
	}
}
