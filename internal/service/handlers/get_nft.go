package handlers

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"gitlab.com/tokend/bridge/core/internal/service/models"
	"gitlab.com/tokend/bridge/core/internal/service/requests"
	"gitlab.com/tokend/bridge/core/resources"
	"net/http"
)

func GetNft(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetNftRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	token, err := TokensQ(r).FilterByID(req.TokenId).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get token")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if token == nil {
		// TODO: not found when will be able to specify custom error
		Log(r).Debug("token not found")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data/token_id": errors.New("token that you have sent does not exist"),
		})...)
		return
	}
	if token.Type != resources.NON_FUNGIBLE {
		Log(r).Debug("token is not non-fungible")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data/token_id": errors.New("token that you have sent is fungible"),
		})...)
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

	metadata, err := ProxyRepo(r).Get(tokenChain.ChainID).GetNftMetadata(*tokenChain, req.NftId)
	if err != nil {
		if err == types.ErrNotFound {
			Log(r).WithError(err).Debug("nft not found")
			ape.RenderErr(w, problems.NotFound())
			return
		}
		Log(r).WithError(err).Error("failed to get nft metadata")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewNftMetadataResponse(req.NftId, *metadata))
}
