// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libp2p

import (
	"errors"
	"time"

	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/libp2p/go-libp2p/core/network"
)

var (
	closeDeadline  = 30 * time.Second
	errExpectedEof = errors.New("read: expected eof")
)
var _ p2p.Stream = (*stream)(nil)

type stream struct {
	network.Stream
	headers         map[string][]byte
	responseHeaders map[string][]byte
	metrics         metrics
}

func newStream(s network.Stream, metrics metrics) *stream { _ = "STUB: not implemented"; return nil }

func (s *stream) Headers() p2p.Headers { _ = "STUB: not implemented"; return *new(p2p.Headers) }

func (s *stream) ResponseHeaders() p2p.Headers { _ = "STUB: not implemented"; return *new(p2p.Headers) }

func (s *stream) Reset() error { _ = "STUB: not implemented"; return nil }

func (s *stream) FullClose() error { _ = "STUB: not implemented"; return nil }

// close the stream to make sure it is gc'd

// So we don't wait forever

// We *have* to observe the EOF. Otherwise, we leak the stream.
// Now, technically, we should do this *before*
// returning from SendMessage as the message
// hasn't really been sent yet until we see the
// EOF but we don't actually *know* what
// protocol the other side is speaking.
