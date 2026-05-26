// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postage

import (
	"math/big"
)

// Batch represents a postage batch, a payment on the blockchain.
type Batch struct {
	ID          []byte   // batch ID
	Value       *big.Int // normalised balance of the batch
	Start       uint64   // block number the batch was created
	Owner       []byte   // owner's ethereum address
	Depth       uint8    // batch depth, i.e., size = 2^{depth}
	BucketDepth uint8    // the depth of neighbourhoods t
	Immutable   bool     // if the batch allows adding new capacity (dilution)
}

// MarshalBinary implements BinaryMarshaller. It will attempt to serialize the
// postage batch to a byte slice.
// serialised as ID(32)|big endian value(32)|start block(8)|owner addr(20)|BucketDepth(1)|depth(1)|immutable(1)
func (b *Batch) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary implements BinaryUnmarshaller. It will attempt deserialize
// the given byte slice into the batch.
func (b *Batch) UnmarshalBinary(buf []byte) error { _ = "STUB: not implemented"; return nil }
