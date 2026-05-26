// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transaction

import (
	"context"
	"errors"
	"io"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/storage"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "transaction"

const (
	storedTransactionPrefix  = "transaction_stored_"
	pendingTransactionPrefix = "transaction_pending_"
)

var (
	// ErrTransactionReverted denotes that the sent transaction has been
	// reverted.
	ErrTransactionReverted = errors.New("transaction reverted")
	ErrUnknownTransaction  = errors.New("unknown transaction")
	ErrAlreadyImported     = errors.New("already imported")
)

const (
	DefaultGasLimit        = 1_000_000 // Used for contract operations when setGasLimit flag is enabled
	DefaultTipBoostPercent = 25
	MaxGasLimit            = 10_000_000 // Maximum allowed gas limit to prevent excessive values
	MinGasLimit            = 21_000     // Minimum gas for any transaction
	GasBufferPercent       = 33         // Add 33% buffer to estimated gas
	FallbackGasLimit       = 500_000    // Fallback when estimation fails and no minimum is set
)

// TxRequest describes a request for a transaction that can be executed.
type TxRequest struct {
	To                   *common.Address // recipient of the transaction
	Data                 []byte          // transaction data
	GasPrice             *big.Int        // gas price or nil if suggested gas price should be used
	GasLimit             uint64          // gas limit or 0 if it should be estimated
	MinEstimatedGasLimit uint64          // minimum gas limit to use if the gas limit was estimated; it will not apply when this value is 0 or when GasLimit is not 0
	GasFeeCap            *big.Int        // adds a cap to maximum fee user is willing to pay
	Value                *big.Int        // amount of wei to send
	Description          string          // optional description
}

type StoredTransaction struct {
	To          *common.Address // recipient of the transaction
	Data        []byte          // transaction data
	GasPrice    *big.Int        // used gas price
	GasLimit    uint64          // used gas limit
	GasTipBoost int             // adds a tip for the miner for prioritizing transaction
	GasTipCap   *big.Int        // adds a cap to the tip
	GasFeeCap   *big.Int        // adds a cap to maximum fee user is willing to pay
	Value       *big.Int        // amount of wei to send
	Nonce       uint64          // used nonce
	Created     int64           // creation timestamp
	Description string          // description
}

// Service is the service to send transactions. It takes care of gas price, gas
// limit and nonce management.
type Service interface {
	io.Closer
	// Send creates a transaction based on the request (with gasprice increased by provided percentage) and sends it.
	Send(ctx context.Context, request *TxRequest, tipCapBoostPercent int) (txHash common.Hash, err error)
	// Call simulate a transaction based on the request.
	Call(ctx context.Context, request *TxRequest) (result []byte, err error)
	// WaitForReceipt waits until either the transaction with the given hash has been mined or the context is cancelled.
	// This is only valid for transaction sent by this service.
	WaitForReceipt(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error)
	// WatchSentTransaction start watching the given transaction.
	// This wraps the monitors watch function by loading the correct nonce from the store.
	// This is only valid for transaction sent by this service.
	WatchSentTransaction(txHash common.Hash) (<-chan types.Receipt, <-chan error, error)
	// StoredTransaction retrieves the stored information for the transaction
	StoredTransaction(txHash common.Hash) (*StoredTransaction, error)
	// PendingTransactions retrieves the list of all pending transaction hashes
	PendingTransactions() ([]common.Hash, error)
	// ResendTransaction resends a previously sent transaction
	// This operation can be useful if for some reason the transaction vanished from the eth networks pending pool
	ResendTransaction(ctx context.Context, txHash common.Hash) error
	// CancelTransaction cancels a previously sent transaction by double-spending its nonce with zero-transfer one
	CancelTransaction(ctx context.Context, originalTxHash common.Hash) (common.Hash, error)
	// TransactionFee retrieves the transaction fee
	TransactionFee(ctx context.Context, txHash common.Hash) (*big.Int, error)
	// UnwrapABIError tries to unwrap the ABI error if the given error is not nil.
	// The original error is wrapped together with the ABI error if it exists.
	UnwrapABIError(ctx context.Context, req *TxRequest, err error, abiErrors map[string]abi.Error) error
}

type transactionService struct {
	wg     sync.WaitGroup
	lock   sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc

	logger           log.Logger
	backend          Backend
	signer           crypto.Signer
	sender           common.Address
	store            storage.StateStorer
	chainID          *big.Int
	monitor          Monitor
	fallbackGasLimit uint64
}

// NewService creates a new transaction service.
func NewService(logger log.Logger, overlayEthAddress common.Address, backend Backend, signer crypto.Signer, store storage.StateStorer, chainID *big.Int, monitor Monitor, fallbackGasLimit uint64) (Service, error) {
	_ = "STUB: not implemented"
	return *new(Service), nil
}

func (t *transactionService) waitForAllPendingTx() error { _ = "STUB: not implemented"; return nil }

// Send creates and signs a transaction based on the request and sends it.
func (t *transactionService) Send(ctx context.Context, request *TxRequest, boostPercent int) (txHash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (t *transactionService) waitForPendingTx(txHash common.Hash) {
	_ = "STUB: not implemented"
	return
}

func (t *transactionService) Call(ctx context.Context, request *TxRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *transactionService) StoredTransaction(txHash common.Hash) (*StoredTransaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// prepareTransaction creates a signable transaction based on a request.
func (t *transactionService) prepareTransaction(ctx context.Context, request *TxRequest, nonce uint64, boostPercent int) (tx *types.Transaction, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Estimate gas using pending state for consistency with PendingNonceAt

// Contract call - use configured fallback

// Simple transfer - use minimum

// Estimation succeeded - add buffer for state changes

// Apply minimum if specified

// Cap at maximum

// Ensure absolute minimum

// Use provided gas limit with bounds validation

/*
	Transactions are EIP 1559 dynamic transactions where there are three fee related fields:
		1. base fee is the price that will be burned as part of the transaction.
		2. max fee is the max price we are willing to spend as gas price.
		3. max priority fee is max price want to give to the miner to prioritize the transaction.
	as an example:
	if base fee is 15, max fee is 20, and max priority is 3, gas price will be 15 + 3 = 18
	if base is 15, max fee is 20, and max priority fee is 10,
	gas price will be 15 + 10 = 25, but since 25 > 20, gas price is 20.
	notice that gas price does not exceed 20 as defined by max fee.
*/

func storedTransactionKey(txHash common.Hash) string { _ = "STUB: not implemented"; return "" }

func pendingTransactionKey(txHash common.Hash) string { _ = "STUB: not implemented"; return "" }

func (t *transactionService) nextNonce(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// PendingNonceAt returns the nonce we should use, but we will
// compare this to our pending tx list, therefore the -1.

// WaitForReceipt waits until either the transaction with the given hash has
// been mined or the context is cancelled.
func (t *transactionService) WaitForReceipt(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// don't wait longer than the context that was passed in

func (t *transactionService) WatchSentTransaction(txHash common.Hash) (<-chan types.Receipt, <-chan error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// loading the tx here guarantees it was in fact sent from this transaction service
// also it allows us to avoid having to load the transaction during the watch loop

func (t *transactionService) PendingTransactions() ([]common.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filterPendingTransactions will filter supplied transaction hashes removing those that are not pending anymore.
// Removed transactions will be also removed from store.
// Returns the pending transactions keyed by hash.
func (t *transactionService) filterPendingTransactions(ctx context.Context, txHashes []common.Hash) map[common.Hash]*types.Transaction {
	_ = "STUB: not implemented"
	return nil
}

// When error occurres consider transaction as pending (so this transaction won't be filtered out),
// unless it was not found

func (t *transactionService) ResendTransaction(ctx context.Context, txHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *transactionService) CancelTransaction(ctx context.Context, originalTxHash common.Hash) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (t *transactionService) Close() error { _ = "STUB: not implemented"; return nil }

func (t *transactionService) TransactionFee(ctx context.Context, txHash common.Hash) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *transactionService) UnwrapABIError(ctx context.Context, req *TxRequest, err error, abiErrors map[string]abi.Error) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errorlint
