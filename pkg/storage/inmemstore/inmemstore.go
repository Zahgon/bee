// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package inmemstore

import (
	"sync"

	"github.com/armon/go-radix"
	"github.com/ethersphere/bee/v2/pkg/storage"
)

const (
	separator = "/"
)

// Store implements an in-memory Store. We will use the hashicorp/go-radix implementation.
// This pkg provides a mutable radix which gives O(k) lookup and ordered iteration.
type Store struct {
	st *radix.Tree
	mu sync.RWMutex
}

func New() *Store { _ = "STUB: not implemented"; return nil }

func key(i storage.Key) string { _ = "STUB: not implemented"; return "" }

func idFromKey(key, pfx string) string { _ = "STUB: not implemented"; return "" }

func (s *Store) Get(i storage.Item) error { _ = "STUB: not implemented"; return nil }

func (s *Store) Has(k storage.Key) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (s *Store) GetSize(k storage.Key) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Store) Put(i storage.Item) error { _ = "STUB: not implemented"; return nil }

func (s *Store) put(i storage.Item) error { _ = "STUB: not implemented"; return nil }

func (s *Store) Delete(i storage.Item) error { _ = "STUB: not implemented"; return nil }

func (s *Store) delete(i storage.Item) error { _ = "STUB: not implemented"; return nil }

func (s *Store) Count(k storage.Key) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Store) Iterate(q storage.Query, fn storage.IterateFn) error {
	_ = "STUB: not implemented"
	return nil
}

// currently there is no optimal way to do reverse iteration. We can efficiently do forward
// iteration. So we have two options, first is to reduce time complexity by compromising
// on space complexity. So we keep track of keys and values during forward iteration
// to do a simple reverse iteration. Other option is to reduce space complexity by keeping
// track of only keys during forward iteration, then use Get to read the value on reverse
// iteration. This would involve additional complexity of doing a Get on reverse iteration.
// For now, inmem implementation is not meant to work for large datasets, so first option
// is chosen.

func (s *Store) Close() error { _ = "STUB: not implemented"; return nil }
