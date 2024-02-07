package enums

import "gitlab.com/tokend/bridge/core/internal/data"

type Erc20BridgingType uint8

const (
	Erc20BridgingTypeLP Erc20BridgingType = iota
	Erc20BridgingTypeWrapped
	Erc20BridgingTypeUSDC
)

func ToErc20BridgingType(t data.BridgingType) Erc20BridgingType {
	switch t {
	case data.BridgingTypeLP:
		return Erc20BridgingTypeLP
	case data.BridgingTypeWrapped:
		return Erc20BridgingTypeWrapped
	case data.BridgeTypeUSDC:
		return Erc20BridgingTypeUSDC
	default:
		panic("unsupported bridging type")
	}
}

type Erc721BridgingType uint8

const (
	Erc721BridgingTypeLP Erc721BridgingType = iota
	Erc721BridgingTypeWrapped
)

func ToErc721BridgingType(t data.BridgingType) Erc721BridgingType {
	switch t {
	case data.BridgingTypeLP:
		return Erc721BridgingTypeLP
	case data.BridgingTypeWrapped:
		return Erc721BridgingTypeWrapped
	default:
		panic("unsupported bridging type")
	}
}

type Erc1155BridgingType uint8

const (
	Erc1155BridgingTypeLP Erc1155BridgingType = iota
	Erc1155BridgingTypeWrapped
)

func ToErc1155BridgingType(t data.BridgingType) Erc1155BridgingType {
	switch t {
	case data.BridgingTypeLP:
		return Erc1155BridgingTypeLP
	case data.BridgingTypeWrapped:
		return Erc1155BridgingTypeWrapped
	default:
		panic("unsupported bridging type")
	}
}
