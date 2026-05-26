// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mockstorer

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/storer"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type chunksResponse struct {
	chunks []*storer.BinC
	err    error
}

// WithSubscribeResp mocks a desired response when calling IntervalChunks method.
// Different possible responses for subsequent responses in multi-call scenarios
// are possible (i.e. first call yields a,b,c, second call yields d,e,f).
// Mock maintains state of current call using chunksCalls counter.
func WithSubscribeResp(chunks []*storer.BinC, err error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithChunks mocks the set of chunks that the store is aware of (used in Get and Has calls).
func WithChunks(chs ...swarm.Chunk) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEvilChunk allows to inject a malicious chunk (request a certain address
// of a chunk, but get another), in order to mock unsolicited chunk delivery.
func WithEvilChunk(addr swarm.Address, ch swarm.Chunk) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCursors(c []uint64, e uint64) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCursorsErr(e error) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRadius(r uint8) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReserveSize(s int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCapacityDoubling(s int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPutHook(f func(swarm.Chunk) error) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSample(s storer.Sample) Option { _ = "STUB: not implemented"; return *new(Option) }

var _ storer.ReserveStore = (*ReserveStore)(nil)

type ReserveStore struct {
	mtx         sync.Mutex
	chunksCalls int
	putCalls    int
	setCalls    int

	chunks    map[string]swarm.Chunk
	evilAddr  swarm.Address
	evilChunk swarm.Chunk

	cursors    []uint64
	cursorsErr error
	epoch      uint64

	radius           uint8
	reservesize      int
	capacityDoubling int

	subResponses []chunksResponse
	putHook      func(swarm.Chunk) error

	sample storer.Sample
}

// NewReserve returns a new Reserve mock.
func NewReserve(opts ...Option) *ReserveStore { _ = "STUB: not implemented"; return nil }

func (s *ReserveStore) EvictBatch(ctx context.Context, batchID []byte) error {
	_ = "STUB: not implemented"
	return nil
}
func (s *ReserveStore) IsWithinStorageRadius(addr swarm.Address) bool {
	_ = "STUB: not implemented"
	return false
}
func (s *ReserveStore) IsFullySynced() bool  { _ = "STUB: not implemented"; return false }
func (s *ReserveStore) StorageRadius() uint8 { _ = "STUB: not implemented"; return 0 }

func (s *ReserveStore) SetStorageRadius(r uint8) { _ = "STUB: not implemented"; return }

func (s *ReserveStore) CommittedDepth() uint8 { _ = "STUB: not implemented"; return 0 }

func (s *ReserveStore) CapacityDoubling() uint8 { _ = "STUB: not implemented"; return 0 }

// IntervalChunks returns a set of chunk in a requested interval.
func (s *ReserveStore) SubscribeBin(ctx context.Context, bin uint8, start uint64) (<-chan *storer.BinC, func(), <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *ReserveStore) ReserveSize() int { _ = "STUB: not implemented"; return 0 }

func (s *ReserveStore) ReserveLastBinIDs() (curs []uint64, epoch uint64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// PutCalls returns the amount of times Put was called.
func (s *ReserveStore) PutCalls() int { _ = "STUB: not implemented"; return 0 }

// SetCalls returns the amount of times Set was called.
func (s *ReserveStore) SetCalls() int { _ = "STUB: not implemented"; return 0 }

// Get chunks.
func (s *ReserveStore) ReserveGet(ctx context.Context, addr swarm.Address, batchID []byte, stampHash []byte) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// inject the malicious chunk instead

// Put chunks.
func (s *ReserveStore) ReservePutter() storage.Putter {
	_ = "STUB: not implemented"
	return *new(storage.Putter)
}

// Put chunks.
func (s *ReserveStore) put(_ context.Context, chs ...swarm.Chunk) error {
	_ = "STUB: not implemented"
	return nil
}

// Has chunks.
func (s *ReserveStore) ReserveHas(addr swarm.Address, batchID []byte, stampHash []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *ReserveStore) ReserveSample(context.Context, []byte, uint8, uint64, *big.Int) (storer.Sample, error) {
	_ = "STUB: not implemented"
	return *new(storer.Sample), nil
}

type Option interface {
	apply(*ReserveStore)
}
type optionFunc func(*ReserveStore)

func (f optionFunc) apply(r *ReserveStore) { _ = "STUB: not implemented"; return }
