package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"gitlab.com/tokend/bridge/core/internal/signature"
	"math/big"
)

const (
	tokenTypeNative  = "native"
	tokenTypeErc20   = "erc20"
	tokenTypeErc721  = "erc721"
	tokenTypeErc1155 = "erc1155"
)

func NewProxy(rpc string, signer signature.Signer, bridgeContract string) (types.Proxy, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.TODO())
	if err != nil {
		return nil, err
	}

	return &evmProxy{
		client:         client,
		signer:         signer,
		chainID:        chainID,
		bridgeContract: common.HexToAddress(bridgeContract),
	}, nil
}

type evmProxy struct {
	client         *ethclient.Client
	signer         signature.Signer
	chainID        *big.Int
	bridgeContract common.Address
}
