// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package swarm

import (
	"testing"
)

// RandAddress generates a random address.
func RandAddress(tb testing.TB) Address { _ = "STUB: not implemented"; return *new(Address) }

// RandAddressAt generates a random address at proximity order prox relative to address.
func RandAddressAt(tb testing.TB, self Address, prox int) Address {
	_ = "STUB: not implemented"
	return *new(Address)
}

// RandAddresses generates slice with a random address.
func RandAddresses(tb testing.TB, count int) []Address { _ = "STUB: not implemented"; return nil }

// RandBatchID generates a random BatchID.
func RandBatchID(tb testing.TB) []byte { _ = "STUB: not implemented"; return nil }
