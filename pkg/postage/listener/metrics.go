// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package listener

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	// aggregate events handled
	EventsProcessed prometheus.Counter
	EventErrors     prometheus.Counter
	PagesProcessed  prometheus.Counter

	// individual event counters
	CreatedCounter prometheus.Counter
	TopupCounter   prometheus.Counter
	DepthCounter   prometheus.Counter
	PriceCounter   prometheus.Counter

	// total calls to chain backend
	BackendCalls  prometheus.Counter
	BackendErrors prometheus.Counter

	// processing durations
	PageProcessDuration  prometheus.Counter
	EventProcessDuration prometheus.Counter
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

// aggregate events handled

// individual event counters

// total call

// processing durations

func (l *listener) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
