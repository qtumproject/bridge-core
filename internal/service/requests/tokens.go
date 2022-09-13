package requests

import (
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/bridge/core/resources"
	"net/http"
)

type TokensRequest struct {
	FilterType    []resources.TokenType `filter:"token_type"`
	IncludeChains bool                  `include:"chain"`
}

func NewTokensRequest(r *http.Request) (TokensRequest, error) {
	request := TokensRequest{}
	err := urlval.DecodeSilently(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}
