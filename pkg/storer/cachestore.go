// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/transaction"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const (
	cacheOverCapacity = "cacheOverCapacity"
)

func (db *DB) cacheWorker(ctx context.Context) { _ = "STUB: not implemented"; return }

// evict at least a min count

// Lookup is the implementation of the CacheStore.Lookup method.
func (db *DB) Lookup() storage.Getter { _ = "STUB: not implemented"; return *new(storage.Getter) }

// here we would ideally have nothing to do but just to return this
// error to the client. The commit is mainly done to end the txn.

// if we are here, it means there was some unexpected error, in which
// case we need to rollback any changes that were already made.

// Cache is the implementation of the CacheStore.Cache method.
func (db *DB) Cache() storage.Putter { _ = "STUB: not implemented"; return *new(storage.Putter) }

// CacheShallowCopy creates cache entries with the expectation that the chunk already exists in the chunkstore.
func (db *DB) CacheShallowCopy(ctx context.Context, store transaction.Storage, addrs ...swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) triggerCacheEviction() { _ = "STUB: not implemented"; return }
