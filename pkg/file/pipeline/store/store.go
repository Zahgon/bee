// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package store

import (
	"context"
	"errors"

	"github.com/ethersphere/bee/v2/pkg/file/pipeline"
	storage "github.com/ethersphere/bee/v2/pkg/storage"
)

var errInvalidData = errors.New("store: invalid data")

type storeWriter struct {
	l    storage.Putter
	ctx  context.Context
	next pipeline.ChainWriter
}

// NewStoreWriter returns a storeWriter. It just writes the given data
// to a given storage.Putter.
func NewStoreWriter(ctx context.Context, l storage.Putter, next pipeline.ChainWriter) pipeline.ChainWriter {
	_ = "STUB: not implemented"
	return *new(pipeline.ChainWriter)
}

func (w *storeWriter) ChainWrite(p *pipeline.PipeWriteArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *storeWriter) Sum() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
