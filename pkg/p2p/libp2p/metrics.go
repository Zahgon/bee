// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libp2p

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	// all metrics fields must be exported
	// to be able to return them by Metrics()
	// using reflection
	CreatedConnectionCount     prometheus.Counter
	HandledConnectionCount     prometheus.Counter
	CreatedStreamCount         prometheus.Counter
	ClosedStreamCount          prometheus.Counter
	StreamResetCount           prometheus.Counter
	HandledStreamCount         prometheus.Counter
	BlocklistedPeerCount       prometheus.Counter
	BlocklistedPeerErrCount    prometheus.Counter
	DisconnectCount            prometheus.Counter
	ConnectBreakerCount        prometheus.Counter
	UnexpectedProtocolReqCount prometheus.Counter
	KickedOutPeersCount        prometheus.Counter
	StreamHandlerErrResetCount prometheus.Counter
	HeadersExchangeDuration    prometheus.Histogram
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (s *Service) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }

// StatusMetrics exposes metrics that are exposed on the status protocol.
func (s *Service) StatusMetrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
