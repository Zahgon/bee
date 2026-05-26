// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pullsync

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	Offered              prometheus.Counter     // number of chunks offered
	Wanted               prometheus.Counter     // number of chunks wanted
	MissingChunks        prometheus.Counter     // number of reserve get errs
	ReceivedZeroAddress  prometheus.Counter     // number of delivered chunks with invalid address
	ReceivedInvalidChunk prometheus.Counter     // number of delivered chunks with invalid address
	Delivered            prometheus.Counter     // number of chunk deliveries
	SentOffered          prometheus.Counter     // number of chunks offered
	SentWanted           prometheus.Counter     // number of chunks wanted
	Sent                 prometheus.Counter     // number of chunks sent
	DuplicateRuid        prometheus.Counter     // number of duplicate RUID requests we got
	LastReceived         *prometheus.CounterVec // last timestamp of the received chunks per bin
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (s *Syncer) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
