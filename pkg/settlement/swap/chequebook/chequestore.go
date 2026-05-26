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
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/transaction"
)

const (
	// prefix for the persistence key
	lastReceivedChequePrefix = "swap_chequebook_last_received_cheque_"
)

var (
	// ErrNoCheque is the error returned if there is no prior cheque for a chequebook or beneficiary.
	ErrNoCheque = errors.New("no cheque")
	// ErrChequeNotIncreasing is the error returned if the cheque amount is the same or lower.
	ErrChequeNotIncreasing = errors.New("cheque cumulativePayout is not increasing")
	// ErrChequeInvalid is the error returned if the cheque itself is invalid.
	ErrChequeInvalid = errors.New("invalid cheque")
	// ErrWrongBeneficiary is the error returned if the cheque has the wrong beneficiary.
	ErrWrongBeneficiary = errors.New("wrong beneficiary")
	// ErrBouncingCheque is the error returned if the chequebook is demonstrably illiquid.
	ErrBouncingCheque = errors.New("bouncing cheque")
	// ErrChequeValueTooLow is the error returned if the after deduction value of a cheque did not cover 1 accounting credit
	ErrChequeValueTooLow = errors.New("cheque value lower than acceptable")
)

// ChequeStore handles the verification and storage of received cheques
type ChequeStore interface {
	// ReceiveCheque verifies and stores a cheque. It returns the total amount earned.
	ReceiveCheque(ctx context.Context, cheque *SignedCheque, exchangeRate, deduction *big.Int) (*big.Int, error)
	// LastCheque returns the last cheque we received from a specific chequebook.
	LastCheque(chequebook common.Address) (*SignedCheque, error)
	// LastCheques returns the last received cheques from every known chequebook.
	LastCheques() (map[common.Address]*SignedCheque, error)
}

type chequeStore struct {
	lock               sync.Mutex
	store              storage.StateStorer
	factory            Factory
	chaindID           int64
	transactionService transaction.Service
	beneficiary        common.Address // the beneficiary we expect in cheques sent to us
	recoverChequeFunc  RecoverChequeFunc
}

type RecoverChequeFunc func(cheque *SignedCheque, chainID int64) (common.Address, error)

// NewChequeStore creates new ChequeStore
func NewChequeStore(
	store storage.StateStorer,
	factory Factory,
	chainID int64,
	beneficiary common.Address,
	transactionService transaction.Service,
	recoverChequeFunc RecoverChequeFunc,
) ChequeStore {
	_ = "STUB: not implemented"
	return *new(ChequeStore)
}

// lastReceivedChequeKey computes the key where to store the last cheque received from a chequebook.
func lastReceivedChequeKey(chequebook common.Address) string { _ = "STUB: not implemented"; return "" }

// LastCheque returns the last cheque we received from a specific chequebook.
func (s *chequeStore) LastCheque(chequebook common.Address) (*SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReceiveCheque verifies and stores a cheque. It returns the totam amount earned.
func (s *chequeStore) ReceiveCheque(ctx context.Context, cheque *SignedCheque, exchangeRate, deduction *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	// verify we are the beneficiary
	return nil, nil
}

// don't allow concurrent processing of cheques
// this would be sufficient on a per chequebook basis

// load the lastCumulativePayout for the cheques chequebook

// if this is the first cheque from this chequebook, verify with the factory.

// check this cheque is actually increasing in value

// blockchain calls below

// this does not change for the same chequebook

// verify the cheque signature

// basic liquidity check
// could be omitted as it is not particularly useful

// store the accepted cheque

// RecoverCheque recovers the issuer ethereum address from a signed cheque
func RecoverCheque(cheque *SignedCheque, chaindID int64) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// keyChequebook computes the chequebook a store entry is for.
func keyChequebook(key []byte, prefix string) (chequebook common.Address, err error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// LastCheques returns the last received cheques from every known chequebook.
func (s *chequeStore) LastCheques() (map[common.Address]*SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
