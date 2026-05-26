// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package skippeers

import (
	"math"
	"sync"
	"time"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const maxDuration time.Duration = math.MaxInt64

type List struct {
	mtx sync.Mutex

	durC chan time.Duration
	quit chan struct{}
	// key is chunk address, value is map of peer address to expiration
	skip map[string]map[string]int64

	wg sync.WaitGroup
}

func NewList(workerWakeUpDur time.Duration) *List { _ = "STUB: not implemented"; return nil }

func (l *List) worker(d time.Duration) { _ = "STUB: not implemented"; return }

func (l *List) Forever(chunk, peer swarm.Address) { _ = "STUB: not implemented"; return }

func (l *List) Add(chunk, peer swarm.Address, expire time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (l *List) ChunkPeers(ch swarm.Address) (peers []swarm.Address) {
	_ = "STUB: not implemented"
	return nil
}

func (l *List) PruneExpiresAfter(ch swarm.Address, d time.Duration) int {
	_ = "STUB: not implemented"
	return 0
}

func (l *List) prune() { _ = "STUB: not implemented"; return }

// Must be called under lock
func (l *List) pruneChunk(ch string, now int64) int { _ = "STUB: not implemented"; return 0 }

func (l *List) Close() error { _ = "STUB: not implemented"; return nil }
