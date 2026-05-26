// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package monitormock

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/transaction"
)

type transactionMonitorMock struct {
	watchTransaction func(txHash common.Hash, nonce uint64) (<-chan types.Receipt, <-chan error, error)
	waitBlock        func(ctx context.Context, block *big.Int) (*types.Block, error)
}

func (m *transactionMonitorMock) WatchTransaction(txHash common.Hash, nonce uint64) (<-chan types.Receipt, <-chan error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (m *transactionMonitorMock) WaitBlock(ctx context.Context, block *big.Int) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *transactionMonitorMock) Close() error {
	_ = "STUB: not implemented"

	// Option is the option passed to the mock Chequebook service
	return nil
}

type Option interface {
	apply(*transactionMonitorMock)
}

type optionFunc func(*transactionMonitorMock)

func (f optionFunc) apply(r *transactionMonitorMock) { _ = "STUB: not implemented"; return }

func WithWatchTransactionFunc(f func(txHash common.Hash, nonce uint64) (<-chan types.Receipt, <-chan error, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithWaitBlockFunc(f func(ctx context.Context, block *big.Int) (*types.Block, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func New(opts ...Option) transaction.Monitor {
	_ = "STUB: not implemented"
	return *new(transaction.Monitor)
}
