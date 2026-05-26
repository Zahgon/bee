// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package swarm

// ContainsAddress reports whether a is present in addrs.
func ContainsAddress(addrs []Address, a Address) bool { _ = "STUB: not implemented"; return false }

// RemoveAddress removes first occurrence of a in addrs, returning the modified slice.
func RemoveAddress(addrs []Address, a Address) []Address { _ = "STUB: not implemented"; return nil }

// IndexOfAddress returns the index of the first occurrence of a in addrs,
// or -1 if not present.
func IndexOfAddress(addrs []Address, a Address) int { _ = "STUB: not implemented"; return 0 }

// IndexOfChunkWithAddress returns the index of the first occurrence of
// Chunk with Address a in chunks, or -1 if not present.
func IndexOfChunkWithAddress(chunks []Chunk, a Address) int { _ = "STUB: not implemented"; return 0 }

// ContainsChunkWithAddress reports whether Chunk with Address a is present in chunks.
func ContainsChunkWithAddress(chunks []Chunk, a Address) bool {
	_ = "STUB: not implemented"
	return false
}

// ContainsChunkWithData reports whether Chunk with data d is present in chunks.
func ContainsChunkWithData(chunks []Chunk, d []byte) bool { _ = "STUB: not implemented"; return false }

// FindStampWithBatchID returns the first occurrence of Stamp having the same batchID.
func FindStampWithBatchID(stamps []Stamp, batchID []byte) (Stamp, bool) {
	_ = "STUB: not implemented"
	return *new(Stamp), false
}
