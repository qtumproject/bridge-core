package types

import (
	"gitlab.com/tokend/bridge/core/internal/data"
)

type Proxy interface {
	Approve(chain data.TokenChain, approveFrom string) (interface{}, error)
}
