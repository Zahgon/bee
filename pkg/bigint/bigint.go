// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bigint

import (
	"math/big"
)

type BigInt struct {
	*big.Int
}

func (i *BigInt) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *BigInt) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Wrap wraps big.Int pointer into BigInt struct.
func Wrap(i *big.Int) *BigInt { _ = "STUB: not implemented"; return nil }

// MarshalBinary implements encoding.BinaryMarshaler using Gob encoding.
// Panics if the underlying *big.Int is nil, as this indicates a programmer error.
func (i *BigInt) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary implements encoding.BinaryUnmarshaler using Gob decoding.
func (i *BigInt) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }
