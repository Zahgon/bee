// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package backendsimulation

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/transaction"
)

var ErrNotImplemented = errors.New("not implemented")

type AccountAtKey struct {
	BlockNumber uint64
	Account     common.Address
}

type simulatedBackend struct {
	blockNumber uint64

	receipts map[common.Hash]*types.Receipt
	noncesAt map[AccountAtKey]uint64

	blocks []Block
	step   uint64
}

type Block struct {
	Number   uint64
	Receipts map[common.Hash]*types.Receipt
	NoncesAt map[AccountAtKey]uint64
}

type Option interface {
	apply(*simulatedBackend)
}

type optionFunc func(*simulatedBackend)

func (f optionFunc) apply(r *simulatedBackend) { _ = "STUB: not implemented"; return }

func WithBlocks(blocks ...Block) Option { _ = "STUB: not implemented"; return *new(Option) }

func New(options ...Option) transaction.Backend {
	_ = "STUB: not implemented"
	return *new(transaction.Backend)
}

func (m *simulatedBackend) advanceBlock() { _ = "STUB: not implemented"; return }

func (*simulatedBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *simulatedBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *simulatedBackend) SuggestedFeeAndTip(ctx context.Context, gasPrice *big.Int, boostPercent int) (*big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (m *simulatedBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *simulatedBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

func (*simulatedBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *simulatedBackend) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *simulatedBackend) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *simulatedBackend) BlockNumber(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *simulatedBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *simulatedBackend) BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *simulatedBackend) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *simulatedBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *simulatedBackend) ChainID(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *simulatedBackend) Close() { _ = "STUB: not implemented"; return }
