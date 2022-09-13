package proxy

import (
	"fmt"
	"github.com/pkg/errors"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/mock"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"gitlab.com/tokend/bridge/core/resources"
)

type ProxyRepo interface {
	Get(chainID string) types.Proxy
}

func NewProxyRepo(chains []data.Chain) (ProxyRepo, error) {
	repo := proxyRepo{
		proxies: make(map[string]types.Proxy),
	}

	for _, c := range chains {
		switch c.Type {
		case resources.EVM:
			repo.proxies[c.ID] = mock.NewProxy()
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
