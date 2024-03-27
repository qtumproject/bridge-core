package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/tokend/bridge/core/internal/data/mem"
	"gitlab.com/tokend/bridge/core/internal/ipfs"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/signature"
)

type Config interface {
	comfig.Logger
	comfig.Listenerer
	signature.Signerer
	mem.Chainer
	ipfs.IpfsClienter
}

type config struct {
	comfig.Logger
	comfig.Listenerer
	signature.Signerer
	mem.Chainer
	ipfs.IpfsClienter
	getter kv.Getter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:       getter,
		Listenerer:   comfig.NewListenerer(getter),
		Logger:       comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Signerer:     signature.NewSignerer(getter),
		Chainer:      mem.NewChainer(getter),
		IpfsClienter: ipfs.NewIpfsClienter(getter),
	}
}
