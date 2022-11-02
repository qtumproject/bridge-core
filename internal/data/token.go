package data

import (
	"gitlab.com/tokend/bridge/core/resources"
)

type TokensQ interface {
	New() TokensQ
	Select() ([]Token, error)
	Get() (*Token, error)
	PageTokens(pageParams Token) TokensQ
	FilterByID(ids ...string) TokensQ
	FilterByType(types ...resources.TokenType) TokensQ
}

type Token struct {
	ID     string              `fig:"id,required"`
	Name   string              `fig:"name,required"`
	Symbol string              `fig:"symbol,required"`
	Icon   *string             `fig:"icon"`
	Type   resources.TokenType `fig:"type,required"`
	// Relation
	Chains []TokenChain `fig:"chains,required"`
}
