// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"context"
)

type UploadStat struct {
	TotalUploaded uint64
	TotalSynced   uint64
	PendingUpload uint64
}

type PinningStat struct {
	TotalCollections int
	TotalChunks      int
}

type CacheStat struct {
	Size     int
	Capacity int
}

type ReserveStat struct {
	SizeWithinRadius int
	TotalSize        int
	Capacity         int
	LastBinIDs       []uint64
	Epoch            uint64
}

type ChunkStoreStat struct {
	TotalChunks    int
	SharedSlots    int
	ReferenceCount int
}

type Info struct {
	Upload     UploadStat
	Pinning    PinningStat
	Cache      CacheStat
	Reserve    ReserveStat
	ChunkStore ChunkStoreStat
}

func (db *DB) DebugInfo(ctx context.Context) (Info, error) {
	_ = "STUB: not implemented"
	return *new(Info), nil
}
