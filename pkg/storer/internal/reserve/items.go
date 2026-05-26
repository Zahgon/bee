// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reserve

import (
	"errors"

	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var (
	errMarshalInvalidAddress = errors.New("marshal: invalid address")
	errUnmarshalInvalidSize  = errors.New("unmarshal: invalid size")
)

// BatchRadiusItem allows iteration of the chunks with respect to bin and batchID.
// Used for batch evictions of certain bins.
type BatchRadiusItem struct {
	Bin       uint8
	BatchID   []byte
	StampHash []byte
	Address   swarm.Address
	BinID     uint64
}

func (b *BatchRadiusItem) Namespace() string { _ = "STUB: not implemented"; return "" }

// batchID/bin/ChunkAddr/stampHash
func (b *BatchRadiusItem) ID() string { _ = "STUB: not implemented"; return "" }

func (b *BatchRadiusItem) String() string { _ = "STUB: not implemented"; return "" }

func (b *BatchRadiusItem) Clone() storage.Item {
	_ = "STUB: not implemented"
	return *new(storage.Item)
}

const batchRadiusItemSize = 1 + swarm.HashSize + swarm.HashSize + 8 + swarm.HashSize

func (b *BatchRadiusItem) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *BatchRadiusItem) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }

// ChunkBinItem allows for iterating on ranges of bin and binIDs for chunks.
// BinIDs come in handy when syncing the reserve contents with other peers.
type ChunkBinItem struct {
	Bin       uint8
	BinID     uint64
	Address   swarm.Address
	BatchID   []byte
	StampHash []byte
	ChunkType swarm.ChunkType
}

func (c *ChunkBinItem) Namespace() string {
	_ = "STUB: not implemented"

	// bin/binID
	return ""
}

func (c *ChunkBinItem) ID() string { _ = "STUB: not implemented"; return "" }

func binIDToString(bin uint8, binID uint64) string { _ = "STUB: not implemented"; return "" }

func (c *ChunkBinItem) String() string { _ = "STUB: not implemented"; return "" }

func (c *ChunkBinItem) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

const chunkBinItemSize = 1 + 8 + swarm.HashSize + swarm.HashSize + 1 + swarm.HashSize

func (c *ChunkBinItem) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *ChunkBinItem) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }

// BinItem stores the latest binIDs for each bin between 0 and swarm.MaxBins
type BinItem struct {
	Bin   uint8
	BinID uint64
}

func (b *BinItem) Namespace() string { _ = "STUB: not implemented"; return "" }

func (b *BinItem) ID() string { _ = "STUB: not implemented"; return "" }

func (c *BinItem) String() string { _ = "STUB: not implemented"; return "" }

func (b *BinItem) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

const binItemSize = 8

func (c *BinItem) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *BinItem) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }

// EpochItem stores the timestamp in seconds of the initial creation of the reserve.
type EpochItem struct {
	Timestamp uint64
}

func (e *EpochItem) Namespace() string   { _ = "STUB: not implemented"; return "" }
func (e *EpochItem) ID() string          { _ = "STUB: not implemented"; return "" }
func (e *EpochItem) String() string      { _ = "STUB: not implemented"; return "" }
func (e *EpochItem) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

const epochItemSize = 8

func (e *EpochItem) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EpochItem) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }

// radiusItem stores the current storage radius of the reserve.
type radiusItem struct {
	Radius uint8
}

func (r *radiusItem) Namespace() string { _ = "STUB: not implemented"; return "" }

func (r *radiusItem) ID() string { _ = "STUB: not implemented"; return "" }

func (r *radiusItem) String() string { _ = "STUB: not implemented"; return "" }

func (r *radiusItem) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

func (r *radiusItem) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *radiusItem) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }

func copyBytes(src []byte) []byte { _ = "STUB: not implemented"; return nil }
