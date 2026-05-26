// Copyright 2026 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reacher

import (
	"github.com/prometheus/client_golang/prometheus"
)

// metrics groups reacher related prometheus counters.
type metrics struct {
	Peers            prometheus.Gauge
	PingAttemptCount prometheus.Counter
	PingErrorCount   prometheus.Counter
	PingDuration     prometheus.Histogram
}

// newMetrics is a convenient constructor for creating new metrics.
func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

// Metrics returns set of prometheus collectors.
func (r *reacher) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
