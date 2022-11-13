package evm

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/signature"
	proxytypes "gitlab.com/tokend/bridge/core/internal/proxy/types"
)

func ExportCheckTxDataAndSign(p proxytypes.Proxy) func(opts *bind.TransactOpts, tx *types.Transaction, rawData []byte) (*types.Transaction, int64, error) {
	return p.(*evmProxy).checkTxDataAndSign
}

func ExportRedeemNative(p proxytypes.Proxy) func(params proxytypes.FungibleRedeemParams, sender common.Address) (*types.Transaction, error) {
	return p.(*evmProxy).redeemNative
}

func ExportBuildTxOpts() func(from common.Address) *bind.TransactOpts {
	return buildTransactOpts
}

func ChangeSigner(p proxytypes.Proxy, signer signature.Signer) proxytypes.Proxy {
	proxy := p.(*evmProxy)
	proxy.signer = signer
	return proxy
}
