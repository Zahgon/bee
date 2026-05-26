// Copyright 2025 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package backendnoop

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/transaction"
	"github.com/prometheus/client_golang/prometheus"
)

var _ transaction.Backend = (*Backend)(nil)

// Backend is a no-op implementation for transaction.Backend interface.
// It's used when the blockchain functionality is disabled.
type Backend struct {
	chainID int64
}

// New creates a new no-op backend with the specified chain ID.
func New(chainID int64) transaction.Backend {
	_ = "STUB: not implemented"
	return *new(transaction.Backend)
}

func (b *Backend) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }

func (b *Backend) CallContract(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Backend) HeaderByNumber(context.Context, *big.Int) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Backend) PendingNonceAt(context.Context, common.Address) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *Backend) SuggestedFeeAndTip(ctx context.Context, gasPrice *big.Int, boostPercent int) (*big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (b *Backend) SuggestGasTipCap(context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Backend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *Backend) SendTransaction(context.Context, *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Backend) TransactionReceipt(context.Context, common.Hash) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Backend) TransactionByHash(context.Context, common.Hash) (tx *types.Transaction, isPending bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (b *Backend) BlockNumber(context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *Backend) BalanceAt(context.Context, common.Address, *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Backend) NonceAt(context.Context, common.Address, *big.Int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *Backend) FilterLogs(context.Context, ethereum.FilterQuery) ([]types.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Backend) ChainID(context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Backend) Close() { _ = "STUB: not implemented"; return }
