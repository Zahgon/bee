// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handshake

import (
	"github.com/prometheus/client_golang/prometheus"
)

// metrics groups handshake related prometheus counters.
type metrics struct {
	SynRx          prometheus.Counter
	SynRxFailed    prometheus.Counter
	SynAckTx       prometheus.Counter
	SynAckTxFailed prometheus.Counter
	AckRx          prometheus.Counter
	AckRxFailed    prometheus.Counter
}

// newMetrics is a convenient constructor for creating new metrics.
func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

// Metrics returns set of prometheus collectors.
func (s *Service) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
