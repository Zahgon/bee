// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/transaction"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"resenje.org/multex"
)

var now = time.Now

// exported for migration
type CacheEntryItem = cacheEntry

const cacheEntrySize = swarm.HashSize + 8

var _ storage.Item = (*cacheEntry)(nil)

var (
	errMarshalCacheEntryInvalidAddress   = errors.New("marshal cacheEntry: invalid address")
	errMarshalCacheEntryInvalidTimestamp = errors.New("marshal cacheEntry: invalid timestamp")
	errUnmarshalCacheEntryInvalidSize    = errors.New("unmarshal cacheEntry: invalid size")
)

// Cache is the part of the localstore which keeps track of the chunks that are not
// part of the reserve but are potentially useful to store for obtaining bandwidth
// incentives.
type Cache struct {
	size     atomic.Int64
	capacity int
	glock    *multex.Multex // blocks Get and Put ops while shallow copy is running.
}

// New creates a new Cache component with the specified capacity. The store is used
// here only to read the initial state of the cache before shutdown if there was
// any.
func New(ctx context.Context, store storage.Reader, capacity uint64) (*Cache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Size returns the current size of the cache.
func (c *Cache) Size() int64 { _ = "STUB: not implemented"; return 0 }

// Capacity returns the capacity of the cache.
func (c *Cache) Capacity() int64 { _ = "STUB: not implemented"; return 0 }

// Putter returns a Storage.Putter instance which adds the chunk to the underlying
// chunkstore and also adds a Cache entry for the chunk.
func (c *Cache) Putter(store transaction.Storage) storage.Putter {
	_ = "STUB: not implemented"
	return *new(storage.Putter)
}

// if chunk is already part of cache, return found.

// Getter returns a Storage.Getter instance which checks if the chunks accessed are
// part of cache it will update the cache indexes. If the operation to update the
// cache indexes fail, we need to fail the operation as this should signal the user
// of this getter to rollback the operation.
func (c *Cache) Getter(store transaction.Storage) storage.Getter {
	_ = "STUB: not implemented"
	return *new(storage.Getter)
}

// check if there is an entry in Cache. As this is the download path, we do
// a best-effort operation. So in case of any error we return the chunk.

// RemoveOldest removes the oldest cache entries from the store. The count
// specifies the number of entries to remove.
func (c *Cache) RemoveOldest(ctx context.Context, st transaction.Storage, count uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// ShallowCopy creates cache entries with the expectation that the chunk already exists in the chunkstore.
func (c *Cache) ShallowCopy(
	ctx context.Context,
	store transaction.Storage,
	addrs ...swarm.Address,
) (err error) {
	_ = "STUB: not implemented"
	// TODO: add proper mutex locking before usage
	return nil
}

// Since the caller has previously referenced the chunk (+1 refCnt), and if the chunk is already referenced
// by the cache store (+1 refCnt), then we must decrement the refCnt by one ( -1 refCnt to bring the total to +1).
// See https://github.com/ethersphere/bee/issues/4530.

// consider only the amount that can fit, the rest should be deleted from the chunkstore.

type cacheEntry struct {
	Address         swarm.Address
	AccessTimestamp int64
}

func (c *cacheEntry) ID() string { _ = "STUB: not implemented"; return "" }

func (cacheEntry) Namespace() string { _ = "STUB: not implemented"; return "" }

func (c *cacheEntry) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *cacheEntry) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }

func (c *cacheEntry) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

func (c cacheEntry) String() string { _ = "STUB: not implemented"; return "" }

var _ storage.Item = (*cacheOrderIndex)(nil)

type cacheOrderIndex struct {
	AccessTimestamp int64
	Address         swarm.Address
}

func keyFromID(ts int64, addr swarm.Address) string { _ = "STUB: not implemented"; return "" }

func idFromKey(key string) (int64, swarm.Address, error) {
	_ = "STUB: not implemented"
	return 0, *new(swarm.Address), nil
}

func (c *cacheOrderIndex) ID() string { _ = "STUB: not implemented"; return "" }

func (cacheOrderIndex) Namespace() string { _ = "STUB: not implemented"; return "" }

func (cacheOrderIndex) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (cacheOrderIndex) Unmarshal(_ []byte) error { _ = "STUB: not implemented"; return nil }

func (c *cacheOrderIndex) Clone() storage.Item {
	_ = "STUB: not implemented"
	return *new(storage.Item)
}

func (c cacheOrderIndex) String() string { _ = "STUB: not implemented"; return "" }
