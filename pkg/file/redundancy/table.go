// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redundancy

type erasureTable struct {
	shards   []int
	parities []int
}

// newErasureTable initializes a shards<->parities table
//
//	 the value order must be strictly descending in both arrays
//		example usage:
//			shards := []int{94, 68, 46, 28, 14, 5, 1}
//			parities := []int{9, 8, 7, 6, 5, 4, 3}
//			var et = newErasureTable(shards, parities)
func newErasureTable(shards, parities []int) erasureTable {
	_ = "STUB: not implemented"
	return *new(erasureTable)
}

// getParities gives back the optimal parity number for a given shard
func (et *erasureTable) getParities(maxShards int) int { _ = "STUB: not implemented"; return 0 }
