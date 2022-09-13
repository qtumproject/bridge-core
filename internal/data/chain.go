package data

import (
	"encoding/json"
	"gitlab.com/tokend/bridge/core/resources"
)

type ChainsQ interface {
	New() ChainsQ
	Select() ([]Chain, error)
	Get() (*Chain, error)

	FilterByID(ids ...string) ChainsQ
	FilterByType(types ...resources.ChainType) ChainsQ
}

type Chain struct {
	ID             string              `fig:"id,required"`
	Name           string              `fig:"name,required"`
	Icon           *string             `fig:"icon"`
	BridgeContract string              `fig:"bridge_contract,required"`
	Type           resources.ChainType `fig:"type,required"`
	ChainParams    json.RawMessage     `fig:"chain_params"`
	ProxyEndpoint  string              `fig:"proxy_endpoint,required"`
	// Relation
	Tokens []TokenChain
}
