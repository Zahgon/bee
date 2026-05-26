// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type peerConnectResponse struct {
	Address string `json:"address"`
}

func (s *Service) peerConnectHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) peerDisconnectHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Peer holds information about a Peer.
type Peer struct {
	Address  swarm.Address `json:"address"`
	FullNode bool          `json:"fullNode"`
}

type BlockListedPeer struct {
	Peer
	Reason   string `json:"reason"`
	Duration int    `json:"duration"`
}

type peersResponse struct {
	Peers []Peer `json:"peers"`
}

type blockListedPeersResponse struct {
	Peers []BlockListedPeer `json:"peers"`
}

func (s *Service) peersHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) blocklistedPeersHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func mapPeers(peers []p2p.Peer) (out []Peer) { _ = "STUB: not implemented"; return nil }

func mapBlockListedPeers(peers []p2p.BlockListedPeer) []BlockListedPeer {
	_ = "STUB: not implemented"
	return nil
}
