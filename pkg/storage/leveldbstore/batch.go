// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldbstore

import (
	"context"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/storage"
	ldb "github.com/syndtr/goleveldb/leveldb"
)

// Batch implements storage.BatchedStore interface Batch method.
func (s *Store) Batch(ctx context.Context) storage.Batch {
	_ = "STUB: not implemented"
	return *new(storage.Batch)
}

type Batch struct {
	ctx context.Context

	mu    sync.Mutex // mu guards batch and done.
	batch *ldb.Batch
	store *Store
	done  bool
}

// Put implements storage.Batch interface Put method.
func (i *Batch) Put(item storage.Item) error { _ = "STUB: not implemented"; return nil }

// Delete implements storage.Batch interface Delete method.
func (i *Batch) Delete(item storage.Item) error { _ = "STUB: not implemented"; return nil }

// Commit implements storage.Batch interface Commit method.
func (i *Batch) Commit() error { _ = "STUB: not implemented"; return nil }
