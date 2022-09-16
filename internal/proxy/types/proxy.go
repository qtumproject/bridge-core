package types

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/amount"
	"gitlab.com/tokend/bridge/core/internal/data"
)

var ErrTxNotConfirmed = errors.New("tx not confirmed yet")
var ErrTxFailed = errors.New("tx failed")
var ErrEventNotFound = errors.New("log not found")
var ErrWrongLockEvent = errors.New("metadata is incorrect")
var ErrAlreadyRedeemed = errors.New("transaction is already redeemed")

type Proxy interface {
	Approve(tokenChain data.TokenChain, approveFrom string) (interface{}, error)
	LockFungible(params FungibleLockParams) (interface{}, error)
	LockNonFungible(params NonFungibleLockParams) (interface{}, error)
	CheckFungibleLockEvent(txHash string, eventIndex int, tokenChain data.TokenChain) (*FungibleLockEvent, error)
	CheckNonFungibleLockEvent(txHash string, eventIndex int, tokenChain data.TokenChain) (*NonFungibleLockEvent, error)
	RedeemFungible(params FungibleRedeemParams) (interface{}, error)
	RedeemNonFungible(params NonFungibleRedeemParams) (interface{}, error)
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

type FungibleLockEvent struct {
	Receiver         string
	DestinationChain string
	Amount           amount.Amount
}

type NonFungibleLockEvent struct {
	Receiver         string
	DestinationChain string
	NftId            string
	NftUri           string
}

type FungibleRedeemParams struct {
	TokenChain data.TokenChain
	Sender     string
	Receiver   string
	TxHash     string
	EventIndex int
	Amount     amount.Amount
}

type NonFungibleRedeemParams struct {
	TokenChain data.TokenChain
	Sender     string
	Receiver   string
	TxHash     string
	EventIndex int
	NftId      string
	NftUri     string
}
