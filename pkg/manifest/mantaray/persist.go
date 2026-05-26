// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mantaray

import (
	"context"
	"errors"
)

var (
	// ErrNoSaver saver interface not given
	ErrNoSaver = errors.New("Node is not persisted but no saver")
	// ErrNoLoader saver interface not given
	ErrNoLoader = errors.New("Node is reference but no loader")
)

// Loader defines a generic interface to retrieve nodes
// from a persistent storage
// for read only operations only
type Loader interface {
	Load(ctx context.Context, reference []byte) (data []byte, err error)
}

// Saver defines a generic interface to persist nodes
// for write operations
type Saver interface {
	Save(ctx context.Context, data []byte) (reference []byte, err error)
}

// LoadSaver is a composite interface of Loader and Saver
// it is meant to be implemented as thin wrappers around persistent storage like Swarm
type LoadSaver interface {
	Loader
	Saver
}

func (n *Node) load(ctx context.Context, l Loader) error { _ = "STUB: not implemented"; return nil }

// Save persists a trie recursively  traversing the nodes
func (n *Node) Save(ctx context.Context, s Saver) error { _ = "STUB: not implemented"; return nil }

func (n *Node) save(ctx context.Context, s Saver) error { _ = "STUB: not implemented"; return nil }
