package handlers

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"gitlab.com/tokend/bridge/core/internal/service/models"
	"gitlab.com/tokend/bridge/core/internal/service/requests"
	"net/http"
)

func Deposit(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewDepositRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	chain, err := ChainsQ(r).FilterByID(req.NetworkTo).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get chain")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if chain == nil {
		Log(r).WithError(err).Debug("provided chain not found")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"network_from": errors.New("such network does not exist"),
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
		if tc.ChainID == req.NetworkTo {
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

	originalChain, err := TokenChainsQ(r).
		FilterByTokenID(req.TokenId).
		FilterByIsOriginal(true).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to find original chain")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	metadata, err := ProxyRepo(r).
		Get(originalChain.ChainID).
		GetNftMetadata(originalChain.ContractAddress, req.NftId)
	if err != nil {
		Log(r).WithError(err).Error("failed to get nft metadata")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tx, err := ProxyRepo(r).Get(req.NetworkTo).CreateNonFungibleDepositTx(types.NonFungibleDepositParams{
		IsOriginal:           tokenChain.IsOriginal,
		Sender:               req.AddressFrom,
		Receiver:             req.AddressTo,
		NftID:                req.NftId,
		TokenAddress:         tokenChain.ContractAddress,
		BridgeAddress:        chain.BridgeContract,
		WithdrawTxHash:       req.TxHash,
		OriginalTokenAddress: originalChain.ContractAddress,
		NftMetadata:          *metadata,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to build tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewTransactionResponse(tx))
}
