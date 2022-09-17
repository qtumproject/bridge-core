package types

type NftMetadata struct {
	MetadataUrl  string         `json:"metadata_url"`
	Name         string         `json:"name"`
	IconURL      string         `json:"image"`
	Description  *string        `json:"description"`
	AnimationUrl *string        `json:"animation_url"`
	ExternalUrl  *string        `json:"external_url"`
	Attributes   []NftAttribute `json:"attributes"`
}

type NftAttribute struct {
	Trait string `json:"trait_type"`
	Value string `json:"value"`
}
