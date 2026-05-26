// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"sync"
	"time"

	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/topology"
)

type mock struct {
	peers           []swarm.Address
	depth           uint8
	closestPeer     swarm.Address
	closestPeerErr  error
	peersErr        error
	addPeersErr     error
	isWithinFunc    func(c swarm.Address) bool
	marshalJSONFunc func() ([]byte, error)
	mtx             sync.Mutex
	health          map[string]bool
	lastSelect      topology.Select
	selectRecorder  *topology.Select
}

var _ topology.Driver = (*mock)(nil)

func WithPeers(peers ...swarm.Address) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithAddPeersErr(err error) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithNeighborhoodDepth(dd uint8) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithClosestPeer(addr swarm.Address) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithClosestPeerErr(err error) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMarshalJSONFunc(f func() ([]byte, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithIsWithinFunc(f func(swarm.Address) bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSelectRecorder records the topology.Select most recently passed to
// EachConnectedPeer or EachConnectedPeerRev into out. Tests use this to assert
// callers opt in to filtering flags (e.g. IncludeBootnodes).
func WithSelectRecorder(out *topology.Select) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func NewTopologyDriver(opts ...Option) *mock { _ = "STUB: not implemented"; return nil }

func (d *mock) AddPeers(addrs ...swarm.Address) { _ = "STUB: not implemented"; return }

func (d *mock) Connected(ctx context.Context, peer p2p.Peer, _ bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *mock) Disconnected(peer p2p.Peer) { _ = "STUB: not implemented"; return }

func (d *mock) Announce(_ context.Context, _ swarm.Address, _ bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *mock) AnnounceTo(_ context.Context, _, _ swarm.Address, _ bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *mock) UpdatePeerHealth(peer swarm.Address, health bool, pingDur time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (d *mock) PeersHealth() map[string]bool { _ = "STUB: not implemented"; return nil }

func (d *mock) Peers() []swarm.Address { _ = "STUB: not implemented"; return nil }

func (d *mock) ClosestPeer(addr swarm.Address, wantSelf bool, _ topology.Select, skipPeers ...swarm.Address) (peerAddr swarm.Address, err error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func (m *mock) IsReachable() bool { _ = "STUB: not implemented"; return false }

func (d *mock) SubscribeTopologyChange() (c <-chan struct{}, unsubscribe func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *mock) NeighborhoodDepth() uint8 { _ = "STUB: not implemented"; return 0 }

func (m *mock) SetStorageRadius(uint8) {
	_ = "STUB: not implemented"

	// EachConnectedPeer implements topology.PeerIterator interface.
	return
}

func (d *mock) EachConnectedPeer(f topology.EachPeerFunc, s topology.Select) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// EachConnectedPeerRev implements topology.PeerIterator interface.
func (d *mock) EachConnectedPeerRev(f topology.EachPeerFunc, s topology.Select) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// LastSelect returns the topology.Select most recently passed to
// EachConnectedPeer or EachConnectedPeerRev. Intended for tests that assert
// callers opt in to filtering flags (e.g. IncludeBootnodes).
func (d *mock) LastSelect() topology.Select {
	_ = "STUB: not implemented"
	return *new(topology.Select)
}

func (d *mock) Snapshot() *topology.KadParams { _ = "STUB: not implemented"; return nil }

func (d *mock) Halt()        { _ = "STUB: not implemented"; return }
func (d *mock) Close() error { _ = "STUB: not implemented"; return nil }

type Option interface {
	apply(*mock)
}

type optionFunc func(*mock)

func (f optionFunc) apply(r *mock) { _ = "STUB: not implemented"; return }
