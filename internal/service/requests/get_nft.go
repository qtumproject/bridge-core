package requests

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type GetNftRequest struct {
	TokenId string `url:"-"`
	NftId   string `url:"-"`
	Chain   string `url:"chain"`
}

func NewGetNftRequest(r *http.Request) (GetNftRequest, error) {
	var req GetNftRequest
	if err := urlval.DecodeSilently(r.URL.Query(), &req); err != nil {
		return req, errors.Wrap(err, "failed to decode request")
	}

	req.TokenId = chi.URLParam(r, "id")
	req.NftId = chi.URLParam(r, "nft_id")

	return req, nil
}
