package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/service/models"
	"gitlab.com/tokend/bridge/core/internal/service/requests"
	"net/http"
)

func Chains(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewChainsRequest(r)
	if err != nil {
		Log(r).WithError(err).Debug("request is incorrect")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	chainsQ := ChainsQ(r)

	if len(request.FilterType) > 0 {
		chainsQ.FilterByType(request.FilterType...)
	}
	chains, err := chainsQ.Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to get chains")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	chainsPage, err := chainsQ.Page(r.URL.Query().Get("limit"), r.URL.Query().Get("page_number"), r.URL.Path, chains)
	if err != nil {
		Log(r).WithError(err).Error("failed to pagination")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	var tokens []data.Token
	if request.IncludeTokens {
		tokens, err = TokensQ(r).FilterByID(tokensId(chains)...).Select()
		if err != nil {
			Log(r).WithError(err).Error("failed to get tokens")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	ape.Render(w, models.NewChainsResponse(chainsPage, tokens)) //todo change docs and add new response model
	//ape.Render(w, models.NewChainListResponse(chainsPage.Items, tokens)) //todo change docs and add new response model
}

func tokensId(chains []data.Chain) []string {
	result := make(map[string]struct{})
	for _, chain := range chains {
		for _, token := range chain.Tokens {
			result[token.TokenID] = struct{}{}
		}
	}

	keys := make([]string, 0, len(result))
	for key, _ := range result {
		keys = append(keys, key)
	}

	return keys
}
