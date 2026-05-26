// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package breaker

import (
	"errors"
	"sync"
	"time"
)

const (
	// defaults
	limit        = 100
	failInterval = 30 * time.Minute
	maxBackoff   = time.Hour
	backoff      = 2 * time.Minute
)

var (
	_ Interface = (*breaker)(nil)

	// ErrClosed is the special error type that indicates that breaker is closed and that is not executing functions at the moment.
	ErrClosed = errors.New("breaker closed")
)

type Interface interface {
	// Execute runs f() if the limit number of consecutive failed calls is not reached within fail interval.
	// f() call is not locked so it can still be executed concurrently.
	// Returns `ErrClosed` if the limit is reached or f() result otherwise.
	Execute(f func() error) error

	// ClosedUntil returns the timestamp when the breaker will become open again.
	ClosedUntil() time.Time
}

type currentTimeFn = func() time.Time

type breaker struct {
	limit                int // breaker will not execute any more tasks after limit number of consecutive failures happen
	consFailedCalls      int // current number of consecutive fails
	firstFailedTimestamp time.Time
	closedTimestamp      time.Time
	backoff              time.Duration // initial backoff duration
	maxBackoff           time.Duration
	failInterval         time.Duration // consecutive failures are counted if they happen within this interval
	currentTimeFn        currentTimeFn
	mtx                  sync.Mutex
}

type Options struct {
	Limit        int
	FailInterval time.Duration
	StartBackoff time.Duration
	MaxBackoff   time.Duration
}

func NewBreaker(o Options) Interface { _ = "STUB: not implemented"; return *new(Interface) }

func newBreakerWithCurrentTimeFn(o Options, currentTimeFn currentTimeFn) Interface {
	_ = "STUB: not implemented"
	return *new(Interface)
}

func (b *breaker) Execute(f func() error) error { _ = "STUB: not implemented"; return nil }

func (b *breaker) ClosedUntil() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (b *breaker) beforef() error { _ = "STUB: not implemented"; return nil }

// use currentTimeFn().Sub() instead of time.Since() so it can be deterministically mocked in tests

func (b *breaker) afterf(err error) error { _ = "STUB: not implemented"; return nil }

func (b *breaker) resetFailed() { _ = "STUB: not implemented"; return }
