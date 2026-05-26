// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storagetest

import (
	"flag"
	"math/rand"
	"testing"

	storage "github.com/ethersphere/bee/v2/pkg/storage"
)

var (
	valueSize        = flag.Int("value_size", 100, "Size of each value")
	compressionRatio = flag.Float64("compression_ratio", 0.5, "")
	maxConcurrency   = flag.Int("max_concurrency", 2048, "Max concurrency in concurrent benchmark")
	batchSize        = flag.Int("batch_size", 1000, "Max number of records that would trigger commit")
)

var keyLen = 16

const (
	hitKeyFormat     = "1%015d"
	missingKeyFormat = "0%015d"
)

func randomBytes(r *rand.Rand, n int) []byte { _ = "STUB: not implemented"; return nil }

func compressibleBytes(r *rand.Rand, ratio float64, valueSize int) []byte {
	_ = "STUB: not implemented"
	return nil
}

type randomValueGenerator struct {
	b []byte
	k int
}

func (g *randomValueGenerator) Value(i int) []byte { _ = "STUB: not implemented"; return nil }

func makeRandomValueGenerator(r *rand.Rand, ratio float64, valueSize int) randomValueGenerator {
	_ = "STUB: not implemented"
	return *new(randomValueGenerator)
}

type entryGenerator interface {
	keyGenerator
	Value(i int) []byte
}

type pairedEntryGenerator struct {
	keyGenerator
	randomValueGenerator
}

type startAtEntryGenerator struct {
	entryGenerator
	start int
}

var _ entryGenerator = (*startAtEntryGenerator)(nil)

func (g *startAtEntryGenerator) NKey() int { _ = "STUB: not implemented"; return 0 }

func (g *startAtEntryGenerator) Key(i int) []byte { _ = "STUB: not implemented"; return nil }

func newStartAtEntryGenerator(start int, g entryGenerator) entryGenerator {
	_ = "STUB: not implemented"
	return *new(entryGenerator)
}

func newSequentialKeys(size int, start int, keyFormat string) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

func newRandomKeys(n int, format string) [][]byte { _ = "STUB: not implemented"; return nil }

func newFullRandomKeys(size int, start int, format string) [][]byte {
	_ = "STUB: not implemented"
	return nil
}

func newFullRandomEntryGenerator(start, size int) entryGenerator {
	_ = "STUB: not implemented"
	return *new(entryGenerator)
}

func newSequentialEntryGenerator(size int) entryGenerator {
	_ = "STUB: not implemented"
	return *new(entryGenerator)
}

type keyGenerator interface {
	NKey() int
	Key(i int) []byte
}

type reversedKeyGenerator struct {
	keyGenerator
}

var _ keyGenerator = (*reversedKeyGenerator)(nil)

func (g *reversedKeyGenerator) Key(i int) []byte { _ = "STUB: not implemented"; return nil }

func newReversedKeyGenerator(g keyGenerator) keyGenerator {
	_ = "STUB: not implemented"
	return *new(keyGenerator)
}

type roundKeyGenerator struct {
	keyGenerator
}

var _ keyGenerator = (*roundKeyGenerator)(nil)

func (g *roundKeyGenerator) Key(i int) []byte { _ = "STUB: not implemented"; return nil }

func newRoundKeyGenerator(g keyGenerator) keyGenerator {
	_ = "STUB: not implemented"
	return *new(keyGenerator)
}

type predefinedKeyGenerator struct {
	keys [][]byte
}

func (g *predefinedKeyGenerator) NKey() int { _ = "STUB: not implemented"; return 0 }

func (g *predefinedKeyGenerator) Key(i int) []byte { _ = "STUB: not implemented"; return nil }

func newRandomKeyGenerator(n int) keyGenerator {
	_ = "STUB: not implemented"
	return *new(keyGenerator)
}

func newRandomMissingKeyGenerator(n int) keyGenerator {
	_ = "STUB: not implemented"
	return *new(keyGenerator)
}

func newFullRandomKeyGenerator(start, n int) keyGenerator {
	_ = "STUB: not implemented"
	return *new(keyGenerator)
}

func newSequentialKeyGenerator(n int) keyGenerator {
	_ = "STUB: not implemented"
	return *new(keyGenerator)
}

func maxInt(a int, b int) int { _ = "STUB: not implemented"; return 0 }

func doRead(b *testing.B, db storage.Store, g keyGenerator, allowNotFound bool) {
	_ = "STUB: not implemented"
	return
}

type singularDBWriter struct {
	db storage.Store
}

func (w *singularDBWriter) Put(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (w *singularDBWriter) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func newDBWriter(db storage.Store) *singularDBWriter { _ = "STUB: not implemented"; return nil }

func doWrite(b *testing.B, db storage.Store, g entryGenerator) { _ = "STUB: not implemented"; return }

func doDelete(b *testing.B, db storage.Store, g keyGenerator) { _ = "STUB: not implemented"; return }

func resetBenchmark(b *testing.B) { _ = "STUB: not implemented"; return }

func populate(b *testing.B, db storage.Store) { _ = "STUB: not implemented"; return }

// chunk
func doDeleteChunk(b *testing.B, db storage.ChunkStore, g keyGenerator) {
	_ = "STUB: not implemented"
	return
}

func doWriteChunk(b *testing.B, db storage.Putter, g entryGenerator) {
	_ = "STUB: not implemented"
	return
}

func doReadChunk(b *testing.B, db storage.ChunkStore, g keyGenerator, allowNotFound bool) {
	_ = "STUB: not implemented"
	return
}

// fixed size batch
type batchDBWriter struct {
	db    storage.Batcher
	batch storage.Batch
	max   int
	count int
}

func (w *batchDBWriter) commit(maxValue int) { _ = "STUB: not implemented"; return }

func (w *batchDBWriter) Put(key, value []byte) { _ = "STUB: not implemented"; return }

func (w *batchDBWriter) Delete(key []byte) { _ = "STUB: not implemented"; return }

func newBatchDBWriter(db storage.Batcher) *batchDBWriter { _ = "STUB: not implemented"; return nil }
