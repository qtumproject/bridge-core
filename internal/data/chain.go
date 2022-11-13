package data

import (
	"encoding/json"
	"gitlab.com/tokend/bridge/core/resources"
)

type ChainsQ interface {
	New() ChainsQ
	Select() ([]Chain, error)
	Get() (*Chain, error)
	Page(limitStr, currentPageStr, path string, chains []Chain) (ChainsQList, error)
	FilterByID(ids ...string) ChainsQ
	FilterByType(types ...resources.ChainType) ChainsQ
}

type ChainsQList struct {
	Items      []Chain `json:"items"`
	NextPageID string  `json:"next_page_id,omitempty" example:"10"`
}
type Chain struct {
	ID          string              `fig:"id,required"`
	Name        string              `fig:"name,required"`
	Icon        *string             `fig:"icon"`
	Type        resources.ChainType `fig:"type,required"`
	ChainParams json.RawMessage     `fig:"chain_params"`
	
	BridgeContract string `fig:"bridge_contract,required"`
	RpcEndpoint    string `fig:"rpc_endpoint,required"`
	// Relation
	Tokens []TokenChain
}
