package evm

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/bridge/core/internal/amount"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc1155"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc20"
	"gitlab.com/tokend/bridge/core/internal/proxy/evm/generated/erc721"
	"gitlab.com/tokend/bridge/core/internal/proxy/types"
	"math/big"
)

func (p *evmProxy) Approve(params types.ApproveParams) (interface{}, error) {
	fromAddress := common.HexToAddress(params.ApproveFrom)

	var tx *ethTypes.Transaction
	var err error
	switch params.TokenChain.TokenType {
	case TokenTypeNative:
		// Approve not needed for native token
		return nil, nil
	case TokenTypeErc20:
		tx, err = p.approveErc20(common.HexToAddress(*params.TokenChain.ContractAddress), fromAddress, *params.Amount)
	case TokenTypeErc721:
		tx, err = p.approveErc721(common.HexToAddress(*params.TokenChain.ContractAddress), fromAddress, *params.NftId)
	case TokenTypeErc1155:
		tx, err = p.approveErc1155(common.HexToAddress(*params.TokenChain.ContractAddress), fromAddress)
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

	return encodeTx(tx, fromAddress, p.chainID, params.TokenChain.ChainID, nil)
}

func (p *evmProxy) approveErc20(tokenAddress common.Address, approveFrom common.Address, requiredAmount amount.Amount) (*ethTypes.Transaction, error) {
	token, err := erc20.NewErc20(tokenAddress, p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transactor")
	}

	approved, err := token.Allowance(&bind.CallOpts{}, approveFrom, p.bridgeContract)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check allowance")
	}

	decimals, err := token.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get token decimals")
	}

	approvedAmount := amount.NewFromIntWithPrecision(approved, int(decimals))
	if amount.Cmp(approvedAmount, requiredAmount) >= 0 {
		// Token is already approved
		return nil, nil
	}

	tx, err := token.Approve(buildTransactOpts(approveFrom), p.bridgeContract,
		requiredAmount.IntWithPrecision(int(decimals)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to build token approve tx")
	}

	return tx, nil
}

func (p *evmProxy) approveErc721(tokenAddress common.Address, approveFrom common.Address,
	nftId string) (*ethTypes.Transaction, error) {
	token, err := erc721.NewErc721(tokenAddress, p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transactor")
	}

	nft, isSet := new(big.Int).SetString(nftId, 10)
	if !isSet {
		return nil, errors.New("failed to parse nft id")
	}
	address, err := token.GetApproved(&bind.CallOpts{}, nft)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check approval")
	}
	if compareAddresses(address, p.bridgeContract) {
		// Token is already approved
		return nil, nil
	}

	tx, err := token.Approve(buildTransactOpts(approveFrom), p.bridgeContract, nft)
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
