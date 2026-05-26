// Copyright 2026 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && amd64 && !purego

package bmt

import (
	"hash"
)

// simdConf is the internal configuration for the SIMD BMT pool.
type simdConf struct {
	segmentSize  int
	segmentCount int
	capacity     int
	depth        int
	maxSize      int
	zerohashes   [][]byte
	prefix       []byte
	batchWidth   int
}

func (c *simdConf) baseHasher() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

// simdPool is the SIMD-batched BMT hasher pool.
type simdPool struct {
	c chan *simdTree
	*simdConf
}

func newSIMDConf(prefix []byte, segmentCount, capacity int) *simdConf {
	_ = "STUB: not implemented"
	return nil
}

// newSIMDPool creates a SIMD BMT pool from a public Conf.
func newSIMDPool(c *Conf) *simdPool { _ = "STUB: not implemented"; return nil }

func (p *simdPool) Get() Hasher { _ = "STUB: not implemented"; return *new(Hasher) }

func (p *simdPool) Put(h Hasher) { _ = "STUB: not implemented"; return }

// simdTree is the tree structure used by the SIMD hasher.
type simdTree struct {
	leaves     []*simdNode
	levels     [][]*simdNode
	buffer     []byte
	concat     [8][]byte
	leafConcat [8][]byte
	// hasher is the scalar keccak (prefix-aware when configured) shared across
	// every scalar hash performed during Sum: the root combine and the outer
	// span wrap. The SIMD batch paths use the stateless keccak.Sum256xN
	// primitives instead, so this hasher is the only Hash instance the tree
	// ever needs. Safe to share because simdHasher is single-threaded per
	// contract and the pool hands out one tree at a time.
	hasher hash.Hash
}

// simdNode is a reusable segment hasher node in the SIMD BMT.
type simdNode struct {
	isLeft      bool
	parent      *simdNode
	left, right []byte
}

func newSIMDNode(index int, parent *simdNode, size int) *simdNode {
	_ = "STUB: not implemented"
	return nil
}

func newSIMDTree(maxsize, depth int, hashfunc func() hash.Hash, prefix []byte) *simdTree {
	_ = "STUB: not implemented"
	return nil
}

// reverse so levels[0]=leaves, levels[len-1]=root
