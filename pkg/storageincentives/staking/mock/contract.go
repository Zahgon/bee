// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/storageincentives/staking"
)

type stakingContractMock struct {
	depositStake     func(ctx context.Context, stakedAmount *big.Int) (common.Hash, error)
	getStake         func(ctx context.Context) (*big.Int, error)
	withdrawAllStake func(ctx context.Context) (common.Hash, error)
	migrateStake     func(ctx context.Context) (common.Hash, error)
	isFrozen         func(ctx context.Context, block uint64) (bool, error)
	updateHeight     func(ctx context.Context) (common.Hash, bool, error)
}

func (s *stakingContractMock) DepositStake(ctx context.Context, stakedAmount *big.Int) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (s *stakingContractMock) ChangeStakeOverlay(_ context.Context, h common.Hash) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (s *stakingContractMock) UpdateHeight(ctx context.Context) (common.Hash, bool, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), false, nil
}

func (s *stakingContractMock) GetPotentialStake(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *stakingContractMock) GetWithdrawableStake(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *stakingContractMock) WithdrawStake(ctx context.Context) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (s *stakingContractMock) MigrateStake(ctx context.Context) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (s *stakingContractMock) IsOverlayFrozen(ctx context.Context, block uint64) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// Option is an option passed to New
		nil
}

type Option func(mock *stakingContractMock)

// New creates a new mock BatchStore.
func New(opts ...Option) staking.Contract { _ = "STUB: not implemented"; return *new(staking.Contract) }

func WithDepositStake(f func(ctx context.Context, stakedAmount *big.Int) (common.Hash, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithGetStake(f func(ctx context.Context) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithWithdrawStake(f func(ctx context.Context) (common.Hash, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMigrateStake(f func(ctx context.Context) (common.Hash, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithIsFrozen(f func(ctx context.Context, block uint64) (bool, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithUpdateHeight(f func(ctx context.Context) (common.Hash, bool, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
