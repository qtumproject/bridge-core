package handlers

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/service/models"
	"gitlab.com/tokend/bridge/core/internal/service/requests"
	"gitlab.com/tokend/bridge/core/resources"
	"net/http"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetBalanceRequest(r)
	if err != nil {
		Log(r).WithError(err).Debug("failed to decode request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	tokenChain, err := TokenChainsQ(r).
		FilterByTokenID(req.TokenId).
		FilterByChainID(req.Chain).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get token chain")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if tokenChain == nil {
		Log(r).WithError(err).Debug("token chain not found")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data": errors.New("token that you have sent does not connected to this network or does not exist"),
		})...)
		return
	}

	token, err := TokensQ(r).FilterByID(tokenChain.TokenID).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get token")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if token.Type == resources.NON_FUNGIBLE && req.Nft == nil {
		Log(r).WithError(err).Debug("nft id is required")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"nft": errors.New("nft id is required for balance of non-fungible token"),
		})...)
		return
	}

	balance, err := ProxyRepo(r).Get(tokenChain.ChainID).Balance(*tokenChain, req.Address, req.Nft)
	if err != nil {
		Log(r).WithError(err).Error("failed to get balance")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewBalanceResponse(req.TokenId, req.Address, req.Nft, balance))
}
