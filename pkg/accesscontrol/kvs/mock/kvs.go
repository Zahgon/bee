// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mock provides an in-memory key-value store implementation.
package mock

import (
	"context"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/accesscontrol/kvs"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var (
	lock       = &sync.Mutex{}
	lockGetPut = &sync.Mutex{}
)

type single struct {
	memoryMock map[string]map[string][]byte
}

var singleInMemorySwarm *single

func getInMemorySwarm() *single { _ = "STUB: not implemented"; return nil }

func getMemory() map[string]map[string][]byte { _ = "STUB: not implemented"; return nil }

type mockKeyValueStore struct {
	address swarm.Address
}

var _ kvs.KeyValueStore = (*mockKeyValueStore)(nil)

func (m *mockKeyValueStore) Get(_ context.Context, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *mockKeyValueStore) Put(_ context.Context, key []byte, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mockKeyValueStore) Save(ctx context.Context) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func New() kvs.KeyValueStore { _ = "STUB: not implemented"; return *new(kvs.KeyValueStore) }

func NewReference(address swarm.Address) kvs.KeyValueStore {
	_ = "STUB: not implemented"
	return *new(kvs.KeyValueStore)
}
