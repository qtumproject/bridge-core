package evm_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"gitlab.com/tokend/bridge/core/internal/amount"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/data/mem"
	"gitlab.com/tokend/bridge/core/internal/ipfs"
	"gitlab.com/tokend/bridge/core/internal/proxy"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/signature"
	proxytypes "gitlab.com/tokend/bridge/core/internal/proxy/types"
	"gitlab.com/tokend/bridge/core/resources"
	"os"
	"testing"
)

type appConfig struct {
	ChainsQ      data.ChainsQ
	TokensQ      data.TokensQ
	TokenChainsQ data.TokenChainsQ
	ProxyRepo    proxy.ProxyRepo
	Signer       signature.Signer
	IPFS         ipfs.Client
}

func setupTest() (*appConfig, error) {
	goerliRpc := os.Getenv("GOERLI_RPC")
	sepoliaRpc := os.Getenv("SEPOLIA_RPC")

	tokens := []data.Token{
		{
			ID:     "1",
			Name:   "Ether",
			Symbol: "ETH",
			Icon:   nil,
			Type:   resources.FUNGIBLE,
			Chains: []data.TokenChain{
				{
					ChainID:      "goerli",
					TokenType:    evm.TokenTypeNative,
					BridgingType: data.BridgingTypeLP,
					AutoSend:     false,
				},
				{
					ChainID:      "goerli",
					TokenType:    evm.TokenTypeErc20,
					BridgingType: data.BridgingTypeLP,
					AutoSend:     false,
				},
			},
		},
	}

	chains := []data.Chain{
		{
			ID:             "goerli",
			Name:           "Ethereum Goerli",
			Icon:           nil,
			Type:           "evm",
			ChainParams:    nil,
			Confirmations:  12,
			BridgeContract: "",
			RpcEndpoint:    goerliRpc,
			Tokens:         nil,
		},
		{
			ID:             "sepolia",
			Name:           "Ethereum Sepolia",
			Icon:           nil,
			Type:           "evm",
			ChainParams:    nil,
			Confirmations:  12,
			BridgeContract: "",
			RpcEndpoint:    sepoliaRpc,
			Tokens:         nil,
		},
	}

	for i, chain := range chains {
		chain.Tokens = make([]data.TokenChain, 0)
		chains[i] = chain
	}

	tokenChains := make([]data.TokenChain, 0)
	for _, token := range tokens {
		for i, tokenChain := range token.Chains {
			tokenChain.TokenID = token.ID
			token.Chains[i] = tokenChain
			tokenChains = append(tokenChains, tokenChain)
			for k, chain := range chains {
				if chain.ID == tokenChain.ChainID {
					chain.Tokens = append(chain.Tokens, tokenChain)
					chains[k] = chain
				}
			}
		}
	}

	ipfs := ipfs.NewClient("https://ipfs.io")
	signer := "40c6d35d0a66d130fd9e89a70ebd1985f1449f1d8962b9e3e09740ac397f39ba"
	signerPk, err := crypto.HexToECDSA(signer)
	if err != nil {
		return nil, err
	}

	proxyRepo, err := proxy.NewProxyRepo(chains, signature.NewSigner(signerPk), ipfs)
	if err != nil {
		return nil, err
	}

	config := appConfig{
		ChainsQ:      mem.NewChainsQ(chains),
		TokensQ:      mem.NewTokenQ(tokens),
		TokenChainsQ: mem.NewTokenChainsQ(tokenChains),
		ProxyRepo:    proxyRepo,
		IPFS:         ipfs,
		Signer:       signature.NewSigner(signerPk),
	}

	return &config, nil
}

func PrepareRedeemMultisig() (*types.Transaction, error) {
	config, err := setupTest()
	if err != nil {
		return nil, err
	}

	proxy := config.ProxyRepo.Get("sepolia")

	signerPkHex := "6581d07a9c8c30d3a0fe0ae4a4df45dd140bf2e3ab78a101aae3b36995007fc2"

	signerPk, err := crypto.HexToECDSA(signerPkHex)
	if err != nil {
		return nil, err
	}
	signer := signature.NewSigner(signerPk)

	sender := signer.Address()

	params := proxytypes.FungibleRedeemParams{
		Sender:     config.Signer.Address().String(),
		Receiver:   config.Signer.Address().String(),
		TxHash:     "0xf3243b4c5836ae62f2bece84b9ed2c43f82a3babc1185d862e305195b494b57a",
		EventIndex: 0,
		Amount:     amount.NewFromInt(1),
	}

	proxy = evm.ChangeSigner(proxy, signer)
	redeemNativeFn := evm.ExportRedeemNative(proxy)
	tx, err := redeemNativeFn(params, sender)

	return tx, err
}

func TestCheckTxDataAndSignSuccess(t *testing.T) {
	tx1, err := PrepareRedeemMultisig()
	require.NoError(t, err)

	config, err := setupTest()
	require.NoError(t, err)

	proxy := config.ProxyRepo.Get("sepolia")

	signerPkHex := "afe3de29ec28caa94070e6eb89034bb2120a8a40fa6df977f9c74a7c92c318a1"

	signerPk, err := crypto.HexToECDSA(signerPkHex)
	require.NoError(t, err)

	signer := signature.NewSigner(signerPk)

	sender := signer.Address()

	params := proxytypes.FungibleRedeemParams{
		Sender:     config.Signer.Address().String(),
		Receiver:   config.Signer.Address().String(),
		TxHash:     "0xf3243b4c5836ae62f2bece84b9ed2c43f82a3babc1185d862e305195b494b57a",
		EventIndex: 0,
		Amount:     amount.NewFromInt(1),
	}

	proxy = evm.ChangeSigner(proxy, signer)
	redeemNativeFn := evm.ExportRedeemNative(proxy)
	tx2, err := redeemNativeFn(params, sender)
	require.NoError(t, err)

	checkTxFn := evm.ExportCheckTxDataAndSign(proxy)
	_, signNum, err := checkTxFn(evm.ExportBuildTxOpts()(sender), tx2, tx1.Data())
	require.NoError(t, err)
	require.Equal(t, int64(2), signNum)
}

func TestCheckTxDataAndSignFail(t *testing.T) {
	tx1, err := PrepareRedeemMultisig()
	require.NoError(t, err)

	config, err := setupTest()
	require.NoError(t, err)

	proxy := config.ProxyRepo.Get("sepolia")

	signerPkHex := "afe3de29ec28caa94070e6eb89034bb2120a8a40fa6df977f9c74a7c92c318a1"

	signerPk, err := crypto.HexToECDSA(signerPkHex)
	require.NoError(t, err)

	signer := signature.NewSigner(signerPk)

	sender := signer.Address()

	params := proxytypes.FungibleRedeemParams{
		Sender:     config.Signer.Address().String(),
		Receiver:   config.Signer.Address().String(),
		TxHash:     "0xf3243b4c5836ae62f2bece84b9ed2c43f82a3babc1185d862e305195b494b57a",
		EventIndex: 0,
		Amount:     amount.NewFromInt(2),
	}

	proxy = evm.ChangeSigner(proxy, signer)
	redeemNativeFn := evm.ExportRedeemNative(proxy)
	tx2, err := redeemNativeFn(params, sender)
	require.NoError(t, err)

	checkTxFn := evm.ExportCheckTxDataAndSign(proxy)
	_, _, err = checkTxFn(evm.ExportBuildTxOpts()(sender), tx2, tx1.Data())
	require.Error(t, err)
}

func TestCheckTxDataAndSignDoubleSigFail(t *testing.T) {
	tx1, err := PrepareRedeemMultisig()
	require.NoError(t, err)

	config, err := setupTest()
	require.NoError(t, err)

	proxy := config.ProxyRepo.Get("sepolia")

	signerPkHex := "6581d07a9c8c30d3a0fe0ae4a4df45dd140bf2e3ab78a101aae3b36995007fc2"

	signerPk, err := crypto.HexToECDSA(signerPkHex)
	require.NoError(t, err)

	signer := signature.NewSigner(signerPk)

	sender := signer.Address()

	params := proxytypes.FungibleRedeemParams{
		Sender:     config.Signer.Address().String(),
		Receiver:   config.Signer.Address().String(),
		TxHash:     "0xf3243b4c5836ae62f2bece84b9ed2c43f82a3babc1185d862e305195b494b57a",
		EventIndex: 0,
		Amount:     amount.NewFromInt(1),
	}

	proxy = evm.ChangeSigner(proxy, signer)
	redeemNativeFn := evm.ExportRedeemNative(proxy)
	tx2, err := redeemNativeFn(params, sender)
	require.NoError(t, err)

	checkTxFn := evm.ExportCheckTxDataAndSign(proxy)
	_, signNum, err := checkTxFn(evm.ExportBuildTxOpts()(sender), tx2, tx1.Data())
	require.Error(t, err)
	require.Equal(t, int64(1), signNum)
}
