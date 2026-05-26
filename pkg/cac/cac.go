// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cac

import (
	"fmt"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var (
	ErrChunkSpanShort = fmt.Errorf("chunk span must have exactly length of %d", swarm.SpanSize)
	ErrChunkDataLarge = fmt.Errorf("chunk data exceeds maximum allowed length")
)

// New creates a new content address chunk by initializing a span and appending the data to it.
func New(data []byte) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// NewWithDataSpan creates a new chunk assuming that the span precedes the actual data.
func NewWithDataSpan(data []byte) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// validateDataLength validates if data length (without span) is correct.
func validateDataLength(dataLength int) error {
	_ = "STUB: not implemented"
	// dataLength could be negative when span size is subtracted
	return nil
}

// newWithSpan creates a new chunk prepending the given span to the data.
func newWithSpan(data, span []byte) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// Valid checks whether the given chunk is a valid content-addressed chunk.
func Valid(c swarm.Chunk) bool { _ = "STUB: not implemented"; return false }

func DoHash(data, span []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
