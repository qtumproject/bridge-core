/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Nft struct {
	Key
	Attributes NftAttributes `json:"attributes"`
}
type NftResponse struct {
	Data     Nft      `json:"data"`
	Included Included `json:"included"`
}

type NftListResponse struct {
	Data     []Nft    `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustNft - returns Nft from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustNft(key Key) *Nft {
	var nFT Nft
	if c.tryFindEntry(key, &nFT) {
		return &nFT
	}
	return nil
}
