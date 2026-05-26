// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package transaction provides transaction support for localstore operations.
All writes to the localstore (both indexstore and chunkstore) must be made using a transaction.
The transaction must be committed for the writes to be stored on the disk.

The rules of the transaction is as follows:

-sharky_write 		-> write to disk, keep sharky location in memory
-sharky_release		-> keep location in memory, do not release from the disk
-indexstore write	-> write to batch
-on commit			-> if batch_commit succeeds, release sharky_release locations from the disk
					-> if batch_commit fails or is not called, release all sharky_write location from the disk, do nothing for sharky_release

See the NewTransaction method for more details.
*/

package transaction

import (
	"context"
	"time"

	"github.com/ethersphere/bee/v2/pkg/sharky"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/prometheus/client_golang/prometheus"
	"resenje.org/multex"
)

type Transaction interface {
	Store
	Commit() error
}

type Store interface {
	ChunkStore() storage.ChunkStore
	IndexStore() storage.IndexStore
}

type ReadOnlyStore interface {
	IndexStore() storage.Reader
	ChunkStore() storage.ReadOnlyChunkStore
}

type Storage interface {
	ReadOnlyStore
	NewTransaction(context.Context) (Transaction, func())
	Run(context.Context, func(Store) error) error
	Close() error
}

type store struct {
	sharky      *sharky.Store
	bstore      storage.BatchStore
	metrics     metrics
	chunkLocker *multex.Multex
}

func NewStorage(sharky *sharky.Store, bstore storage.BatchStore) Storage {
	_ = "STUB: not implemented"
	return *new(Storage)
}

type transaction struct {
	start      time.Time
	batch      storage.Batch
	indexstore storage.IndexStore
	chunkStore *chunkStoreTrx
	sharkyTrx  *sharkyTrx
	metrics    metrics
}

// NewTransaction returns a new storage transaction.
// Commit must be called to persist data to the disk.
// The callback function must be the final call of the transaction whether or not any errors
// were returned from the storage ops or commit. Safest option is to do a defer call immediately after
// creating the transaction.
// By design, it is best to not batch too many writes to a single transaction, including multiple chunks writes.
// Calls made to the transaction are NOT thread-safe.
func (s *store) NewTransaction(ctx context.Context) (Transaction, func()) {
	_ = "STUB: not implemented"
	return *new(Transaction), nil
}

// for whatever reason, commit was not called
// release uncommitted but written sharky locations
// unlock the locked addresses

func (s *store) IndexStore() storage.Reader { _ = "STUB: not implemented"; return *new(storage.Reader) }

func (s *store) ChunkStore() storage.ReadOnlyChunkStore {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyChunkStore)
}

// Run creates a new transaction and gives the caller access to the transaction
// in the form of a callback function. After the callback returns, the transaction
// is committed to the disk. See the NewTransaction method for more details on how transactions operate internally.
// By design, it is best to not batch too many writes to a single transaction, including multiple chunks writes.
// Calls made to the transaction are NOT thread-safe.
func (s *store) Run(ctx context.Context, f func(Store) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Metrics returns set of prometheus collectors.
func (s *store) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }

// StatusMetrics exposes metrics that are exposed on the status protocol.
func (s *store) StatusMetrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }

func (s *store) Close() error { _ = "STUB: not implemented"; return nil }

func (t *transaction) Commit() (err error) { _ = "STUB: not implemented"; return nil }

// since the batch commit has failed, we must release the written chunks from sharky.

// the batch commit was successful, we can now release the accumulated locations from sharky.

// IndexStore gives access to the index store of the transaction.
// Note that no writes are persisted to the disk until the commit is called.
func (t *transaction) IndexStore() storage.IndexStore {
	_ = "STUB: not implemented"
	return *

	// ChunkStore gives access to the chunkstore of the transaction.
	// Note that no writes are persisted to the disk until the commit is called.
	new(storage.IndexStore)
}

func (t *transaction) ChunkStore() storage.ChunkStore {
	_ = "STUB: not implemented"
	return *new(storage.ChunkStore)
}

type chunkStoreTrx struct {
	indexStore   storage.IndexStore
	sharkyTrx    *sharkyTrx
	globalLocker *multex.Multex
	lockedAddrs  map[string]struct{}
	metrics      metrics
	readOnly     bool
}

func (c *chunkStoreTrx) Get(ctx context.Context, addr swarm.Address) (ch swarm.Chunk, err error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

func (c *chunkStoreTrx) Has(ctx context.Context, addr swarm.Address) (_ bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *chunkStoreTrx) Put(ctx context.Context, ch swarm.Chunk) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *chunkStoreTrx) Delete(ctx context.Context, addr swarm.Address) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *chunkStoreTrx) Iterate(ctx context.Context, fn storage.IterateChunkFn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *chunkStoreTrx) Replace(ctx context.Context, ch swarm.Chunk, emplace bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *chunkStoreTrx) lock(addr swarm.Address) func() {
	_ = "STUB: not implemented"
	// directly lock
	return nil
}

// lock chunk only once in the same transaction

// unlocking the chunk will be done in the Commit()

type indexTrx struct {
	store   storage.Reader
	batch   storage.Batch
	metrics metrics
}

func (s *indexTrx) Get(i storage.Item) error           { _ = "STUB: not implemented"; return nil }
func (s *indexTrx) Has(k storage.Key) (bool, error)    { _ = "STUB: not implemented"; return false, nil }
func (s *indexTrx) GetSize(k storage.Key) (int, error) { _ = "STUB: not implemented"; return 0, nil }
func (s *indexTrx) Iterate(q storage.Query, f storage.IterateFn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *indexTrx) Count(k storage.Key) (int, error) { _ = "STUB: not implemented"; return 0, nil }
func (s *indexTrx) Put(i storage.Item) error         { _ = "STUB: not implemented"; return nil }
func (s *indexTrx) Delete(i storage.Item) error      { _ = "STUB: not implemented"; return nil }

type sharkyTrx struct {
	sharky       *sharky.Store
	metrics      metrics
	writtenLocs  []sharky.Location
	releasedLocs []sharky.Location
}

func (s *sharkyTrx) Read(ctx context.Context, loc sharky.Location, buf []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *sharkyTrx) Write(ctx context.Context, data []byte) (_ sharky.Location, err error) {
	_ = "STUB: not implemented"
	return *new(sharky.Location), nil
}

func (s *sharkyTrx) Release(ctx context.Context, loc sharky.Location) error {
	_ = "STUB: not implemented"
	return nil
}

func handleMetric(key string, m metrics) func(*error) { _ = "STUB: not implemented"; return nil }
