// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pushsync

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	TotalSent               prometheus.Counter
	TotalReceived           prometheus.Counter
	TotalHandlerErrors      prometheus.Counter
	TotalRequests           prometheus.Counter
	TotalSendAttempts       prometheus.Counter
	TotalFailedSendAttempts prometheus.Counter
	TotalOutgoing           prometheus.Counter
	TotalOutgoingErrors     prometheus.Counter
	InvalidStampErrors      prometheus.Counter
	StampValidationTime     *prometheus.HistogramVec
	Forwarder               prometheus.Counter
	Storer                  prometheus.Counter
	TotalHandlerTime        *prometheus.HistogramVec
	PushToPeerTime          *prometheus.HistogramVec

	ReceiptDepth        *prometheus.CounterVec
	ShallowReceiptDepth *prometheus.CounterVec
	ShallowReceipt      prometheus.Counter
	OverdraftRefresh    prometheus.Counter
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (s *PushSync) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
