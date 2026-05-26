// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/pushsync"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type mock struct {
	sendChunk func(ctx context.Context, chunk swarm.Chunk) (*pushsync.Receipt, error)
}

func New(sendChunk func(ctx context.Context, chunk swarm.Chunk) (*pushsync.Receipt, error)) pushsync.PushSyncer {
	_ = "STUB: not implemented"
	return *new(pushsync.PushSyncer)
}

func (s *mock) PushChunkToClosest(ctx context.Context, chunk swarm.Chunk) (*pushsync.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *mock) Close() error { _ = "STUB: not implemented"; return nil }
