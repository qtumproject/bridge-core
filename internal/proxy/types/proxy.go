package types

import (
	"encoding/json"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/resources"
)

type Proxy interface {
	GetNftMetadata(tokenAddress string, nftID string) (*data.NFTMetadata, error)
	CreateNonFungibleWithdrawTx(params NonFungibleWithdrawParams) (json.RawMessage, error)
	CreateNonFungibleDepositTx(params NonFungibleDepositParams) (json.RawMessage, error)
	//CreateApprovalTx(tokenId *big.Int, tokenAddress, ownerAddress, receiverAddress string) (json.RawMessage, error)

	//GetBalance(tokenAddress string, address string) (resources.Amount, error)
	//GetAllowance(tokenAddress string, address string) (resources.Amount, error)
	//CreateAllowanceTx(tokenAddress string, ownerAddress string, spenderAddress string, amount resources.Amount) (json.RawMessage, error)
	//CreateFungibleWithdrawTx(params FungibleWithdrawParams) (json.RawMessage, error)
	//CreateFungibleDepositTx(params FungibleDepositParams) (json.RawMessage, error)
	//GetFungibleWithdrawMetadata(txHash string) (FungibleWithdrawMetadata, error)
}

type NonFungibleWithdrawParams struct {
	IsOriginal           bool
	Sender               string
	Receiver             string
	NftID                string
	OriginalTokenAddress string
	TokenAddress         string
	BridgeAddress        string
}

type NonFungibleDepositParams struct {
	IsOriginal           bool
	Sender               string
	Receiver             string
	NftID                string
	TokenAddress         string
	BridgeAddress        string
	OriginalTokenAddress string
	WithdrawTxHash       string
	NftMetadata          data.NFTMetadata
}

type FungibleWithdrawParams struct {
	IsOriginal    bool
	Sender        string
	Receiver      string
	Amount        *resources.Amount
	TokenAddress  string
	BridgeAddress string
}

type FungibleDepositParams struct {
	IsOriginal     bool
	Sender         string
	Receiver       string
	Amount         *resources.Amount
	TokenAddress   string
	BridgeAddress  string
	WithdrawTxHash string
}

//type FungibleWithdrawMetadata struct {
//	TxHash       string
//	Receiver     string
//	TokenAddress string
//	Amount       *resources.Amount
//}
