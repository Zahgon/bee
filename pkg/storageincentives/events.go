// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storageincentives

import (
	"context"
	"sync"
)

type PhaseType int

const (
	commit PhaseType = iota + 1
	reveal
	claim
	sample
)

func (p PhaseType) String() string { _ = "STUB: not implemented"; return "" }

type events struct {
	mtx sync.Mutex
	ev  map[PhaseType]*event
}

type event struct {
	funcs  []func(context.Context)
	ctx    context.Context
	cancel context.CancelFunc
}

func newEvents() *events { _ = "STUB: not implemented"; return nil }

func (e *events) On(phase PhaseType, f func(context.Context)) { _ = "STUB: not implemented"; return }

func (e *events) Publish(phase PhaseType) { _ = "STUB: not implemented"; return }

func (e *events) Cancel(phases ...PhaseType) { _ = "STUB: not implemented"; return }

func (e *events) Close() { _ = "STUB: not implemented"; return }
