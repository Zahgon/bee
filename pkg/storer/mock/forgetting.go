// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mockstorer

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	storage "github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type DelayedStore struct {
	storage.ChunkStore
	cache map[string]time.Duration
	mu    sync.Mutex
}

func NewDelayedStore(s storage.ChunkStore) *DelayedStore { _ = "STUB: not implemented"; return nil }

func (d *DelayedStore) Delay(addr swarm.Address, delay time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (d *DelayedStore) Get(ctx context.Context, addr swarm.Address) (ch swarm.Chunk, err error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

type ForgettingStore struct {
	storage.ChunkStore
	record atomic.Bool
	mu     sync.Mutex
	n      atomic.Int64
	missed map[string]struct{}
}

func NewForgettingStore(s storage.ChunkStore) *ForgettingStore {
	_ = "STUB: not implemented"
	return nil
}

func (f *ForgettingStore) Stored() int64 { _ = "STUB: not implemented"; return 0 }

func (f *ForgettingStore) Record() { _ = "STUB: not implemented"; return }

func (f *ForgettingStore) Unrecord() { _ = "STUB: not implemented"; return }

func (f *ForgettingStore) Miss(addr swarm.Address) { _ = "STUB: not implemented"; return }

func (f *ForgettingStore) Unmiss(addr swarm.Address) { _ = "STUB: not implemented"; return }

func (f *ForgettingStore) miss(addr swarm.Address) { _ = "STUB: not implemented"; return }

func (f *ForgettingStore) unmiss(addr swarm.Address) { _ = "STUB: not implemented"; return }

func (f *ForgettingStore) isMiss(addr swarm.Address) bool { _ = "STUB: not implemented"; return false }

func (f *ForgettingStore) Reset() { _ = "STUB: not implemented"; return }

func (f *ForgettingStore) Missed() int { _ = "STUB: not implemented"; return 0 }

// Get implements the ChunkStore interface.
// if in recording phase, record the chunk address as miss and returns Get on the embedded store
// if in forgetting phase, returns ErrNotFound if the chunk address is recorded as miss
func (f *ForgettingStore) Get(ctx context.Context, addr swarm.Address) (ch swarm.Chunk, err error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// Put implements the ChunkStore interface.
func (f *ForgettingStore) Put(ctx context.Context, ch swarm.Chunk) (err error) {
	_ = "STUB: not implemented"
	return nil
}
