package relayer

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/signature"
	"io"
	"math/big"
	"net/http"
	"strings"
)

func NewQtumRelayer(proxy string, signer signature.Signer) Relayer {
	return &qtumRelayer{proxy, signer}
}

type qtumRelayer struct {
	proxy  string
	signer signature.Signer
}

func (r *qtumRelayer) SendTx(tx *ethTypes.Transaction, chainID *big.Int) (common.Hash, error) {
	body, err := json.Marshal(map[string]interface{}{
		"to":       tx.To().String(),
		"chain_id": chainID.String(),
		"value":    tx.Value().String(),
		"data":     "0x" + common.Bytes2Hex(tx.Data()),
	})
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to marshal body")
	}
	resp, err := http.Post(r.proxy, "application/json", strings.NewReader(string(body)))
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to send tx")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return common.Hash{}, errors.New(fmt.Sprintf("failed to send tx: %s", string(body)))
	}
	var result map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to decode response")
	}

	return common.HexToHash(result["tx_hash"].(string)), nil
}
