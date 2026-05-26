// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/chequebook"
)

// Service is the mock chequebook service.
type Service struct {
	chequebookBalanceFunc          func(context.Context) (*big.Int, error)
	chequebookAvailableBalanceFunc func(context.Context) (*big.Int, error)
	chequebookAddressFunc          func() common.Address
	chequebookIssueFunc            func(ctx context.Context, beneficiary common.Address, amount *big.Int, sendChequeFunc chequebook.SendChequeFunc) (*big.Int, error)
	chequebookWithdrawFunc         func(ctx context.Context, amount *big.Int) (hash common.Hash, err error)
	chequebookDepositFunc          func(ctx context.Context, amount *big.Int) (hash common.Hash, err error)
	lastChequeFunc                 func(common.Address) (*chequebook.SignedCheque, error)
	lastChequesFunc                func() (map[common.Address]*chequebook.SignedCheque, error)
}

// WithChequebook*Functions set the mock chequebook functions
func WithChequebookBalanceFunc(f func(ctx context.Context) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithChequebookAvailableBalanceFunc(f func(ctx context.Context) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithChequebookAddressFunc(f func() common.Address) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithChequebookDepositFunc(f func(ctx context.Context, amount *big.Int) (hash common.Hash, err error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithChequebookIssueFunc(f func(ctx context.Context, beneficiary common.Address, amount *big.Int, sendChequeFunc chequebook.SendChequeFunc) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithChequebookWithdrawFunc(f func(ctx context.Context, amount *big.Int) (hash common.Hash, err error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLastChequeFunc(f func(beneficiary common.Address) (*chequebook.SignedCheque, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLastChequesFunc(f func() (map[common.Address]*chequebook.SignedCheque, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// NewChequebook creates the mock chequebook implementation
func NewChequebook(opts ...Option) chequebook.Service {
	_ = "STUB: not implemented"
	return *new(chequebook.Service)
}

// Balance mocks the chequebook .Balance function
func (s *Service) Balance(ctx context.Context) (bal *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) AvailableBalance(ctx context.Context) (bal *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deposit mocks the chequebook .Deposit function
func (s *Service) Deposit(ctx context.Context, amount *big.Int) (hash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// WaitForDeposit mocks the chequebook .WaitForDeposit function
func (s *Service) WaitForDeposit(ctx context.Context, txHash common.Hash) error {
	_ = "STUB: not implemented"
	return nil

	// Address mocks the chequebook .Address function
}

func (s *Service) Address() common.Address { _ = "STUB: not implemented"; return *new(common.Address) }

func (s *Service) Issue(ctx context.Context, beneficiary common.Address, amount *big.Int, sendChequeFunc chequebook.SendChequeFunc) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) LastCheque(beneficiary common.Address) (*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) LastCheques() (map[common.Address]*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) Withdraw(ctx context.Context, amount *big.Int) (hash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// Option is the option passed to the mock Chequebook service
type Option interface {
	apply(*Service)
}

type optionFunc func(*Service)

func (f optionFunc) apply(r *Service) { _ = "STUB: not implemented"; return }
