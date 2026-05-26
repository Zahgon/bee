// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transaction

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/transaction/backend"
)

// Backend is the minimum of blockchain backend functions we need.
type Backend interface {
	backend.Geth
	SuggestedFeeAndTip(ctx context.Context, gasPrice *big.Int, boostPercent int) (*big.Int, *big.Int, error)
}

// IsSynced will check if we are synced with the given blockchain backend. This
// is true if the current wall clock is after the block time of last block
// with the given maxDelay as the maximum duration we can be behind the block
// time.
func IsSynced(ctx context.Context, backend Backend, maxDelay time.Duration) (bool, time.Time, error) {
	_ = "STUB: not implemented"
	return false, *new(time.Time), nil
}

// WaitSynced will wait until we are synced with the given blockchain backend,
// with the given maxDelay duration as the maximum time we can be behind the
// last block.
func WaitSynced(ctx context.Context, logger log.Logger, backend Backend, maxDelay time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitBlockAfterTransaction(ctx context.Context, backend Backend, pollingInterval time.Duration, txHash common.Hash, additionalConfirmations uint64) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// in the case where we cannot find the block even though we already saw a higher number we keep on trying
