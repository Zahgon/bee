// Copyright 2025 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

type subscriber struct {
	stable bool
}

func NewSubscriber(stable bool) *subscriber { _ = "STUB: not implemented"; return nil }

func (s *subscriber) Subscribe() (<-chan struct{}, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *subscriber) IsStabilized() bool { _ = "STUB: not implemented"; return false }
