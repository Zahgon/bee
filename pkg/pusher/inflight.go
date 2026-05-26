// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pusher

import (
	"sync"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type inflight struct {
	mtx      sync.Mutex
	inflight map[[64]byte]struct{}
}

func newInflight() *inflight { _ = "STUB: not implemented"; return nil }

func (i *inflight) delete(idAddress swarm.Address, batchID []byte) {
	_ = "STUB: not implemented"
	return
}

func (i *inflight) set(idAddress swarm.Address, batchID []byte) bool {
	_ = "STUB: not implemented"
	return false
}

type attempts struct {
	mtx        sync.Mutex
	retryCount int
	attempts   map[string]int
}

// try to log a chunk sync attempt. returns false when
// maximum amount of attempts have been reached.
func (a *attempts) try(idAddress swarm.Address) bool { _ = "STUB: not implemented"; return false }

func (a *attempts) delete(idAddress swarm.Address) { _ = "STUB: not implemented"; return }
