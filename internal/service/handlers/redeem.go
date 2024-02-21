package handlers

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"gitlab.com/tokend/bridge/core/internal/service/models"
	"gitlab.com/tokend/bridge/core/internal/service/requests"
	"gitlab.com/tokend/bridge/core/resources"
	"net/http"
)

func Redeem(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRedeemRequest(r)
	if err != nil {
		Log(r).WithError(err).Debug("failed to parse request")
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

	sourceChain, err := TokenChainsQ(r).
		FilterByTokenID(req.TokenId).
		FilterByChainID(req.ChainFrom).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get token chain")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if sourceChain == nil {
		Log(r).WithError(err).Debug("token chain not found")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data/chain_from": errors.New("token that you have sent does not connected to this chain"),
		})...)
		return
	}

	switch token.Type {
	case resources.FUNGIBLE:
		redeemFungibleToken(w, r, req, *sourceChain)
	case resources.NON_FUNGIBLE:
		redeemNonFungibleToken(w, r, req, *sourceChain)
	default:
		Log(r).WithError(err).
			WithField("token", token.ID).
			WithField("type", token.Type).
			Error("token type not supported")
		ape.RenderErr(w, problems.InternalError())
	}
}

func redeemFungibleToken(w http.ResponseWriter, r *http.Request, req resources.RedeemRequest, sourceChain data.TokenChain) {
	event, err := ProxyRepo(r).Get(sourceChain.ChainID).CheckFungibleLockEvent(req.TxHash, *req.EventIndex, sourceChain)
	if err != nil {
		renderCheckEventError(w, r, err)
		return
	}

	destChain, err := TokenChainsQ(r).
		FilterByTokenID(req.TokenId).
		FilterByChainID(event.DestinationChain).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get token chain")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if destChain == nil {
		Log(r).WithError(err).Debug("token chain not found")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data/chain_to": errors.New("token that you have sent does not connected to this chain"),
		})...)
		return
	}

	sender := event.Receiver
	if req.Sender != nil {
		sender = *req.Sender
	}

	params := types.FungibleRedeemParams{
		TokenChain: *destChain,
		Amount:     event.Amount,
		Receiver:   event.Receiver,
		Sender:     sender,
		TxHash:     req.TxHash,
		EventIndex: *req.EventIndex,
	}

	if req.RawTxData != nil {
		rawTxData, err := hexutil.Decode(*req.RawTxData)
		if err != nil {
			return
		}

		params.RawTxData = &rawTxData
	}

	tx, err := ProxyRepo(r).Get(destChain.ChainID).RedeemFungible(params)
	if err != nil {
		renderRedeemError(w, r, err)
		return
	}

	chain, err := ChainsQ(r).FilterByID(destChain.ChainID).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get chain")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewTxResponse(tx, *chain))
}

func redeemNonFungibleToken(w http.ResponseWriter, r *http.Request, req resources.RedeemRequest, sourceChain data.TokenChain) {
	event, err := ProxyRepo(r).Get(sourceChain.ChainID).CheckNonFungibleLockEvent(req.TxHash, *req.EventIndex, sourceChain)
	if err != nil {
		renderCheckEventError(w, r, err)
		return
	}

	destChain, err := TokenChainsQ(r).
		FilterByTokenID(req.TokenId).
		FilterByChainID(event.DestinationChain).
		Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get token chain")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if destChain == nil {
		Log(r).WithError(err).Debug("token chain not found")
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data/chain_to": errors.New("token that you have sent does not connected to this chain"),
		})...)
		return
	}

	uri, err := getOriginalUri(r, req.TokenId, event.NftId)
	if err != nil {
		Log(r).WithError(err).Error("failed to get original uri")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	sender := event.Receiver
	if req.Sender != nil {
		sender = *req.Sender
	}

	params := types.NonFungibleRedeemParams{
		TokenChain: *destChain,
		Receiver:   event.Receiver,
		Sender:     sender,
		TxHash:     req.TxHash,
		EventIndex: *req.EventIndex,
		NftId:      event.NftId,
		NftUri:     *uri,
	}

	if req.RawTxData != nil {
		rawTxData, err := hexutil.Decode(*req.RawTxData)
		if err != nil {
			return
		}

		params.RawTxData = &rawTxData
	}

	tx, err := ProxyRepo(r).Get(destChain.ChainID).RedeemNonFungible(params)
	if err != nil {
		renderRedeemError(w, r, err)
		return
	}

	chain, err := ChainsQ(r).FilterByID(destChain.ChainID).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get chain")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewTxResponse(tx, *chain))
}

func getOriginalUri(r *http.Request, tokenId string, nftId string) (*string, error) {
	chain, err := TokenChainsQ(r).
		FilterByTokenID(tokenId).
		// TODO: Add original chain?
		FilterByBridgingType(data.BridgingTypeLP).
		Get()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get original chain")
	}
	if chain == nil {
		return nil, errors.New("original chain not found")
	}

	uri, err := ProxyRepo(r).Get(chain.ChainID).GetNftMetadataUri(*chain, nftId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get nft metadata uri")
	}

	return &uri, nil
}

func renderCheckEventError(w http.ResponseWriter, r *http.Request, err error) {
	if err == types.ErrTxNotFound {
		Log(r).WithError(err).Debug("tx not found")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusBadRequest),
			Status: fmt.Sprintf("%d", http.StatusBadRequest),
			Detail: "tx with this hash not found, try different network or tx hash",
			Code:   "tx_not_found",
		})
		return
	}
	if err == types.ErrWrongToken {
		Log(r).WithError(err).Debug("wrong token")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusBadRequest),
			Status: fmt.Sprintf("%d", http.StatusBadRequest),
			Detail: "transaction locks different token from the token specified in request",
			Code:   "wrong_token",
		})
		return
	}
	if err == types.ErrTxNotConfirmed {
		Log(r).WithError(err).Debug("tx not confirmed")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusBadRequest),
			Status: fmt.Sprintf("%d", http.StatusBadRequest),
			Detail: "tx not confirmed",
			Code:   "not_confirmed",
		})
		return
	}
	if err == types.ErrEventNotFound || err == types.ErrWrongLockEvent || err == types.ErrTxFailed {
		Log(r).WithError(err).Debug("invalid transaction")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusBadRequest),
			Status: fmt.Sprintf("%d", http.StatusBadRequest),
			Detail: "invalid transaction",
			Code:   "invalid_transaction",
		})
		return
	}
	Log(r).WithError(err).Error("failed to check lock event")
	ape.RenderErr(w, problems.InternalError())
	return
}

func renderRedeemError(w http.ResponseWriter, r *http.Request, err error) {
	if err == types.ErrAlreadyRedeemed {
		Log(r).WithError(err).Debug("already redeemed")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusBadRequest),
			Status: fmt.Sprintf("%d", http.StatusBadRequest),
			Detail: "already redeemed",
			Code:   "already_redeemed",
		})
		return
	}
	if err == types.ErrWrongSignedTx {
		Log(r).WithError(err).Debug("wrong signed tx")
		ape.RenderErr(w, &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusBadRequest),
			Status: fmt.Sprintf("%d", http.StatusBadRequest),
			Detail: "Params that was provided in transaction for signing and in event is different",
			Code:   "wrong_signed_tx",
		})
		return
	}
	Log(r).WithError(err).Error("failed to redeem token")
	ape.RenderErr(w, problems.InternalError())
	return
}
