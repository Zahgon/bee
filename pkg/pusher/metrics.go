// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pusher

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	TotalToPush      prometheus.Counter
	TotalSynced      prometheus.Counter
	TotalErrors      prometheus.Counter
	MarkAndSweepTime prometheus.Histogram
	SyncTime         prometheus.Histogram
	ErrorTime        prometheus.Histogram
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (s *Service) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
