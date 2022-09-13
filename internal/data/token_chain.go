package data

type TokenChainsQ interface {
	New() TokenChainsQ
	Select() ([]TokenChain, error)
	Get() (*TokenChain, error)

	FilterByTokenID(ids ...string) TokenChainsQ
	FilterByChainID(ids ...string) TokenChainsQ
	FilterByIsOriginal(isOriginal bool) TokenChainsQ
}

type TokenChain struct {
	TokenID         string
	ChainID         string `fig:"chain_id,required"`
	ContractAddress string `fig:"contract_address,required"`
	IsOriginal      bool   `fig:"is_original"`
}
