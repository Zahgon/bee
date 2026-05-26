// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"github.com/prometheus/client_golang/prometheus"
)

// metrics groups various metrics counters for statistical reasons.
type metrics struct {
	ErrorCount prometheus.Counter
	WarnCount  prometheus.Counter
	InfoCount  prometheus.Counter
	DebugCount prometheus.Counter
	TraceCount prometheus.Counter
}

// Fire implements Hook interface.
func (m metrics) Fire(v Level) error { _ = "STUB: not implemented"; return nil }

// newLogMetrics returns pointer to a new metrics instance ready to use.
func newLogMetrics() *metrics { _ = "STUB: not implemented"; return nil }
