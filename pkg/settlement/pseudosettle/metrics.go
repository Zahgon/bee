// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pseudosettle

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	// all metrics fields must be exported
	// to be able to return them by Metrics()
	// using reflection
	TotalReceivedPseudoSettlements  prometheus.Counter
	TotalSentPseudoSettlements      prometheus.Counter
	ReceivedPseudoSettlements       prometheus.Counter
	SentPseudoSettlements           prometheus.Counter
	ReceivedPseudoSettlementsErrors prometheus.Counter
	SentPseudoSettlementsErrors     prometheus.Counter
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (s *Service) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
