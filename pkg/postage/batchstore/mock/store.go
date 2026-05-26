// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"math/big"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/postage"
)

var _ postage.Storer = (*BatchStore)(nil)

// BatchStore is a mock BatchStorer
type BatchStore struct {
	radius                uint8
	cs                    *postage.ChainState
	isWithinStorageRadius bool
	id                    []byte
	batch                 *postage.Batch
	getErr                error
	getErrDelayCnt        int
	updateErr             error
	saveErr               error
	updateErrDelayCnt     int
	resetCallCount        int

	existsFn func([]byte) (bool, error)

	mtx sync.Mutex
}

func (bs *BatchStore) SetBatchExpiryHandler(eh postage.BatchExpiryHandler) {
	_ = "STUB: not implemented"

	// Option is an option passed to New.
	return
}

type Option func(*BatchStore)

// New creates a new mock BatchStore
func New(opts ...Option) *BatchStore { _ = "STUB: not implemented"; return nil }

// WithReserveState will set the initial reservestate in the ChainStore mock.
func WithRadius(radius uint8) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithChainState will set the initial chainstate in the ChainStore mock.
func WithChainState(cs *postage.ChainState) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGetErr will set the get error returned by the ChainStore mock. The error
// will be returned on each subsequent call after delayCnt calls to Get have
// been made.
func WithGetErr(err error, delayCnt int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUpdateErr will set the put error returned by the ChainStore mock.
// The error will be returned on each subsequent call after delayCnt
// calls to Update have been made.
func WithUpdateErr(err error, delayCnt int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSaveError(err error) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBatch will set batch to the one provided by user.
// This will be returned in the next Get.
func WithBatch(b *postage.Batch) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithExistsFunc(f func([]byte) (bool, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithAcceptAllExistsFunc() Option { _ = "STUB: not implemented"; return *new(Option) }

// Get mocks the Get method from the BatchStore.
func (bs *BatchStore) Get(id []byte) (*postage.Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Iterate mocks the Iterate method from the BatchStore
func (bs *BatchStore) Iterate(f func(*postage.Batch) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// Save mocks the Save method from the BatchStore.
func (bs *BatchStore) Save(batch *postage.Batch) error { _ = "STUB: not implemented"; return nil }

// Update mocks the Update method from the BatchStore.
func (bs *BatchStore) Update(batch *postage.Batch, newValue *big.Int, newDepth uint8) error {
	_ = "STUB: not implemented"
	return nil
}

// GetChainState mocks the GetChainState method from the BatchStore
func (bs *BatchStore) GetChainState() *postage.ChainState {
	_ = "STUB: not implemented"

	// GetChainState mocks the GetChainState method from the BatchStore
	return nil
}

func (bs *BatchStore) Commitment() (uint64, error) {
	_ = "STUB: not implemented"

	// PutChainState mocks the PutChainState method from the BatchStore
	return 0, nil
}

func (bs *BatchStore) PutChainState(cs *postage.ChainState) error {
	_ = "STUB: not implemented"
	return nil
}

func (bs *BatchStore) Radius() uint8 { _ = "STUB: not implemented"; return 0 }

// Exists reports whether batch referenced by the give id exists.
func (bs *BatchStore) Exists(id []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (bs *BatchStore) Reset() error { _ = "STUB: not implemented"; return nil }

func (bs *BatchStore) ResetCalls() int { _ = "STUB: not implemented"; return 0 }

type MockEventUpdater struct {
	inProgress bool
	err        error
}

func NewNotReady() *MockEventUpdater           { _ = "STUB: not implemented"; return nil }
func NewWithError(err error) *MockEventUpdater { _ = "STUB: not implemented"; return nil }

func (s *MockEventUpdater) GetSyncStatus() (isDone bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}
