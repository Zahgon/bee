// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package salud monitors the connected peers, calculates certain thresholds, and marks peers as unhealthy that
// fall short of the thresholds to maintain network salud (health).
package salud

import (
	"context"
	"sync"
	"time"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/stabilization"
	"github.com/ethersphere/bee/v2/pkg/status"
	"github.com/ethersphere/bee/v2/pkg/storer"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/topology"
	"go.uber.org/atomic"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "salud"

const (
	requestTimeout         = time.Second * 10
	initialBackoffDelay    = 10 * time.Second
	maxBackoffDelay        = 5 * time.Minute
	backoffFactor          = 2
	DefaultDurPercentile   = 0.4 // consider 40% as healthy, lower percentile = stricter duration check
	DefaultConnsPercentile = 0.8 // consider 80% as healthy, lower percentile = stricter conns check
)

type topologyDriver interface {
	UpdatePeerHealth(peer swarm.Address, health bool, dur time.Duration)
	topology.PeerIterator
}

type peerStatus interface {
	PeerSnapshot(ctx context.Context, peer swarm.Address) (*status.Snapshot, error)
}

type service struct {
	wg            sync.WaitGroup
	quit          chan struct{}
	logger        log.Logger
	topology      topologyDriver
	status        peerStatus
	metrics       metrics
	isSelfHealthy *atomic.Bool
	reserve       storer.RadiusChecker

	radiusSubsMtx sync.Mutex
	radiusC       []chan uint8
}

func New(
	status peerStatus,
	topology topologyDriver,
	reserve storer.RadiusChecker,
	logger log.Logger,
	startupStabilizer stabilization.Subscriber,
	mode string,
	durPercentile float64,
	connsPercentile float64,
) *service {
	_ = "STUB: not implemented"
	return nil
}

func (s *service) worker(startupStabilizer stabilization.Subscriber, mode string, durPercentile float64, connsPercentile float64) {
	_ = "STUB: not implemented"
	return
}

func (s *service) Close() error { _ = "STUB: not implemented"; return nil }

type peer struct {
	status   *status.Snapshot
	dur      time.Duration
	addr     swarm.Address
	bin      uint8
	neighbor bool
}

// salud acquires the status snapshot of every peer and computes an nth percentile of response duration and connected
// per count, the most common storage radius, and the batch commitment, and based on these values, marks peers as unhealhy that fall beyond
// the allowed thresholds.
func (s *service) salud(mode string, durPercentile float64, connsPercentile float64) {
	_ = "STUB: not implemented"
	return
}

// sort peers by duration, highest first to give priority to the fastest peers

// descending

func (s *service) IsHealthy() bool { _ = "STUB: not implemented"; return false }

func (s *service) publishRadius(r uint8) { _ = "STUB: not implemented"; return }

func (s *service) SubscribeNetworkStorageRadius() (<-chan uint8, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// percentileDur finds the p percentile of response duration.
// Less is better.
func percentileDur(peers []peer, p float64) float64 { _ = "STUB: not implemented"; return 0 }

// ascending

// percentileConns finds the p percentile of connection count.
// More is better.
func percentileConns(peers []peer, p float64) uint64 { _ = "STUB: not implemented"; return 0 }

// descending

// radius finds the most common radius.
func (s *service) committedDepth(peers []peer) (uint8, uint8) {
	_ = "STUB: not implemented"
	return 0, 0
}

// commitment finds the most common batch commitment.
func commitment(peers []peer) uint64 { _ = "STUB: not implemented"; return 0 }

func maxIndex(n []int) int { _ = "STUB: not implemented"; return 0 }
