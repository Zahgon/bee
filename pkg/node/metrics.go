// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package node

import (
	"github.com/prometheus/client_golang/prometheus"
)

type nodeMetrics struct {
	// WarmupDuration measures time in seconds for the node warmup to complete
	WarmupDuration prometheus.Histogram
	// FullSyncDuration measures time in seconds for the full sync to complete
	FullSyncDuration prometheus.Histogram
}

func newMetrics() nodeMetrics { _ = "STUB: not implemented"; return *new(nodeMetrics) }

// middle range should be more infrequent (because of addressbook)

// middle range should be more frequent

// 2-3 hours range

func getMetrics(nodeMetrics nodeMetrics) []prometheus.Collector {
	_ = "STUB: not implemented"
	return nil
}
