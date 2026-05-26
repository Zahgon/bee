// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapped

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/transaction"
	"github.com/ethersphere/bee/v2/pkg/transaction/backend"
	"github.com/ethersphere/bee/v2/pkg/transaction/wrapped/cache"
)

var _ transaction.Backend = (*wrappedBackend)(nil)

type blockNumberAnchor struct {
	number    uint64
	timestamp time.Time
}

type wrappedBackend struct {
	backend           backend.Geth
	metrics           metrics
	minimumGasTipCap  int64
	blockTime         time.Duration
	blockSyncInterval uint64
	blockNumberCache  *cache.SingleFlightCache[blockNumberAnchor]
}

func NewBackend(
	backend backend.Geth,
	minimumGasTipCap uint64,
	blockTime time.Duration,
	blockSyncInterval uint64,
) transaction.Backend {
	_ = "STUB: not implemented"
	return *new(transaction.Backend)
}

func (b *wrappedBackend) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *wrappedBackend) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (b *wrappedBackend) BlockNumber(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *wrappedBackend) estimatedBlockNumber(anchor blockNumberAnchor, now time.Time) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (b *wrappedBackend) estimatedBlockNumberWithElapsed(anchor blockNumberAnchor, now time.Time) (uint64, uint64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (b *wrappedBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *wrappedBackend) BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *wrappedBackend) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *wrappedBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *wrappedBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *wrappedBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *wrappedBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *wrappedBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *wrappedBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *wrappedBackend) ChainID(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *wrappedBackend) Close() { _ = "STUB: not implemented"; return }
