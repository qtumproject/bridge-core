package signature

import (
	sha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
)

type DepositFungibleLog struct {
	TxHash       string
	ChainID      *big.Int
	TokenAddress string
	Receiver     string
	Amount       *big.Int
}

func (log DepositFungibleLog) Hash() []byte {
	return sha3.SoliditySHA3(
		sha3.Uint256(log.ChainID),
		sha3.Address(log.TokenAddress),
		sha3.Address(log.Receiver),
		sha3.String(log.TxHash),
		sha3.Uint256(log.Amount),
	)
}
