// Copyright 2025 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"sync"
	"time"
)

type mockClock struct {
	mu   sync.Mutex
	time time.Time
}

func NewClock(t time.Time) *mockClock { _ = "STUB: not implemented"; return nil }

// Now returns the current mock time.
func (mc *mockClock) Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Advance advances the mock time by the given duration.
func (mc *mockClock) Advance(d time.Duration) { _ = "STUB: not implemented"; return }
