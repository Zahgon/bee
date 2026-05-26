// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postage

import (
	"errors"
	"math/big"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var _ Storer = (*NoOpBatchStore)(nil)

var ErrChainDisabled = errors.New("chain disabled")

// NoOpBatchStore is a placeholder implementation for postage.Storer
type NoOpBatchStore struct{}

func (b *NoOpBatchStore) SetBatchExpiryHandler(BatchExpiryHandler) {
	_ = "STUB: not implemented"
	return
}

func (b *NoOpBatchStore) Get([]byte) (*Batch, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *NoOpBatchStore) Exists([]byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (b *NoOpBatchStore) Iterate(func(*Batch) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *NoOpBatchStore) Save(*Batch) error { _ = "STUB: not implemented"; return nil }

func (b *NoOpBatchStore) Update(*Batch, *big.Int, uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *NoOpBatchStore) GetChainState() *ChainState { _ = "STUB: not implemented"; return nil }

func (b *NoOpBatchStore) PutChainState(*ChainState) error { _ = "STUB: not implemented"; return nil }

func (b *NoOpBatchStore) Radius() uint8 { _ = "STUB: not implemented"; return 0 }

func (b *NoOpBatchStore) IsWithinStorageRadius(swarm.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func (b *NoOpBatchStore) StorageRadius() uint8 { _ = "STUB: not implemented"; return 0 }

func (b *NoOpBatchStore) SetStorageRadius(func(uint8) uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *NoOpBatchStore) Commitment() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *NoOpBatchStore) Reset() error { _ = "STUB: not implemented"; return nil }
