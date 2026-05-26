// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package staking

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/transaction"
	"github.com/ethersphere/bee/v2/pkg/util/abiutil"
	"github.com/ethersphere/go-sw3-abi/sw3abi"
)

var (
	MinimumStakeAmount = big.NewInt(100000000000000000)

	erc20ABI = abiutil.MustParseABI(sw3abi.ERC20ABIv0_6_9)

	ErrInsufficientStakeAmount = errors.New("insufficient stake amount")
	ErrInsufficientFunds       = errors.New("insufficient token balance")
	ErrInsufficientStake       = errors.New("insufficient stake")
	ErrNotImplemented          = errors.New("not implemented")
	ErrNotPaused               = errors.New("contract is not paused")
	ErrUnexpectedLength        = errors.New("unexpected results length")

	approveDescription       = "Approve tokens for stake deposit operations"
	depositStakeDescription  = "Deposit Stake"
	withdrawStakeDescription = "Withdraw stake"
	migrateStakeDescription  = "Migrate stake"
)

type Contract interface {
	DepositStake(ctx context.Context, stakedAmount *big.Int) (common.Hash, error)
	ChangeStakeOverlay(ctx context.Context, nonce common.Hash) (common.Hash, error)
	GetPotentialStake(ctx context.Context) (*big.Int, error)
	GetWithdrawableStake(ctx context.Context) (*big.Int, error)
	WithdrawStake(ctx context.Context) (common.Hash, error)
	MigrateStake(ctx context.Context) (common.Hash, error)
	UpdateHeight(ctx context.Context) (common.Hash, bool, error)
	RedistributionStatuser
}

type RedistributionStatuser interface {
	IsOverlayFrozen(ctx context.Context, block uint64) (bool, error)
}

type contract struct {
	owner                  common.Address
	stakingContractAddress common.Address
	stakingContractABI     abi.ABI
	bzzTokenAddress        common.Address
	transactionService     transaction.Service
	overlayNonce           common.Hash
	gasLimit               uint64
	height                 uint8
}

func New(
	owner common.Address,
	stakingContractAddress common.Address,
	stakingContractABI abi.ABI,
	bzzTokenAddress common.Address,
	transactionService transaction.Service,
	nonce common.Hash,
	gasLimit uint64,
	height uint8,
) Contract {
	_ = "STUB: not implemented"
	return *new(Contract)
}

func (c *contract) DepositStake(ctx context.Context, stakedAmount *big.Int) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// ChangeStakeOverlay only changes the overlay address used in the redistribution game.
func (c *contract) ChangeStakeOverlay(ctx context.Context, nonce common.Hash) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// UpdateHeight submits the reserve doubling height to the contract only if the height is a new value.
func (c *contract) UpdateHeight(ctx context.Context) (common.Hash, bool, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), false, nil
}

func (c *contract) GetPotentialStake(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *contract) GetWithdrawableStake(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *contract) WithdrawStake(ctx context.Context) (txHash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (c *contract) MigrateStake(ctx context.Context) (txHash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (c *contract) IsOverlayFrozen(ctx context.Context, block uint64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *contract) sendApproveTransaction(ctx context.Context, amount *big.Int) (receipt *types.Receipt, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *contract) sendTransaction(ctx context.Context, callData []byte, desc string) (receipt *types.Receipt, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *contract) sendManageStakeTransaction(ctx context.Context, stakedAmount *big.Int) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *contract) getPotentialStake(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// overlay bytes32,
// committedStake uint256,
// potentialStake uint256,
// lastUpdatedBlockNumber uint256,

func (c *contract) getWithdrawableStake(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *contract) getBalance(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *contract) migrateStake(ctx context.Context) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *contract) withdrawFromStake(ctx context.Context) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *contract) paused(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *contract) getHeight(ctx context.Context) (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
