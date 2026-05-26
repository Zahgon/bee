// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reserve

import (
	"context"
	"sync/atomic"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/transaction"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/topology"
	"resenje.org/multex"
)

const reserveScope = "reserve"

type Reserve struct {
	baseAddr     swarm.Address
	radiusSetter topology.SetStorageRadiuser
	logger       log.Logger

	capacity int
	size     atomic.Int64
	radius   atomic.Uint32

	multx *multex.Multex
	st    transaction.Storage
}

func New(
	baseAddr swarm.Address,
	st transaction.Storage,
	capacity int,
	radiusSetter topology.SetStorageRadiuser,
	logger log.Logger,
) (*Reserve, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reserve Put has to handle multiple possible scenarios.
//  1. Since the same chunk may belong to different postage stamp indices, the reserve will support one chunk to many postage
//     stamp indices relationship.
//  2. A new chunk that shares the same stamp index belonging to the same batch with an already stored chunk will overwrite
//     the existing chunk if the new chunk has a higher stamp timestamp (regardless of batch type).
//  3. A new chunk that has the same address belonging to the same stamp index with an already stored chunk will overwrite the existing chunk
//     if the new chunk has a higher stamp timestamp (regardless of batch type and chunk type, eg CAC & SOC).
func (r *Reserve) Put(ctx context.Context, chunk swarm.Chunk) error {
	_ = "STUB: not implemented"
	// batchID lock, Put vs Eviction
	return nil
}

// check if the chunk with the same batch, stamp timestamp and index is already stored

// bin lock

// index collision

// same chunk address

// load item to get the binID

// delete old chunk index items

// An older and different chunk with the same batchID and stamp index has been previously
// saved to the reserve. We must do the below before saving the new chunk:
// 1. Delete the old chunk from the chunkstore.
// 2. Delete the old chunk's stamp data.
// 3. Delete ALL old chunk related items from the reserve.
// 4. Update the stamp index.

// replace old stamp index.

func (r *Reserve) Has(addr swarm.Address, batchID []byte, stampHash []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Reserve) Get(ctx context.Context, addr swarm.Address, batchID []byte, stampHash []byte) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// EvictBatchBin evicts all chunks from bins upto the bin provided.
// Pinned chunks are protected from eviction to maintain data integrity.
func (r *Reserve) EvictBatchBin(
	ctx context.Context,
	batchID []byte,
	count int,
	bin uint8,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Check if the chunk is pinned in any collection

func (r *Reserve) removeChunk(
	ctx context.Context,
	trx transaction.Store,
	chunkAddress swarm.Address,
	batchID []byte,
	stampHash []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

func RemoveChunkWithItem(
	ctx context.Context,
	trx transaction.Store,
	item *BatchRadiusItem,
) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveChunkMetaData removes chunk reserve metadata from reserve indexes but keeps the cunks in the chunkstore.
// used at pinned data eviction
func RemoveChunkMetaData(
	ctx context.Context,
	trx transaction.Store,
	item *BatchRadiusItem,
) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteCorruptedChunkMetadata removes all reserve index entries for a chunk
// whose Sharky data was found to be corrupted during recovery. It is intended
// to be called from the recovery path, where only a storage.IndexStore (not a
// full transaction.Store) is available. If the chunk has no reserve metadata
// (e.g. it belongs to the upload store or cache), the function is a no-op.
func DeleteCorruptedChunkMetadata(store storage.IndexStore, baseAddr swarm.Address, addr swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reserve) IterateBin(bin uint8, startBinID uint64, cb func(swarm.Address, uint64, []byte, []byte) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reserve) IterateChunks(startBin uint8, cb func(swarm.Chunk) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reserve) IterateChunksItems(startBin uint8, cb func(*ChunkBinItem) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset removes all the entries in the reserve. Must be done before any calls to the reserve.
func (r *Reserve) Reset(ctx context.Context) error {
	_ = "STUB: not implemented"

	// step 1: delete epoch timestamp
	return nil
}

// step 2: delete batchRadiusItem, chunkBinItem, and the chunk data

// step 3: delete stampindex and chunkstamp

// step 4: delete binItems

func (r *Reserve) Radius() uint8 { _ = "STUB: not implemented"; return 0 }

func (r *Reserve) Size() int { _ = "STUB: not implemented"; return 0 }

func (r *Reserve) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (r *Reserve) IsWithinCapacity() bool { _ = "STUB: not implemented"; return false }

func (r *Reserve) EvictionTarget() int { _ = "STUB: not implemented"; return 0 }

func (r *Reserve) SetRadius(rad uint8) error { _ = "STUB: not implemented"; return nil }

func (r *Reserve) LastBinIDs() ([]uint64, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (r *Reserve) IncBinID(store storage.IndexStore, bin uint8) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
