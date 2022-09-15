package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/resources"
	"net/http"
)

func NewRedeemRequest(r *http.Request) (resources.RedeemRequest, error) {
	request := struct {
		Data resources.RedeemRequest
	}{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request.Data, errors.Wrap(err, "failed to decode request")
	}

	if request.Data.EventIndex == nil {
		i := 0
		request.Data.EventIndex = &i
	}

	return request.Data, nil
}
