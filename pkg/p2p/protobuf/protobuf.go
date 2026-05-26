// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protobuf

import (
	"context"
	"errors"
	"io"

	"github.com/ethersphere/bee/v2/pkg/p2p"
	ggio "github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
)

const delimitedReaderMaxSize = 128 * 1024 // max message size

var ErrTimeout = errors.New("timeout")

type Message = proto.Message

func NewWriterAndReader(s p2p.Stream) (Writer, Reader) {
	_ = "STUB: not implemented"
	return *new(Writer), *new(Reader)
}

func NewReader(r io.Reader) Reader { _ = "STUB: not implemented"; return *new(Reader) }

func NewWriter(w io.Writer) Writer { _ = "STUB: not implemented"; return *new(Writer) }

func ReadMessages(r io.Reader, newMessage func() Message) (m []Message, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Reader struct {
	ggio.Reader
}

func newReader(r ggio.Reader) Reader { _ = "STUB: not implemented"; return *new(Reader) }

func (r Reader) ReadMsgWithContext(ctx context.Context, msg proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

type Writer struct {
	ggio.Writer
}

func newWriter(r ggio.Writer) Writer { _ = "STUB: not implemented"; return *new(Writer) }

func (w Writer) WriteMsgWithContext(ctx context.Context, msg proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}
