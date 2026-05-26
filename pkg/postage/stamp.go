// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postage

import (
	"errors"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// StampSize is the number of bytes in the serialisation of a stamp
const (
	StampSize   = 113
	IndexSize   = 8
	BucketDepth = 16
)

var (
	// ErrOwnerMismatch is the error given for invalid signatures.
	ErrOwnerMismatch = errors.New("owner mismatch")
	// ErrInvalidIndex the error given for invalid stamp index.
	ErrInvalidIndex = errors.New("invalid index")
	// ErrStampInvalid is the error given if stamp cannot deserialise.
	ErrStampInvalid = errors.New("invalid stamp")
	// ErrBucketMismatch is the error given if stamp index bucket verification fails.
	ErrBucketMismatch = errors.New("bucket mismatch")
	// ErrInvalidBatchID is the error returned if the batch ID is incorrect
	ErrInvalidBatchID = errors.New("invalid batch ID")
	// ErrInvalidBatchIndex is the error returned if the batch index is incorrect
	ErrInvalidBatchIndex = errors.New("invalid batch index")
	// ErrInvalidBatchTimestamp is the error returned if the batch timestamp is incorrect
	ErrInvalidBatchTimestamp = errors.New("invalid batch timestamp")
	// ErrInvalidBatchSignature is the error returned if the batch signature is incorrect
	ErrInvalidBatchSignature = errors.New("invalid batch signature")
)

var _ swarm.Stamp = (*Stamp)(nil)

// Stamp represents a postage stamp as attached to a chunk.
type Stamp struct {
	batchID   []byte // postage batch ID
	index     []byte // index of the batch
	timestamp []byte // to signal order when assigning the indexes to multiple chunks
	sig       []byte // common r[32]s[32]v[1]-style 65 byte ECDSA signature of batchID|index|address by owner or grantee
}

// NewStamp constructs a new stamp from a given batch ID, index and signatures.
func NewStamp(batchID, index, timestamp, sig []byte) *Stamp { _ = "STUB: not implemented"; return nil }

// BatchID returns the batch ID of the stamp.
func (s *Stamp) BatchID() []byte {
	_ = "STUB: not implemented"

	// Index returns the within-batch index of the stamp.
	return nil
}

func (s *Stamp) Index() []byte {
	_ = "STUB: not implemented"

	// Sig returns the signature of the stamp by the user
	return nil
}

func (s *Stamp) Sig() []byte {
	_ = "STUB: not implemented"

	// Timestamp returns the timestamp of the stamp
	return nil
}

func (s *Stamp) Timestamp() []byte { _ = "STUB: not implemented"; return nil }

func (s *Stamp) Clone() swarm.Stamp { _ = "STUB: not implemented"; return *new(swarm.Stamp) }

// Hash returns the hash of the stamp.
func (s *Stamp) Hash() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarshalBinary gives the byte slice serialisation of a stamp:
// batchID[32]|index[8]|timestamp[8]|Signature[65].
func (s *Stamp) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary parses a serialised stamp into id and signature.
func (s *Stamp) UnmarshalBinary(buf []byte) error { _ = "STUB: not implemented"; return nil }

type stampJson struct {
	BatchID   []byte `json:"batchID"`
	Index     []byte `json:"index"`
	Timestamp []byte `json:"timestamp"`
	Sig       []byte `json:"sig"`
}

func (s *Stamp) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *Stamp) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// ToSignDigest creates a digest to represent the stamp which is to be signed by the owner.
func ToSignDigest(addr, batchId, index, timestamp []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ValidStampFn func(chunk swarm.Chunk) (swarm.Chunk, error)

// ValidStamp returns a stampvalidator function passed to protocols with chunk entrypoints.
func ValidStamp(batchStore Storer) ValidStampFn {
	_ = "STUB: not implemented"
	return *new(ValidStampFn)
}

// Valid checks the validity of the postage stamp; in particular:
// - authenticity - check batch is valid on the blockchain
// - authorisation - the batch owner is the stamp signer
// the validity  check is only meaningful in its association of a chunk
// this chunk address needs to be given as argument
func (s *Stamp) Valid(chunkAddr swarm.Address, ownerAddr []byte, depth, bucketDepth uint8, immutable bool) error {
	_ = "STUB: not implemented"
	return nil
}

// RecoverBatchOwner returns ethereum address that signed postage batch of supplied stamp.
func RecoverBatchOwner(chunkAddr swarm.Address, stamp swarm.Stamp) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
