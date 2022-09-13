package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/urlval"
	"gitlab.com/tokend/bridge/core/resources"
	"net/http"
)

func NewDepositRequest(r *http.Request) (resources.TransferConfirmationRequest, error) {
	request := resources.TransferConfirmationRequest{}
	err := urlval.DecodeSilently(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, ValidateDepositRequest(request)
}

func ValidateDepositRequest(r resources.TransferConfirmationRequest) error {
	errs := validation.Errors{
		"/address_from": validation.Validate(&r.AddressFrom, validation.Required),
		"/address_to":   validation.Validate(&r.AddressTo, validation.Required),
		"/network_from": validation.Validate(&r.NetworkFrom, validation.Required),
		"/network_to":   validation.Validate(&r.NetworkTo, validation.Required),
		"/nft_id":       validation.Validate(&r.NftId, validation.Required),
		"/token_id":     validation.Validate(&r.TokenId, validation.Required),
		"/tx_hash":      validation.Validate(&r.TxHash, validation.Required),
	}

	return errs.Filter()
}
