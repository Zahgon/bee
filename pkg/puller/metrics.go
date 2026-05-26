// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package puller

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	SyncWorkerIterCounter prometheus.Counter     // counts the number of syncing iterations
	SyncWorkerCounter     prometheus.Gauge       // count number of syncing jobs
	SyncedCounter         *prometheus.CounterVec // number of synced chunks
	SyncWorkerErrCounter  prometheus.Counter     // count number of errors
	MaxUintErrCounter     prometheus.Counter     // how many times we got maxuint as topmost
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (p *Puller) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
