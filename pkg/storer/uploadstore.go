// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"context"

	storage "github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const uploadsLock = "pin-upload-store"

// Report implements the storage.PushReporter by wrapping the internal reporter
// with a transaction.
func (db *DB) Report(ctx context.Context, chunk swarm.Chunk, state storage.ChunkState) error {
	_ = "STUB: not implemented"
	return nil
}

// Upload is the implementation of UploadStore.Upload method.
func (db *DB) Upload(ctx context.Context, pin bool, tagID uint64) (PutterSession, error) {
	_ = "STUB: not implemented"
	return *new(PutterSession), nil
}

// NewSession is the implementation of UploadStore.NewSession method.
func (db *DB) NewSession() (SessionInfo, error) {
	_ = "STUB: not implemented"
	return *new(SessionInfo), nil
}

// Session is the implementation of the UploadStore.Session method.
func (db *DB) Session(tagID uint64) (SessionInfo, error) {
	_ = "STUB: not implemented"
	return *new(SessionInfo), nil
}

// DeleteSession is the implementation of the UploadStore.DeleteSession method.
func (db *DB) DeleteSession(tagID uint64) error { _ = "STUB: not implemented"; return nil }

// ListSessions is the implementation of the UploadStore.ListSessions method.
func (db *DB) ListSessions(offset, limit int) ([]SessionInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
