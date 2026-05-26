// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/erc20"
)

type Service struct {
	balanceOfFunc func(ctx context.Context, address common.Address) (*big.Int, error)
	transferFunc  func(ctx context.Context, address common.Address, value *big.Int) (common.Hash, error)
}

func WithBalanceOfFunc(f func(ctx context.Context, address common.Address) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTransferFunc(f func(ctx context.Context, address common.Address, value *big.Int) (common.Hash, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func New(opts ...Option) erc20.Service { _ = "STUB: not implemented"; return *new(erc20.Service) }

func (s *Service) BalanceOf(ctx context.Context, address common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) Transfer(ctx context.Context, address common.Address, value *big.Int) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// Option is the option passed to the mock Chequebook service
type Option interface {
	apply(*Service)
}

type optionFunc func(*Service)

func (f optionFunc) apply(r *Service) { _ = "STUB: not implemented"; return }
