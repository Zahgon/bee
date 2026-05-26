// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storagetest

import (
	"testing"

	"github.com/ethersphere/bee/v2/pkg/encryption"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/google/go-cmp/cmp"
)

var (
	// MinAddressBytes represents bytes that can be used to represent a min. address.
	MinAddressBytes = [swarm.HashSize]byte{swarm.HashSize - 1: 0x00}

	// MaxAddressBytes represents bytes that can be used to represent a max. address.
	MaxAddressBytes = [swarm.HashSize]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	MaxEncryptedRefBytes = [encryption.ReferenceSize]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	// MaxStampIndexBytes represents bytes that can be used to represent a max. stamp index.
	MaxStampIndexBytes = [swarm.StampIndexSize]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	// MaxBatchTimestampBytes represents bytes that can be used to represent a max. batch timestamp.
	MaxBatchTimestampBytes = [swarm.StampTimestampSize]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
)

var _ storage.Item = (*ItemStub)(nil)

// ItemStub is a stub for storage.Item.
type ItemStub struct {
	MarshalBuf   []byte
	MarshalErr   error
	UnmarshalBuf []byte
}

// ID implements the storage.Item interface.
func (im ItemStub) ID() string { _ = "STUB: not implemented"; return "" }

// Namespace implements the storage.Item interface.
func (im ItemStub) Namespace() string {
	_ = "STUB: not implemented"

	// Marshal implements the storage.Item interface.
	return ""
}

func (im ItemStub) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal implements the storage.Item interface.
func (im *ItemStub) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// Clone implements the storage.Item interface.
func (im *ItemStub) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

// Clone implements the storage.Item interface.
func (im ItemStub) String() string { _ = "STUB: not implemented"; return "" }

type obj1 struct {
	Id      string
	SomeInt uint64
	Buf     []byte
}

func (o *obj1) ID() string { _ = "STUB: not implemented"; return "" }

func (obj1) Namespace() string { _ = "STUB: not implemented"; return "" }

func (o *obj1) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (o *obj1) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }

func (o *obj1) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

func (o obj1) String() string { _ = "STUB: not implemented"; return "" }

type obj2 struct {
	Id        int
	SomeStr   string
	SomeFloat float64
}

func (o *obj2) ID() string { _ = "STUB: not implemented"; return "" }

func (obj2) Namespace() string { _ = "STUB: not implemented"; return "" }

func (o *obj2) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (o *obj2) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }

func (o *obj2) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

func (o obj2) String() string { _ = "STUB: not implemented"; return "" }

func randBytes(count int) []byte { _ = "STUB: not implemented"; return nil }

func checkTestItemEqual(t *testing.T, a, b storage.Item) { _ = "STUB: not implemented"; return }

// TestStore provides correctness testsuite for Store interface.
func TestStore(t *testing.T, s storage.Store) { _ = "STUB: not implemented"; return }

// ItemMarshalAndUnmarshalTest represents a test case
// for the TestItemMarshalAndUnmarshal function.
type ItemMarshalAndUnmarshalTest struct {
	Item         storage.Item
	Factory      func() storage.Item
	MarshalErr   error // Expected error from Marshal.
	UnmarshalErr error // Expected error from Unmarshal.
	CmpOpts      []cmp.Option
}

// TestItemMarshalAndUnmarshal provides correctness testsuite
// for storage.Item serialization and deserialization.
func TestItemMarshalAndUnmarshal(t *testing.T, test *ItemMarshalAndUnmarshalTest) {
	_ = "STUB: not implemented"
	return
}

// ItemCloneTest represents a test case for the TestItemClone function.
type ItemCloneTest struct {
	Item    storage.Item
	CmpOpts []cmp.Option
}

// TestItemClone provides correctness testsuite for storage.Item clone capabilities.
func TestItemClone(t *testing.T, test *ItemCloneTest) { _ = "STUB: not implemented"; return }

func BenchmarkStore(b *testing.B, s storage.Store) { _ = "STUB: not implemented"; return }

// BenchmarkBatchedStore provides a benchmark suite for the
// storage.BatchedStore. Only the Write and Delete methods are tested.
func BenchmarkBatchedStore(b *testing.B, bs storage.BatchStore) { _ = "STUB: not implemented"; return }

func BenchmarkReadRandom(b *testing.B, db storage.Store) { _ = "STUB: not implemented"; return }

func BenchmarkReadRandomMissing(b *testing.B, db storage.Store) { _ = "STUB: not implemented"; return }

func BenchmarkReadSequential(b *testing.B, db storage.Store) { _ = "STUB: not implemented"; return }

func BenchmarkReadReverse(b *testing.B, db storage.Store) { _ = "STUB: not implemented"; return }

func BenchmarkReadHot(b *testing.B, db storage.Store) { _ = "STUB: not implemented"; return }

func BenchmarkIterateSequential(b *testing.B, db storage.Store) { _ = "STUB: not implemented"; return }

func BenchmarkIterateReverse(b *testing.B, db storage.Store) { _ = "STUB: not implemented"; return }

func BenchmarkWriteSequential(b *testing.B, db storage.Store) { _ = "STUB: not implemented"; return }

func BenchmarkWriteInBatches(b *testing.B, bs storage.BatchStore) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkWriteInFixedSizeBatches(b *testing.B, bs storage.BatchStore) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkWriteRandom(b *testing.B, db storage.Store) { _ = "STUB: not implemented"; return }

func BenchmarkDeleteRandom(b *testing.B, db storage.Store) { _ = "STUB: not implemented"; return }

func BenchmarkDeleteSequential(b *testing.B, db storage.Store) { _ = "STUB: not implemented"; return }

func BenchmarkDeleteInBatches(b *testing.B, bs storage.BatchStore) {
	_ = "STUB: not implemented"
	return
}

func BenchmarkDeleteInFixedSizeBatches(b *testing.B, bs storage.BatchStore) {
	_ = "STUB: not implemented"
	return
}
