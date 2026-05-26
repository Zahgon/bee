// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Service struct {
	rate   *big.Int
	deduct *big.Int
}

func New(rate, deduct *big.Int) Service { _ = "STUB: not implemented"; return *new(Service) }

func (s Service) Start() { _ = "STUB: not implemented"; return }

func (s Service) GetPrice(ctx context.Context) (*big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s Service) CurrentRates() (exchangeRate, deduction *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s Service) Close() error { _ = "STUB: not implemented"; return nil }

func DiscoverPriceOracleAddress(chainID int64) (priceOracleAddress common.Address, found bool) {
	_ = "STUB: not implemented"
	return *new(common.Address), false
}

func (s *Service) SetValues(rate, deduct *big.Int) { _ = "STUB: not implemented"; return }
