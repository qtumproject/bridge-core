package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/amount"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc1155"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc20"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc721"
	"math/big"
)

func (p *evmProxy) BridgeBalance(tokenChain data.TokenChain, nftId *string) (amount.Amount, error) {
	return p.Balance(tokenChain, p.bridgeContract.String(), nftId)
}

func (p *evmProxy) Balance(tokenChain data.TokenChain, address string, nftId *string) (amount.Amount, error) {
	switch tokenChain.TokenType {
	case TokenTypeNative:
		return p.nativeBalance(address)
	case TokenTypeErc20:
		return p.erc20Balance(*tokenChain.ContractAddress, address)
	case TokenTypeErc721:
		return p.erc721Balance(*tokenChain.ContractAddress, address, *nftId)
	case TokenTypeErc1155:
		return p.erc1155Balance(*tokenChain.ContractAddress, address, *nftId)
	default:
		return amount.Amount{}, errors.Errorf("unknown token type: %s, token: %s", tokenChain.TokenType, tokenChain.TokenID)
	}
}

func (p *evmProxy) nativeBalance(address string) (amount.Amount, error) {
	balance, err := p.client.BalanceAt(context.TODO(), common.HexToAddress(address), nil)
	if err != nil {
		return amount.Amount{}, errors.Wrap(err, "failed to get balance")
	}

	return amount.NewFromBigInt(balance), nil
}

func (p *evmProxy) erc20Balance(tokenAddress string, address string) (amount.Amount, error) {
	token, err := erc20.NewErc20(common.HexToAddress(tokenAddress), p.client)
	if err != nil {
		return amount.Amount{}, errors.Wrap(err, "failed to create erc20 instance")
	}

	balance, err := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		return amount.Amount{}, errors.Wrap(err, "failed to get balance")
	}

	decimals, err := p.getDecimals(tokenAddress)
	if err != nil {
		return amount.Amount{}, errors.Wrap(err, "failed to get decimals")
	}

	return amount.NewFromIntWithPrecision(balance, decimals), nil
}

func (p *evmProxy) erc721Balance(tokenAddress string, address string, nftId string) (amount.Amount, error) {
	token, err := erc721.NewErc721(common.HexToAddress(tokenAddress), p.client)
	if err != nil {
		return amount.Amount{}, errors.Wrap(err, "failed to create erc721 instance")
	}

	tokenId, ok := big.NewInt(0).SetString(nftId, 10)
	if !ok {
		return amount.Amount{}, errors.Wrap(err, "failed to parse nft id")
	}
	owner, err := token.OwnerOf(&bind.CallOpts{}, tokenId)
	if err != nil {
		return amount.Amount{}, errors.Wrap(err, "failed to get owner")
	}

	if compareAddresses(owner, common.HexToAddress(address)) {
		return amount.MustNewFromString("1"), nil
	}

	return amount.NewFromInt(0), nil
}

func (p *evmProxy) erc1155Balance(tokenAddress string, address string, nftId string) (amount.Amount, error) {
	token, err := erc1155.NewErc1155(common.HexToAddress(tokenAddress), p.client)
	if err != nil {
		return amount.Amount{}, errors.Wrap(err, "failed to create erc721 instance")
	}

	tokenId, ok := big.NewInt(0).SetString(nftId, 10)
	if !ok {
		return amount.Amount{}, errors.Wrap(err, "failed to parse nft id")
	}
	balance, err := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address), tokenId)
	if err != nil {
		return amount.Amount{}, errors.Wrap(err, "failed to get balance")
	}

	if balance.Cmp(big.NewInt(0)) == 1 {
		return amount.MustNewFromString("1"), nil
	}

	return amount.NewFromInt(0), nil
}
