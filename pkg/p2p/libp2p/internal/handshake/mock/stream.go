// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"bytes"

	"github.com/ethersphere/bee/v2/pkg/p2p"
)

type Stream struct {
	readBuffer        *bytes.Buffer
	writeBuffer       *bytes.Buffer
	writeCounter      int
	readCounter       int
	readError         error
	writeError        error
	readErrCheckmark  int
	writeErrCheckmark int
}

func NewStream(readBuffer, writeBuffer *bytes.Buffer) *Stream {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) SetReadErr(err error, checkmark int) { _ = "STUB: not implemented"; return }

func (s *Stream) SetWriteErr(err error, checkmark int) { _ = "STUB: not implemented"; return }

func (s *Stream) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stream) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stream) Headers() p2p.Headers { _ = "STUB: not implemented"; return *new(p2p.Headers) }

func (s *Stream) ResponseHeaders() p2p.Headers { _ = "STUB: not implemented"; return *new(p2p.Headers) }

func (s *Stream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Stream) FullClose() error { _ = "STUB: not implemented"; return nil }

func (s *Stream) Reset() error { _ = "STUB: not implemented"; return nil }
