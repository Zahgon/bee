// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postage

import (
	"context"
	"errors"
	"io"
	"math/big"
	"sync"
	"time"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/storage"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "postage"

const (
	// blockThreshold is used to allow threshold no of blocks to be synced before a
	// batch is usable.
	blockThreshold = 10
)

const (
	// stampIssuerSaveInterval is how often dirty stamp issuers are flushed to disk.
	stampIssuerSaveInterval = time.Minute
)

var (
	// ErrNotFound is the error returned when issuer with given batch ID does not exist.
	ErrNotFound = errors.New("not found")
	// ErrNotUsable is the error returned when issuer with given batch ID is not usable.
	ErrNotUsable = errors.New("not usable")
)

// Service is the postage service interface.
type Service interface {
	Add(*StampIssuer) error
	StampIssuers() []*StampIssuer
	GetStampIssuer([]byte) (*StampIssuer, func() error, error)
	IssuerUsable(*StampIssuer) bool
	BatchEventListener
	BatchExpiryHandler
	io.Closer
}

// service handles postage batches
// stores the active batches.
type service struct {
	logger       log.Logger
	mtx          sync.Mutex
	store        storage.Store
	postageStore Storer
	chainID      int64
	issuers      []*StampIssuer

	quit chan struct{}
	done chan struct{}
}

// NewService constructs a new Service. wasClean indicates whether the previous
// shutdown was graceful; if false, bucket counts are recovered from the stamp store.
func NewService(logger log.Logger, store storage.Store, postageStore Storer, chainID int64, wasClean bool) (Service, error) {
	_ = "STUB: not implemented"
	return *new(Service), nil
}

func (s *service) recoverBuckets() error { _ = "STUB: not implemented"; return nil }

func (s *service) run() {
	_ = "STUB: not implemented"

	// using 1 minute to significantly reduce disk writes
	return
}

// Add adds a stamp issuer to the active issuers.
func (ps *service) Add(st *StampIssuer) error { _ = "STUB: not implemented"; return nil }

// HandleCreate implements the BatchEventListener interface. This is fired on receiving
// a batch creation event from the blockchain listener to ensure that if a stamp
// issuer was not created initially, we will create it here.
func (ps *service) HandleCreate(b *Batch, amount *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// HandleTopUp implements the BatchEventListener interface. This is fired on receiving
// a batch topup event from the blockchain to update stampissuer details
func (ps *service) HandleTopUp(batchID []byte, amount *big.Int) { _ = "STUB: not implemented"; return }

func (ps *service) HandleDepthIncrease(batchID []byte, newDepth uint8) {
	_ = "STUB: not implemented"
	return
}

// StampIssuers returns the currently active stamp issuers.
func (ps *service) StampIssuers() []*StampIssuer { _ = "STUB: not implemented"; return nil }

func (ps *service) IssuerUsable(st *StampIssuer) bool { _ = "STUB: not implemented"; return false }

// this checks at least threshold blocks are seen on the blockchain after
// the batch creation, before we start using a stamp issuer. The threshold
// is meant to allow enough time for upstream peers to see the batch and
// hence validate the stamps issued

// GetStampIssuer finds a stamp issuer by batch ID.
func (ps *service) GetStampIssuer(batchID []byte) (*StampIssuer, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// save persists the specified stamp issuer to the stamperstore.
func (ps *service) save(st *StampIssuer) error { _ = "STUB: not implemented"; return nil }

func (ps *service) Close() error { _ = "STUB: not implemented"; return nil }

// HandleStampExpiry handles stamp expiry for a given id.
func (ps *service) HandleStampExpiry(ctx context.Context, id []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// removeStampItems removes all stamp items belonging to the given batch.
func (ps *service) removeStampItems(ctx context.Context, batchID []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// SetExpired removes all expired batches from the stamp issuers.
func (ps *service) removeIssuer(batchID []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// add adds a stamp issuer to the active issuers and returns false if it is already present.
// Must be mutex locked before usage.
func (ps *service) add(st *StampIssuer) bool { _ = "STUB: not implemented"; return false }
