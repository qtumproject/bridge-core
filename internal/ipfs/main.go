package ipfs

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client interface {
	Get(result interface{}, url string) error
	WrapUrl(url string) string
	IsIpfsUrl(url string) bool
}

func NewClient(url string) Client {
	// TODO: make better
	if strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}
	return &client{
		proxyUrl: url,
	}
}

type client struct {
	proxyUrl string
}

func (c *client) Get(result interface{}, url string) error {
	resp, err := http.Get(c.WrapUrl(url))
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response")
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return errors.New(fmt.Sprintf("failed to get data from ipfs status %d, body %s",
			resp.StatusCode, string(body)))
	}

	return json.Unmarshal(body, result)
}

func (c *client) WrapUrl(url string) string {
	if !c.IsIpfsUrl(url) {
		return url
	}

	url = strings.Replace(url, "ipfs://ipfs/", "", 1)
	url = strings.Replace(url, "ipfs://", "", 1)
	return fmt.Sprintf("%s/ipfs/%s", c.proxyUrl, url)
}

func (c *client) IsIpfsUrl(url string) bool {
	return strings.HasPrefix(url, "ipfs://")
}
