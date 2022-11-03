/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type EvmTransactionAttributes struct {
	// The chain ID encoded as hex
	ChainId string `json:"chain_id"`
	// Is enough signatures to send tx
	Confirmed *bool `json:"confirmed,omitempty"`
	// transaction call data encoded as hex
	Data string `json:"data"`
	// The address of the sender
	From string `json:"from"`
	// The address of transaction recipient
	To string `json:"to"`
	// The amount of wei to send
	Value string `json:"value"`
}
