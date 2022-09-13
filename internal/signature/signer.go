package signature

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
)

type Hasher interface {
	Hash() []byte
}

type Signer interface {
	Sign(Hasher) (*Parameters, error)
	// function needed for bind.OptsTransacion
	SignTx(*types.Transaction, *big.Int) (*types.Transaction, error)
	Opts(*big.Int) (*bind.TransactOpts, error)
	Address() common.Address
}

type signer struct {
	privKey *ecdsa.PrivateKey
	address common.Address
}

func NewSigner(privKey *ecdsa.PrivateKey) Signer {
	pubKeyECDSA, ok := privKey.Public().(*ecdsa.PublicKey)
	if !ok {
		panic(errors.New("cannot assign public key to ecdsa public key"))
	}

	return &signer{
		privKey: privKey,
		address: crypto.PubkeyToAddress(*pubKeyECDSA),
	}
}

func (s *signer) Sign(hasher Hasher) (*Parameters, error) {
	signature, err := crypto.Sign(hasher.Hash(), s.privKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign hash")
	}

	return ParseSignatureParameters(signature)
}

func (s *signer) SignTx(tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), s.privKey)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func (s *signer) Address() common.Address {
	return s.address
}

func (s *signer) Opts(chainID *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(s.privKey, chainID)
}
