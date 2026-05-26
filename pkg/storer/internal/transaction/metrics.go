// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transaction

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	MethodCalls    *prometheus.CounterVec
	MethodDuration *prometheus.HistogramVec
}

// newMetrics is a convenient constructor for creating new metrics.
func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }
