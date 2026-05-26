// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package soc provides the single-owner chunk implementation
// and validator.
package soc

import (
	"errors"

	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var (
	errInvalidAddress = errors.New("soc: invalid address")
	errWrongChunkSize = errors.New("soc: chunk length is less than minimum")
)

// ID is a SOC identifier
type ID []byte

// SOC wraps a content-addressed chunk.
type SOC struct {
	id        ID
	owner     []byte // owner is the address in bytes of SOC owner.
	signature []byte
	chunk     swarm.Chunk // wrapped chunk.
}

// New creates a new SOC representation from arbitrary id and
// a content-addressed chunk.
func New(id ID, ch swarm.Chunk) *SOC { _ = "STUB: not implemented"; return nil }

// NewSigned creates a single-owner chunk based on already signed data.
func NewSigned(id ID, ch swarm.Chunk, owner, sig []byte) (*SOC, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Address returns the SOC chunk address.
func (s *SOC) Address() (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// WrappedChunk returns the chunk wrapped by the SOC.
func (s *SOC) WrappedChunk() swarm.Chunk {
	_ = "STUB: not implemented"

	// Chunk returns the SOC chunk.
	return *new(swarm.Chunk)
}

func (s *SOC) Chunk() (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// Signature returns the SOC signature.
func (s *SOC) Signature() []byte {
	_ = "STUB: not implemented"

	// OwnerAddress returns the ethereum address of the SOC owner.
	return nil
}

func (s *SOC) OwnerAddress() []byte {
	_ = "STUB: not implemented"

	// ID returns the SOC id.
	return nil
}

func (s *SOC) ID() []byte {
	_ = "STUB: not implemented"

	// toBytes is a helper function to convert the SOC data to bytes.
	return nil
}

func (s *SOC) toBytes() []byte { _ = "STUB: not implemented"; return nil }

// Sign signs a SOC using the given signer.
// It returns a signed SOC chunk ready for submission to the network.
func (s *SOC) Sign(signer crypto.Signer) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	// create owner
	return *new(swarm.Chunk), nil
}

// generate the data to sign

// sign the chunk

// UnwrapCAC extracts the CAC inside the SOC.
func UnwrapCAC(sch swarm.Chunk) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// FromChunk recreates a SOC representation from swarm.Chunk data.
func FromChunk(sch swarm.Chunk) (*SOC, error) { _ = "STUB: not implemented"; return nil, nil }

// add all the data fields to the SOC

// recover owner information

// CreateAddress creates a new SOC address from the id and
// the ethereum address of the owner.
func CreateAddress(id ID, owner []byte) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// hash hashes the given values in order.
func hash(values ...[]byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// recoverAddress returns the ethereum address of the owner of a SOC.
func recoverAddress(signature, digest []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
