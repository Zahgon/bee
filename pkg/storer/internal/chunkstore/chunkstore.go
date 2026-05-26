// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chunkstore

import (
	"context"
	"errors"

	"github.com/ethersphere/bee/v2/pkg/sharky"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var (
	// errMarshalInvalidRetrievalIndexAddress is returned if the RetrievalIndexItem address is zero during marshaling.
	errMarshalInvalidRetrievalIndexAddress = errors.New("marshal RetrievalIndexItem: address is zero")
	// errMarshalInvalidRetrievalIndexLocation is returned if the RetrievalIndexItem location is invalid during marshaling.
	errMarshalInvalidRetrievalIndexLocation = errors.New("marshal RetrievalIndexItem: location is invalid")
	// errUnmarshalInvalidRetrievalIndexSize is returned during unmarshaling if the passed buffer is not the expected size.
	errUnmarshalInvalidRetrievalIndexSize = errors.New("unmarshal RetrievalIndexItem: invalid size")
	// errUnmarshalInvalidRetrievalIndexLocationBytes is returned during unmarshaling if the location buffer is invalid.
	errUnmarshalInvalidRetrievalIndexLocationBytes = errors.New("unmarshal RetrievalIndexItem: invalid location bytes")
)

const RetrievalIndexItemSize = swarm.HashSize + 8 + sharky.LocationSize + 4

var _ storage.Item = (*RetrievalIndexItem)(nil)

// Sharky provides an abstraction for the sharky.Store operations used in the
// chunkstore. This allows us to be more flexible in passing in the sharky instance
// to chunkstore. For eg, check the TxChunkStore implementation in this pkg.
type Sharky interface {
	Read(context.Context, sharky.Location, []byte) error
	Write(context.Context, []byte) (sharky.Location, error)
	Release(context.Context, sharky.Location) error
}

func Get(ctx context.Context, r storage.Reader, s storage.Sharky, addr swarm.Address) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// helper to read chunk from retrievalIndex.
func readChunk(ctx context.Context, s storage.Sharky, rIdx *RetrievalIndexItem) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

func Has(_ context.Context, r storage.Reader, addr swarm.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func Put(ctx context.Context, s storage.IndexStore, sh storage.Sharky, ch swarm.Chunk) error {
	_ = "STUB: not implemented"
	return nil
}

// if this is the first instance of this address, we should store the chunk
// in sharky and create the new indexes.

func Replace(ctx context.Context, s storage.IndexStore, sh storage.Sharky, ch swarm.Chunk, emplace bool) error {
	_ = "STUB: not implemented"
	return nil
}

func Delete(ctx context.Context, s storage.IndexStore, sh storage.Sharky, addr swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// If there are more references for this we don't delete it from sharky.

func Iterate(ctx context.Context, s storage.IndexStore, sh storage.Sharky, fn storage.IterateChunkFn) error {
	_ = "STUB: not implemented"
	return nil
}

func IterateChunkEntries(st storage.Reader, fn func(swarm.Address, uint32) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

type LocationResult struct {
	Err      error
	Location sharky.Location
}

type IterateResult struct {
	Err  error
	Item *RetrievalIndexItem
}

// IterateLocations iterates over entire retrieval index and plucks only sharky location.
func IterateLocations(
	ctx context.Context,
	st storage.Reader,
) <-chan LocationResult {
	_ = "STUB: not implemented"
	return nil
}

// Iterate iterates over entire retrieval index with a call back.
func IterateItems(st storage.Store, callBackFunc func(*RetrievalIndexItem) error) error {
	_ = "STUB: not implemented"
	return nil
}

// RetrievalIndexItem is the index which gives us the sharky location from the swarm.Address.
// The RefCnt stores the reference of each time a Put operation is issued on this Address.
type RetrievalIndexItem struct {
	Address   swarm.Address
	Timestamp uint64
	Location  sharky.Location
	RefCnt    uint32
}

func (r *RetrievalIndexItem) ID() string { _ = "STUB: not implemented"; return "" }

func (RetrievalIndexItem) Namespace() string { _ = "STUB: not implemented"; return "" }

// Stored in bytes as:
// |--Address(32)--|--Timestamp(8)--|--Location(7)--|--RefCnt(4)--|
func (r *RetrievalIndexItem) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *RetrievalIndexItem) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }

func (r *RetrievalIndexItem) Clone() storage.Item {
	_ = "STUB: not implemented"
	return *new(storage.Item)
}

func (r RetrievalIndexItem) String() string { _ = "STUB: not implemented"; return "" }
