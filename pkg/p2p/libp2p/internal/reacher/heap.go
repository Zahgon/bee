// Copyright 2026 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reacher

// peerHeap is a min-heap of peers ordered by retryAfter time.
type peerHeap []*peer

func (h peerHeap) Len() int           { _ = "STUB: not implemented"; return 0 }
func (h peerHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (h peerHeap) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func (h *peerHeap) Push(x any) { _ = "STUB: not implemented"; return }

func (h *peerHeap) Pop() any { _ = "STUB: not implemented"; return *new(any) }

// avoid memory leak
// for safety
