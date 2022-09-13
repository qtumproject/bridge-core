package handlers

import (
	"net/http"
)

func Nft(w http.ResponseWriter, r *http.Request) {
	//req, err := requests.NewNftDetailsRequest(r)
	//if err != nil {
	//	Log(r).WithError(err).Debug("request is incorrect")
	//	ape.RenderErr(w, problems.BadRequest(err)...)
	//	return
	//}
	//
	//token, err := TokensQ(r).FilterByID(req.TokenID).Get()
	//if err != nil {
	//	Log(r).WithError(err).Error("failed to get token")
	//	ape.RenderErr(w, problems.InternalError())
	//	return
	//}
	//if token == nil {
	//	Log(r).WithError(err).Debug("provided token not found")
	//	ape.RenderErr(w, problems.NotFound())
	//	return
	//}
	//var tokenChain *data.TokenChain
	//for _, tc := range token.Chains {
	//	if tc.IsOriginal {
	//		tokenChain = &tc
	//		break
	//	}
	//}
	//if tokenChain == nil {
	//	Log(r).WithError(err).Error("token don't have original chain")
	//	ape.RenderErr(w, problems.InternalError())
	//	return
	//}
	//
	//attr, err := ProxyRepo(r).
	//	Get(tokenChain.ChainID).
	//	GetNftMetadata(tokenChain.ContractAddress, req.NftID)
	//if err != nil {
	//	Log(r).WithError(err).Error("failed to get token details")
	//	ape.RenderErr(w, problems.InternalError())
	//	return
	//}
	//
	//ape.Render(w, models.NewNftMetadataResponse(req.NftID, *attr))
}
