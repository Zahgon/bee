// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package kvs provides functionalities needed
// for storing key-value pairs on Swarm.
//
//nolint:ireturn
package kvs

import (
	"context"
	"errors"

	"github.com/ethersphere/bee/v2/pkg/file"
	"github.com/ethersphere/bee/v2/pkg/manifest"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var (
	// ErrNothingToSave indicates that no new key-value pair was added to the store.
	ErrNothingToSave = errors.New("nothing to save")
	// ErrNotFound is returned when an Entry is not found in the storage.
	ErrNotFound = errors.New("kvs entry not found")
)

// KeyValueStore represents a key-value store.
type KeyValueStore interface {
	// Get retrieves the value associated with the given key.
	Get(ctx context.Context, key []byte) ([]byte, error)
	// Put stores the given key-value pair in the store.
	Put(ctx context.Context, key, value []byte) error
	// Save saves key-value pair to the underlying storage and returns the reference.
	Save(ctx context.Context) (swarm.Address, error)
}

type keyValueStore struct {
	manifest manifest.Interface
	putCnt   int
}

var _ KeyValueStore = (*keyValueStore)(nil)

// Get retrieves the value associated with the given key.
func (s *keyValueStore) Get(ctx context.Context, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put stores the given key-value pair in the store.
func (s *keyValueStore) Put(ctx context.Context, key []byte, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Save saves key-value pair to the underlying storage and returns the reference.
func (s *keyValueStore) Save(ctx context.Context) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// New creates a new key-value store with a simple manifest.
func New(ls file.LoadSaver) (KeyValueStore, error) {
	_ = "STUB: not implemented"
	return *new(KeyValueStore), nil
}

// NewReference loads a key-value store with a simple manifest.
func NewReference(ls file.LoadSaver, ref swarm.Address) (KeyValueStore, error) {
	_ = "STUB: not implemented"
	return *new(KeyValueStore), nil
}
