// Copyright 2025 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libp2ptest

import (
	"bytes"
	"sync"
)

// SafeBuffer is a thread-safe bytes.Buffer.
type SafeBuffer struct {
	b bytes.Buffer
	m sync.Mutex
}

func (s *SafeBuffer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *SafeBuffer) String() string { _ = "STUB: not implemented"; return "" }
