// Copyright 2026 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bmt

import (
	"hash"
)

// goroutineConf is the internal configuration for the goroutine BMT pool.
type goroutineConf struct {
	segmentSize  int
	segmentCount int
	capacity     int
	depth        int
	maxSize      int
	zerohashes   [][]byte
	prefix       []byte
	hasherFunc   func() hash.Hash
}

func (c *goroutineConf) baseHasher() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

// goroutinePool is the goroutine-based BMT hasher pool.
type goroutinePool struct {
	c chan *goroutineTree
	*goroutineConf
}

func newGoroutineConf(prefix []byte, segmentCount, capacity int) *goroutineConf {
	_ = "STUB: not implemented"
	return nil
}

// newGoroutinePool creates a goroutine BMT pool from a public Conf.
func newGoroutinePool(c *Conf) *goroutinePool { _ = "STUB: not implemented"; return nil }

// Get returns a BMT hasher, possibly reusing a tree from the pool.
func (p *goroutinePool) Get() Hasher { _ = "STUB: not implemented"; return *new(Hasher) }

// Put returns a hasher's tree to the pool for reuse.
func (p *goroutinePool) Put(h Hasher) { _ = "STUB: not implemented"; return }

// goroutineProverPool is a pool of goroutine-backed Provers. Trees are reused
// via a channel; the hasher shell is reallocated per GetProver (cheap relative
// to a tree).
type goroutineProverPool struct {
	c chan *goroutineTree
	*goroutineConf
}

func newGoroutineProverPool(c *Conf) *goroutineProverPool { _ = "STUB: not implemented"; return nil }

// GetProver returns a goroutine-backed Prover, reusing a tree from the pool.
func (p *goroutineProverPool) GetProver() *Prover { _ = "STUB: not implemented"; return nil }

// PutProver returns a Prover's tree to the pool for reuse.
func (p *goroutineProverPool) PutProver(pr *Prover) {
	_ = "STUB: not implemented"

	// goroutineTree is the tree structure used by the goroutine hasher.
	return
}

type goroutineTree struct {
	leaves []*goroutineNode
	buffer []byte
}

// goroutineNode is a reusable segment hasher node in the goroutine BMT.
type goroutineNode struct {
	isLeft      bool
	parent      *goroutineNode
	state       int32
	left, right []byte
	hasher      hash.Hash
}

func newGoroutineNode(index int, parent *goroutineNode, hasher hash.Hash) *goroutineNode {
	_ = "STUB: not implemented"
	return nil
}

func newGoroutineTree(maxsize, depth int, hashfunc func() hash.Hash) *goroutineTree {
	_ = "STUB: not implemented"
	return nil
}

// toggle implements atomic bool toggle for goroutineNode coordination.
func (n *goroutineNode) toggle() bool { _ = "STUB: not implemented"; return false }

// getSister returns a copy of the sibling section on the opposite side.
func (n *goroutineNode) getSister(isLeft bool) []byte { _ = "STUB: not implemented"; return nil }
