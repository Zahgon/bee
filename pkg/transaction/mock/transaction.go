// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/transaction"
)

type transactionServiceMock struct {
	send                 func(ctx context.Context, request *transaction.TxRequest, boost int) (txHash common.Hash, err error)
	waitForReceipt       func(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error)
	watchSentTransaction func(txHash common.Hash) (chan types.Receipt, chan error, error)
	call                 func(ctx context.Context, request *transaction.TxRequest) (result []byte, err error)
	pendingTransactions  func() ([]common.Hash, error)
	resendTransaction    func(ctx context.Context, txHash common.Hash) error
	storedTransaction    func(txHash common.Hash) (*transaction.StoredTransaction, error)
	cancelTransaction    func(ctx context.Context, originalTxHash common.Hash) (common.Hash, error)
	transactionFee       func(ctx context.Context, txHash common.Hash) (*big.Int, error)
}

func (m *transactionServiceMock) Send(ctx context.Context, request *transaction.TxRequest, boostPercent int) (txHash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (m *transactionServiceMock) WaitForReceipt(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *transactionServiceMock) WatchSentTransaction(txHash common.Hash) (<-chan types.Receipt, <-chan error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (m *transactionServiceMock) Call(ctx context.Context, request *transaction.TxRequest) (result []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *transactionServiceMock) PendingTransactions() ([]common.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *transactionServiceMock) ResendTransaction(ctx context.Context, txHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *transactionServiceMock) StoredTransaction(txHash common.Hash) (*transaction.StoredTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *transactionServiceMock) CancelTransaction(ctx context.Context, originalTxHash common.Hash) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (m *transactionServiceMock) Close() error {
	_ = "STUB: not implemented"

	// TransactionFee returns fee of transaction
	return nil
}

func (m *transactionServiceMock) TransactionFee(ctx context.Context, txHash common.Hash) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *transactionServiceMock) UnwrapABIError(_ context.Context, _ *transaction.TxRequest, err error, _ map[string]abi.Error) error {
	_ = "STUB: not implemented"

	// Option is the option passed to the mock Chequebook service
	return nil
}

type Option interface {
	apply(*transactionServiceMock)
}

type optionFunc func(*transactionServiceMock)

func (f optionFunc) apply(r *transactionServiceMock) { _ = "STUB: not implemented"; return }

func WithSendFunc(f func(context.Context, *transaction.TxRequest, int) (txHash common.Hash, err error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithWaitForReceiptFunc(f func(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCallFunc(f func(ctx context.Context, request *transaction.TxRequest) (result []byte, err error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithStoredTransactionFunc(f func(txHash common.Hash) (*transaction.StoredTransaction, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPendingTransactionsFunc(f func() ([]common.Hash, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithResendTransactionFunc(f func(ctx context.Context, txHash common.Hash) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCancelTransactionFunc(f func(ctx context.Context, originalTxHash common.Hash) (common.Hash, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTransactionFeeFunc(f func(ctx context.Context, txHash common.Hash) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func New(opts ...Option) transaction.Service {
	_ = "STUB: not implemented"
	return *new(transaction.Service)
}

type Call struct {
	abi    *abi.ABI
	to     common.Address
	result []byte
	method string
	params []any
}

func ABICall(abi *abi.ABI, to common.Address, result []byte, method string, params ...any) Call {
	_ = "STUB: not implemented"
	return *new(Call)
}

func WithABICallSequence(calls ...Call) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithABICall(abi *abi.ABI, to common.Address, result []byte, method string, params ...any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithABISend(abi *abi.ABI, txHash common.Hash, expectedAddress common.Address, expectedValue *big.Int, method string, params ...any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
