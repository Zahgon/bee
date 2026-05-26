// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	CacheHit  prometheus.Counter
	CacheMiss prometheus.Counter
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (c *Cache) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
