// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

const bytesInKB = 1000

var fileSizeBucketsKBytes = []int64{100, 500, 2500, 4999, 5000, 10000}

type metrics struct {
	// all metrics fields must be exported
	// to be able to return them by Metrics()
	// using reflection
	RequestCount       prometheus.Counter
	ResponseDuration   prometheus.Histogram
	PingRequestCount   prometheus.Counter
	ResponseCodeCounts *prometheus.CounterVec

	ContentApiDuration *prometheus.HistogramVec
	UploadSpeed        *prometheus.HistogramVec
	DownloadSpeed      *prometheus.HistogramVec
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func toFileSizeBucket(bytes int64) int64 { _ = "STUB: not implemented"; return 0 }

func (s *Service) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }

// StatusMetrics exposes metrics that are exposed on the status protocol.
func (s *Service) StatusMetrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }

func (s *Service) pageviewMetricsHandler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (s *Service) responseCodeMetricsHandler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// UpgradedResponseWriter adds more functionality on top of ResponseWriter
type UpgradedResponseWriter interface {
	http.ResponseWriter
	http.Pusher
	http.Hijacker
	http.Flusher
}

type responseWriter struct {
	UpgradedResponseWriter
	statusCode  int
	wroteHeader bool
	size        int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	_ = "STUB: not implemented"
	// StatusOK is called by default if nothing else is called
	return nil
}

func (rw *responseWriter) Status() int { _ = "STUB: not implemented"; return 0 }

func (rr *responseWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (rw *responseWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

func newDebugMetrics() (r *prometheus.Registry) { _ = "STUB: not implemented"; return nil }

// register standard metrics

func (s *Service) MetricsRegistry() *prometheus.Registry { _ = "STUB: not implemented"; return nil }

func (s *Service) MustRegisterMetrics(cs ...prometheus.Collector) {
	_ = "STUB: not implemented"
	return
}
