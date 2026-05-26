// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pusher provides protocol-orchestrating functionality
// over the pushsync protocol. It makes sure that chunks meant
// to be distributed over the network are sent used using the
// pushsync protocol.
package pusher

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/pushsync"
	"github.com/ethersphere/bee/v2/pkg/stabilization"
	storage "github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/opentracing/opentracing-go"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "pusher"

type Op struct {
	Chunk  swarm.Chunk
	Err    chan error
	Direct bool
	Span   opentracing.Span

	identityAddress swarm.Address
}

type OpChan <-chan *Op

type Storer interface {
	storage.PushReporter
	storage.PushSubscriber
	ReservePutter() storage.Putter
}

type Service struct {
	networkID         uint64
	storer            Storer
	pushSyncer        pushsync.PushSyncer
	batchExist        postage.BatchExist
	logger            log.Logger
	metrics           metrics
	quit              chan struct{}
	chunksWorkerQuitC chan struct{}
	inflight          *inflight
	attempts          *attempts
	smuggler          chan OpChan
}

const (
	ConcurrentPushes  = swarm.Branches // how many chunks to push simultaneously
	DefaultRetryCount = 6
)

func New(
	networkID uint64,
	storer Storer,
	pushSyncer pushsync.PushSyncer,
	batchExist postage.BatchExist,
	logger log.Logger,
	startupStabilizer stabilization.Subscriber,
	retryCount int,
) *Service {
	_ = "STUB: not implemented"
	return nil
}

// chunksWorker is a loop that keeps looking for chunks that are locally uploaded ( by monitoring pushIndex )
// and pushes them to the closest peer and get a receipt.
func (s *Service) chunksWorker(startupStabilizer stabilization.Subscriber) {
	_ = "STUB: not implemented"
	return
}

// inflight.set handles the backpressure for the maximum amount of inflight chunks
// and duplicate handling.

// no peer was found which may mean that the node is suffering from connections issues
// we must slow down the pusher to prevent constant retries

func (s *Service) pushDeferred(ctx context.Context, logger log.Logger, op *Op) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// store the chunk

func (s *Service) pushDirect(ctx context.Context, logger log.Logger, op *Op) error {
	_ = "STUB: not implemented"
	return nil
}

// store the chunk

// out of attempts for retry, swallow error

func (s *Service) shallowReceipt(idAddress swarm.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Service) AddFeed(c <-chan *Op) { _ = "STUB: not implemented"; return }

func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }

// Wait for chunks worker to finish
