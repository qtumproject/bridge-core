package solana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"io"
	"math/big"
	"net/http"
)

const contentType = "application/vnd.api+json"

func NewProxy(endpoint string) types.Proxy {
	return &proxy{
		endpoint: endpoint,
	}
}

type proxy struct {
	endpoint string
}

func (p *proxy) GetNftMetadata(tokenAddress string, nftID string) (*data.NFTMetadata, error) {
	resp, err := http.Post(fmt.Sprintf("%s/integrations/solana-proxy/v1/metadata", p.endpoint),
		contentType,
		encodeBody(map[string]interface{}{
			"data": map[string]interface{}{
				"attributes": map[string]interface{}{
					"address": tokenAddress,
					"tokenId": nftID,
				},
			},
		}))
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request proxy")
	}

	parsedResp, err := parseResponse(resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse response")
	}

	var metadata data.NFTMetadata
	err = json.Unmarshal(parsedResp, &metadata)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal json")
	}

	return &metadata, nil
}

func (p *proxy) CreateApprovalTx(tokenId *big.Int, tokenAddress, ownerAddress, receiverAddress string) (json.RawMessage, error) {
	return nil, nil
}

func (p *proxy) CreateNonFungibleWithdrawTx(params types.NonFungibleWithdrawParams) (json.RawMessage, error) {
	resp, err := http.Post(fmt.Sprintf("%s/integrations/solana-proxy/v1/deposit", p.endpoint),
		contentType,
		encodeBody(map[string]interface{}{
			"data": map[string]interface{}{
				"attributes": map[string]interface{}{
					"address":         params.TokenAddress,
					"network":         "ethereum", // FIXME
					"originalAddress": params.OriginalTokenAddress,
					"sender":          params.Sender,
					"receiver":        params.Receiver,
					"tokenId":         params.NftID,
				},
			},
		}))
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request proxy")
	}

	return parseResponse(resp)
}

func (p *proxy) CreateNonFungibleDepositTx(params types.NonFungibleDepositParams) (json.RawMessage, error) {
	resp, err := http.Post(fmt.Sprintf("%s/integrations/solana-proxy/v1/withdraw", p.endpoint),
		contentType,
		encodeBody(map[string]interface{}{
			"data": map[string]interface{}{
				"attributes": map[string]interface{}{
					"address": params.TokenAddress,
					"metadata": map[string]interface{}{
						"name":    params.NftMetadata.Name,
						"symbol":  "TNFT", // FIXME
						"creator": "2Vj6PTAAkjw2H6EnKhZoGd1RHhPWF243LGceQUwj4v5m",
						"uri":     "", // FIXME
					},
					"network":         "ethereum", // FIXME
					"originalAddress": params.OriginalTokenAddress,
					"sender":          params.Sender,
					"receiver":        params.Receiver,
					"tokenId":         params.NftID,
					"tx":              params.WithdrawTxHash,
				},
			},
		}))
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request proxy")
	}

	return parseResponse(resp)
}

func parseResponse(r *http.Response) (json.RawMessage, error) {
	defer r.Body.Close()
	rawResponse, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response")
	}

	if r.StatusCode < 200 || r.StatusCode >= 300 {
		return nil, errors.New(fmt.Sprintf("request failed status code - %d, body - %s", r.StatusCode,
			string(rawResponse)))
	}

	var data struct {
		Data tx `json:"data"`
	}
	err = json.Unmarshal(rawResponse, &data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response body")
	}

	return json.Marshal(map[string]string{
		"base64": data.Data.Attributes.Tx,
	})
}

func encodeBody(value interface{}) io.Reader {
	rawVal, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(rawVal)
}

type tx struct {
	ID         string       `json:"id"`
	Attributes txAttributes `json:"attributes"`
}

type txAttributes struct {
	Tx string `json:"tx"`
}
