/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ChainAttributes struct {
	// Address of bridge contract in specific chain
	BridgeContract string      `json:"bridge_contract"`
	ChainParams    interface{} `json:"chain_params"`
	ChainType      ChainType   `json:"chain_type"`
	// Link to network icon
	Icon *string `json:"icon,omitempty"`
	Name string  `json:"name"`
}
