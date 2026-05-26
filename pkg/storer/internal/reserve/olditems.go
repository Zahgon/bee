// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reserve

import (
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// BatchRadiusItemV1 allows iteration of the chunks with respect to bin and batchID.
// Used for batch evictions of certain bins.
type BatchRadiusItemV1 struct {
	Bin     uint8
	BatchID []byte
	Address swarm.Address
	BinID   uint64
}

func (b *BatchRadiusItemV1) Namespace() string { _ = "STUB: not implemented"; return "" }

func (b *BatchRadiusItemV1) ID() string { _ = "STUB: not implemented"; return "" }

func (b *BatchRadiusItemV1) String() string { _ = "STUB: not implemented"; return "" }

func (b *BatchRadiusItemV1) Clone() storage.Item {
	_ = "STUB: not implemented"
	return *new(storage.Item)
}

const batchRadiusItemSizeV1 = 1 + swarm.HashSize + swarm.HashSize + 8

func (b *BatchRadiusItemV1) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *BatchRadiusItemV1) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }

// ChunkBinItemV1 allows for iterating on ranges of bin and binIDs for chunks.
// BinIDs come in handy when syncing the reserve contents with other peers.
type ChunkBinItemV1 struct {
	Bin       uint8
	BinID     uint64
	Address   swarm.Address
	BatchID   []byte
	ChunkType swarm.ChunkType
}

func (c *ChunkBinItemV1) Namespace() string { _ = "STUB: not implemented"; return "" }

func (c *ChunkBinItemV1) ID() string { _ = "STUB: not implemented"; return "" }

func (c *ChunkBinItemV1) String() string { _ = "STUB: not implemented"; return "" }

func (c *ChunkBinItemV1) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

const chunkBinItemSizeV1 = 1 + 8 + swarm.HashSize + swarm.HashSize + 1

func (c *ChunkBinItemV1) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *ChunkBinItemV1) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }
