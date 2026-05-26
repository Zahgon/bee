// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"io"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const (
	maxBufferSize = swarm.ChunkSize * 2
)

// ChunkPipe ensures that only the last read is smaller than the chunk size,
// regardless of size of individual writes.
type ChunkPipe struct {
	io.ReadCloser
	writer io.WriteCloser
	data   []byte
	cursor int
}

// Creates a new ChunkPipe
func NewChunkPipe() io.ReadWriteCloser { _ = "STUB: not implemented"; return *new(io.ReadWriteCloser) }

// Read implements io.Reader
func (c *ChunkPipe) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Writer implements io.Writer
func (c *ChunkPipe) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// NOTE: the Write method contract requires all sent data to be
// written before returning (without error)

// Close implements io.Closer
func (c *ChunkPipe) Close() error { _ = "STUB: not implemented"; return nil }
