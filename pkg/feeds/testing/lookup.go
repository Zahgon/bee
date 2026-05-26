// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package testing provides tests for update  and resolution of time-based feeds
package testing

import (
	"context"
	"testing"
	"time"

	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/feeds"
	storage "github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type Timeout struct {
	storage.ChunkStore
}

var searchTimeout = 30 * time.Millisecond

// Get overrides the mock storer and introduces latency
func (t *Timeout) Get(ctx context.Context, addr swarm.Address) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// nolint:tparallel
func TestFinderBasic(t *testing.T, finderf func(storage.Getter, *feeds.Feed) feeds.Lookup, updaterf func(putter storage.Putter, signer crypto.Signer, topic []byte) (feeds.Updater, error)) {
	_ = "STUB: not implemented"
	return
}

// nolint:tparallel
func TestFinderFixIntervals(t *testing.T, nextf func() (bool, int64), finderf func(storage.Getter, *feeds.Feed) feeds.Lookup, updaterf func(putter storage.Putter, signer crypto.Signer, topic []byte) (feeds.Updater, error)) {
	_ = "STUB: not implemented"
	return
}

func TestFinderIntervals(t *testing.T, nextf func() (bool, int64), finderf func(storage.Getter, *feeds.Feed) feeds.Lookup, updaterf func(putter storage.Putter, signer crypto.Signer, topic []byte) (feeds.Updater, error)) {
	_ = "STUB: not implemented"
	return
}

func TestFinderRandomIntervals(t *testing.T, finderf func(storage.Getter, *feeds.Feed) feeds.Lookup, updaterf func(putter storage.Putter, signer crypto.Signer, topic []byte) (feeds.Updater, error)) {
	_ = "STUB: not implemented"
	return
}
