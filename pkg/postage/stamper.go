// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postage

import (
	"errors"

	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// ErrBucketFull is the error when a collision bucket is full.
var ErrBucketFull = errors.New("bucket full")

// Stamper can issue stamps from the given address of chunk.
type Stamper interface {
	// addr is the request address of the chunk and idAddr is the identity address of the chunk.
	Stamp(addr, idAddr swarm.Address) (*Stamp, error)
	BatchId() []byte
}

// stamper connects a stampissuer with a signer.
// A stamper is created for each upload session.
type stamper struct {
	store  storage.Store
	issuer *StampIssuer
	signer crypto.Signer
}

// NewStamper constructs a Stamper.
func NewStamper(store storage.Store, issuer *StampIssuer, signer crypto.Signer) Stamper {
	_ = "STUB: not implemented"
	return *new(Stamper)
}

// Stamp takes chunk, see if the chunk can be included in the batch and
// signs it with the owner of the batch of this Stamp issuer.
func (st *stamper) Stamp(addr, idAddr swarm.Address) (*Stamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BatchId gives back batch id of stamper
func (st *stamper) BatchId() []byte { _ = "STUB: not implemented"; return nil }

type presignedStamper struct {
	stamp *Stamp
	owner []byte
}

func NewPresignedStamper(stamp *Stamp, owner []byte) Stamper {
	_ = "STUB: not implemented"
	return *new(Stamper)
}

func (st *presignedStamper) Stamp(addr, _ swarm.Address) (*Stamp, error) {
	_ = "STUB: not implemented"
	// check stored stamp is against the chunk address
	// Recover the public key from the signature
	return nil, nil
}

func (st *presignedStamper) BatchId() []byte { _ = "STUB: not implemented"; return nil }
