// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blocker

import (
	"sync"
	"time"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"go.uber.org/atomic"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "blocker"

// sequencerResolution represents monotonic sequencer resolution.
// It must be in the time.Duration base form without a multiplier.
var sequencerResolution = time.Second

type peer struct {
	blockAfter uint64
	address    swarm.Address
}

type Blocker struct {
	sequence          atomic.Uint64 // Monotonic clock.
	mu                sync.Mutex
	blocklister       p2p.Blocklister
	flagTimeout       time.Duration // how long before blocking a flagged peer
	blockDuration     time.Duration // how long to blocklist a bad peer
	peers             map[string]*peer
	logger            log.Logger
	wakeupCh          chan struct{}
	quit              chan struct{}
	closeWg           sync.WaitGroup
	blocklistCallback func(swarm.Address)
}

func New(blocklister p2p.Blocklister, flagTimeout, blockDuration, wakeUpTime time.Duration, callback func(swarm.Address), logger log.Logger) *Blocker {
	_ = "STUB: not implemented"
	return nil
}

func (b *Blocker) block() { _ = "STUB: not implemented"; return }

func (b *Blocker) Flag(addr swarm.Address) { _ = "STUB: not implemented"; return }

func (b *Blocker) Unflag(addr swarm.Address) { _ = "STUB: not implemented"; return }

func (b *Blocker) PruneUnseen(seen []swarm.Address) { _ = "STUB: not implemented"; return }

// Close will exit the worker loop.
// must be called only once.
func (b *Blocker) Close() error { _ = "STUB: not implemented"; return nil }
