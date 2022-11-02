package data

import (
	"encoding/json"
	"gitlab.com/tokend/bridge/core/resources"
)

type ChainsQ interface {
	New() ChainsQ
	Select() ([]Chain, error)
	Get() (*Chain, error)
	Paginate(limitStr, currentPageStr string, chains []Chain) (ChainsQList, error)
	FilterByID(ids ...string) ChainsQ
	FilterByType(types ...resources.ChainType) ChainsQ
}

type ChainsQList struct {
	// A list of articles
	Items []Chain `json:"items"`
	// The id to query the next page
	NextPageID int `json:"next_page_id,omitempty" example:"10"`
}
type Chain struct {
	ID          string              `fig:"id,required"`
	Name        string              `fig:"name,required"`
	Icon        *string             `fig:"icon"`
	Type        resources.ChainType `fig:"type,required"`
	ChainParams json.RawMessage     `fig:"chain_params"`

	Confirmations  int    `fig:"confirmations,required"`
	BridgeContract string `fig:"bridge_contract,required"`
	RpcEndpoint    string `fig:"rpc_endpoint,required"`
	// Relation
	Tokens []TokenChain
}
