// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storageincentives

import (
	"errors"

	"github.com/ethersphere/bee/v2/pkg/storageincentives/redistribution"
	storer "github.com/ethersphere/bee/v2/pkg/storer"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var errProofCreation = errors.New("reserve commitment hasher: failure in proof creation")

// spanOffset returns the byte index of chunkdata where the spansize starts
func spanOffset(sampleItem storer.SampleItem) uint8 { _ = "STUB: not implemented"; return 0 }

// makeInclusionProofs creates transaction data for claim method.
// In the document this logic, result data, is also called Proof of entitlement (POE).
func makeInclusionProofs(
	reserveSampleItems []storer.SampleItem,
	anchor1 []byte,
	anchor2 []byte,
) (redistribution.ChunkInclusionProofs, error) {
	_ = "STUB: not implemented"
	return *new(redistribution.ChunkInclusionProofs), nil
}

// Sample chunk proofs

// Witness1 proofs

// OG chunk proof

// TR chunk proof

// cleanup

// Witness2 proofs
// OG Chunk proof

// TR Chunk proof

// cleanup

// Witness3 proofs
// OG Chunk proof

// TR Chunk Proof

// cleanup

// map to output and add SOC related data if it is necessary

func sampleChunk(items []storer.SampleItem) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

func sampleHash(items []storer.SampleItem) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}
