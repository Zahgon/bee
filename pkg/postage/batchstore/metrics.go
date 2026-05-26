// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package batchstore

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	Commitment        prometheus.Gauge
	Radius            prometheus.Gauge
	UnreserveDuration *prometheus.HistogramVec
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (s *store) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
