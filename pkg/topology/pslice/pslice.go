// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pslice

import (
	"sync"

	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/topology"
)

// PSlice maintains a list of addresses, indexing them by their different proximity orders.
type PSlice struct {
	peers     [][]swarm.Address // the slice of peers
	baseBytes []byte
	mu        sync.RWMutex
	maxBins   int
}

// New creates a new PSlice.
func New(maxBins int, base swarm.Address) *PSlice { _ = "STUB: not implemented"; return nil }

// Add a peer at a certain PO.
func (s *PSlice) Add(addrs ...swarm.Address) { _ = "STUB: not implemented"; return }

// bypass unnecessary allocations below if address count is one

// iterates over all peers from deepest bin to shallowest.
func (s *PSlice) EachBin(pf topology.EachPeerFunc) error { _ = "STUB: not implemented"; return nil }

// EachBinRev iterates over all peers from shallowest bin to deepest.
func (s *PSlice) EachBinRev(pf topology.EachPeerFunc) error { _ = "STUB: not implemented"; return nil }

func (s *PSlice) BinSize(bin uint8) int { _ = "STUB: not implemented"; return 0 }

func (s *PSlice) BinPeers(bin uint8) []swarm.Address { _ = "STUB: not implemented"; return nil }

// Length returns the number of peers in the Pslice.
func (s *PSlice) Length() int { _ = "STUB: not implemented"; return 0 }

// ShallowestEmpty returns the shallowest empty bin if one exists.
// If such bin does not exists, returns true as bool value.
func (s *PSlice) ShallowestEmpty() (uint8, bool) { _ = "STUB: not implemented"; return 0, false }

// Exists checks if a peer exists.
func (s *PSlice) Exists(addr swarm.Address) bool { _ = "STUB: not implemented"; return false }

// Remove a peer at a certain PO.
func (s *PSlice) Remove(addr swarm.Address) { _ = "STUB: not implemented"; return }

// Since order of elements does not matter, the optimized removing process
// below replaces the index to be removed with the last element of the array,
// and shortens the array by one.

// make copy of the bin slice with one fewer element

// if the index is the last element, then assign slice and return early

// replace index being removed with last element

// assign the copy with the index removed back to the original array

func (s *PSlice) po(peer []byte) uint8 { _ = "STUB: not implemented"; return 0 }

// index returns if a peer exists and the index in the slice.
func (s *PSlice) index(addr swarm.Address, po uint8) (bool, int) {
	_ = "STUB: not implemented"
	return false, 0
}
