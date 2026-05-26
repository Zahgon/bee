// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
)

type statusSnapshotResponse struct {
	Overlay                 string  `json:"overlay"`
	Proximity               uint    `json:"proximity"`
	BeeMode                 string  `json:"beeMode"`
	ReserveSize             uint64  `json:"reserveSize"`
	ReserveSizeWithinRadius uint64  `json:"reserveSizeWithinRadius"`
	PullsyncRate            float64 `json:"pullsyncRate"`
	StorageRadius           uint8   `json:"storageRadius"`
	ConnectedPeers          uint64  `json:"connectedPeers"`
	NeighborhoodSize        uint64  `json:"neighborhoodSize"`
	RequestFailed           bool    `json:"requestFailed,omitempty"`
	BatchCommitment         uint64  `json:"batchCommitment"`
	IsReachable             bool    `json:"isReachable"`
	LastSyncedBlock         uint64  `json:"lastSyncedBlock"`
	CommittedDepth          uint8   `json:"committedDepth"`
	IsWarmingUp             bool    `json:"isWarmingUp"`
}

type statusResponse struct {
	Snapshots []statusSnapshotResponse `json:"snapshots"`
}

type statusNeighborhoodResponse struct {
	Neighborhood            string `json:"neighborhood"`
	ReserveSizeWithinRadius int    `json:"reserveSizeWithinRadius"`
	Proximity               uint8  `json:"proximity"`
}

type neighborhoodsResponse struct {
	Neighborhoods []statusNeighborhoodResponse `json:"neighborhoods"`
}

// statusAccessHandler is a middleware that limits the number of simultaneous
// status requests.
func (s *Service) statusAccessHandler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// statusGetHandler returns the current node status.
func (s *Service) statusGetHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

// statusGetPeersHandler returns the status of currently connected peers.
func (s *Service) statusGetPeersHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// mu protects snapshots.

// statusGetHandler returns the current node status.
func (s *Service) statusGetNeighborhoods(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
