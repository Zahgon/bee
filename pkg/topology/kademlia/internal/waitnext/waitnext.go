// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package metrics provides service for collecting various metrics about peers.
// It is intended to be used with the kademlia where the metrics are collected.
package waitnext

import (
	"sync"
	"time"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type next struct {
	tryAfter       time.Time
	failedAttempts int
}

type WaitNext struct {
	next map[string]*next
	sync.Mutex
}

func New() *WaitNext { _ = "STUB: not implemented"; return nil }

func (r *WaitNext) Set(addr swarm.Address, tryAfter time.Time, attempts int) {
	_ = "STUB: not implemented"
	return
}

func (r *WaitNext) SetTryAfter(addr swarm.Address, tryAfter time.Time) {
	_ = "STUB: not implemented"
	return
}

func (r *WaitNext) Waiting(addr swarm.Address) bool { _ = "STUB: not implemented"; return false }

func (r *WaitNext) Attempts(addr swarm.Address) int { _ = "STUB: not implemented"; return 0 }

func (r *WaitNext) Remove(addr swarm.Address) { _ = "STUB: not implemented"; return }
