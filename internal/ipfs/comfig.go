package ipfs

import (
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type IpfsClienter interface {
	IpfsClient() Client
}

func NewIpfsClienter(getter kv.Getter) IpfsClienter {
	return &ipfsClienter{
		getter: getter,
	}
}

type ipfsClienter struct {
	once   comfig.Once
	getter kv.Getter
}

func (c *ipfsClienter) IpfsClient() Client {
	return c.once.Do(func() interface{} {
		config := struct {
			Endpoint string `fig:"endpoint,required"`
		}{}

		err := figure.
			Out(&config).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(c.getter, "ipfs")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out ipfs"))
		}

		return NewClient(config.Endpoint)
	}).(Client)
}
