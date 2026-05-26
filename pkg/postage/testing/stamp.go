// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

import (
	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const signatureSize = 65

// MustNewSignature will create a new random signature (65 byte slice). Panics on errors.
func MustNewSignature() []byte { _ = "STUB: not implemented"; return nil }

// MustNewValidSignature will create a new valid signature. Panics on errors.
func MustNewValidSignature(signer crypto.Signer, addr swarm.Address, id, index, timestamp []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// MustNewStamp will generate a invalid postage stamp with random data. Panics on errors.
func MustNewStamp() *postage.Stamp { _ = "STUB: not implemented"; return nil }

// MustNewValidStamp will generate a valid postage stamp with random data. Panics on errors.
func MustNewValidStamp(signer crypto.Signer, addr swarm.Address) *postage.Stamp {
	_ = "STUB: not implemented"
	return nil
}

// MustNewBatchStamp will generate a postage stamp with the provided batch ID and assign
// random data to other fields. Panics on error
func MustNewBatchStamp(batch []byte) *postage.Stamp { _ = "STUB: not implemented"; return nil }

// MustNewBatchStamp will generate a postage stamp with the provided batch ID and assign
// random data to other fields. Panics on error
func MustNewFields(batch []byte, index, ts uint64) *postage.Stamp {
	_ = "STUB: not implemented"
	return nil
}

// MustNewStampWithTimestamp will generate a postage stamp with provided timestamp and
// random data for other fields. Panics on errors.
func MustNewStampWithTimestamp(ts uint64) *postage.Stamp { _ = "STUB: not implemented"; return nil }
