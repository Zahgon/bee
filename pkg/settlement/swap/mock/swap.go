// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethersphere/bee/v2/pkg/settlement/swap"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/chequebook"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/swapprotocol"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type Service struct {
	settlementsSent map[string]*big.Int
	settlementsRecv map[string]*big.Int

	settlementSentFunc func(swarm.Address) (*big.Int, error)
	settlementRecvFunc func(swarm.Address) (*big.Int, error)

	settlementsSentFunc func() (map[string]*big.Int, error)
	settlementsRecvFunc func() (map[string]*big.Int, error)

	deductionByPeers  map[string]struct{}
	deductionForPeers map[string]struct{}

	receiveChequeFunc   func(context.Context, swarm.Address, *chequebook.SignedCheque, *big.Int, *big.Int) error
	payFunc             func(context.Context, swarm.Address, *big.Int)
	handshakeFunc       func(swarm.Address, common.Address) error
	lastSentChequeFunc  func(swarm.Address) (*chequebook.SignedCheque, error)
	lastSentChequesFunc func() (map[string]*chequebook.SignedCheque, error)

	lastReceivedChequeFunc  func(swarm.Address) (*chequebook.SignedCheque, error)
	lastReceivedChequesFunc func() (map[string]*chequebook.SignedCheque, error)

	cashChequeFunc    func(ctx context.Context, peer swarm.Address) (common.Hash, error)
	cashoutStatusFunc func(ctx context.Context, peer swarm.Address) (*chequebook.CashoutStatus, error)
}

// WithSettlementSentFunc sets the mock settlement function
func WithSettlementSentFunc(f func(swarm.Address) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSettlementRecvFunc(f func(swarm.Address) (*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSettlementsSentFunc sets the mock settlements function
func WithSettlementsSentFunc(f func() (map[string]*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSettlementsRecvFunc(f func() (map[string]*big.Int, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithReceiveChequeFunc(f func(context.Context, swarm.Address, *chequebook.SignedCheque, *big.Int, *big.Int) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPayFunc(f func(context.Context, swarm.Address, *big.Int)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithHandshakeFunc(f func(swarm.Address, common.Address) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLastSentChequeFunc(f func(swarm.Address) (*chequebook.SignedCheque, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLastSentChequesFunc(f func() (map[string]*chequebook.SignedCheque, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLastReceivedChequeFunc(f func(swarm.Address) (*chequebook.SignedCheque, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLastReceivedChequesFunc(f func() (map[string]*chequebook.SignedCheque, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCashChequeFunc(f func(ctx context.Context, peer swarm.Address) (common.Hash, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithCashoutStatusFunc(f func(ctx context.Context, peer swarm.Address) (*chequebook.CashoutStatus, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// New creates the mock swap implementation
func New(opts ...Option) swap.Interface { _ = "STUB: not implemented"; return *new(swap.Interface) }

func NewSwap(opts ...Option) swapprotocol.Swap {
	_ = "STUB: not implemented"
	return *new(swapprotocol.Swap)
}

// Pay is the mock Pay function of swap.
func (s *Service) Pay(ctx context.Context, peer swarm.Address, amount *big.Int) {
	_ = "STUB: not implemented"
	return
}

// TotalSent is the mock TotalSent function of swap.
func (s *Service) TotalSent(peer swarm.Address) (totalSent *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalReceived is the mock TotalReceived function of swap.
func (s *Service) TotalReceived(peer swarm.Address) (totalReceived *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SettlementsSent is the mock SettlementsSent function of swap.
func (s *Service) SettlementsSent() (map[string]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SettlementsReceived is the mock SettlementsReceived function of swap.
func (s *Service) SettlementsReceived() (map[string]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handshake is called by the swap protocol when a handshake is received.
func (s *Service) Handshake(peer swarm.Address, beneficiary common.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) LastSentCheque(address swarm.Address) (*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) LastSentCheques() (map[string]*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) LastReceivedCheque(address swarm.Address) (*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) LastReceivedCheques() (map[string]*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) CashCheque(ctx context.Context, peer swarm.Address) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (s *Service) CashoutStatus(ctx context.Context, peer swarm.Address) (*chequebook.CashoutStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) ReceiveCheque(ctx context.Context, peer swarm.Address, cheque *chequebook.SignedCheque, exchangeRate, deduction *big.Int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) GetDeductionForPeer(peer swarm.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Service) GetDeductionByPeer(peer swarm.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Service) AddDeductionByPeer(peer swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// Option is the option passed to the mock settlement service
type Option interface {
	apply(*Service)
}

type optionFunc func(*Service)

func (f optionFunc) apply(r *Service) { _ = "STUB: not implemented"; return }
