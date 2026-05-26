// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package api

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/storageincentives/redistribution"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type RCHashResponse struct {
	Hash            swarm.Address        `json:"hash"`
	Proofs          ChunkInclusionProofs `json:"proofs"`
	DurationSeconds float64              `json:"durationSeconds"`
}

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
	ProofSegments  []string     `json:"proofSegments"`
	ProveSegment   string       `json:"proveSegment"`
	ProofSegments2 []string     `json:"proofSegments2"`
	ProveSegment2  string       `json:"proveSegment2"`
	ChunkSpan      uint64       `json:"chunkSpan"`
	ProofSegments3 []string     `json:"proofSegments3"`
	PostageProof   PostageProof `json:"postageProof"`
	SocProof       []SOCProof   `json:"socProof"`
}

// SOCProof structure must exactly match
// corresponding structure (of the same name) in Redistribution.sol smart contract.
type PostageProof struct {
	Signature string `json:"signature"`
	PostageId string `json:"postageId"`
	Index     string `json:"index"`
	TimeStamp string `json:"timeStamp"`
}

// SOCProof structure must exactly match
// corresponding structure (of the same name) in Redistribution.sol smart contract.
type SOCProof struct {
	Signer     string `json:"signer"`
	Signature  string `json:"signature"`
	Identifier string `json:"identifier"`
	ChunkAddr  string `json:"chunkAddr"`
}

func renderChunkInclusionProofs(proofs redistribution.ChunkInclusionProofs) ChunkInclusionProofs {
	_ = "STUB: not implemented"
	return *new(ChunkInclusionProofs)
}

func renderChunkInclusionProof(proof redistribution.ChunkInclusionProof) ChunkInclusionProof {
	_ = "STUB: not implemented"
	return *new(ChunkInclusionProof)
}

func renderCommonHash(proofSegments []common.Hash) []string { _ = "STUB: not implemented"; return nil }

// This API is kept for testing the sampler. As a result, no documentation or tests are added here.
func (s *Service) rchash(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }
