// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sharky

import (
	"github.com/prometheus/client_golang/prometheus"
)

// metrics groups sharky related prometheus counters.
type metrics struct {
	TotalWriteCalls        prometheus.Counter
	TotalWriteCallsErr     prometheus.Counter
	TotalReadCalls         prometheus.Counter
	TotalReadCallsErr      prometheus.Counter
	TotalReleaseCalls      prometheus.Counter
	TotalReleaseCallsErr   prometheus.Counter
	ShardCount             prometheus.Gauge
	CurrentShardSize       *prometheus.GaugeVec
	ShardFragmentation     *prometheus.GaugeVec
	LastAllocatedShardSlot *prometheus.GaugeVec
	LastReleasedShardSlot  *prometheus.GaugeVec
}

// newMetrics is a convenient constructor for creating new metrics.
func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

// Metrics returns set of prometheus collectors.
func (s *Store) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
