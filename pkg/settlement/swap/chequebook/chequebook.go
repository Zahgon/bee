// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chequebook

import (
	"context"
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/erc20"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/transaction"
	"github.com/ethersphere/bee/v2/pkg/util/abiutil"
	"github.com/ethersphere/go-sw3-abi/sw3abi"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "chequebook"

// SendChequeFunc is a function to send cheques.
type SendChequeFunc func(cheque *SignedCheque) error

const (
	lastIssuedChequeKeyPrefix = "swap_chequebook_last_issued_cheque_"
	totalIssuedKey            = "swap_chequebook_total_issued_"
)

var (
	// ErrOutOfFunds is the error when the chequebook has not enough free funds for a cheque
	ErrOutOfFunds = errors.New("chequebook out of funds")
	// ErrInsufficientFunds is the error when the chequebook has not enough free funds for a user action
	ErrInsufficientFunds = errors.New("insufficient token balance")

	chequebookABI          = abiutil.MustParseABI(sw3abi.ERC20SimpleSwapABIv0_6_9)
	chequeCashedEventType  = chequebookABI.Events["ChequeCashed"]
	chequeBouncedEventType = chequebookABI.Events["ChequeBounced"]
)

// Service is the main interface for interacting with the nodes chequebook.
type Service interface {
	// Deposit starts depositing erc20 token into the chequebook. This returns once the transactions has been broadcast.
	Deposit(ctx context.Context, amount *big.Int) (hash common.Hash, err error)
	// Withdraw starts withdrawing erc20 token from the chequebook. This returns once the transactions has been broadcast.
	Withdraw(ctx context.Context, amount *big.Int) (hash common.Hash, err error)
	// WaitForDeposit waits for the deposit transaction to confirm and verifies the result.
	WaitForDeposit(ctx context.Context, txHash common.Hash) error
	// Balance returns the token balance of the chequebook.
	Balance(ctx context.Context) (*big.Int, error)
	// AvailableBalance returns the token balance of the chequebook which is not yet used for uncashed cheques.
	AvailableBalance(ctx context.Context) (*big.Int, error)
	// Address returns the address of the used chequebook contract.
	Address() common.Address
	// Issue a new cheque for the beneficiary with an cumulativePayout amount higher than the last.
	Issue(ctx context.Context, beneficiary common.Address, amount *big.Int, sendChequeFunc SendChequeFunc) (*big.Int, error)
	// LastCheque returns the last cheque we issued for the beneficiary.
	LastCheque(beneficiary common.Address) (*SignedCheque, error)
	// LastCheques returns the last cheques for all beneficiaries.
	LastCheques() (map[common.Address]*SignedCheque, error)
}

type service struct {
	lock               sync.Mutex
	transactionService transaction.Service

	address      common.Address
	contract     *chequebookContract
	ownerAddress common.Address

	erc20Service erc20.Service

	store               storage.StateStorer
	chequeSigner        ChequeSigner
	totalIssuedReserved *big.Int
}

// New creates a new chequebook service for the provided chequebook contract.
func New(transactionService transaction.Service, address, ownerAddress common.Address, store storage.StateStorer, chequeSigner ChequeSigner, erc20Service erc20.Service) (Service, error) {
	_ = "STUB: not implemented"
	return *new(Service), nil
}

// Address returns the address of the used chequebook contract.
func (s *service) Address() common.Address {
	_ = "STUB: not implemented"

	// Deposit starts depositing erc20 token into the chequebook. This returns once the transactions has been broadcast.
	return *new(common.Address)
}

func (s *service) Deposit(ctx context.Context, amount *big.Int) (hash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// check we can afford this so we don't waste gas

// Balance returns the token balance of the chequebook.
func (s *service) Balance(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// AvailableBalance returns the token balance of the chequebook which is not yet used for uncashed cheques.
}

func (s *service) AvailableBalance(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// balance plus totalPaidOut is the total amount ever put into the chequebook (ignoring deposits and withdrawals which cancelled out)
// minus the total amount we issued from this chequebook this gives use the portion of the balance not covered by any cheques

// WaitForDeposit waits for the deposit transaction to confirm and verifies the result.
func (s *service) WaitForDeposit(ctx context.Context, txHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// lastIssuedChequeKey computes the key where to store the last cheque for a beneficiary.
func lastIssuedChequeKey(beneficiary common.Address) string { _ = "STUB: not implemented"; return "" }

func (s *service) reserveTotalIssued(ctx context.Context, amount *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *service) unreserveTotalIssued(amount *big.Int) { _ = "STUB: not implemented"; return }

// Issue issues a new cheque and passes it to sendChequeFunc.
// The cheque is considered sent and saved when sendChequeFunc succeeds.
// The available balance which is available after sending the cheque is passed
// to the caller for it to be communicated over metrics.
func (s *service) Issue(ctx context.Context, beneficiary common.Address, amount *big.Int, sendChequeFunc SendChequeFunc) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// increase cumulativePayout by amount

// create and sign the new cheque

// actually send the check before saving to avoid double payment

// returns the total amount in cheques issued so far
func (s *service) totalIssued() (*big.Int, error) { _ = "STUB: not implemented"; return nil, nil }

// LastCheque returns the last cheque we issued for the beneficiary.
func (s *service) LastCheque(beneficiary common.Address) (*SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func keyBeneficiary(key []byte, prefix string) (beneficiary common.Address, err error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// LastCheques returns the last cheques for all beneficiaries.
func (s *service) LastCheques() (map[common.Address]*SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *service) Withdraw(ctx context.Context, amount *big.Int) (hash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// check we can afford this so we don't waste gas and don't risk bouncing cheques
