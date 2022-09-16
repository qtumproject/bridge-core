package requests

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/urlval"
	"net/http"
)

type GetBalanceRequest struct {
	TokenId string  `url:"-"`
	Address string  `url:"address"`
	Chain   string  `url:"chain"`
	Nft     *string `url:"nft"`
}

func NewGetBalanceRequest(r *http.Request) (GetBalanceRequest, error) {
	var req GetBalanceRequest
	if err := urlval.DecodeSilently(r.URL.Query(), &req); err != nil {
		return req, errors.Wrap(err, "failed to decode request")
	}

	req.TokenId = chi.URLParam(r, "id")

	return req, nil
}
