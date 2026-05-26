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

type AddrTuple struct {
	Addr swarm.Address // the peer address
	PO   uint8         // the po
}

func WithEachPeerRevCalls(addrs ...AddrTuple) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithDepth(d uint8) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDepthCalls(d ...uint8) Option { _ = "STUB: not implemented"; return *new(Option) }

type Mock struct {
	mtx          sync.Mutex
	peers        []swarm.Address
	eachPeerRev  []AddrTuple
	depth        uint8
	depthReplies []uint8
	depthCalls   int
	trigs        []chan struct{}
	trigMtx      sync.Mutex
}

func NewMockKademlia(o ...Option) *Mock { _ = "STUB: not implemented"; return nil }

// AddPeers is called when a peers are added to the topology backlog
// for further processing by connectivity strategy.
func (m *Mock) AddPeers(addr ...swarm.Address) { _ = "STUB: not implemented"; return }

// TODO: Implement

func (m *Mock) ClosestPeer(addr swarm.Address, _ bool, _ topology.Select, skipPeers ...swarm.Address) (peerAddr swarm.Address, err error) {
	_ = "STUB: not implemented"
	return *
	// TODO: Implement
	new(swarm.Address), nil
}

func (m *Mock) EachNeighbor(topology.EachPeerFunc) error { _ = "STUB: not implemented"; return nil }

// TODO: Implement

func (m *Mock) EachNeighborRev(topology.EachPeerFunc) error { _ = "STUB: not implemented"; return nil }

// TODO: Implement

func (m *Mock) UpdatePeerHealth(swarm.Address, bool, time.Duration) {
	_ = "STUB: not implemented"
	return
	// TODO: Implement
}

// PeerIterator iterates from closest bin to farthest
func (m *Mock) SetStorageRadius(uint8) { _ = "STUB: not implemented"; return }

func (m *Mock) AddRevPeers(addrs ...AddrTuple) { _ = "STUB: not implemented"; return }

// EachConnectedPeer iterates from closest bin to farthest
func (m *Mock) EachConnectedPeer(f topology.EachPeerFunc, _ topology.Select) error {
	_ = "STUB: not implemented"
	return nil
}

// EachConnectedPeerRev iterates from farthest bin to closest
func (m *Mock) EachConnectedPeerRev(f topology.EachPeerFunc, _ topology.Select) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mock) IsReachable() bool { _ = "STUB: not implemented"; return false }

func (m *Mock) NeighborhoodDepth() uint8 { _ = "STUB: not implemented"; return 0 }

// Connected is called when a peer dials in.
func (m *Mock) Connected(_ context.Context, peer p2p.Peer, _ bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Disconnected is called when a peer disconnects.
func (m *Mock) Disconnected(peer p2p.Peer) { _ = "STUB: not implemented"; return }

func (m *Mock) Announce(_ context.Context, _ swarm.Address, _ bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mock) AnnounceTo(_ context.Context, _, _ swarm.Address, _ bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mock) SubscribeTopologyChange() (c <-chan struct{}, unsubscribe func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Mock) Trigger() { _ = "STUB: not implemented"; return }

func (m *Mock) ResetPeers() { _ = "STUB: not implemented"; return }

func (d *Mock) Halt()        { _ = "STUB: not implemented"; return }
func (m *Mock) Close() error { _ = "STUB: not implemented"; return nil }

func (m *Mock) Snapshot() *topology.KadParams { _ = "STUB: not implemented"; return nil }

// TODO: Implement

type Option interface {
	apply(*Mock)
}
type optionFunc func(*Mock)

func (f optionFunc) apply(r *Mock) { _ = "STUB: not implemented"; return }
