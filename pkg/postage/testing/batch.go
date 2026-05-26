// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

import (
	"math/big"
	"testing"

	"github.com/ethersphere/bee/v2/pkg/postage"
)

const (
	defaultBucketDepth = 12
	defaultDepth       = 16
)

// BatchOption is an optional parameter for NewBatch
type BatchOption func(c *postage.Batch)

// MustNewID will generate a new random ID (32 byte slice). Panics on errors.
func MustNewID() []byte { _ = "STUB: not implemented"; return nil }

// MustNewAddress will generate a new random address (20 byte slice). Panics on
// errors.
func MustNewAddress() []byte { _ = "STUB: not implemented"; return nil }

// NewBigInt will generate a new random big int (uint64 base value).
func NewBigInt() *big.Int { _ = "STUB: not implemented"; return nil }

// MustNewBatch will create a new test batch. Fields that are not supplied will
// be filled with random data. Panics on errors.
func MustNewBatch(opts ...BatchOption) *postage.Batch { _ = "STUB: not implemented"; return nil }

// WithOwner will set the batch owner on a randomized batch.
func WithOwner(owner []byte) BatchOption { _ = "STUB: not implemented"; return *new(BatchOption) }

// WithValue will set the batch Value to the given value.
func WithValue(value int64) BatchOption { _ = "STUB: not implemented"; return *new(BatchOption) }

// WithDepth will set the batch Depth to the given depth.
func WithDepth(depth uint8) BatchOption { _ = "STUB: not implemented"; return *new(BatchOption) }

// WithStart will set the batch Start to the given start.
func WithStart(start uint64) BatchOption { _ = "STUB: not implemented"; return *new(BatchOption) }

// CompareBatches is a testing helper that compares two batches and fails the
// test if all fields are not equal.
// Fails on first different value and prints the comparison.
func CompareBatches(t *testing.T, want, got *postage.Batch) { _ = "STUB: not implemented"; return }
