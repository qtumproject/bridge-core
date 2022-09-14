package types

import (
	"gitlab.com/tokend/bridge/core/internal/amount"
	"gitlab.com/tokend/bridge/core/internal/data"
)

type Proxy interface {
	Approve(tokenChain data.TokenChain, approveFrom string) (interface{}, error)
	LockFungible(params FungibleLockParams) (interface{}, error)
	LockNonFungible(params NonFungibleLockParams) (interface{}, error)
}

type FungibleLockParams struct {
	TokenChain       data.TokenChain
	Sender           string
	Receiver         string
	DestinationChain string
	Amount           amount.Amount
}

type NonFungibleLockParams struct {
	TokenChain       data.TokenChain
	Sender           string
	Receiver         string
	DestinationChain string
	NftId            string
}
