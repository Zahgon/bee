// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package inmemstore

import (
	"context"
	"sync"

	storage "github.com/ethersphere/bee/v2/pkg/storage"
)

// batchOp represents a batch operations.
type batchOp interface {
	Item() storage.Item
}

// batchOpBase is a base type for batch operations holding data.
type batchOpBase struct{ item storage.Item }

// Item implements batchOp interface Item method.
func (b batchOpBase) Item() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

type (
	batchOpPut    struct{ batchOpBase }
	batchOpDelete struct{ batchOpBase }
)

type Batch struct {
	ctx context.Context

	mu    sync.Mutex // mu guards batch, ops, and done.
	ops   map[string]batchOp
	store *Store
	done  bool
}

// Batch implements storage.BatchedStore interface Batch method.
func (s *Store) Batch(ctx context.Context) storage.Batch {
	_ = "STUB: not implemented"
	return *new(storage.Batch)
}

// Put implements storage.Batch interface Put method.
func (i *Batch) Put(item storage.Item) error { _ = "STUB: not implemented"; return nil }

// Delete implements storage.Batch interface Delete method.
func (i *Batch) Delete(item storage.Item) error { _ = "STUB: not implemented"; return nil }

// Commit implements storage.Batch interface Commit method.
func (i *Batch) Commit() error { _ = "STUB: not implemented"; return nil }
