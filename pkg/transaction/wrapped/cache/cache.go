// Copyright 2026 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"context"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"resenje.org/singleflight"
)

type (
	Loader[T any]         func() (T, error)
	ReuseEvaluator[T any] func(value T) bool
)

type SingleFlightCache[T any] struct {
	mu    sync.RWMutex
	value T

	group   singleflight.Group[string, any]
	key     string
	metrics metricSet
}

func NewSingleFlightCache[T any](metricsPrefix string) *SingleFlightCache[T] {
	_ = "STUB: not implemented"
	return nil
}

func (c *SingleFlightCache[T]) Collectors() []prometheus.Collector {
	_ = "STUB: not implemented"
	return nil
}

func (c *SingleFlightCache[T]) Set(value T) { _ = "STUB: not implemented"; return }

func (c *SingleFlightCache[T]) PeekOrLoad(ctx context.Context, canReuse ReuseEvaluator[T], loader Loader[T]) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
