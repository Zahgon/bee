// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redistribution

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/bmt"
	"github.com/ethersphere/bee/v2/pkg/storer"
)

type ChunkInclusionProofs struct {
	A ChunkInclusionProof `json:"proof1"`
	B ChunkInclusionProof `json:"proof2"`
	C ChunkInclusionProof `json:"proofLast"`
}

// ChunkInclusionProof structure must exactly match
// corresponding structure (of the same name) in Redistribution.sol smart contract.
// github.com/ethersphere/storage-incentives/blob/ph_f2/src/Redistribution.sol
// github.com/ethersphere/storage-incentives/blob/master/src/Redistribution.sol (when merged to master)
type ChunkInclusionProof struct {
	ProofSegments  []common.Hash `json:"proofSegments"`
	ProveSegment   common.Hash   `json:"proveSegment"`
	ProofSegments2 []common.Hash `json:"proofSegments2"`
	ProveSegment2  common.Hash   `json:"proveSegment2"`
	ChunkSpan      uint64        `json:"chunkSpan"`
	ProofSegments3 []common.Hash `json:"proofSegments3"`
	PostageProof   PostageProof  `json:"postageProof"`
	SocProof       []SOCProof    `json:"socProof"`
}

// SOCProof structure must exactly match
// corresponding structure (of the same name) in Redistribution.sol smart contract.
type PostageProof struct {
	Signature []byte      `json:"signature"`
	PostageId common.Hash `json:"postageId"`
	Index     uint64      `json:"index"`
	TimeStamp uint64      `json:"timeStamp"`
}

// SOCProof structure must exactly match
// corresponding structure (of the same name) in Redistribution.sol smart contract.
type SOCProof struct {
	Signer     common.Address `json:"signer"`
	Signature  []byte         `json:"signature"`
	Identifier common.Hash    `json:"identifier"`
	ChunkAddr  common.Hash    `json:"chunkAddr"`
}

// NewChunkInclusionProof transforms arguments to ChunkInclusionProof object
func NewChunkInclusionProof(proofp1, proofp2, proofp3 bmt.Proof, sampleItem storer.SampleItem) (ChunkInclusionProof, error) {
	_ = "STUB: not implemented"
	return *new(ChunkInclusionProof), nil
}

// should be uint64 on the other size; copied from pkg/api/bytes.go

func toCommonHash(hashes [][]byte) []common.Hash { _ = "STUB: not implemented"; return nil }

func makeSOCProof(sampleItem storer.SampleItem) ([]SOCProof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
