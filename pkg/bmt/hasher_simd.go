// Copyright 2026 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && amd64 && !purego

package bmt

var _ Hasher = (*simdHasher)(nil)

// simdHasher is a BMT hasher implementation that buffers all data and defers hashing
// to Hash(), using SIMD-accelerated Keccak for multi-level trees.
//
// Single-threaded: SIMD provides intra-call parallelism (4-way or 8-way), replacing
// the goroutine-per-section model of the goroutine hasher.
//
// The same hasher instance must not be called concurrently on more than one chunk.
// The same hasher instance is synchronously reusable after calling Reset.
type simdHasher struct {
	*simdConf
	bmt  *simdTree
	size int
	span []byte
}

func newSIMDHasher() *simdHasher { _ = "STUB: not implemented"; return nil }

func newSIMDPrefixHasher(prefix []byte) *simdHasher { _ = "STUB: not implemented"; return nil }

// Capacity returns the maximum number of bytes this hasher can process.
func (h *simdHasher) Capacity() int {
	_ = "STUB: not implemented"

	// Size returns the digest size.
	return 0
}

func (h *simdHasher) Size() int { _ = "STUB: not implemented"; return 0 }

// BlockSize returns the optimal write block size.
func (h *simdHasher) BlockSize() int { _ = "STUB: not implemented"; return 0 }

// SetHeaderInt64 sets the span preamble from an int64.
func (h *simdHasher) SetHeaderInt64(length int64) { _ = "STUB: not implemented"; return }

// SetHeader copies the span preamble from the argument.
func (h *simdHasher) SetHeader(span []byte) {
	_ = "STUB: not implemented"

	// Write buffers input to be hashed. All hashing is deferred to Sum().
	return
}

func (h *simdHasher) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Sum returns the BMT root hash of the buffer written so far appended to b,
// satisfying the standard library hash.Hash interface.
func (h *simdHasher) Sum(b []byte) []byte {
	_ = "STUB: not implemented"
	// empty input: no data was ever written, so the BMT root is the all-zero
	// subtree hash at depth h.depth. Still prepend the span so the output
	// shape matches a normal chunk hash.
	return nil
}

// zero-fill the tail of the buffer: every leaf section must carry
// deterministic bytes, because SIMD batches hash whole sections at a time
// without a "is this section occupied?" check. clear() lowers to memclr.

// degenerate single-level tree: the whole buffer is already one section,
// so there is nothing for SIMD to batch — just run the scalar hasher.

// general case: fan out through the tree via hashSIMD, which processes
// each level bottom-up in batches of 4 or 8 hashes per SIMD call.

// prepend the span and hash once more to produce the chunk address.
// When a prefix is configured, h.bmt.hasher is a PrefixHasher whose Reset
// re-absorbs the prefix, so the final digest is
// keccak(prefix || span || rootHash) — the same wrap-per-level rule used
// by hashSIMD at every internal level.

// Reset prepares the Hasher for reuse. The internal data buffer is not zeroed;
// any stale bytes past h.size are overwritten on the next Write and zero-filled
// on demand by Hash, so post-Reset inspection of the buffer may still see
// previous-chunk contents.
func (h *simdHasher) Reset() { _ = "STUB: not implemented"; return }

// hashSIMD computes the BMT root hash using SIMD-accelerated Keccak hashing.
// It processes the tree level by level from leaves to root, using batched
// SIMD calls instead of goroutine-per-section. A single thread handles all
// levels since SIMD already provides intra-call parallelism (4-way or 8-way).
func (h *simdHasher) hashSIMD() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Leaf level: hash each section and write results to parent nodes.

// Internal levels: process each level single-threaded (diminishing work).

// Root level: hash using the tree's shared scalar hasher.

// hashLeavesBatch hashes leaf sections in the range [start, end) using SIMD batches.
func (h *simdHasher) hashLeavesBatch(start, end, bw, secsize, prefixLen int) {
	_ = "STUB: not implemented"
	// the chunk buffer holds raw leaf-level bytes back-to-back: section i lives at [i*secsize, (i+1)*secsize).
	return
}

// AVX-512 path: each SIMD call hashes up to 8 sections in lockstep.

// the trailing batch may be short — clamp so we never read past `end`.

// stage each lane's input. With a configured prefix we must materialise
// prefix||section into a scratch buffer because the prefix bytes are not
// stored in `buf` itself; without a prefix we hand the section slice
// straight to the SIMD primitive (zero-copy).

// nil out unused lanes so XKCP treats them as no-op fillers, see keccak.Sum256x8 docs.

// single 8-way SIMD permutation produces all 8 digests at once.

// each digest is the value of the parent's left or right child slot;
// write it directly into the parent so the next level can read it without copying.

// AVX2 path: identical structure to the AVX-512 branch but at width 4.

// clamp the trailing batch so we never read past `end`.

// stage prefix||section per lane (or a zero-copy slice when no prefix is configured).

// nil out unused lanes so XKCP treats them as no-op fillers.

// 4-way SIMD permutation produces all digests in one call.

// place each digest directly into the parent's left/right child slot.

// hashNodesBatch hashes a level of internal nodes using SIMD batches.
// Each node's left||right (64 bytes) is hashed to produce the input for its parent.
func (h *simdHasher) hashNodesBatch(nodes []*simdNode, bw, prefixLen int) {
	_ = "STUB: not implemented"
	return
}

// `concat` is the per-lane scratch buffer reused across calls; its leading prefixLen
// bytes already hold the configured prefix (set up at tree construction time), so
// we only ever rewrite the trailing left||right region.

// AVX-512 path: 8 nodes hashed per SIMD call.

// clamp the trailing batch so we never read past `count`.

// for each active lane, stage prefix||left||right into its scratch slot.
// the `prefixLen` bytes at the head are already populated, so we only
// overwrite the [prefixLen, prefixLen+2*segSize) range.

// nil out unused lanes so XKCP treats them as no-op fillers.

// one 8-way permutation produces all parent digests for this batch.

// drop each digest into its parent's left/right slot for the next level.

// AVX2 path: same shape as the AVX-512 branch but at width 4.

// clamp trailing batch to the remaining nodes.

// stage prefix||left||right per active lane in the per-lane scratch buffer.

// nil out unused lanes so XKCP treats them as no-op fillers.

// 4-way SIMD permutation produces all parent digests for this batch.

// place each digest directly into the parent's left/right child slot.
