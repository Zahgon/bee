// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storageincentives

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	// phase gauge and counter
	CurrentPhase            prometheus.Gauge
	RevealPhase             prometheus.Counter
	CommitPhase             prometheus.Counter
	ClaimPhase              prometheus.Counter
	Winner                  prometheus.Counter
	NeighborhoodSelected    prometheus.Counter
	SampleDuration          prometheus.Gauge
	Round                   prometheus.Gauge
	InsufficientFundsToPlay prometheus.Counter

	// total calls to chain backend
	BackendCalls  prometheus.Counter
	BackendErrors prometheus.Counter

	// metrics for err processing
	ErrReveal         prometheus.Counter
	ErrCommit         prometheus.Counter
	ErrClaim          prometheus.Counter
	ErrWinner         prometheus.Counter
	ErrCheckIsPlaying prometheus.Counter
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

// total call

// phase errors

// TODO: register metric
func (a *Agent) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
