// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chequebook

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/transaction"
)

// ErrNoCashout is the error if there has not been any cashout action for the chequebook
var ErrNoCashout = errors.New("no prior cashout")

// CashoutService is the service responsible for managing cashout actions
type CashoutService interface {
	// CashCheque sends a cashing transaction for the last cheque of the chequebook
	CashCheque(ctx context.Context, chequebook common.Address, recipient common.Address) (common.Hash, error)
	// CashoutStatus gets the status of the latest cashout transaction for the chequebook
	CashoutStatus(ctx context.Context, chequebookAddress common.Address) (*CashoutStatus, error)
}

type cashoutService struct {
	store              storage.StateStorer
	backend            transaction.Backend
	transactionService transaction.Service
	chequeStore        ChequeStore
}

// LastCashout contains information about the last cashout
type LastCashout struct {
	TxHash   common.Hash
	Cheque   SignedCheque // the cheque that was used to cashout which may be different from the latest cheque
	Result   *CashChequeResult
	Reverted bool
}

// CashoutStatus is information about the last cashout and uncashed amounts
type CashoutStatus struct {
	Last           *LastCashout // last cashout for a chequebook
	UncashedAmount *big.Int     // amount not yet cashed out
}

// CashChequeResult summarizes the result of a CashCheque or CashChequeBeneficiary call
type CashChequeResult struct {
	Beneficiary      common.Address // beneficiary of the cheque
	Recipient        common.Address // address which received the funds
	Caller           common.Address // caller of cashCheque
	TotalPayout      *big.Int       // total amount that was paid out in this call
	CumulativePayout *big.Int       // cumulative payout of the cheque that was cashed
	CallerPayout     *big.Int       // payout for the caller of cashCheque
	Bounced          bool           // indicates whether parts of the cheque bounced
}

// cashoutAction is the data we store for a cashout
type cashoutAction struct {
	TxHash common.Hash
	Cheque SignedCheque // the cheque that was used to cashout which may be different from the latest cheque
}
type chequeCashedEvent struct {
	Beneficiary      common.Address
	Recipient        common.Address
	Caller           common.Address
	TotalPayout      *big.Int
	CumulativePayout *big.Int
	CallerPayout     *big.Int
}

// NewCashoutService creates a new CashoutService
func NewCashoutService(
	store storage.StateStorer,
	backend transaction.Backend,
	transactionService transaction.Service,
	chequeStore ChequeStore,
) CashoutService {
	_ = "STUB: not implemented"
	return *new(CashoutService)
}

// cashoutActionKey computes the store key for the last cashout action for the chequebook
func cashoutActionKey(chequebook common.Address) string { _ = "STUB: not implemented"; return "" }

func (s *cashoutService) paidOut(ctx context.Context, chequebook, beneficiary common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CashCheque sends a cashout transaction for the last cheque of the chequebook
func (s *cashoutService) CashCheque(ctx context.Context, chequebook, recipient common.Address) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// CashoutStatus gets the status of the latest cashout transaction for the chequebook
func (s *cashoutService) CashoutStatus(ctx context.Context, chequebookAddress common.Address) (*CashoutStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if we never cashed out, assume everything is uncashed

// treat not found as pending

// uncashed is the difference since the last sent cashout. we assume that the entire cheque will clear in the pending transaction.

// if a tx failed (should be almost impossible in practice) we no longer have the necessary information to compute uncashed locally
// assume there are no pending transactions and that the on-chain paidOut is the last cashout action

// uncashed is the difference since the last sent (and confirmed) cashout.

// parseCashChequeBeneficiaryReceipt processes the receipt from a CashChequeBeneficiary transaction
func (s *cashoutService) parseCashChequeBeneficiaryReceipt(chequebookAddress common.Address, receipt *types.Receipt) (*CashChequeResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Equal compares to CashChequeResults
func (r *CashChequeResult) Equal(o *CashChequeResult) bool { _ = "STUB: not implemented"; return false }
