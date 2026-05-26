// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package swap

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	TotalReceived    prometheus.Counter
	TotalSent        prometheus.Counter
	ChequesReceived  prometheus.Counter
	ChequesSent      prometheus.Counter
	ChequesRejected  prometheus.Counter
	AvailableBalance prometheus.Gauge
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (s *Service) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
