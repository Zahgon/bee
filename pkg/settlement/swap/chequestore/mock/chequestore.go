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

// Service is the mock chequeStore service.
type Service struct {
	receiveCheque func(ctx context.Context, cheque *chequebook.SignedCheque, exchangeRate *big.Int, deduction *big.Int) (*big.Int, error)
	lastCheque    func(chequebook common.Address) (*chequebook.SignedCheque, error)
	lastCheques   func() (map[common.Address]*chequebook.SignedCheque, error)
}

func WithReceiveChequeFunc(f func(ctx context.Context, cheque *chequebook.SignedCheque, exchangeRate *big.Int, deduction *big.Int) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLastChequeFunc(f func(chequebook common.Address) (*chequebook.SignedCheque, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLastChequesFunc(f func() (map[common.Address]*chequebook.SignedCheque, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// NewChequeStore creates the mock chequeStore implementation
func NewChequeStore(opts ...Option) chequebook.ChequeStore {
	_ = "STUB: not implemented"
	return *new(chequebook.ChequeStore)
}

func (s *Service) ReceiveCheque(ctx context.Context, cheque *chequebook.SignedCheque, exchangeRate, deduction *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) LastCheque(chequebook common.Address) (*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) LastCheques() (map[common.Address]*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil,

		// Option is the option passed to the mock ChequeStore service
		nil
}

type Option interface {
	apply(*Service)
}

type optionFunc func(*Service)

func (f optionFunc) apply(r *Service) { _ = "STUB: not implemented"; return }
