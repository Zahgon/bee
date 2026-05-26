// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package leveldb

import (
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/storage"

	"github.com/syndtr/goleveldb/leveldb"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "leveldb"

var _ storage.StateStorer = (*Store)(nil)

// Store uses LevelDB to store values.
type Store struct {
	db     *leveldb.DB
	logger log.Logger
}

func NewInMemoryStateStore(l log.Logger) (*Store, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewStateStore creates a new persistent state storage.
func NewStateStore(path string, l log.Logger) (*Store, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get retrieves a value of the requested key. If no results are found,
// storage.ErrNotFound will be returned.
func (s *Store) Get(key string, i any) error { _ = "STUB: not implemented"; return nil }

// Put stores a value for an arbitrary key. BinaryMarshaler
// interface method will be called on the provided value
// with fallback to JSON serialization.
func (s *Store) Put(key string, i any) (err error) { _ = "STUB: not implemented"; return nil }

// Delete removes entries stored under a specific key.
func (s *Store) Delete(key string) (err error) { _ = "STUB: not implemented"; return nil }

// Iterate entries that match the supplied prefix.
func (s *Store) Iterate(prefix string, iterFunc storage.StateIterFunc) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Close releases the resources used by the store.
func (s *Store) Close() error { _ = "STUB: not implemented"; return nil }
