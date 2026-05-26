// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ratelimit provides a mechanism to rate limit requests based on a string key,
// refill rate and burst amount. Under the hood, it's a token bucket of size burst amount,
// that refills at the refill rate.
package ratelimit

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Limiter struct {
	mtx     sync.Mutex
	limiter map[string]*rate.Limiter
	rate    rate.Limit
	burst   int
}

// New returns a new Limiter object with refresh rate and burst amount
func New(r time.Duration, burst int) *Limiter { _ = "STUB: not implemented"; return nil }

// Allow checks if the limiter that belongs to 'key' has not exceeded the limit.
func (l *Limiter) Allow(key string, count int) bool { _ = "STUB: not implemented"; return false }

// Wait blocks until the limiter permits n events to happen. Returns the time duration
// the limiter waited for to allow the number of events to occur.
func (l *Limiter) Wait(ctx context.Context, key string, count int) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// Clear deletes the limiter that belongs to 'key'
func (l *Limiter) getLimiter(key string) *rate.Limiter { _ = "STUB: not implemented"; return nil }

// Clear deletes the limiter that belongs to 'key'
func (l *Limiter) Clear(key string) { _ = "STUB: not implemented"; return }
