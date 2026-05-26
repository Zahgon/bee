// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import "context"

type mockSyncer struct{ rate float64 }

func NewMockRateReporter(r float64) *mockSyncer { _ = "STUB: not implemented"; return nil }
func (m *mockSyncer) SyncRate() float64         { _ = "STUB: not implemented"; return 0 }
func (m *mockSyncer) Start(context.Context)     { _ = "STUB: not implemented"; return }
