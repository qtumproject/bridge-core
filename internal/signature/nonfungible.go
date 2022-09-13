package signature

import (
	sha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
)

type DepositNonFungibleLog struct {
	TxHash       string
	ChainID      *big.Int
	TokenAddress string
	Receiver     string
	TokenID      *big.Int
}

func (log DepositNonFungibleLog) Hash() []byte {
	return sha3.SoliditySHA3(
		sha3.Uint256(log.ChainID),
		sha3.Address(log.TokenAddress),
		sha3.Address(log.Receiver),
		sha3.String(log.TxHash),
		sha3.Uint256(log.TokenID),
	)
}
