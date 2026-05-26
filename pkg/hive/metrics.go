// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hive

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	BroadcastPeers      prometheus.Counter
	BroadcastPeersPeers prometheus.Counter
	BroadcastPeersSends prometheus.Counter

	PeersHandler      prometheus.Counter
	PeersHandlerPeers prometheus.Counter
	UnreachablePeers  prometheus.Counter

	PingTime        prometheus.Histogram
	PingFailureTime prometheus.Histogram

	PeerConnectAttempts prometheus.Counter
	PeerUnderlayErr     prometheus.Counter
	StorePeerErr        prometheus.Counter
	ReachablePeers      prometheus.Counter
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (s *Service) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
