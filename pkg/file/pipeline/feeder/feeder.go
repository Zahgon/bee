// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feeder

import (
	"github.com/ethersphere/bee/v2/pkg/file/pipeline"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const span = swarm.SpanSize

type chunkFeeder struct {
	size      int
	next      pipeline.ChainWriter
	buffer    []byte
	bufferIdx int
	wrote     int64
}

// NewChunkFeederWriter creates a new chunkFeeder that allows for partial
// writes into the pipeline. Any pending data in the buffer is flushed to
// subsequent writers when Sum() is called.
func NewChunkFeederWriter(size int, next pipeline.ChainWriter) pipeline.Interface {
	_ = "STUB: not implemented"
	return *new(pipeline.Interface)
}

// Write writes data to the chunk feeder. It returns the number of bytes written
// to the feeder. The number of bytes written does not necessarily reflect how many
// bytes were actually flushed to subsequent writers, since the feeder is buffered
// and works in chunk-size quantiles.
func (f *chunkFeeder) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	// data length
	return 0, nil
}

// written

// write the data into the buffer and return

// if we are here it means we have to do at least one write

// span of current write

// copy from existing buffer to this one

// don't account what was already in the buffer when returning
// number of written bytes

// if we can't fill a whole write, buffer the rest and return

// fill stuff up from the incoming write

// Sum flushes any pending data to subsequent writers and returns
// the cryptographic root-hash representing the data written to
// the feeder.
func (f *chunkFeeder) Sum() ([]byte, error) {
	_ = "STUB: not implemented"
	// flush existing data in the buffer
	return nil, nil
}

// this is an empty file, we should write the span of
// an empty file (0).
