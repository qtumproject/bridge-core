package qtum

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/qtum/generated/erc1155"
	"gitlab.com/tokend/bridge/core/internal/proxy/qtum/generated/erc721"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

func (p *evmProxy) GetNftMetadata(tokenChain data.TokenChain, nftId string) (*types.NftMetadata, error) {
	metadataUri, err := p.GetNftMetadataUri(tokenChain, nftId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get token uri")
	}

	return p.getMetadata(metadataUri, nftId)
}

func (p *evmProxy) GetNftMetadataUri(tokenChain data.TokenChain, nftId string) (string, error) {
	nft, ok := big.NewInt(0).SetString(nftId, 10)
	if !ok {
		return "", errors.New("failed to parse nft id")
	}

	var metadataUri string
	var err error
	switch tokenChain.TokenType {
	case TokenTypeErc721:
		metadataUri, err = p.getErc721TokenUri(*tokenChain.ContractAddress, nft)
	case TokenTypeErc1155:
		metadataUri, err = p.getErc1155TokenUri(*tokenChain.ContractAddress, nft)
	default:
		return "", errors.New("unsupported token type")
	}
	if err != nil {
		return "nil", errors.Wrap(err, "failed to get token uri")
	}

	return metadataUri, nil
}

func (p *evmProxy) getErc721TokenUri(tokenAddress string, nftId *big.Int) (string, error) {
	token, err := erc721.NewErc721(common.HexToAddress(tokenAddress), p.client)
	if err != nil {
		return "", errors.Wrap(err, "failed to create erc721 token")
	}

	uri, err := token.TokenURI(&bind.CallOpts{}, nftId)
	if err != nil {
		return "", errors.Wrap(err, "failed to get token uri")
	}

	return uri, nil
}

func (p *evmProxy) getErc1155TokenUri(tokenAddress string, nftId *big.Int) (string, error) {
	token, err := erc1155.NewErc1155(common.HexToAddress(tokenAddress), p.client)
	if err != nil {
		return "", errors.Wrap(err, "failed to create erc1155 token")
	}

	uri, err := token.Uri(&bind.CallOpts{}, nftId)
	if err != nil {
		return "", errors.Wrap(err, "failed to get token uri")
	}

	return uri, nil
}

func (p *evmProxy) getMetadata(originalUrl string, nftId string) (*types.NftMetadata, error) {
	// Wrap some custom metadata urls
	url := p.ipfsClient.WrapUrl(originalUrl)
	url = wrapOpenseaUrl(url, nftId)
	url = interpolateUrl(url, nftId)

	metadata, err := getMetadata(url)
	if err != nil {
		return nil, err
	}

	metadata.MetadataUrl = originalUrl
	// Wrap ipfs urls for easier frontend access
	metadata.IconURL = p.ipfsClient.WrapUrl(metadata.IconURL)
	metadata.AnimationUrl = p.wrapPointerIpfsUrl(metadata.AnimationUrl)
	metadata.ExternalUrl = p.wrapPointerIpfsUrl(metadata.ExternalUrl)

	return metadata, nil
}

func (p *evmProxy) wrapPointerIpfsUrl(url *string) *string {
	if url == nil {
		return nil
	}

	wrapped := p.ipfsClient.WrapUrl(*url)
	return &wrapped
}

func wrapOpenseaUrl(url string, nftId string) string {
	if !strings.HasPrefix(url, "https://api.opensea.io") {
		return url
	}

	// URL with 0x don't work
	return strings.Replace(url, "0x{id}", nftId, 1)
}

func interpolateUrl(url string, nftId string) string {
	return strings.Replace(url, "{id}", nftId, 1)
}

func getMetadata(url string) (*types.NftMetadata, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get metadata")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read metadata")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, types.ErrNotFound
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, errors.Errorf("failed to get metadata status %d, body %s", resp.StatusCode, string(body))
	}

	metadata := &types.NftMetadata{}
	err = json.Unmarshal(body, metadata)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode metadata")
	}

	return metadata, nil
}
