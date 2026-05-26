// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashtrie

import (
	"context"
	"errors"

	"github.com/ethersphere/bee/v2/pkg/file/pipeline"
	"github.com/ethersphere/bee/v2/pkg/file/redundancy"
	"github.com/ethersphere/bee/v2/pkg/storage"
)

var (
	errInconsistentRefs = errors.New("inconsistent references")
	errTrieFull         = errors.New("trie full")
)

const maxLevel = 8

type hashTrieWriter struct {
	ctx                    context.Context // context for put function of dispersed replica chunks
	refSize                int
	cursors                []int  // level cursors, key is level. level 0 is data level holds how many chunks were processed. Intermediate higher levels will always have LOWER cursor values.
	buffer                 []byte // keeps intermediate level data
	full                   bool   // indicates whether the trie is full. currently we support (128^7)*4096 = 2305843009213693952 bytes
	pipelineFn             pipeline.PipelineFunc
	rParams                redundancy.RedundancyParams
	parityChunkFn          redundancy.ParityChunkCallback
	chunkCounters          []uint8        // counts the chunk references in intermediate chunks. key is the chunk level.
	effectiveChunkCounters []uint8        // counts the effective  chunk references in intermediate chunks. key is the chunk level.
	maxChildrenChunks      uint8          // maximum number of chunk references in intermediate chunks.
	replicaPutter          storage.Putter // putter to save dispersed replicas of the root chunk
}

func NewHashTrieWriter(ctx context.Context, refLen int, rParams redundancy.RedundancyParams, pipelineFn pipeline.PipelineFunc, replicaPutter storage.Putter, rLevel redundancy.Level) pipeline.ChainWriter {
	_ = "STUB: not implemented"
	return *new(pipeline.ChainWriter)
}

// double size as temp workaround for weak calculation of needed buffer space

// accepts writes of hashes from the previous writer in the chain, by definition these writes
// are on level 1
func (h *hashTrieWriter) ChainWrite(p *pipeline.PipeWriteArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *hashTrieWriter) writeToIntermediateLevel(level int, parityChunk bool, span, ref, key []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// update counters

// at this point the erasure coded chunks have been written

// writeToDataLevel caches data chunks and call writeToIntermediateLevel
func (h *hashTrieWriter) writeToDataLevel(span, ref, key, data []byte) error {
	_ = "STUB: not implemented"
	// write dataChunks to the level above
	return nil
}

// wrapFullLevel wraps an existing level and writes the resulting hash to the following level
// then truncates the current level data by shifting the cursors.
// Steps are performed in the following order:
//   - take all of the data in the current level
//   - break down span and hash data
//   - sum the span size, concatenate the hash to the buffer
//   - call the short pipeline with the span and the buffer
//   - get the hash that was created, append it one level above, and if necessary, wrap that level too
//   - remove already hashed data from buffer
//
// assumes that h.chunkCounters[level] has reached h.maxChildrenChunks at fullchunk
// or redundancy.Encode was called in case of rightmost chunks
func (h *hashTrieWriter) wrapFullLevel(level int) error { _ = "STUB: not implemented"; return nil }

// sum up the spans of the level, then we need to bmt them and store it as a chunk
// then write the chunk address to the next level up

// we do not add span of parity chunks to the common because that is gibberish

// parity reference has always hash length

// this "truncates" the current level that was wrapped
// by setting the cursors to the cursors of one level above

// Sum returns the Swarm merkle-root content-addressed hash
// of an arbitrary-length binary data.
// The algorithm it uses is as follows:
//   - From level 1 till maxLevel 8, iterate:
//     -- If level data length equals 0 then continue to next level
//     -- If level data length equals 1 reference then carry over level data to next
//     -- If level data length is bigger than 1 reference then sum the level and
//     write the result to the next level
//   - Return the hash in level 8
//
// the cases are as follows:
//   - one hash in a given level, in which case we _do not_ perform a hashing operation, but just move
//     the hash to the next level, potentially resulting in a level wrap
//   - more than one hash, in which case we _do_ perform a hashing operation, appending the hash to
//     the next level
func (h *hashTrieWriter) Sum() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// level empty, continue to the next.

// this case is possible and necessary due to the carry over
// in the next switch case statement. normal writes done
// through writeToLevel will automatically wrap a full level.
// erasure encoding call is not necessary since ElevateCarrierChunk solves that

// this cursor assignment basically means:
// take the hash|span|key from this level, and append it to
// the data of the next level. you may wonder how this works:
// every time we sum a level, the sum gets written into the next level
// and the level cursor gets set to the next level's cursor (see the
// truncating at the end of wrapFullLevel). there might (or not) be
// a hash at the next level, and the cursor of the next level is
// necessarily _smaller_ than the cursor of this level, so in fact what
// happens is that due to the shifting of the cursors, the data of this
// level will appear to be concatenated with the data of the next level.
// we therefore get a "carry-over" behavior between intermediate levels
// that might or might not have data. the eventual result is that the last
// hash generated will always be carried over to the last level (8), then returned.

// replace cached chunk to the level as well

// update counters, subtracting from current level is not necessary

// call erasure encoding before writing the last chunk on the level

// more than 0 but smaller than chunk size - wrap the level to the one above it

// return the hash in the highest level, that's all we need

// save disperse replicas of the root chunk
