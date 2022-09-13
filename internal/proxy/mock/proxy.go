package mock

import (
	"encoding/json"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
)

func NewProxy() types.Proxy {
	return &proxy{}
}

type proxy struct{}

func (p *proxy) GetNftMetadata(tokenAddress string, nftID string) (*data.NFTMetadata, error) {
	return &data.NFTMetadata{}, nil
}

func (p *proxy) CreateNonFungibleWithdrawTx(params types.NonFungibleWithdrawParams) (json.RawMessage, error) {
	return nil, nil
}

func (p *proxy) CreateNonFungibleDepositTx(params types.NonFungibleDepositParams) (json.RawMessage, error) {
	return nil, nil
}
