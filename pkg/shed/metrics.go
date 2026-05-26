// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shed

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	// all metrics fields must be exported
	// to be able to return them by Metrics()
	// using reflection
	PutCounter            prometheus.Counter
	PutFailCounter        prometheus.Counter
	GetCounter            prometheus.Counter
	GetFailCounter        prometheus.Counter
	GetNotFoundCounter    prometheus.Counter
	HasCounter            prometheus.Counter
	HasFailCounter        prometheus.Counter
	DeleteCounter         prometheus.Counter
	DeleteFailCounter     prometheus.Counter
	IteratorCounter       prometheus.Counter
	WriteBatchCounter     prometheus.Counter
	WriteBatchFailCounter prometheus.Counter
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (s *DB) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
