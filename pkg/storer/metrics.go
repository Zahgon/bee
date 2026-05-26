// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"context"
	"time"

	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/prometheus/client_golang/prometheus"
)

// metrics groups storer related prometheus counters.
type metrics struct {
	MethodCalls                   *prometheus.CounterVec
	MethodCallsDuration           *prometheus.HistogramVec
	ReserveSize                   prometheus.Gauge
	ReserveSizeWithinRadius       prometheus.Gauge
	ReserveCleanup                prometheus.Counter
	StorageRadius                 prometheus.Gauge
	CacheSize                     prometheus.Gauge
	EvictedChunkCount             prometheus.Counter
	ExpiredChunkCount             prometheus.Counter
	OverCapTriggerCount           prometheus.Counter
	ExpiredBatchCount             prometheus.Counter
	LevelDBStats                  *prometheus.HistogramVec
	ExpiryTriggersCount           prometheus.Counter
	ExpiryRunsCount               prometheus.Counter
	ReserveMissingBatch           prometheus.Gauge
	ReserveSampleDuration         *prometheus.HistogramVec
	ReserveSampleRunSummary       *prometheus.GaugeVec
	ReserveSampleLastRunTimestamp prometheus.Gauge
	RecoveryPrunedChunkCount      prometheus.Counter
}

// newMetrics is a convenient constructor for creating new metrics.
func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

var _ storage.Putter = (*putterWithMetrics)(nil)

// putterWithMetrics wraps storage.Putter and adds metrics.
type putterWithMetrics struct {
	storage.Putter

	metrics   metrics
	component string
}

func (m putterWithMetrics) Put(ctx context.Context, chunk swarm.Chunk) error {
	_ = "STUB: not implemented"
	return nil
}

var _ storage.Getter = (*getterWithMetrics)(nil)

// getterWithMetrics wraps storage.Getter and adds metrics.
type getterWithMetrics struct {
	storage.Getter

	metrics   metrics
	component string
}

func (m getterWithMetrics) Get(ctx context.Context, address swarm.Address) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// captureDuration returns a function that returns the duration since the given start.
func captureDuration(start time.Time) func() float64 { _ = "STUB: not implemented"; return nil }
