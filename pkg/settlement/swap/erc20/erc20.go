// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package erc20

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/transaction"
	"github.com/ethersphere/bee/v2/pkg/util/abiutil"
	"github.com/ethersphere/go-sw3-abi/sw3abi"
)

var (
	erc20ABI     = abiutil.MustParseABI(sw3abi.ERC20ABIv0_6_9)
	errDecodeABI = errors.New("could not decode abi data")
)

type Service interface {
	BalanceOf(ctx context.Context, address common.Address) (*big.Int, error)
	Transfer(ctx context.Context, address common.Address, value *big.Int) (common.Hash, error)
}

type erc20Service struct {
	transactionService transaction.Service
	address            common.Address
}

func New(transactionService transaction.Service, address common.Address) Service {
	_ = "STUB: not implemented"
	return *new(Service)
}

func (c *erc20Service) BalanceOf(ctx context.Context, address common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *erc20Service) Transfer(ctx context.Context, address common.Address, value *big.Int) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}
