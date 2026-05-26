// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/pullsync"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var _ pullsync.Interface = (*PullSyncMock)(nil)

func WithSyncError(err error) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCursors(v []uint64, e uint64) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReplies(replies ...SyncReply) Option { _ = "STUB: not implemented"; return *new(Option) }

func toID(a swarm.Address, bin uint8, start uint64) string { _ = "STUB: not implemented"; return "" }

type SyncReply struct {
	Peer    swarm.Address
	Bin     uint8
	Start   uint64
	Topmost uint64
	Count   int
}

type PullSyncMock struct {
	mtx             sync.Mutex
	syncCalls       []SyncReply
	syncErr         error
	cursors         []uint64
	epoch           uint64
	getCursorsPeers []swarm.Address
	replies         map[string][]SyncReply

	quit chan struct{}
}

func NewPullSync(opts ...Option) *PullSyncMock { _ = "STUB: not implemented"; return nil }

func (p *PullSyncMock) Sync(ctx context.Context, peer swarm.Address, bin uint8, start uint64) (topmost uint64, count int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (p *PullSyncMock) GetCursors(_ context.Context, peer swarm.Address) ([]uint64, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (p *PullSyncMock) ResetCalls(peer swarm.Address) { _ = "STUB: not implemented"; return }

func (p *PullSyncMock) SyncCalls(peer swarm.Address) (res []SyncReply) {
	_ = "STUB: not implemented"
	return nil
}

func (p *PullSyncMock) CursorsCalls(peer swarm.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *PullSyncMock) SetEpoch(epoch uint64) { _ = "STUB: not implemented"; return }

type Option interface {
	apply(*PullSyncMock)
}
type optionFunc func(*PullSyncMock)

func (f optionFunc) apply(r *PullSyncMock) { _ = "STUB: not implemented"; return }
