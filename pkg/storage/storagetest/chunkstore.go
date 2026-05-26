// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storagetest

import (
	"testing"

	storage "github.com/ethersphere/bee/v2/pkg/storage"
)

// TestChunkStore runs a correctness test suite on a given ChunkStore.
func TestChunkStore(t *testing.T, st storage.ChunkStore) { _ = "STUB: not implemented"; return }

// Delete all even numbered indexes along with 0

// delete twice as it was put twice

// Check even numbered indexes are deleted

// Check rest of the entries are intact

func RunChunkStoreBenchmarkTests(b *testing.B, s storage.ChunkStore) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkChunkStoreWriteSequential(b *testing.B, s storage.Putter) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkChunkStoreWriteRandom(b *testing.B, s storage.Putter) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkChunkStoreReadSequential(b *testing.B, s storage.ChunkStore) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkChunkStoreReadRandom(b *testing.B, s storage.ChunkStore) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkChunkStoreReadRandomMissing(b *testing.B, s storage.ChunkStore) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkChunkStoreReadReverse(b *testing.B, db storage.ChunkStore) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkChunkStoreReadHot(b *testing.B, s storage.ChunkStore) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkChunkStoreIterateSequential(b *testing.B, s storage.ChunkStore) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkChunkStoreIterateReverse(b *testing.B, s storage.ChunkStore) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkChunkStoreDeleteRandom(b *testing.B, s storage.ChunkStore) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkChunkStoreDeleteSequential(b *testing.B, s storage.ChunkStore) {
	_ = "STUB: not implemented"
	return
}
