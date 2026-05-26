// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

import (
	"testing"

	"github.com/ethersphere/bee/v2/pkg/postage"
)

// NewChainState will create a new ChainState with random values.
func NewChainState() *postage.ChainState { _ = "STUB: not implemented"; return nil }

// CompareChainState is a test helper that compares two ChainStates and fails
// the test if they are not exactly equal.
// Fails on first difference and returns a descriptive comparison.
func CompareChainState(t *testing.T, want, got *postage.ChainState) {
	_ = "STUB: not implemented"
	return
}
