// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/transaction"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// PutterCloserWithReference provides a Putter which can be closed with a root
// swarm reference associated with this session.
type PutterCloserWithReference interface {
	Put(context.Context, transaction.Store, swarm.Chunk) error
	Close(storage.IndexStore, swarm.Address) error
	Cleanup(transaction.Storage) error
}

var emptyAddr = make([]byte, swarm.HashSize)

// AddressOrZero returns swarm.ZeroAddress if the buf is of zero bytes. The Zero byte
// buffer is used by the items to serialize their contents and if valid swarm.ZeroAddress
// entries are allowed.
func AddressOrZero(buf []byte) swarm.Address { _ = "STUB: not implemented"; return *new(swarm.Address) }

// AddressBytesOrZero is a helper which creates a zero buffer of swarm.HashSize. This
// is required during storing the items in the Store as their serialization formats
// are strict.
func AddressBytesOrZero(addr swarm.Address) []byte { _ = "STUB: not implemented"; return nil }

// NewInmemStorage constructs a inmem Storage implementation which can be used
// for the tests in the internal packages.
func NewInmemStorage() transaction.Storage {
	_ = "STUB: not implemented"
	return *new(transaction.Storage)
}

type inmemStorage struct {
	indexStore storage.IndexStore
	chunkStore storage.ChunkStore
}

func (t *inmemStorage) NewTransaction(ctx context.Context) (transaction.Transaction, func()) {
	_ = "STUB: not implemented"
	return *new(transaction.Transaction), nil
}

type inmemTrx struct {
	indexStore storage.IndexStore
	chunkStore storage.ChunkStore
}

func (t *inmemStorage) IndexStore() storage.Reader {
	_ = "STUB: not implemented"
	return *new(storage.Reader)
}
func (t *inmemStorage) ChunkStore() storage.ReadOnlyChunkStore {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyChunkStore)
}

func (t *inmemTrx) IndexStore() storage.IndexStore {
	_ = "STUB: not implemented"
	return *new(storage.IndexStore)
}
func (t *inmemTrx) ChunkStore() storage.ChunkStore {
	_ = "STUB: not implemented"
	return *new(storage.ChunkStore)
}
func (t *inmemTrx) Commit() error { _ = "STUB: not implemented"; return nil }

func (t *inmemStorage) Close() error { _ = "STUB: not implemented"; return nil }
func (t *inmemStorage) Run(ctx context.Context, f func(s transaction.Store) error) error {
	_ = "STUB: not implemented"
	return nil
}
