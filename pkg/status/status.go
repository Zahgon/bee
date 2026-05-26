// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package status

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/status/internal/pb"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/topology"
	"github.com/prometheus/client_golang/prometheus"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "status"

const (
	protocolName    = "status"
	protocolVersion = "1.1.3"
	streamName      = "status"
)

// Snapshot is the current snapshot of the system.
type Snapshot pb.Snapshot

// SyncReporter defines the interface to report syncing rate.
type SyncReporter interface {
	SyncRate() float64
}

// Reserve defines the reserve storage related information required.
type Reserve interface {
	ReserveSize() int
	ReserveSizeWithinRadius() uint64
	StorageRadius() uint8
	CommittedDepth() uint8
}

type topologyDriver interface {
	topology.PeerIterator
	IsReachable() bool
}

// Service is the status service.
type Service struct {
	logger         log.Logger
	streamer       p2p.Streamer
	topologyDriver topologyDriver

	beeMode         string
	reserve         Reserve
	sync            SyncReporter
	chainState      postage.ChainStateGetter
	metricsRegistry *prometheus.Registry
}

type Metrics struct {
	UploadSpeed   prometheus.Histogram
	DownloadSpeed prometheus.Histogram
}

// NewService creates a new status service.
func NewService(
	logger log.Logger,
	streamer p2p.Streamer,
	topology topologyDriver,
	beeMode string,
	chainState postage.ChainStateGetter,
	reserve Reserve,
	metricsRegistry *prometheus.Registry,
) *Service {
	_ = "STUB: not implemented"
	return nil
}

// LocalSnapshot returns the current status snapshot of this node.
func (s *Service) LocalSnapshot() (*Snapshot, error) { _ = "STUB: not implemented"; return nil, nil }

// include self

// PeerSnapshot sends request for status snapshot to the peer.
func (s *Service) PeerSnapshot(ctx context.Context, peer swarm.Address) (*Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Protocol returns the protocol specification.
func (s *Service) Protocol() p2p.ProtocolSpec {
	_ = "STUB: not implemented"
	return *new(p2p.ProtocolSpec)
}

// handler handles the status stream request/response.
func (s *Service) handler(ctx context.Context, _ p2p.Peer, stream p2p.Stream) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) SetSync(sync SyncReporter) { _ = "STUB: not implemented"; return }

func (s *Service) encodeMetrics() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
