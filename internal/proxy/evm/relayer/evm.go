package relayer

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/signature"
	"math/big"
)

func NewEvmRelayer(rpc string, signer signature.Signer) Relayer {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}
	return &evmRelayer{client, signer}
}

type evmRelayer struct {
	client *ethclient.Client
	signer signature.Signer
}

func (r *evmRelayer) SendTx(tx *ethTypes.Transaction, chainID *big.Int) (common.Hash, error) {
	tx, err := r.signer.SignTx(tx, chainID)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to sign tx")
	}

	err = r.client.SendTransaction(context.TODO(), tx)

	return tx.Hash(), errors.Wrap(err, "failed to send tx")
}
