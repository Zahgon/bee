// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// the code below implements the integration of dispersed replicas in chunk upload.
// using storage.Putter interface.
package replicas

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/file/redundancy"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// putter is the private implementation of the public storage.Putter interface
// putter extends the original putter to a concurrent multiputter
type putter struct {
	putter storage.Putter
	rLevel redundancy.Level
}

// NewPutter is the putter constructor
func NewPutter(p storage.Putter, rLevel redundancy.Level) storage.Putter {
	_ = "STUB: not implemented"
	return *new(storage.Putter)
}

// Put makes the getter satisfy the storage.Getter interface
func (p *putter) Put(ctx context.Context, ch swarm.Chunk) (err error) {
	_ = "STUB: not implemented"
	return nil
}
