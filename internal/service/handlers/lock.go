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

func Lock(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewLockRequest(r)
	if err != nil {
		Log(r).WithError(err).Debug("failed to decode request")
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
		Log(r).Debug("token not found")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data/token_id": errors.New("token that you have sent does not exist"),
		})...)
		return
	}

	tokenChain, err := TokenChainsQ(r).
		FilterByTokenID(req.TokenId).
		FilterByChainID(req.ChainFrom).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get token chain")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if tokenChain == nil {
		Log(r).WithError(err).Debug("token chain not found")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data/chain_from": errors.New("token that you have sent does not connected to this chain"),
		})...)
		return
	}

	destTokenChain, err := TokenChainsQ(r).
		FilterByTokenID(req.TokenId).
		FilterByChainID(req.ChainTo).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get token chain")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if destTokenChain == nil {
		Log(r).WithError(err).Debug("token chain not found")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data/chain_to": errors.New("token that you have sent does not connected to this chain"),
		})...)
		return
	}

	var tx interface{}
	switch token.Type {
	case resources.FUNGIBLE:
		if req.Amount == nil {
			Log(r).WithError(err).Debug("amount is not set")
			ape.RenderErr(w, problems.BadRequest(validation.Errors{
				"data/amount": errors.New("amount is not set"),
			})...)
			return
		}
		tx, err = ProxyRepo(r).Get(tokenChain.ChainID).LockFungible(types.FungibleLockParams{
			TokenChain:       *tokenChain,
			Sender:           req.Sender,
			Receiver:         req.Receiver,
			DestinationChain: req.ChainTo,
			Amount:           *req.Amount,
		})
	case resources.NON_FUNGIBLE:
		if req.NftId == nil {
			Log(r).WithError(err).Debug("nft_id is not set")
			ape.RenderErr(w, problems.BadRequest(validation.Errors{
				"data/nft_id": errors.New("nft_id is not set"),
			})...)
			return
		}
		tx, err = ProxyRepo(r).Get(tokenChain.ChainID).LockNonFungible(types.NonFungibleLockParams{
			TokenChain:       *tokenChain,
			Sender:           req.Sender,
			Receiver:         req.Receiver,
			DestinationChain: req.ChainTo,
			NftId:            *req.NftId,
		})
	default:
		Log(r).Errorf("token type is not supported %s, token id - %s", token.Type, token.ID)
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if err != nil {
		Log(r).WithError(err).Error("failed to create lock tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	chain, err := ChainsQ(r).FilterByID(tokenChain.ChainID).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get chain")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewTxResponse(tx, *chain))
}
