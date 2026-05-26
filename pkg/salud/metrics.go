// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package salud

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	AvgDur                prometheus.Gauge
	PDur                  prometheus.Gauge
	PConns                prometheus.Gauge
	NetworkRadius         prometheus.Gauge
	NeighborhoodRadius    prometheus.Gauge
	Commitment            prometheus.Gauge
	ReserveSizePercentErr prometheus.Gauge
	Healthy               prometheus.Counter
	Unhealthy             prometheus.Counter
	NeighborhoodAvgDur    prometheus.Gauge
	NeighborCount         prometheus.Gauge
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

// Neighborhood-specific metrics

func (s *service) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
