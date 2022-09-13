package resources

type ApproveRequest struct {
	Address   string `json:"address"`
	Network   string `json:"network"`
	TokenId   string `json:"token_id"`
	AddressTo string `json:"address_to"`
}
