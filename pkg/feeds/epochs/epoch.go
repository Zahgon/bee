// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package epochs implements time-based feeds using epochs as index
// and provide sequential as well as concurrent lookup algorithms
package epochs

import (
	"github.com/ethersphere/bee/v2/pkg/feeds"
)

const (
	maxLevel = 32
)

var _ feeds.Index = (*epoch)(nil)

// epoch is referencing a slot in the epoch grid and represents an update
// it  implements the feeds.Index interface
type epoch struct {
	start uint64
	level uint8
}

func (e *epoch) String() string { _ = "STUB: not implemented"; return "" }

// MarshalBinary implements the BinaryMarshaler interface
func (e *epoch) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func next(e feeds.Index, last int64, at uint64) feeds.Index {
	_ = "STUB: not implemented"
	return *new(feeds.Index)
}

// Next implements feeds.Index advancement
func (e *epoch) Next(last int64, at uint64) feeds.Index {
	_ = "STUB: not implemented"
	return *new(feeds.Index)
}

// lca calculates the lowest common ancestor epoch given two unix times
func lca(at, after uint64) *epoch { _ = "STUB: not implemented"; return nil }

// parent returns the ancestor of an epoch
// the call is unsafe in that it must not be called on a toplevel epoch
func (e *epoch) parent() *epoch { _ = "STUB: not implemented"; return nil }

// left returns the left sister of an epoch
// it is unsafe in that it must not be called on a left sister epoch
func (e *epoch) left() *epoch { _ = "STUB: not implemented"; return nil }

// at returns the left of right child epoch of an epoch depending on where `at` falls
// it is unsafe in that it must not be called with an at that does not fall within the epoch
func (e *epoch) childAt(at uint64) *epoch { _ = "STUB: not implemented"; return nil }

// isLeft returns true if epoch is a left sister of its parent
func (e *epoch) isLeft() bool { _ = "STUB: not implemented"; return false }

// length returns the span of the epoch
func (e *epoch) length() uint64 { _ = "STUB: not implemented"; return 0 }
