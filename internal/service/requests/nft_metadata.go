package requests

import (
	"github.com/go-chi/chi"
	"net/http"
)

type NftMetadataRequest struct {
	NftID   string
	TokenID string
}

func NewNftDetailsRequest(r *http.Request) (NftMetadataRequest, error) {
	request := NftMetadataRequest{}
	request.TokenID = chi.URLParam(r, "token_id")
	request.NftID = chi.URLParam(r, "nft_id")

	return request, nil
}
