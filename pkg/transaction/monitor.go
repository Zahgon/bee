// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transaction

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/log"
)

var (
	ErrTransactionCancelled = errors.New("transaction cancelled")
	ErrMonitorClosed        = errors.New("monitor closed")
)

// Monitor is a nonce-based watcher for transaction confirmations.
// Instead of watching transactions individually, the senders nonce is monitored and transactions are checked based on this.
// The idea is that if the nonce is still lower than that of a pending transaction, there is no point in actually checking the transaction for a receipt.
// At the same time if the nonce was already used and this was a few blocks ago we can reasonably assume that it will never confirm.
type Monitor interface {
	io.Closer
	// WatchTransaction watches the transaction until either there is 1 confirmation or a competing transaction with cancellationDepth confirmations.
	WatchTransaction(txHash common.Hash, nonce uint64) (<-chan types.Receipt, <-chan error, error)
}
type transactionMonitor struct {
	lock       sync.Mutex
	ctx        context.Context    // context which is used for all backend calls
	cancelFunc context.CancelFunc // function to cancel the above context
	wg         sync.WaitGroup

	logger  log.Logger
	backend Backend
	sender  common.Address // sender of transactions which this instance can monitor

	pollingInterval   time.Duration // time between checking for new blocks
	cancellationDepth uint64        // number of blocks until considering a tx cancellation final

	watchesByNonce map[uint64]map[common.Hash][]transactionWatch // active watches grouped by nonce and tx hash
	watchAdded     chan struct{}                                 // channel to trigger instant pending check
}

type transactionWatch struct {
	start    time.Time
	receiptC chan types.Receipt // channel to which the receipt will be written once available
	errC     chan error         // error channel (primarily for cancelled transactions)
}

func NewMonitor(logger log.Logger, backend Backend, sender common.Address, pollingInterval time.Duration, cancellationDepth uint64) Monitor {
	_ = "STUB: not implemented"
	return *new(Monitor)
}

func (tm *transactionMonitor) WatchTransaction(txHash common.Hash, nonce uint64) (<-chan types.Receipt, <-chan error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// these channels will be written to at most once
// buffer size is 1 to avoid blocking in the watch loop

// main watch loop
func (tm *transactionMonitor) watchPending() { _ = "STUB: not implemented"; return }

// flag if this iteration was triggered by the watchAdded channel

// if a new watch has been added check again without waiting

// otherwise wait

// if the main context is cancelled terminate

// if there are no watched transactions there is nothing to do

// switch to new head subscriptions once websockets are the norm

// if the block number is not higher than before there is nothing todo
// unless a watch was added in which case we will do the check anyway
// in the rare case where a block was reorged and the new one is the first to contain our tx we wait an extra block

func (tm *transactionMonitor) hasWatches() bool { _ = "STUB: not implemented"; return false }

func watchStart(watches []transactionWatch) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// checkPending checks the given block for confirmed or cancelled transactions.
func (tm *transactionMonitor) checkPending(block uint64) error {
	_ = "STUB: not implemented"
	// Snapshot nonces and tx hashes to check (releases lock during slow RPC calls).
	return nil
}

// Check receipts without holding lock (RPC calls can be slow).

// Check for cancellations.

// Notify subscribers and cleanup.

func (tm *transactionMonitor) Close() error { _ = "STUB: not implemented"; return nil }
