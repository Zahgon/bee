// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pss

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	TotalMessagesSentCounter prometheus.Counter
	MessageMiningDuration    prometheus.Gauge
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (s *pss) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
