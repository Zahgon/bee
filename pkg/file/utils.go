// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"github.com/ethersphere/bee/v2/pkg/file/redundancy"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var zeroAddress = [32]byte{}

// ChunkPayloadSize returns the effective byte length of an intermediate chunk
// assumes data is always chunk size (without span)
func ChunkPayloadSize(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// ChunkAddresses returns data shards and parities of the intermediate chunk
// assumes data is truncated by ChunkPayloadSize
func ChunkAddresses(data []byte, parities, reflen int) (addrs []swarm.Address, shardCnt int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// ReferenceCount brute-forces the data shard count from which identify the parity count as well in a substree
// assumes span > swarm.chunkSize
// returns data and parity shard number
func ReferenceCount(span uint64, level redundancy.Level, encrytedChunk bool) (int, int) {
	_ = "STUB: not implemented"
	// assume we have a trie of size `span` then we can assume that all of
	// the forks except for the last one on the right are of equal size
	// this is due to how the splitter wraps levels.
	// first the algorithm will search for a BMT level where span can be included
	// then identify how large data one reference can hold on that level
	// then count how many references can satisfy span
	// and finally how many parity shards should be on that level
	return 0, 0
}

// branching factor is how many data shard references can fit into one intermediate chunk

// search for branch level big enough to include span

// span in one full reference

// referenceSize = branching ** (branchLevel - 1)
