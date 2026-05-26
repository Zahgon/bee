// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package events

import (
	"sync"
)

type Subscriber struct {
	mtx  sync.Mutex
	subs map[string][]chan struct{}
}

func NewSubscriber() *Subscriber { _ = "STUB: not implemented"; return nil }

func (b *Subscriber) Subscribe(str string) (<-chan struct{}, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Subscriber) Trigger(str string) { _ = "STUB: not implemented"; return }
