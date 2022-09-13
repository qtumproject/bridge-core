package data

type NFTMetadata struct {
	MetadataURL  string
	Name         string
	IconURL      string
	Description  *string
	AnimationURL *string
	ExternalURL  *string
	Attributes   []NFTAttribute
}

type NFTAttribute struct {
	Trait string
	Value string
}
