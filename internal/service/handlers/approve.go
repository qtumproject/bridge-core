package handlers

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/service/models"
	"gitlab.com/tokend/bridge/core/internal/service/requests"
	"math/big"
	"net/http"
)

func Approve(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewApproveRequest(r)
	if err != nil {
		Log(r).WithError(err).Debug("failed to decode request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	chain, err := ChainsQ(r).FilterByID(req.Network).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get chain")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if chain == nil {
		Log(r).WithError(err).Debug("provided chain not found")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"network": errors.New("such network does not exist"),
		})...)
		return
	}

	token, err := TokensQ(r).FilterByID(req.TokenId).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get token")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if token == nil {
		Log(r).WithError(err).Debug("provided token not found")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"token_id": errors.New("such token does not exist"),
		})...)
		return
	}
	var tokenChain *data.TokenChain
	for _, tc := range token.Chains {
		if tc.ChainID == req.Network {
			tokenChain = &tc
			break
		}
	}
	if tokenChain == nil {
		Log(r).WithError(err).Debug("provided chain not found")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"network_from": errors.New("not supported network by this token"),
		})...)
		return
	}

	tokenId := new(big.Int)
	tokenId, ok := tokenId.SetString(req.TokenId, 10)
	if !ok {
		Log(r).WithError(err).Error("failed to convert token id to int")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tx, err := ProxyRepo(r).Get(req.Network).CreateApprovalTx(tokenId, tokenChain.ContractAddress, req.Address, req.AddressTo)
	if err != nil {
		Log(r).WithError(err).Error("failed to build tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewTransactionResponse(tx))
}
