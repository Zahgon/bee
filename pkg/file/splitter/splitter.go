// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package splitter provides implementations of the file.Splitter interface
package splitter

import (
	"context"
	"io"

	"github.com/ethersphere/bee/v2/pkg/file"
	storage "github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// simpleSplitter wraps a non-optimized implementation of file.Splitter
type simpleSplitter struct {
	putter storage.Putter
}

// NewSimpleSplitter creates a new SimpleSplitter
func NewSimpleSplitter(storePutter storage.Putter) file.Splitter {
	_ = "STUB: not implemented"
	return *new(file.Splitter)
}

// Split implements the file.Splitter interface
//
// It uses a non-optimized internal component that blocks when performing
// multiple levels of hashing when building the file hash tree.
//
// It returns the Swarmhash of the data.
func (s *simpleSplitter) Split(ctx context.Context, r io.ReadCloser, dataLength int64, toEncrypt bool) (addr swarm.Address, err error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}
