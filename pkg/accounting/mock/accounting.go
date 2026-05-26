// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mock provides a mock implementation for the
// accounting interface.
package mock

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/accounting"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// Service is the mock Accounting service.
type Service struct {
	lock                    sync.Mutex
	balances                map[string]*big.Int
	prepareDebitFunc        func(peer swarm.Address, price uint64) (accounting.Action, error)
	prepareCreditFunc       func(peer swarm.Address, price uint64, originated bool) (accounting.Action, error)
	balanceFunc             func(swarm.Address) (*big.Int, error)
	shadowBalanceFunc       func(swarm.Address) (*big.Int, error)
	balancesFunc            func() (map[string]*big.Int, error)
	compensatedBalanceFunc  func(swarm.Address) (*big.Int, error)
	compensatedBalancesFunc func() (map[string]*big.Int, error)
	peerAccountingFunc      func() (map[string]accounting.PeerInfo, error)
	balanceSurplusFunc      func(swarm.Address) (*big.Int, error)
}

type debitAction struct {
	accounting *Service
	price      *big.Int
	peer       swarm.Address
	applied    bool
}

type creditAction struct {
	accounting *Service
	price      *big.Int
	peer       swarm.Address
	applied    bool
}

// WithPrepareDebitFunc sets the mock PrepareDebit function
func WithPrepareDebitFunc(f func(peer swarm.Address, price uint64) (accounting.Action, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPrepareCreditFunc sets the mock PrepareCredit function
func WithPrepareCreditFunc(f func(peer swarm.Address, price uint64, originated bool) (accounting.Action, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithBalanceFunc sets the mock Balance function
func WithBalanceFunc(f func(swarm.Address) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithBalancesFunc sets the mock Balances function
func WithBalancesFunc(f func() (map[string]*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCompensatedBalanceFunc sets the mock Balance function
func WithCompensatedBalanceFunc(f func(swarm.Address) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCompensatedBalancesFunc sets the mock Balances function
func WithCompensatedBalancesFunc(f func() (map[string]*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithBalanceSurplusFunc sets the mock SurplusBalance function
func WithBalanceSurplusFunc(f func(swarm.Address) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPeerAccountingFunc(f func() (map[string]accounting.PeerInfo, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// NewAccounting creates the mock accounting implementation
func NewAccounting(opts ...Option) *Service { _ = "STUB: not implemented"; return nil }

func (s *Service) MakeCreditAction(peer swarm.Address, price uint64) accounting.Action {
	_ = "STUB: not implemented"
	return *new(accounting.Action)
}

// Debit is the mock function wrapper that calls the set implementation
func (s *Service) PrepareDebit(_ context.Context, peer swarm.Address, price uint64) (accounting.Action, error) {
	_ = "STUB: not implemented"
	return *new(accounting.Action), nil
}

func (s *Service) PrepareCredit(_ context.Context, peer swarm.Address, price uint64, originated bool) (accounting.Action, error) {
	_ = "STUB: not implemented"
	return *new(accounting.Action), nil
}

func (a *debitAction) Apply() error { _ = "STUB: not implemented"; return nil }

func (a *creditAction) Cleanup() { _ = "STUB: not implemented"; return }

func (a *creditAction) Apply() error { _ = "STUB: not implemented"; return nil }

func (a *debitAction) Cleanup() {
	_ = "STUB: not implemented"

	// Balance is the mock function wrapper that calls the set implementation
	return
}

func (s *Service) Balance(peer swarm.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) ShadowBalance(peer swarm.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Balances is the mock function wrapper that calls the set implementation
func (s *Service) Balances() (map[string]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CompensatedBalance is the mock function wrapper that calls the set implementation
func (s *Service) CompensatedBalance(peer swarm.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CompensatedBalances is the mock function wrapper that calls the set implementation
func (s *Service) CompensatedBalances() (map[string]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) PeerAccounting() (map[string]accounting.PeerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) Connect(peer swarm.Address, full bool) { _ = "STUB: not implemented"; return }

func (s *Service) Disconnect(peer swarm.Address) { _ = "STUB: not implemented"; return }

func (s *Service) SurplusBalance(peer swarm.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Option is the option passed to the mock accounting service
type Option interface {
	apply(*Service)
}

type optionFunc func(*Service)

func (f optionFunc) apply(r *Service) { _ = "STUB: not implemented"; return }
