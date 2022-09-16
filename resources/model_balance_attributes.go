/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import "gitlab.com/tokend/bridge/core/internal/amount"

type BalanceAttributes struct {
	// Amount of tokens in the balance, for fungible tokens this is the total amount, for non-fungible tokens returns 1 if the token is owned by the account, 0 otherwise.
	Amount amount.Amount `json:"amount"`
}
