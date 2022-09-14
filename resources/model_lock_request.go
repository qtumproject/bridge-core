/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "gitlab.com/tokend/bridge/core/internal/amount"

type LockRequest struct {
	// amount of token to transfer, should be presented only for fungible tokens transfer
	Amount    *amount.Amount `json:"amount,omitempty"`
	ChainFrom string         `json:"chain_from"`
	ChainTo   string         `json:"chain_to"`
	// id of NFT to transfer, should be presented only for NFT transfer
	NftId    *string `json:"nft_id,omitempty"`
	Receiver string  `json:"receiver"`
	Sender   string  `json:"sender"`
	TokenId  string  `json:"token_id"`
}
