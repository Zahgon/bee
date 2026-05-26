// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mockstorer

import (
	"context"
	"sync"
	"time"

	"github.com/ethersphere/bee/v2/pkg/pusher"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/storer"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"go.uber.org/atomic"
)

// now returns the current time.Time; used in testing.
var now = time.Now

type mockStorer struct {
	chunkStore     storage.ChunkStore
	mu             sync.Mutex
	pins           []swarm.Address
	sessionID      atomic.Uint64
	activeSessions map[uint64]*storer.SessionInfo
	chunkPushC     chan *pusher.Op
	debugInfo      storer.Info

	storageRadius  uint8
	committedDepth uint8
}

type putterSession struct {
	chunkStore storage.Putter
	done       func(swarm.Address) error
}

func (p *putterSession) Put(ctx context.Context, ch swarm.Chunk) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *putterSession) Done(address swarm.Address) error { _ = "STUB: not implemented"; return nil }

func (p *putterSession) Cleanup() error {
	_ = "STUB: not implemented"

	// New returns a mock storer implementation that is designed to be used for the
	// unit tests.
	return nil
}

func New() *mockStorer { _ = "STUB: not implemented"; return nil }

func NewWithChunkStore(cs storage.ChunkStore) *mockStorer { _ = "STUB: not implemented"; return nil }

func NewWithDebugInfo(info storer.Info) *mockStorer { _ = "STUB: not implemented"; return nil }

func (m *mockStorer) Upload(_ context.Context, pin bool, tagID uint64) (storer.PutterSession, error) {
	_ = "STUB: not implemented"
	return *new(storer.PutterSession), nil
}

func (m *mockStorer) NewSession() (storer.SessionInfo, error) {
	_ = "STUB: not implemented"
	return *new(storer.SessionInfo), nil
}

func (m *mockStorer) Session(tagID uint64) (storer.SessionInfo, error) {
	_ = "STUB: not implemented"
	return *new(storer.SessionInfo), nil
}

func (m *mockStorer) DeleteSession(tagID uint64) error { _ = "STUB: not implemented"; return nil }

func (m *mockStorer) ListSessions(offset, limit int) ([]storer.SessionInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *mockStorer) DeletePin(_ context.Context, address swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mockStorer) Pins() ([]swarm.Address, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *mockStorer) HasPin(address swarm.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *mockStorer) NewCollection(ctx context.Context) (storer.PutterSession, error) {
	_ = "STUB: not implemented"
	return *new(storer.PutterSession), nil
}

func (m *mockStorer) Lookup() storage.Getter {
	_ = "STUB: not implemented"
	return *new(storage.Getter)
}

func (m *mockStorer) Cache() storage.Putter { _ = "STUB: not implemented"; return *new(storage.Putter) }

func (m *mockStorer) DirectUpload() storer.PutterSession {
	_ = "STUB: not implemented"
	return *new(storer.PutterSession)
}

func (m *mockStorer) Download(_ bool) storage.Getter {
	_ = "STUB: not implemented"
	return *new(storage.Getter)
}

func (m *mockStorer) PusherFeed() <-chan *pusher.Op { _ = "STUB: not implemented"; return nil }

func (m *mockStorer) ChunkStore() storage.ReadOnlyChunkStore {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyChunkStore)
}

func (m *mockStorer) StorageRadius() uint8 { _ = "STUB: not implemented"; return 0 }

func (m *mockStorer) CommittedDepth() uint8 { _ = "STUB: not implemented"; return 0 }

func (m *mockStorer) CapacityDoubling() uint8 { _ = "STUB: not implemented"; return 0 }

func (m *mockStorer) IsWithinStorageRadius(_ swarm.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *mockStorer) DebugInfo(_ context.Context) (storer.Info, error) {
	_ = "STUB: not implemented"
	return *new(storer.Info), nil
}

func (m *mockStorer) NeighborhoodsStat(ctx context.Context) ([]*storer.NeighborhoodStat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *mockStorer) Put(ctx context.Context, ch swarm.Chunk) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mockStorer) SetStorageRadius(radius uint8) { _ = "STUB: not implemented"; return }

func (m *mockStorer) SetCommittedDepth(depth uint8) { _ = "STUB: not implemented"; return }
