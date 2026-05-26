// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lightnode

import (
	"context"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/topology"
	"github.com/ethersphere/bee/v2/pkg/topology/pslice"
)

type Container struct {
	base              swarm.Address
	peerMu            sync.Mutex // peerMu guards connectedPeers and disconnectedPeers.
	connectedPeers    *pslice.PSlice
	disconnectedPeers *pslice.PSlice
	metrics           metrics
}

func NewContainer(base swarm.Address) *Container { _ = "STUB: not implemented"; return nil }

func (c *Container) Connected(ctx context.Context, peer p2p.Peer) {
	_ = "STUB: not implemented"
	return
}

func (c *Container) Disconnected(peer p2p.Peer) { _ = "STUB: not implemented"; return }

func (c *Container) Count() int { _ = "STUB: not implemented"; return 0 }

func (c *Container) RandomPeer(not swarm.Address) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func (c *Container) EachPeer(pf topology.EachPeerFunc) error { _ = "STUB: not implemented"; return nil }

func (c *Container) PeerInfo() topology.BinInfo {
	_ = "STUB: not implemented"
	return *new(topology.BinInfo)
}

func peersInfo(s *pslice.PSlice) []*topology.PeerInfo { _ = "STUB: not implemented"; return nil }
