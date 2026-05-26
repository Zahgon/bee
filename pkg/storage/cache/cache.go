// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/ethersphere/bee/v2/pkg/storage"
	lru "github.com/hashicorp/golang-lru/v2"
)

// key returns a string representation of the given key.
func key(key storage.Key) string { _ = "STUB: not implemented"; return "" }

var _ storage.IndexStore = (*Cache)(nil)

// Cache is a wrapper around a storage.Store that adds a layer
// of in-memory caching for the Get and Has operations.
type Cache struct {
	storage.IndexStore

	lru     *lru.Cache[string, []byte]
	metrics metrics
}

// Wrap adds a layer of in-memory caching to storage.Reader Get and Has operations.
// It returns an error if the capacity is less than or equal to zero or if the
// given store implements storage.Tx
func Wrap(store storage.IndexStore, capacity int) (*Cache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add caches given item.
func (c *Cache) add(i storage.Item) { _ = "STUB: not implemented"; return }

// Get implements storage.Store interface.
// On a call it tries to first retrieve the item from cache.
// If the item does not exist in cache, it tries to retrieve
// it from the underlying store.
func (c *Cache) Get(i storage.Item) error { _ = "STUB: not implemented"; return nil }

// Has implements storage.Store interface.
// On a call it tries to first retrieve the item from cache.
// If the item does not exist in cache, it tries to retrieve
// it from the underlying store.
func (c *Cache) Has(k storage.Key) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Put implements storage.Store interface.
// On a call it also inserts the item into the cache so that the next
// call to Put and Has will be able to retrieve the item from cache.
func (c *Cache) Put(i storage.Item) error { _ = "STUB: not implemented"; return nil }

// Delete implements storage.Store interface.
// On a call it also removes the item from the cache.
func (c *Cache) Delete(i storage.Item) error { _ = "STUB: not implemented"; return nil }

func (c *Cache) Close() error { _ = "STUB: not implemented"; return nil }
