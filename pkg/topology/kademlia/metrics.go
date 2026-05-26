// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kademlia

import (
	"github.com/prometheus/client_golang/prometheus"
)

// metrics groups kademlia related prometheus counters.
type metrics struct {
	PickCalls                             prometheus.Counter
	PickCallsFalse                        prometheus.Counter
	CurrentDepth                          prometheus.Gauge
	CurrentStorageDepth                   prometheus.Gauge
	CurrentlyKnownPeers                   prometheus.Gauge
	CurrentlyConnectedPeers               prometheus.Gauge
	InternalMetricsFlushTime              prometheus.Histogram
	InternalMetricsFlushTotalErrors       prometheus.Counter
	TotalBeforeExpireWaits                prometheus.Counter
	TotalInboundConnections               prometheus.Counter
	TotalInboundDisconnections            prometheus.Counter
	TotalOutboundConnections              prometheus.Counter
	TotalOutboundConnectionAttempts       prometheus.Counter
	TotalOutboundConnectionFailedAttempts prometheus.Counter
	TotalBootNodesConnectionAttempts      prometheus.Counter
	StartAddAddressBookOverlaysTime       prometheus.Histogram
	PeerLatencyEWMA                       prometheus.Histogram
	Blocklist                             prometheus.Counter
	ReachabilityStatus                    *prometheus.GaugeVec
	PeersReachabilityStatus               *prometheus.GaugeVec
}

// newMetrics is a convenient constructor for creating new metrics.
func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

// Metrics returns set of prometheus collectors.
func (k *Kad) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
