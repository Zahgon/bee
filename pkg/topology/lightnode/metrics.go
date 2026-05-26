// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lightnode

import (
	"github.com/prometheus/client_golang/prometheus"
)

// metrics groups lightnode related prometheus counters.
type metrics struct {
	CurrentlyConnectedPeers    prometheus.Gauge
	CurrentlyDisconnectedPeers prometheus.Gauge
}

// newMetrics is a convenient constructor for creating new metrics.
func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

// Metrics returns set of prometheus collectors.
func (c *Container) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
