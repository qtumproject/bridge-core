/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type RedeemRequest struct {
	ChainFrom string `json:"chain_from"`
	// index of event in transaction on chain_from that locked tokens, use it only if lock transaction has a few locking events
	EventIndex *int `json:"event_index,omitempty"`
	// redeem tx data with collected signatures
	RawTxData *string `json:"raw_tx_data,omitempty"`
	// address that will send transaction to chain_to, optional by default it will be the same as receiver on source chain
	Sender  *string `json:"sender,omitempty"`
	TokenId string  `json:"token_id"`
	// hash of transaction on chain_from that locked tokens
	TxHash string `json:"tx_hash"`
}
