// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sharky

import (
	"context"
	"errors"
	"io/fs"
	"sync"
)

var (
	// ErrTooLong returned by Write if the blob length exceeds the max blobsize.
	ErrTooLong = errors.New("data too long")
	// ErrQuitting returned by Write when the store is Closed before the write completes.
	ErrQuitting = errors.New("quitting")
)

// Store models the sharded fix-length blobstore
// Design provides lockless sharding:
// - shard choice responding to backpressure by running operation
// - read prioritisation over writing
// - free slots allow write
type Store struct {
	maxDataSize int             // max length of blobs
	writes      chan write      // shared write operations channel
	shards      []*shard        // shards
	wg          *sync.WaitGroup // count started operations
	quit        chan struct{}   // quit channel
	metrics     metrics
}

// New constructs a sharded blobstore
// arguments:
// - base directory string
// - shard count - positive integer < 256 - cannot be zero or expect panic
// - shard size - positive integer multiple of 8 - for others expect undefined behaviour
// - maxDataSize - positive integer representing the maximum blob size to be stored
func New(basedir fs.FS, shardCnt int, maxDataSize int) (*Store, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes each shard and return incidental errors from each shard
func (s *Store) Close() error { _ = "STUB: not implemented"; return nil }

// create creates a new shard with index, max capacity limit, file within base directory
func (s *Store) create(index uint8, maxDataSize int, basedir fs.FS) (*shard, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read reads the content of the blob found at location into the byte buffer given
// The location is assumed to be obtained by an earlier Write call storing the blob
func (s *Store) Read(ctx context.Context, loc Location, buf []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// it is important that this select would NEVER respect the context
// cancellation. this would result in a deadlock on the shard, since
// the result of the operation must be drained from errc, allowing the
// shard to be able to handle new operations (#2932).

// we need to make sure that the forever loop in shard.go can
// always return due to shutdown in case this goroutine goes away.

// Write stores a new blob and returns its location to be used as a reference
// It can be given to a Read call to return the stored blob.
func (s *Store) Write(ctx context.Context, data []byte) (loc Location, err error) {
	_ = "STUB: not implemented"
	return *new(Location), nil
}

// buffer the channel to avoid blocking in shard.process on quit or context done

// Release gives back the slot to the shard
// From here on the slot can be reused and overwritten
// Release is meant to be called when an entry in the upstream db is removed
// Note that releasing is not safe for obfuscating earlier content, since
// even after reuse, the slot may be used by a very short blob and leaves the
// rest of the old blob bytes untouched
func (s *Store) Release(ctx context.Context, loc Location) error {
	_ = "STUB: not implemented"
	return nil
}
