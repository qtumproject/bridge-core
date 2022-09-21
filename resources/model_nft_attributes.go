/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type NftAttributes struct {
	AnimationUrl *string        `json:"animation_url,omitempty"`
	Attributes   []NftAttribute `json:"attributes"`
	Description  *string        `json:"description,omitempty"`
	ExternalUrl  *string        `json:"external_url,omitempty"`
	// Link to icon
	Image string `json:"image"`
	// original url to metadata stored in the contract
	MetadataUrl string `json:"metadata_url"`
	Name        string `json:"name"`
}
