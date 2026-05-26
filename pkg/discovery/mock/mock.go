// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type Discovery struct {
	mtx           sync.Mutex
	ctr           int // how many ops
	records       map[string][]swarm.Address
	broadcastFunc func(context.Context, swarm.Address, ...swarm.Address) error
}

type Option interface {
	apply(*Discovery)
}
type optionFunc func(*Discovery)

func (f optionFunc) apply(r *Discovery) { _ = "STUB: not implemented"; return }

func WithBroadcastPeers(f func(context.Context, swarm.Address, ...swarm.Address) error) optionFunc {
	_ = "STUB: not implemented"
	return *new(optionFunc)
}

func NewDiscovery(opts ...Option) *Discovery { _ = "STUB: not implemented"; return nil }

func (d *Discovery) BroadcastPeers(ctx context.Context, addressee swarm.Address, peers ...swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Discovery) Broadcasts() int { _ = "STUB: not implemented"; return 0 }

func (d *Discovery) AddresseeRecords(addressee swarm.Address) (peers []swarm.Address, exists bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (d *Discovery) Reset() { _ = "STUB: not implemented"; return }
