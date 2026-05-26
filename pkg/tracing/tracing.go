// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tracing

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/opentracing/opentracing-go"
)

var (
	// ErrContextNotFound is returned when tracing context is not present
	// in p2p Headers or context.
	ErrContextNotFound = errors.New("tracing context not found")

	// noopTracer is the tracer that does nothing to handle a nil Tracer usage.
	noopTracer = &Tracer{tracer: new(opentracing.NoopTracer)}
)

// contextKey is used to reference a tracing context span as context value.
type contextKey struct{}

// LogField is the key in log message field that holds tracing id value.
const LogField = "traceID"

const (
	// TraceContextHeaderName is the http header name used to propagate tracing context.
	TraceContextHeaderName = "swarm-trace-id"

	// TraceBaggageHeaderPrefix is the prefix for http headers used to propagate baggage.
	TraceBaggageHeaderPrefix = "swarmctx-"
)

// Tracer connect to a tracing server and handles tracing spans and contexts
// by using opentracing Tracer.
type Tracer struct {
	tracer opentracing.Tracer
}

// Options are optional parameters for Tracer constructor.
type Options struct {
	Enabled     bool
	Endpoint    string
	ServiceName string
}

// NewTracer creates a new Tracer and returns a closer which needs to be closed
// when the Tracer is no longer used to flush remaining traces.
func NewTracer(o *Options) (*Tracer, io.Closer, error) {
	_ = "STUB: not implemented"
	return nil, *new(io.Closer), nil
}

// StartSpanFromContext starts a new tracing span that is either a root one or a
// child of existing one from the provided Context. If logger is provided, a new
// log Entry will be returned with "traceID" log field.
func (t *Tracer) StartSpanFromContext(ctx context.Context, operationName string, l log.Logger, opts ...opentracing.StartSpanOption) (opentracing.Span, log.Logger, context.Context) {
	_ = "STUB: not implemented"
	return *new(opentracing.Span), *new(log.Logger), *new(context.Context)
}

// FollowSpanFromContext starts a new tracing span that is either a root one or
// follows an existing one from the provided Context. If logger is provided, a new
// log Entry will be returned with "traceID" log field.
func (t *Tracer) FollowSpanFromContext(ctx context.Context, operationName string, l log.Logger, opts ...opentracing.StartSpanOption) (opentracing.Span, log.Logger, context.Context) {
	_ = "STUB: not implemented"
	return *new(opentracing.Span), *new(log.Logger), *new(context.Context)
}

// AddContextHeader adds a tracing span context to provided p2p Headers from
// the go context. If the tracing span context is not present in go context,
// ErrContextNotFound is returned.
func (t *Tracer) AddContextHeader(ctx context.Context, headers p2p.Headers) error {
	_ = "STUB: not implemented"
	return nil
}

// FromHeaders returns tracing span context from p2p Headers. If the tracing
// span context is not present in go context, ErrContextNotFound is returned.
func (t *Tracer) FromHeaders(headers p2p.Headers) (opentracing.SpanContext, error) {
	_ = "STUB: not implemented"
	return *new(opentracing.SpanContext), nil
}

// WithContextFromHeaders returns a new context with injected tracing span
// context if they are found in p2p Headers. If the tracing span context is not
// present in go context, ErrContextNotFound is returned.
func (t *Tracer) WithContextFromHeaders(ctx context.Context, headers p2p.Headers) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// AddContextHTTPHeader adds a tracing span context to provided HTTP headers
// from the go context. If the tracing span context is not present in
// go context, ErrContextNotFound is returned.
func (t *Tracer) AddContextHTTPHeader(ctx context.Context, headers http.Header) error {
	_ = "STUB: not implemented"
	return nil
}

// FromHTTPHeaders returns tracing span context from HTTP headers. If the tracing
// span context is not present in go context, ErrContextNotFound is returned.
func (t *Tracer) FromHTTPHeaders(headers http.Header) (opentracing.SpanContext, error) {
	_ = "STUB: not implemented"
	return *new(opentracing.SpanContext), nil
}

// WithContextFromHTTPHeaders returns a new context with injected tracing span
// context if they are found in HTTP headers. If the tracing span context is not
// present in go context, ErrContextNotFound is returned.
func (t *Tracer) WithContextFromHTTPHeaders(ctx context.Context, headers http.Header) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// WithContext adds tracing span context to go context.
func WithContext(ctx context.Context, c opentracing.SpanContext) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// FromContext return tracing span context from go context. If the tracing span
// context is not present in go context, nil is returned.
func FromContext(ctx context.Context) opentracing.SpanContext {
	_ = "STUB: not implemented"
	return *new(opentracing.SpanContext)
}

// NewLoggerWithTraceID creates a new log Entry with "traceID" field added if it
// exists in tracing span context stored from go context.
func NewLoggerWithTraceID(ctx context.Context, l log.Logger) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

func loggerWithTraceID(sc opentracing.SpanContext, l log.Logger) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}
