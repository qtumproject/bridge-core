package evm

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/data"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc1155"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc20"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc721"
	"math/big"
)

func (p *evmProxy) Approve(tokenChain data.TokenChain, approveFrom string) (interface{}, error) {
	fromAddress := common.HexToAddress(approveFrom)

	var tx *ethTypes.Transaction
	var err error
	switch tokenChain.TokenType {
	case tokenTypeNative:
		// Approve not needed for native token
		return nil, nil
	case tokenTypeErc20:
		tx, err = p.approveErc20(common.HexToAddress(*tokenChain.ContractAddress), fromAddress)
	case tokenTypeErc721:
		tx, err = p.approveErc721(common.HexToAddress(*tokenChain.ContractAddress), fromAddress)
	case tokenTypeErc1155:
		tx, err = p.approveErc1155(common.HexToAddress(*tokenChain.ContractAddress), fromAddress)
	default:
		return nil, errors.New("unknown token type")
	}
	if err != nil {
		return nil, err
	}
	if tx == nil {
		// Token is already approved
		return nil, nil
	}

	return encodeTx(tx, fromAddress, p.chainID, tokenChain.ChainID, nil)
}

func (p *evmProxy) approveErc20(tokenAddress common.Address, approveFrom common.Address) (*ethTypes.Transaction, error) {
	token, err := erc20.NewErc20(tokenAddress, p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transactor")
	}

	amount, err := token.Allowance(&bind.CallOpts{}, approveFrom, p.bridgeContract)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check allowance")
	}
	if amount.Cmp(big.NewInt(0)) == 1 {
		// Token is already approved
		return nil, nil
	}

	uin256max := big.NewInt(0)
	uin256max.SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935", 10)
	tx, err := token.Approve(buildTransactOpts(approveFrom), p.bridgeContract, uin256max)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build token approve tx")
	}

	return tx, nil
}

func (p *evmProxy) approveErc721(tokenAddress common.Address, approveFrom common.Address) (*ethTypes.Transaction, error) {
	token, err := erc721.NewErc721(tokenAddress, p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transactor")
	}

	approved, err := token.IsApprovedForAll(&bind.CallOpts{}, approveFrom, p.bridgeContract)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check approval")
	}
	if approved {
		// Token is already approved
		return nil, nil
	}

	tx, err := token.SetApprovalForAll(buildTransactOpts(approveFrom), p.bridgeContract, true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build token approve tx")
	}

	return tx, err
}

func (p *evmProxy) approveErc1155(tokenAddress common.Address, approveFrom common.Address) (*ethTypes.Transaction, error) {
	token, err := erc1155.NewErc1155(tokenAddress, p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transactor")
	}

	approved, err := token.IsApprovedForAll(&bind.CallOpts{}, approveFrom, p.bridgeContract)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check approval")
	}
	if approved {
		// Token is already approved
		return nil, nil
	}

	tx, err := token.SetApprovalForAll(buildTransactOpts(approveFrom), p.bridgeContract, true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build token approve tx")
	}

	return tx, err
}
