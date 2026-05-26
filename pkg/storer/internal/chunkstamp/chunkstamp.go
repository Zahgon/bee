// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chunkstamp

import (
	"errors"

	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var (
	// errMarshalInvalidChunkStampItemScope is returned during marshaling if the scope is not set.
	errMarshalInvalidChunkStampItemScope = errors.New("marshal chunkstamp.Item: invalid scope")
	// errMarshalInvalidChunkStampAddress is returned during marshaling if the address is zero.
	errMarshalInvalidChunkStampItemAddress = errors.New("marshal chunkstamp.item: invalid address")
	// errUnmarshalInvalidChunkStampAddress is returned during unmarshaling if the address is not set.
	errUnmarshalInvalidChunkStampItemAddress = errors.New("unmarshal chunkstamp.item: invalid address")
	// errMarshalInvalidChunkStamp is returned if the stamp is invalid during marshaling.
	errMarshalInvalidChunkStampItemStamp = errors.New("marshal chunkstamp.item: invalid stamp")
	// errUnmarshalInvalidChunkStampSize is returned during unmarshaling if the passed buffer is not the expected size.
	errUnmarshalInvalidChunkStampItemSize = errors.New("unmarshal chunkstamp.item: invalid size")
)

var _ storage.Item = (*Item)(nil)

// Item is the index used to represent a stamp for a chunk.
//
// Going ahead we will support multiple stamps on chunks. This Item will allow
// mapping multiple stamps to a single address. For this reason, the address is
// part of the Namespace and can be used to iterate on all the stamps for this
// address.
type Item struct {
	scope   []byte // The scope of other related item.
	address swarm.Address
	stamp   swarm.Stamp
}

// ID implements the storage.Item interface.
func (i *Item) ID() string { _ = "STUB: not implemented"; return "" }

// Namespace implements the storage.Item interface.
func (i *Item) Namespace() string { _ = "STUB: not implemented"; return "" }

func (i *Item) SetScope(ns []byte) {
	_ = "STUB: not implemented"

	// Marshal implements the storage.Item interface.
	// address is not part of the payload which is stored, as address is part of the
	// prefix, hence already known before querying this object. This will be reused
	// during unmarshaling.
	return
}

func (i *Item) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	// The address is not part of the payload, but it is used to create the
	// scope so it is better if we check that the address is correctly
	// set here before it is stored in the underlying storage.
	return nil, nil
}

// Unmarshal implements the storage.Item interface.
func (i *Item) Unmarshal(bytes []byte) error { _ = "STUB: not implemented"; return nil }

// Ensure that the address is set already in the item.

// Clone implements the storage.Item interface.
func (i *Item) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

// String implements the storage.Item interface.
func (i Item) String() string { _ = "STUB: not implemented"; return "" }

// Load returns first found swarm.Stamp related to the given address.
func Load(s storage.Reader, scope string, addr swarm.Address) (swarm.Stamp, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Stamp), nil
}

// LoadWithBatchID returns swarm.Stamp related to the given address and batchID.
func LoadWithBatchID(s storage.Reader, scope string, addr swarm.Address, batchID []byte) (swarm.Stamp, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Stamp), nil
}

// LoadWithStampHash returns swarm.Stamp related to the given address and stamphash.
func LoadWithStampHash(s storage.Reader, scope string, addr swarm.Address, hash []byte) (swarm.Stamp, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Stamp), nil
}

// Store creates new or updated an existing stamp index
// record related to the given scope and chunk.
func Store(s storage.IndexStore, scope string, chunk swarm.Chunk) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteAll removes all swarm.Stamp related to the given address.
func DeleteAll(s storage.IndexStore, scope string, addr swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete removes a stamp associated with an chunk and batchID.
func Delete(s storage.IndexStore, scope string, addr swarm.Address, batchId []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteWithStamp(
	writer storage.Writer,
	scope string,
	addr swarm.Address,
	stamp swarm.Stamp,
) error {
	_ = "STUB: not implemented"
	return nil
}
