// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libp2p

import (
	"context"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/libp2p/go-libp2p/core/network"
	libp2ppeer "github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

type peerRegistry struct {
	overlayToPeerID     map[string]libp2ppeer.ID                    // map overlay address to underlay peer id
	overlays            map[libp2ppeer.ID]swarm.Address             // map underlay peer id to overlay address
	full                map[libp2ppeer.ID]bool                      // map to track whether a node is full or light node (true=full)
	bee260Compatibility map[libp2ppeer.ID]bool                      // map to track bee260 backward compatibility
	connections         map[libp2ppeer.ID]map[network.Conn]struct{} // list of connections for safe removal on Disconnect notification
	streams             map[libp2ppeer.ID]map[network.Stream]context.CancelFunc
	mu                  sync.RWMutex

	//nolint:misspell
	disconnecter     disconnecter // peerRegistry notifies libp2p on peer disconnection
	network.Notifiee              // peerRegistry can be the receiver for network.Notify
}

type disconnecter interface {
	disconnected(swarm.Address)
}

func newPeerRegistry() *peerRegistry { _ = "STUB: not implemented"; return nil }

func (r *peerRegistry) Exists(overlay swarm.Address) (found bool) {
	_ = "STUB: not implemented"
	return false
}

// Disconnected removes the peer from registry in disconnect.
// peerRegistry has to be set by network.Network.Notify().
func (r *peerRegistry) Disconnected(_ network.Network, c network.Conn) {
	_ = "STUB: not implemented"
	return
}

// remove only the related connection,
// not eventusally newly created one for the same peer

// if there are multiple libp2p connections, consider the node disconnected only when the last connection is disconnected

func (r *peerRegistry) addStream(peerID libp2ppeer.ID, stream network.Stream, cancel context.CancelFunc) {
	_ = "STUB: not implemented"
	return
}

// it is possible that an addStream will be called after a disconnect

func (r *peerRegistry) removeStream(peerID libp2ppeer.ID, stream network.Stream) {
	_ = "STUB: not implemented"
	return
}

func (r *peerRegistry) peers() []p2p.Peer { _ = "STUB: not implemented"; return nil }

func (r *peerRegistry) addIfNotExists(c network.Conn, overlay swarm.Address, full bool) (exists bool) {
	_ = "STUB: not implemented"
	return false
}

// the connection is added even if the peer already exists in peer registry
// this is solving a case of multiple underlying libp2p connections for the same peer

func (r *peerRegistry) peerID(overlay swarm.Address) (peerID libp2ppeer.ID, found bool) {
	_ = "STUB: not implemented"
	return *new(libp2ppeer.ID), false
}

func (r *peerRegistry) overlay(peerID libp2ppeer.ID) (swarm.Address, bool) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), false
}

func (r *peerRegistry) fullnode(peerID libp2ppeer.ID) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (r *peerRegistry) bee260(peerID libp2ppeer.ID) (compat, found bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (r *peerRegistry) setBee260(peerID libp2ppeer.ID, compat bool) {
	_ = "STUB: not implemented"
	return
}

func (r *peerRegistry) isConnected(peerID libp2ppeer.ID, remoteAddr ma.Multiaddr) (swarm.Address, bool) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), false
}

// check connection remote address

// we ARE connected to the peer on expected address

func (r *peerRegistry) remove(overlay swarm.Address) (found, full bool, peerID libp2ppeer.ID) {
	_ = "STUB: not implemented"
	return false, false, *new(libp2ppeer.ID)
}

func (r *peerRegistry) setDisconnecter(d disconnecter) { _ = "STUB: not implemented"; return }
