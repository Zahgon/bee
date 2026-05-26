// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stampindex

import (
	"errors"

	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var (
	// errStampItemMarshalScopeInvalid is returned when trying to
	// marshal a Item with invalid scope.
	errStampItemMarshalScopeInvalid = errors.New("marshal stampindex.Item: scope is invalid")
	// errStampItemMarshalBatchIDInvalid is returned when trying to
	// marshal a Item with invalid batchID.
	errStampItemMarshalBatchIDInvalid = errors.New("marshal stampindex.Item: batchID is invalid")
	// errStampItemMarshalBatchIndexInvalid is returned when trying
	// to marshal a Item with invalid batchIndex.
	errStampItemMarshalBatchIndexInvalid = errors.New("marshal stampindex.Item: batchIndex is invalid")
	// errStampItemUnmarshalInvalidSize is returned when trying
	// to unmarshal buffer with smaller size then is the size
	// of the Item fields.
	errStampItemUnmarshalInvalidSize = errors.New("unmarshal stampindex.Item: invalid size")
)

var _ storage.Item = (*Item)(nil)

// Item is an store.Item that represents data relevant to stamp.
type Item struct {
	// Keys.
	scope      []byte // The scope of other related item.
	BatchID    []byte
	StampIndex []byte
	StampHash  []byte

	// Values.
	StampTimestamp []byte
	ChunkAddress   swarm.Address
}

// ID implements the storage.Item interface.
func (i Item) ID() string { _ = "STUB: not implemented"; return "" }

// Namespace implements the storage.Item interface.
func (i Item) Namespace() string { _ = "STUB: not implemented"; return "" }

func (i Item) GetScope() []byte { _ = "STUB: not implemented"; return nil }

func (i *Item) SetScope(ns []byte) {
	_ = "STUB: not implemented"

	// Marshal implements the storage.Item interface.
	return
}

func (i Item) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal implements the storage.Item interface.
func (i *Item) Unmarshal(bytes []byte) error { _ = "STUB: not implemented"; return nil }

// Clone  implements the storage.Item interface.
func (i *Item) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

// String implements the fmt.Stringer interface.
func (i Item) String() string { _ = "STUB: not implemented"; return "" }

// LoadOrStore tries to first load a stamp index related record from the store.
// If the record is not found, it will try to create and save a new record and
// return it.
func LoadOrStore(
	s storage.IndexStore,
	scope string,
	chunk swarm.Chunk,
) (item *Item, loaded bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Load returns stamp index record related to the given scope and stamp.
func Load(s storage.Reader, scope string, stamp swarm.Stamp) (*Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Store creates new or updated an existing stamp index
// record related to the given scope and chunk.
func Store(s storage.IndexStore, scope string, chunk swarm.Chunk) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete removes the related stamp index record from the storage.
func Delete(s storage.Writer, scope string, stamp swarm.Stamp) error {
	_ = "STUB: not implemented"
	return nil
}
