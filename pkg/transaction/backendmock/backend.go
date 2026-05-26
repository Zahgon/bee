// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package backendmock

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

type backendMock struct {
	callContract       func(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	sendTransaction    func(ctx context.Context, tx *types.Transaction) error
	suggestedFeeAndTip func(ctx context.Context, gasPrice *big.Int, boostPercent int) (*big.Int, *big.Int, error)
	suggestGasTipCap   func(ctx context.Context) (*big.Int, error)
	estimateGas        func(ctx context.Context, msg ethereum.CallMsg) (gas uint64, err error)
	transactionReceipt func(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	pendingNonceAt     func(ctx context.Context, account common.Address) (uint64, error)
	transactionByHash  func(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	blockNumber        func(ctx context.Context) (uint64, error)
	headerByNumber     func(ctx context.Context, number *big.Int) (*types.Header, error)
	balanceAt          func(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error)
	nonceAt            func(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
}

func (m *backendMock) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *backendMock) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *backendMock) SuggestedFeeAndTip(ctx context.Context, gasPrice *big.Int, boostPercent int) (*big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (m *backendMock) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *backendMock) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

func (*backendMock) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *backendMock) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *backendMock) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *backendMock) BlockNumber(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *backendMock) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *backendMock) BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *backendMock) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *backendMock) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *backendMock) ChainID(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *backendMock) Close() { _ = "STUB: not implemented"; return }

func New(opts ...Option) transaction.Backend {
	_ = "STUB: not implemented"
	return *new(transaction.Backend)
}

// Option is the option passed to the mock Chequebook service
type Option interface {
	apply(*backendMock)
}

type optionFunc func(*backendMock)

func (f optionFunc) apply(r *backendMock) { _ = "STUB: not implemented"; return }

func WithCallContractFunc(f func(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithBalanceAt(f func(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPendingNonceAtFunc(f func(ctx context.Context, account common.Address) (uint64, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSuggestedFeeAndTipFunc(f func(ctx context.Context, gasPrice *big.Int, boostPercent int) (*big.Int, *big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSuggestGasTipCapFunc(f func(ctx context.Context) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithEstimateGasFunc(f func(ctx context.Context, msg ethereum.CallMsg) (gas uint64, err error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTransactionReceiptFunc(f func(ctx context.Context, txHash common.Hash) (*types.Receipt, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTransactionByHashFunc(f func(ctx context.Context, txHash common.Hash) (*types.Transaction, bool, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSendTransactionFunc(f func(ctx context.Context, tx *types.Transaction) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithBlockNumberFunc(f func(context.Context) (uint64, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithHeaderbyNumberFunc(f func(ctx context.Context, number *big.Int) (*types.Header, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithNonceAtFunc(f func(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
