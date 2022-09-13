package models

import (
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/resources"
)

func NewNftMetadataResponse(id string, model data.NFTMetadata) resources.NftResponse {
	return resources.NftResponse{
		Data: resources.Nft{
			Key: resources.Key{
				ID:   id,
				Type: resources.NFT,
			},
			Attributes: resources.NftAttributes{
				Name: model.Name,
				Icon: &model.IconURL,
			},
		},
	}
}
