// Copyright 2026 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metricSet struct {
	Hits        prometheus.Counter
	Misses      prometheus.Counter
	Loads       prometheus.Counter
	SharedLoads prometheus.Counter
	LoadErrors  prometheus.Counter
}

func newMetricSet(prefix string) metricSet { _ = "STUB: not implemented"; return *new(metricSet) }
