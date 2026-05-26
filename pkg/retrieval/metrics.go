// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package retrieval

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	// all metrics fields must be exported
	// to be able to return them by Metrics()
	// using reflection

	RequestCounter        prometheus.Counter
	RequestSuccessCounter prometheus.Counter
	RequestFailureCounter prometheus.Counter
	RequestDurationTime   prometheus.Histogram
	RequestAttempts       prometheus.Histogram
	PeerRequestCounter    prometheus.Counter
	TotalRetrieved        prometheus.Counter
	InvalidChunkRetrieved prometheus.Counter
	ChunkPrice            prometheus.Summary
	TotalErrors           prometheus.Counter
	ChunkRetrieveTime     prometheus.Histogram
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (s *Service) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }

// StatusMetrics exposes metrics that are exposed on the status protocol.
func (s *Service) StatusMetrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
