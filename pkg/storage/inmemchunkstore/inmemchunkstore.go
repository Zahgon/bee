// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package inmemchunkstore

import (
	"context"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type ChunkStore struct {
	mu     sync.Mutex
	chunks map[string]chunkCount
}

type chunkCount struct {
	chunk swarm.Chunk
	count int
}

func New() *ChunkStore { _ = "STUB: not implemented"; return nil }

func (c *ChunkStore) Get(_ context.Context, addr swarm.Address) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

func (c *ChunkStore) Put(_ context.Context, ch swarm.Chunk) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ChunkStore) Has(_ context.Context, addr swarm.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *ChunkStore) Delete(_ context.Context, addr swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ChunkStore) Replace(_ context.Context, ch swarm.Chunk, emplace bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ChunkStore) Iterate(_ context.Context, fn storage.IterateChunkFn) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ChunkStore) Close() error { _ = "STUB: not implemented"; return nil }

func (c *ChunkStore) key(addr swarm.Address) string { _ = "STUB: not implemented"; return "" }
