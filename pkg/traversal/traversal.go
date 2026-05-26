// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package traversal provides abstraction and implementation
// needed to traverse all chunks below a given root hash.
// It tries to parse all manifests and collections in its
// attempt to log all chunk addresses on the way.
package traversal

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/file/redundancy"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// Traverser represents service which traverse through address dependent chunks.
type Traverser interface {
	// Traverse iterates through each address related to the supplied one, if possible.
	Traverse(context.Context, swarm.Address, swarm.AddressIterFunc) error
}

// New constructs for a new Traverser.
func New(getter storage.Getter, putter storage.Putter, rLevel redundancy.Level) Traverser {
	_ = "STUB: not implemented"
	return *new(Traverser)
}

// service is implementation of Traverser using storage.Storer as its storage.
type service struct {
	getter storage.Getter
	putter storage.Putter
	rLevel redundancy.Level
}

// Traverse implements Traverser.Traverse method.
func (s *service) Traverse(ctx context.Context, addr swarm.Address, iterFn swarm.AddressIterFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// skip SOC check for encrypted references

// if this is a SOC, the traversal will be just be the single chunk

// Heuristic determination if the reference represents a manifest reference.
// The assumption is that if the root chunk span is less than or equal to swarm.ChunkSize,
// then the reference is likely a manifest reference. This is because manifest holds metadata
// that points to the actual data file, and this metadata is assumed to be small - Less than or equal to swarm.ChunkSize.

// Based on the returned errors we conclude that it might
// not be a manifest, so we try non-manifest processing.

// Non-manifest processing.
