// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"context"
	"errors"
	"sync/atomic"

	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const (
	reserveOverCapacity = "reserveOverCapacity"
	reserveUnreserved   = "reserveUnreserved"
	batchExpiry         = "batchExpiry"
	batchExpiryDone     = "batchExpiryDone"
)

var (
	errMaxRadius            = errors.New("max radius reached")
	reserveSizeWithinRadius atomic.Uint64
)

type Syncer interface {
	// Number of active historical syncing jobs.
	SyncRate() float64
	Start(context.Context)
}

func threshold(capacity int) int { _ = "STUB: not implemented"; return 0 }

func (db *DB) startReserveWorkers(
	ctx context.Context,
	radius func() (uint8, error),
	ready chan<- struct{},
) {
	_ = "STUB: not implemented"
	return
}

// node shutdown

// syncing can now begin now that the reserver worker is running

func (db *DB) countWithinRadius(ctx context.Context) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (db *DB) reserveWorker(ctx context.Context, ready chan<- struct{}) {
	_ = "STUB: not implemented"
	return
}

func (db *DB) evictExpiredBatches(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (db *DB) getExpiredBatches() ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (db *DB) evictBatch(
	ctx context.Context,
	batchID []byte,
	evictCount int,
	upToBin uint8,
) (evicted int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// EvictBatch evicts all chunks belonging to a batch from the reserve.
func (db *DB) EvictBatch(ctx context.Context, batchID []byte) error {
	_ = "STUB: not implemented"
	return nil

	// if reserve is not configured, do nothing
}

func (db *DB) ReserveGet(ctx context.Context, addr swarm.Address, batchID []byte, stampHash []byte) (ch swarm.Chunk, err error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

func (db *DB) ReserveHas(addr swarm.Address, batchID []byte, stampHash []byte) (has bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ReservePutter returns a Putter for inserting chunks into the reserve.
func (db *DB) ReservePutter() storage.Putter {
	_ = "STUB: not implemented"
	return *new(storage.Putter)
}

func (db *DB) unreserve(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

// evict at least a min count

// eviction happens in batches, so we need to keep track of the total
// number of chunks evicted even if there was an error

// we can only get error here for critical cases, for eg. batch commit
// error, which is not recoverable

// ReserveLastBinIDs returns all of the highest binIDs from all the bins in the reserve and the epoch time of the reserve.
func (db *DB) ReserveLastBinIDs() ([]uint64, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (db *DB) ReserveIterateChunks(cb func(swarm.Chunk) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) StorageRadius() uint8 { _ = "STUB: not implemented"; return 0 }

func (db *DB) CommittedDepth() uint8 { _ = "STUB: not implemented"; return 0 }

func (db *DB) CapacityDoubling() uint8 { _ = "STUB: not implemented"; return 0 }

func (db *DB) ReserveSize() int { _ = "STUB: not implemented"; return 0 }

func (db *DB) ReserveSizeWithinRadius() uint64 { _ = "STUB: not implemented"; return 0 }

func (db *DB) IsWithinStorageRadius(addr swarm.Address) bool {
	_ = "STUB: not implemented"
	return false
}

// BinC is the result returned from the SubscribeBin channel that contains the chunk address and the binID
type BinC struct {
	Address   swarm.Address
	BinID     uint64
	BatchID   []byte
	StampHash []byte
}

// SubscribeBin returns a channel that feeds all the chunks in the reserve from a certain bin between a start and end binIDs.
func (db *DB) SubscribeBin(ctx context.Context, bin uint8, start uint64) (<-chan *BinC, func(), <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type NeighborhoodStat struct {
	Neighborhood            swarm.Neighborhood
	ReserveSizeWithinRadius int
	Proximity               uint8
}

func (db *DB) NeighborhoodsStat(ctx context.Context) ([]*NeighborhoodStat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func neighborhoodPrefixes(base swarm.Address, radius int, suffixLength int) []swarm.Address {
	_ = "STUB: not implemented"
	return nil
}

// copy base address

// set pseudo suffix

// clear rest of the bits

// Clears the bit at pos in n.
func clearBit(n, pos uint8) uint8 { _ = "STUB: not implemented"; return 0 }

// Sets the bit at pos in the integer n.
func setBit(n, pos uint8) uint8 { _ = "STUB: not implemented"; return 0 }

func hasBit(n, pos uint8) bool { _ = "STUB: not implemented"; return false }

// expiredBatchItem is a storage.Item implementation for expired batches.
type expiredBatchItem struct {
	BatchID []byte
}

// ID implements storage.Item.
func (e *expiredBatchItem) ID() string { _ = "STUB: not implemented"; return "" }

// Namespace implements storage.Item.
func (e *expiredBatchItem) Namespace() string { _ = "STUB: not implemented"; return "" }

// Marshal implements storage.Item.
// It is a no-op as expiredBatchItem is not serialized.
func (e *expiredBatchItem) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"

	// Unmarshal implements storage.Item.
	// It is a no-op as expiredBatchItem is not serialized.
	return nil, nil
}

func (e *expiredBatchItem) Unmarshal(_ []byte) error {
	_ = "STUB: not implemented"

	// Clone implements storage.Item.
	return nil
}

func (e *expiredBatchItem) Clone() storage.Item {
	_ = "STUB: not implemented"
	return *new(storage.Item)
}

// String implements storage.Item.
func (e *expiredBatchItem) String() string { _ = "STUB: not implemented"; return "" }

func (db *DB) po(addr swarm.Address) uint8 { _ = "STUB: not implemented"; return 0 }
