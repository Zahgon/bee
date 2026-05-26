// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethersphere/bee/v2/pkg/bmt"
	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const SampleSize = 16

type SampleItem struct {
	TransformedAddress swarm.Address
	ChunkAddress       swarm.Address
	ChunkData          []byte
	Stamp              *postage.Stamp
}

type Sample struct {
	Stats SampleStats
	Items []SampleItem
}

// ReserveSample generates the sample of reserve storage of a node required for the
// storage incentives agent to participate in the lottery round. In order to generate
// this sample we need to iterate through all the chunks in the node's reserve and
// calculate the transformed hashes of all the chunks using the anchor as the salt.
// In order to generate the transformed hashes, we will use the std hmac keyed-hash
// implementation by using the anchor as the key. Nodes need to calculate the sample
// in the most optimal way and there are time restrictions. The lottery round is a
// time based round, so nodes participating in the round need to perform this
// calculation within the round limits.
// In order to optimize this we use a simple pipeline pattern:
// Iterate chunk addresses -> Get the chunk data and calculate transformed hash -> Assemble the sample
// If the node has doubled their capacity by some factor, sampling process need to only pertain to the
// chunks of the selected neighborhood as determined by the anchor and the "committed depth" and NOT the whole reserve.
// The committed depth is the sum of the radius and the doubling factor.
// For example, the committed depth is 11, but the local node has a doubling factor of 3, so the
// local radius will eventually drop to 8. The sampling must only consider chunks with proximity 11 to the anchor.
func (db *DB) ReserveSample(
	ctx context.Context,
	anchor []byte,
	committedDepth uint8,
	consensusTime uint64,
	minBatchBalance *big.Int,
) (Sample, error) {
	_ = "STUB: not implemented"
	return *new(Sample), nil
}

// Phase 1: Iterate chunk addresses

// Phase 2: Get the chunk data and calculate transformed hash

// exclude chunks who's batches balance are below minimum

// Skip chunks if they are not SOC or CAC

// insert function will insert the new item in its correct place. If the sample
// size goes beyond what we need we omit the last item.

// ensuring to pass the check order function of redistribution contract
// replace the chunk at index if the chunk is CAC

// Phase 3: Assemble the sample. Here we need to assemble only the first SampleSize
// no of items from the results of the 2nd phase.
// In this step stamps are loaded and validated only if chunk will be added to sample.

// check if the timestamp on the postage stamp is not later than the consensus time.

// less function uses the byte compare to check for lexicographic ordering
func le(a, b swarm.Address) bool { _ = "STUB: not implemented"; return false }

func (db *DB) batchesBelowValue(until *big.Int) (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func transformedAddress(hasher bmt.Hasher, chunk swarm.Chunk, chType swarm.ChunkType) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func transformedAddressCAC(hasher bmt.Hasher, chunk swarm.Chunk) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func transformedAddressSOC(hasher bmt.Hasher, socChunk swarm.Chunk) (swarm.Address, error) {
	_ = "STUB: not implemented"
	// Calculate transformed address from wrapped chunk
	return *new(swarm.Address), nil
}

// Hash address and transformed address to make transformed address for this SOC

type SampleStats struct {
	TotalDuration             time.Duration
	TotalIterated             int64
	IterationDuration         time.Duration
	SampleInserts             int64
	NewIgnored                int64
	InvalidStamp              int64
	BelowBalanceIgnored       int64
	TaddrDuration             time.Duration
	ValidStampDuration        time.Duration
	BatchesBelowValueDuration time.Duration
	RogueChunk                int64
	ChunkLoadDuration         time.Duration
	ChunkLoadFailed           int64
	StampLoadFailed           int64
}

func (s *SampleStats) add(other SampleStats) { _ = "STUB: not implemented"; return }

// RandSample returns Sample with random values.
func RandSample(t *testing.T, anchor []byte) Sample { _ = "STUB: not implemented"; return *new(Sample) }

// MakeSampleUsingChunks returns Sample constructed using supplied chunks.
func MakeSampleUsingChunks(chunks []swarm.Chunk, anchor []byte) (Sample, error) {
	_ = "STUB: not implemented"
	return *new(Sample), nil
}

func newStamp(s swarm.Stamp) *postage.Stamp { _ = "STUB: not implemented"; return nil }

func getChunkType(chunk swarm.Chunk) swarm.ChunkType {
	_ = "STUB: not implemented"
	return *new(swarm.ChunkType)
}

func (db *DB) recordReserveSampleMetrics(duration time.Duration, stats *SampleStats, workers int, err error) {
	_ = "STUB: not implemented"
	return
}
