package relayer

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Relayer interface {
	SendTx(tx *ethTypes.Transaction, chainID *big.Int) (common.Hash, error)
}
