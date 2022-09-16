/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ProcessedTransaction struct {
	Key
	Relationships ProcessedTransactionRelationships `json:"relationships"`
}
type ProcessedTransactionResponse struct {
	Data     ProcessedTransaction `json:"data"`
	Included Included             `json:"included"`
}

type ProcessedTransactionListResponse struct {
	Data     []ProcessedTransaction `json:"data"`
	Included Included               `json:"included"`
	Links    *Links                 `json:"links"`
}

// MustProcessedTransaction - returns ProcessedTransaction from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustProcessedTransaction(key Key) *ProcessedTransaction {
	var processedTransaction ProcessedTransaction
	if c.tryFindEntry(key, &processedTransaction) {
		return &processedTransaction
	}
	return nil
}
