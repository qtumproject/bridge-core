package signature

import (
	"github.com/ethereum/go-ethereum/common"
	sha3 "github.com/miguelmota/go-solidity-sha3"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/enums"
	"math/big"
)

type NativeLog struct {
	Amount     *big.Int
	Receiver   string
	TxHash     common.Hash
	EventIndex int
	ChainID    *big.Int
}

func (log NativeLog) Hash() []byte {
	return sha3.SoliditySHA3(
		sha3.Uint256(log.Amount),
		sha3.Address(log.Receiver),
		sha3.Bytes32(log.TxHash.String()),
		sha3.Uint256(big.NewInt(int64(log.EventIndex))),
		sha3.Uint256(log.ChainID),
	)
}

type Erc20Log struct {
	TokenAddress string
	Amount       *big.Int
	Receiver     string
	TxHash       common.Hash
	EventIndex   int
	ChainID      *big.Int
	BridgingType enums.Erc20BridgingType
}

func (log Erc20Log) Hash() []byte {
	return sha3.SoliditySHA3(
		sha3.Address(log.TokenAddress),
		sha3.Uint256(log.Amount),
		sha3.Address(log.Receiver),
		sha3.Bytes32(log.TxHash.String()),
		sha3.Uint256(big.NewInt(int64(log.EventIndex))),
		sha3.Uint256(log.ChainID),
		sha3.Uint8(uint8(log.BridgingType)),
	)
}

type Erc721Log struct {
	TokenAddress string
	TokenID      *big.Int
	Receiver     string
	TxHash       common.Hash
	EventIndex   int
	ChainID      *big.Int
	TokenUri     string
	BridgingType enums.Erc721BridgingType
}

func (log Erc721Log) Hash() []byte {
	return sha3.SoliditySHA3(
		sha3.Address(log.TokenAddress),
		sha3.Uint256(log.TokenID),
		sha3.Address(log.Receiver),
		sha3.Bytes32(log.TxHash.String()),
		sha3.Uint256(big.NewInt(int64(log.EventIndex))),
		sha3.Uint256(log.ChainID),
		sha3.String(log.TokenUri),
		sha3.Uint8(uint8(log.BridgingType)),
	)
}

type Erc1155Log struct {
	TokenAddress string
	TokenID      *big.Int
	Amount       *big.Int
	Receiver     string
	TxHash       common.Hash
	EventIndex   int
	ChainID      *big.Int
	TokenUri     string
	BridgingType enums.Erc1155BridgingType
}

func (log Erc1155Log) Hash() []byte {
	return sha3.SoliditySHA3(
		sha3.Address(log.TokenAddress),
		sha3.Uint256(log.TokenID),
		sha3.Uint256(log.Amount),
		sha3.Address(log.Receiver),
		sha3.Bytes32(log.TxHash.String()),
		sha3.Uint256(big.NewInt(int64(log.EventIndex))),
		sha3.Uint256(log.ChainID),
		sha3.String(log.TokenUri),
		sha3.Uint8(uint8(log.BridgingType)),
	)
}
