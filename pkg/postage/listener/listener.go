// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package listener

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/util/syncutil"
	"github.com/prometheus/client_golang/prometheus"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "listener"

const (
	blockPage          = 5000      // how many blocks to sync every time we page
	blockPageSnapshot  = 50000     // how many blocks to sync every time from snapshot
	tailSize           = 4         // how many blocks to tail from the tip of the chain
	defaultBatchFactor = uint64(5) // minimal number of blocks to sync at once
)

// for testing, set externally
var batchFactorOverridePublic = "5"

var (
	ErrPostageSyncingStalled = errors.New("postage syncing stalled")
	ErrPostagePaused         = errors.New("postage contract is paused")
	ErrParseSnapshot         = errors.New("failed to parse snapshot data")
)

type BlockHeightContractFilterer interface {
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	BlockNumber(context.Context) (uint64, error)
}

type listener struct {
	logger    log.Logger
	ev        BlockHeightContractFilterer
	blockTime time.Duration

	postageStampContractAddress common.Address
	postageStampContractABI     abi.ABI
	quit                        chan struct{}
	wg                          sync.WaitGroup
	metrics                     metrics
	stallingTimeout             time.Duration
	backoffTime                 time.Duration
	syncingStopped              *syncutil.Signaler

	// Cached postage stamp contract event topics.
	batchCreatedTopic       common.Hash
	batchTopUpTopic         common.Hash
	batchDepthIncreaseTopic common.Hash
	priceUpdateTopic        common.Hash
	pausedTopic             common.Hash
}

func New(
	syncingStopped *syncutil.Signaler,
	logger log.Logger,
	ev BlockHeightContractFilterer,
	postageStampContractAddress common.Address,
	postageStampContractABI abi.ABI,
	blockTime time.Duration,
	stallingTimeout time.Duration,
	backoffTime time.Duration,
) postage.Listener {
	_ = "STUB: not implemented"
	return *new(postage.Listener)
}

func (l *listener) filterQuery(from, to *big.Int) ethereum.FilterQuery {
	_ = "STUB: not implemented"
	return *new(ethereum.FilterQuery)
}

func (l *listener) processEvent(e types.Log, updater postage.EventUpdater) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *listener) Listen(ctx context.Context, from uint64, updater postage.EventUpdater) <-chan error {
	_ = "STUB: not implemented"
	return nil
}

// if we have a zero value batch - silence & log then move on

// Type assertion to detect if backend is SnapshotLogFilterer

// if for whatever reason we are stuck for too long we terminate
// this can happen because of rpc errors but also because of a stalled backend node
// this does not catch the case were a backend node is actively syncing but not caught up

// if we have a last blocknumber from the backend we can make a good estimate on when we need to requery
// otherwise we just use the backoff time

// in a test blockchain there might be not be enough blocks yet

// consider to-tailSize as the "latest" block we need to sync to

// round down to the largest multiple of batchFactor

// if the blockNumber is actually less than what we already, it might mean the backend is not synced or some reorg scenario

// do some paging (sub-optimal)

// Context cancelled is returned on shutdown, therefore we do nothing here.

// trigger shutdown in start.go

func (l *listener) Close() error { _ = "STUB: not implemented"; return nil }

type batchCreatedEvent struct {
	BatchId           [32]byte
	TotalAmount       *big.Int
	NormalisedBalance *big.Int
	Owner             common.Address
	Depth             uint8
	BucketDepth       uint8
	ImmutableFlag     bool
}

type batchTopUpEvent struct {
	BatchId           [32]byte
	TopupAmount       *big.Int
	NormalisedBalance *big.Int
}

type batchDepthIncreaseEvent struct {
	BatchId           [32]byte
	NewDepth          uint8
	NormalisedBalance *big.Int
}

type priceUpdateEvent struct {
	Price *big.Int
}

func totalTimeMetric(metric prometheus.Counter, start time.Time) { _ = "STUB: not implemented"; return }
