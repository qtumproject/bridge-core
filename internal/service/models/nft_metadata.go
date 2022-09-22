package models

import (
	"fmt"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"gitlab.com/tokend/bridge/core/resources"
)

func NewNftMetadataResponse(nftId string, model types.NftMetadata) resources.NftResponse {
	result := resources.NftResponse{
		Data: resources.Nft{
			Key: resources.Key{
				ID:   fmt.Sprintf("%s", nftId),
				Type: resources.NFT,
			},
			Attributes: resources.NftAttributes{
				Name:         model.Name,
				Image:        model.IconURL,
				MetadataUrl:  model.MetadataUrl,
				Description:  model.Description,
				ExternalUrl:  model.ExternalUrl,
				AnimationUrl: model.AnimationUrl,
				Attributes:   make([]resources.NftAttribute, 0, len(model.Attributes)),
			},
		},
	}

	for _, attr := range model.Attributes {
		result.Data.Attributes.Attributes = append(result.Data.Attributes.Attributes, resources.NftAttribute{
			TraitType: attr.Trait,
			Value:     attr.Value,
		})
	}

	return result
}
