package types

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/amount"
	"gitlab.com/tokend/bridge/core/internal/data"
)

var ErrTxNotFound = errors.New("tx not found")
var ErrTxNotConfirmed = errors.New("tx not confirmed yet")
var ErrTxFailed = errors.New("tx failed")
var ErrEventNotFound = errors.New("log not found")
var ErrWrongLockEvent = errors.New("metadata is incorrect")
var ErrWrongToken = errors.New("token is incorrect")
var ErrAlreadyRedeemed = errors.New("transaction is already redeemed")
var ErrNotFound = errors.New("not found")
var ErrWrongSignedTx = errors.New("signed tx does not match tx log")

type Proxy interface {
	Approve(params ApproveParams) (interface{}, error)
	LockFungible(params FungibleLockParams) (interface{}, error)
	LockNonFungible(params NonFungibleLockParams) (interface{}, error)
	CheckFungibleLockEvent(txHash string, eventIndex int, tokenChain data.TokenChain) (*FungibleLockEvent, error)
	CheckNonFungibleLockEvent(txHash string, eventIndex int, tokenChain data.TokenChain) (*NonFungibleLockEvent, error)
	RedeemFungible(params FungibleRedeemParams) (interface{}, error)
	RedeemNonFungible(params NonFungibleRedeemParams) (interface{}, error)

	// Balance returns balance of the given token for the given address
	// For fungible tokens it returns amount of tokens
	// For non-fungible tokens returns 1 if the token is owned by the account, 0 otherwise
	// nftId should be not nil for non-fungible tokens
	Balance(tokenChain data.TokenChain, address string, nftId *string) (amount.Amount, error)
	BridgeBalance(tokenChain data.TokenChain, nftId *string) (amount.Amount, error)
	GetNftMetadataUri(tokenChain data.TokenChain, nftId string) (string, error)
	GetNftMetadata(tokenChain data.TokenChain, nftId string) (*NftMetadata, error)
}

type ApproveParams struct {
	TokenChain  data.TokenChain
	ApproveFrom string
	Amount      *amount.Amount
	NftId       *string
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
}

type FungibleRedeemParams struct {
	TokenChain data.TokenChain
	Sender     string
	Receiver   string
	TxHash     string
	EventIndex int
	Amount     amount.Amount
	RawTxData  *[]byte
}

type NonFungibleRedeemParams struct {
	TokenChain data.TokenChain
	Sender     string
	Receiver   string
	TxHash     string
	EventIndex int
	NftId      string
	NftUri     string
	RawTxData  *[]byte
}
