package requests

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/resources"
	"net/http"
)

func NewApproveRequest(r *http.Request) (resources.ApproveRequest, error) {
	request := struct {
		Data resources.ApproveRequest
	}{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request.Data, errors.Wrap(err, "failed to decode request")
	}

	return request.Data, nil
}
