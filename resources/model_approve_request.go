/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "gitlab.com/tokend/bridge/core/internal/amount"

type ApproveRequest struct {
	Address string `json:"address"`
	ChainId string `json:"chain_id"`
	TokenId string `json:"token_id"`
	// id of NFT to transfer, should be presented only for NFT transfer
	NftId *string `json:"nft_id,omitempty"`
	// amount of token to transfer, should be presented only for fungible tokens transfer
	Amount *amount.Amount `json:"amount,omitempty"`
}
