// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"context"
	"io"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// simpleReadCloser wraps a byte slice in a io.ReadCloser implementation.
type simpleReadCloser struct {
	buffer io.Reader
	closed bool
}

// NewSimpleReadCloser creates a new simpleReadCloser.
func NewSimpleReadCloser(buffer []byte) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

// Read implements io.Reader.
func (s *simpleReadCloser) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close implements io.Closer.
func (s *simpleReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

// JoinReadAll reads all output from the provided Joiner.
func JoinReadAll(ctx context.Context, j Joiner, outFile io.Writer) (int64, error) {
	_ = "STUB: not implemented"

	// join, rinse, repeat until done
	return 0, nil
}

// SplitWriteAll writes all input from provided reader to the provided splitter
func SplitWriteAll(ctx context.Context, s Splitter, r io.Reader, l int64, toEncrypt bool) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

type Loader interface {
	// Load a reference in byte slice representation and return all content associated with the reference.
	Load(context.Context, []byte) ([]byte, error)
}

type Saver interface {
	// Save an arbitrary byte slice and return the reference byte slice representation.
	Save(context.Context, []byte) ([]byte, error)
}

type LoadSaver interface {
	Loader
	Saver
}
