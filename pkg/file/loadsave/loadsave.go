// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package loadsave provides lightweight persistence abstraction
// for manifest operations.
package loadsave

import (
	"context"
	"errors"

	"github.com/ethersphere/bee/v2/pkg/file"
	"github.com/ethersphere/bee/v2/pkg/file/pipeline"
	"github.com/ethersphere/bee/v2/pkg/file/redundancy"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var errReadonlyLoadSave = errors.New("readonly manifest loadsaver")

// loadSave is needed for manifest operations and provides
// simple wrapping over load and save operations using file
// package abstractions. use with caution since Loader will
// load all of the subtrie of a given hash in memory.
type loadSave struct {
	getter     storage.Getter
	putter     storage.Putter
	pipelineFn func() pipeline.Interface
	rootCh     swarm.Chunk
	rLevel     redundancy.Level
}

// New returns a new read-write load-saver.
func New(getter storage.Getter, putter storage.Putter, pipelineFn func() pipeline.Interface, rLevel redundancy.Level) file.LoadSaver {
	_ = "STUB: not implemented"
	return *new(file.LoadSaver)
}

// NewReadonly returns a new read-only load-saver
// which will error on write.
func NewReadonly(getter storage.Getter, putter storage.Putter, rLevel redundancy.Level) file.LoadSaver {
	_ = "STUB: not implemented"
	return *new(file.LoadSaver)
}

// NewReadonlyWithRootCh returns a new read-only load-saver
// which will error on write.
func NewReadonlyWithRootCh(getter storage.Getter, putter storage.Putter, rootCh swarm.Chunk, rLevel redundancy.Level) file.LoadSaver {
	_ = "STUB: not implemented"
	return *new(file.LoadSaver)
}

func (ls *loadSave) Load(ctx context.Context, ref []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ls *loadSave) Save(ctx context.Context, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
