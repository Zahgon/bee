// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package batchservice

import (
	"context"
	"errors"
	"hash"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/storage"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "batchservice"

const (
	dirtyDBKey    = "batchservice_dirty_db"
	checksumDBKey = "batchservice_checksum"
)

var ErrZeroValueBatch = errors.New("low balance batch")

type batchService struct {
	stateStore    storage.StateStorer
	storer        postage.Storer
	logger        log.Logger
	listener      postage.Listener
	owner         []byte
	batchListener postage.BatchEventListener

	checksum hash.Hash // checksum hasher
	resync   bool

	pendingChainState *postage.ChainState
}

type Interface interface {
	postage.EventUpdater
}

// New will create a new BatchService.
func New(
	stateStore storage.StateStorer,
	storer postage.Storer,
	logger log.Logger,
	listener postage.Listener,
	owner []byte,
	batchListener postage.BatchEventListener,
	checksumFunc func() hash.Hash,
	resync bool,
) (Interface, error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

func (svc *batchService) getChainState() *postage.ChainState { _ = "STUB: not implemented"; return nil }

// Create will create a new batch with the given ID, owner value and depth and
// stores it in the BatchedStore.
func (svc *batchService) Create(id, owner []byte, totalAmout, normalisedBalance *big.Int, depth, bucketDepth uint8, immutable bool, txHash common.Hash) error {
	_ = "STUB: not implemented"
	// dont add batches which have value which equals total cumulative
	// payout or that are going to expire already within the next couple of blocks
	return nil
}

// don't do anything

// TopUp implements the EventUpdater interface. It tops ups a batch with the
// given ID with the given amount.
func (svc *batchService) TopUp(id []byte, totalAmout, normalisedBalance *big.Int, txHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateDepth implements the EventUpdater interface. It sets the new depth of a
// batch with the given ID.
func (svc *batchService) UpdateDepth(id []byte, depth uint8, normalisedBalance *big.Int, txHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdatePrice implements the EventUpdater interface. It sets the current
// price from the chain in the service chain state.
func (svc *batchService) UpdatePrice(price *big.Int, txHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (svc *batchService) UpdateBlockNumber(blockNumber uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (svc *batchService) TransactionStart() error { _ = "STUB: not implemented"; return nil }

func (svc *batchService) TransactionEnd() error { _ = "STUB: not implemented"; return nil }

var ErrInterruped = errors.New("postage sync interrupted")

func (svc *batchService) Start(ctx context.Context, startBlock uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// updateChecksum updates the batchservice checksum once an event gets
// processed. It swaps the existing checksum which is in the hasher
// with the new checksum and persists it in the statestore.
func (svc *batchService) updateChecksum(txHash common.Hash) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
