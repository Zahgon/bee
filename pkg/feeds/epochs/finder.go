// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package epochs

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/feeds"
	storage "github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var (
	_ feeds.Lookup = (*finder)(nil)
	_ feeds.Lookup = (*asyncFinder)(nil)
)

// finder encapsulates a chunk store getter and a feed and provides
// non-concurrent lookup methods
type finder struct {
	getter *feeds.Getter
}

// NewFinder constructs an AsyncFinder
func NewFinder(getter storage.Getter, feed *feeds.Feed) feeds.Lookup {
	_ = "STUB: not implemented"
	return *new(feeds.Lookup)
}

// At looks up the version valid at time `at`
// after is a unix time hint of the latest known update
func (f *finder) At(ctx context.Context, at int64, after uint64) (swarm.Chunk, feeds.Index, feeds.Index, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), *new(feeds.Index), *new(feeds.Index), nil
}

// common returns the lowest common ancestor for which a feed update chunk is found in the chunk store
func (f *finder) common(ctx context.Context, at int64, after uint64) (*epoch, swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return nil, *new(swarm.Chunk), nil
}

// at is a non-concurrent recursive Finder function to find the version update chunk at time `at`
func (f *finder) at(ctx context.Context, at uint64, e *epoch, ch swarm.Chunk) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// error retrieving

// epoch not found on branch
// no lower resolution

// traverse earlier branch

// epoch found
// check if timestamp is later then target

// matching update time or finest resolution

// continue traversing based on at

type result struct {
	path  *path
	chunk swarm.Chunk
	*epoch
}

// asyncFinder encapsulates a chunk store getter and a feed and provides
// non-concurrent lookup methods
type asyncFinder struct {
	getter *feeds.Getter
}

type path struct {
	at     int64
	top    *result
	bottom *result
	cancel chan struct{}
}

func newPath(at int64) *path { _ = "STUB: not implemented"; return nil }

// NewAsyncFinder constructs an AsyncFinder
func NewAsyncFinder(getter storage.Getter, feed *feeds.Feed) feeds.Lookup {
	_ = "STUB: not implemented"
	return *new(feeds.Lookup)
}

func (f *asyncFinder) get(ctx context.Context, at int64, e *epoch) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// at attempts to retrieve all epoch chunks on the path for `at` concurrently
func (f *asyncFinder) at(ctx context.Context, at int64, p *path, e *epoch, c chan<- *result) {
	_ = "STUB: not implemented"
	return
}

func (f *asyncFinder) At(ctx context.Context, at int64, after uint64) (swarm.Chunk, feeds.Index, feeds.Index, error) {
	_ = "STUB: not implemented"
	// TODO: current and next index return values need to be implemented
	return *new(swarm.Chunk), *new(feeds.Index), *new(feeds.Index), nil
}

// At looks up the version valid at time `at`
// after is a unix time hint of the latest known update
func (f *asyncFinder) asyncAt(ctx context.Context, at int64, _ uint64) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// ignore result from paths already  cancelled

// update chunk for epoch found
// return if deepest level epoch

// ignore if higher level than the deepest epoch found

// update chunk for epoch not found
// if top level than return with no update found

// if topmost epoch not found, then set bottom

// found - not found for two consecutive epochs

// cancel path

// recursive call on new path through left sister
