package proxy

import (
	"fmt"
	"github.com/pkg/errors"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"gitlab.com/tokend/bridge/core/internal/signature"
	"gitlab.com/tokend/bridge/core/resources"
)

type ProxyRepo interface {
	Get(chainID string) types.Proxy
}

func NewProxyRepo(chains []data.Chain, signer signature.Signer) (ProxyRepo, error) {
	repo := proxyRepo{
		proxies: make(map[string]types.Proxy),
	}

	for _, c := range chains {
		switch c.Type {
		case resources.EVM:
			proxy, err := evm.NewProxy(c.RpcEndpoint, signer, c.BridgeContract, c.Confirmations)
			if err != nil {
				return nil, errors.Wrap(err, "failed to create evm proxy")
			}
			repo.proxies[c.ID] = proxy
		default:
			return nil, errors.New(fmt.Sprintf("Unsupported network type"))
		}
	}

	return &repo, nil
}

type proxyRepo struct {
	proxies map[string]types.Proxy
}

func (r *proxyRepo) Get(chainID string) types.Proxy {
	return r.proxies[chainID]
}
