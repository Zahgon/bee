// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/postage/postagecontract"
)

type contractMock struct {
	createBatch           func(ctx context.Context, initialBalance *big.Int, depth uint8, immutable bool, label string) (common.Hash, []byte, error)
	topupBatch            func(ctx context.Context, id []byte, amount *big.Int) (common.Hash, error)
	diluteBatch           func(ctx context.Context, id []byte, newDepth uint8) (common.Hash, error)
	expireBatches         func(ctx context.Context) error
	paused                func(ctx context.Context) (bool, error)
	minimumValidityBlocks func(ctx context.Context) (uint64, error)
}

func (c *contractMock) CreateBatch(ctx context.Context, initialBalance *big.Int, depth uint8, immutable bool, label string) (common.Hash, []byte, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil, nil
}

func (c *contractMock) TopUpBatch(ctx context.Context, batchID []byte, topupBalance *big.Int) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (c *contractMock) DiluteBatch(ctx context.Context, batchID []byte, newDepth uint8) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (c *contractMock) ExpireBatches(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *contractMock) Paused(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *contractMock) MinimumValidityBlocks(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Option is an option passed to New
type Option func(*contractMock)

// New creates a new mock BatchStore.
func New(opts ...Option) postagecontract.Interface {
	_ = "STUB: not implemented"
	return *new(postagecontract.Interface)
}

func WithCreateBatchFunc(f func(ctx context.Context, initialBalance *big.Int, depth uint8, immutable bool, label string) (common.Hash, []byte, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTopUpBatchFunc(f func(ctx context.Context, batchID []byte, amount *big.Int) (common.Hash, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithDiluteBatchFunc(f func(ctx context.Context, batchID []byte, newDepth uint8) (common.Hash, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithExpiresBatchesFunc(f func(ctx context.Context) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPaused(f func(ctx context.Context) (bool, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMinimumValidityBlocksFunc(f func(ctx context.Context) (uint64, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
