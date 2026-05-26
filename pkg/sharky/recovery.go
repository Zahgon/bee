// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sharky

import (
	"context"
	"errors"
	"os"
	"sync"
)

// Recovery allows disaster recovery.
type Recovery struct {
	mtx        sync.Mutex
	shards     []*slots
	shardFiles []*os.File
	datasize   int
}

var ErrShardNotFound = errors.New("shard not found")

func NewRecovery(dir string, shardCnt int, datasize int) (*Recovery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add marks a location as used (not free).
func (r *Recovery) Add(loc Location) error { _ = "STUB: not implemented"; return nil }

func (r *Recovery) Read(ctx context.Context, loc Location, buf []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Recovery) Move(ctx context.Context, from Location, to Location) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Recovery) TruncateAt(ctx context.Context, shard uint8, slot uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// Save saves all free slots files of the recovery (without closing).
func (r *Recovery) Save() error { _ = "STUB: not implemented"; return nil }

// Close closes data and free slots files of the recovery (without saving).
func (r *Recovery) Close() error { _ = "STUB: not implemented"; return nil }
