// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpaccess

import (
	"bufio"
	"net"
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/tracing"
)

// NewHTTPAccessSuppressLogHandler creates a
// handler that will suppress access log messages.
func NewHTTPAccessSuppressLogHandler() func(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

// NewHTTPAccessLogHandler creates a handler that
// will log a message after a request has been served.
func NewHTTPAccessLogHandler(logger log.Logger, tracer *tracing.Tracer, message string) func(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

// No need to layer on another responseRecorder.

// responseRecorder is an implementation of
// http.ResponseWriter that records various metrics.
type responseRecorder struct {
	http.ResponseWriter

	// Metrics.
	status int
	size   int
}

// Write implements http.ResponseWriter.
func (rr *responseRecorder) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteHeader implements http.ResponseWriter.
func (rr *responseRecorder) WriteHeader(s int) { _ = "STUB: not implemented"; return }

// CloseNotify implements http.CloseNotifier.
func (rr *responseRecorder) CloseNotify() <-chan bool {
	_ = "STUB: not implemented"
	// staticcheck SA1019 CloseNotifier interface is required by gorilla compress handler.
	// nolint:staticcheck
	return nil
}

// Hijack implements http.Hijacker.
func (rr *responseRecorder) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

// Flush implements http.Flusher.
func (rr *responseRecorder) Flush() { _ = "STUB: not implemented"; return }

// Push implements http.Pusher.
func (rr *responseRecorder) Push(target string, opts *http.PushOptions) error {
	_ = "STUB: not implemented"
	return nil
}
