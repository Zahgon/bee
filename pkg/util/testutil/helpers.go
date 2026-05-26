// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testutil

import (
	"io"
	"testing"

	"github.com/ethersphere/bee/v2/pkg/log"
)

// RandBytes returns bytes slice of specified size filled with random values.
func RandBytes(tb testing.TB, size int) []byte { _ = "STUB: not implemented"; return nil }

// RandBytesWithSeed returns bytes slice of specified size filled with random values generated using seed.
func RandBytesWithSeed(tb testing.TB, size int, seed int64) []byte {
	_ = "STUB: not implemented"
	return nil
}

// CleanupCloser adds Cleanup function to Test which will close supplied Closers.
func CleanupCloser(t *testing.T, closers ...io.Closer) { _ = "STUB: not implemented"; return }

// NewLogger returns a new log.Logger that uses t.Log method
// as the log sink. It is particularly useful for debugging tests.
func NewLogger(t *testing.T) log.Logger { _ = "STUB: not implemented"; return *new(log.Logger) }
