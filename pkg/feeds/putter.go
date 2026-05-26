// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feeds

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/crypto"
	storage "github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// Updater is the generic interface f
type Updater interface {
	Update(ctx context.Context, at int64, payload []byte) error
	Feed() *Feed
}

// Putter encapsulates a chunk store putter and a Feed to store feed updates
type Putter struct {
	putter storage.Putter
	signer crypto.Signer
	*Feed
}

// NewPutter constructs a feed Putter
func NewPutter(putter storage.Putter, signer crypto.Signer, topic []byte) (*Putter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put pushes an update to the feed through the chunk stores
func (u *Putter) Put(ctx context.Context, i Index, payload []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func toChunk(payload []byte) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}
