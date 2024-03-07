package handlers

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/amount"
	"gitlab.com/tokend/bridge/core/internal/data"
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
	if destTokenChain.BridgingType == data.BridgingTypeLP {
		balance, err := ProxyRepo(r).Get(destTokenChain.ChainID).BridgeBalance(*destTokenChain, req.NftId)
		if err != nil {
			Log(r).WithError(err).Error("failed to get bridge balance")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		am := amount.NewFromBigInt(amount.One)
		if req.Amount != nil {
			am = *req.Amount
		}
		if amount.Cmp(balance, am) == -1 {
			Log(r).WithError(err).Debug("insufficient balance in destination liquidity pool")
			ape.RenderErr(w, &jsonapi.ErrorObject{
				Title:  http.StatusText(http.StatusBadRequest),
				Status: fmt.Sprintf("%d", http.StatusBadRequest),
				Detail: "Insufficient balance in destination liquidity pool",
				Code:   "insufficient_lp_balance",
			})
			return
		}
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
			Amount:           req.Amount,
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
