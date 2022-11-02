package models

import (
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/service/requests"
	"gitlab.com/tokend/bridge/core/resources"
)

func newTokenKey(id string) resources.Key {
	return resources.Key{
		ID:   id,
		Type: resources.TOKEN,
	}
}

func newTokenModel(value data.Token) *resources.Token {
	return &resources.Token{
		Key: newTokenKey(value.ID),
		Attributes: resources.TokenAttributes{
			Name:      value.Name,
			Symbol:    value.Symbol,
			Icon:      value.Icon,
			TokenType: value.Type,
		},
	}
}

func newTokenModelWithRelation(value data.Token) *resources.Token {
	model := newTokenModel(value)
	model.Relationships = resources.TokenRelationships{
		Chains: resources.RelationCollection{
			Data: make([]resources.Key, len(value.Chains)),
		},
	}
	for i, relation := range value.Chains {
		model.Relationships.Chains.Data[i] = newChainKey(relation.ChainID)
	}

	return model
}

func NewTokenListResponse(tokens []data.Token, chains []data.Chain) resources.TokenListResponse {
	response := resources.TokenListResponse{
		Data:     make([]resources.Token, len(tokens)),
		Included: resources.Included{},
	}

	for i, token := range tokens {
		response.Data[i] = *newTokenModelWithRelation(token)
	}

	for _, chain := range chains {
		response.Included.Add(newChainModel(chain))
	}

	return response
}

func NewTokenResponse(Data data.TokensQList, chains []data.Chain) requests.TokensResponse {
	response := requests.TokensResponse{
		Data:     Data,
		Included: resources.Included{},
	}

	for _, chain := range chains {
		response.Included.Add(newChainModel(chain))
	}

	return response
}
