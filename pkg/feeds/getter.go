// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feeds

import (
	"context"
	"errors"

	storage "github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var ErrNotLegacyPayload = errors.New("feed update is not in the legacy payload structure")

type WrappedChunkNotFoundError struct {
	Ref []byte
}

func (e WrappedChunkNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

// Lookup is the interface for time based feed lookup
type Lookup interface {
	At(ctx context.Context, at int64, after uint64) (chunk swarm.Chunk, currentIndex, nextIndex Index, err error)
}

// Getter encapsulates a chunk Getter getter and a feed and provides non-concurrent lookup methods
type Getter struct {
	getter storage.Getter
	*Feed
}

// NewGetter constructs a feed Getter
func NewGetter(getter storage.Getter, feed *Feed) *Getter { _ = "STUB: not implemented"; return nil }

// Latest looks up the latest update of the feed
// after is a unix time hint of the latest known update
func Latest(ctx context.Context, l Lookup, after uint64) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// Get creates an update of the underlying feed at the given epoch
// and looks it up in the chunk Getter based on its address
func (f *Getter) Get(ctx context.Context, i Index) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

func GetWrappedChunk(ctx context.Context, getter storage.Getter, ch swarm.Chunk, legacyResolve bool) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// try to split the timestamp and reference
// possible values right now:
// unencrypted ref: span+timestamp+ref => 8+8+32=48
// encrypted ref: span+timestamp+ref+decryptKey => 8+8+64=80

// FromChunk parses out the wrapped chunk
func FromChunk(ch swarm.Chunk) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// legacyPayload returns back the referenced chunk and datetime from the legacy feed payload
func legacyPayload(wrappedChunk swarm.Chunk) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func IsV1Payload(ch swarm.Chunk) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func isV1Length(length int) bool { _ = "STUB: not implemented"; return false }
