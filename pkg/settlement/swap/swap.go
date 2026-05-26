// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package swap

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/settlement"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/chequebook"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/swapprotocol"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "swap"

var (
	// ErrWrongChequebook is the error if a peer uses a different chequebook from before.
	ErrWrongChequebook = errors.New("wrong chequebook")
	// ErrUnknownBeneficary is the error if a peer has never announced a beneficiary.
	ErrUnknownBeneficary = errors.New("unknown beneficiary for peer")
	// ErrChequeValueTooLow is the error a peer issued a cheque not covering 1 accounting credit
	ErrChequeValueTooLow = errors.New("cheque value too low")
	ErrNoChequebook      = errors.New("no chequebook")
)

type Interface interface {
	settlement.Interface
	// LastSentCheque returns the last sent cheque for the peer
	LastSentCheque(peer swarm.Address) (*chequebook.SignedCheque, error)
	// LastSentCheques returns the list of last sent cheques for all peers
	LastSentCheques() (map[string]*chequebook.SignedCheque, error)
	// LastReceivedCheque returns the last received cheque for the peer
	LastReceivedCheque(peer swarm.Address) (*chequebook.SignedCheque, error)
	// LastReceivedCheques returns the list of last received cheques for all peers
	LastReceivedCheques() (map[string]*chequebook.SignedCheque, error)
	// CashCheque sends a cashing transaction for the last cheque of the peer
	CashCheque(ctx context.Context, peer swarm.Address) (common.Hash, error)
	// CashoutStatus gets the status of the latest cashout transaction for the peers chequebook
	CashoutStatus(ctx context.Context, peer swarm.Address) (*chequebook.CashoutStatus, error)
}

// Service is the implementation of the swap settlement layer.
type Service struct {
	proto          swapprotocol.Interface
	logger         log.Logger
	store          storage.StateStorer
	accounting     settlement.Accounting
	metrics        metrics
	chequebook     chequebook.Service
	chequeStore    chequebook.ChequeStore
	cashout        chequebook.CashoutService
	addressbook    Addressbook
	networkID      uint64
	cashoutAddress common.Address
}

// New creates a new swap Service.
func New(proto swapprotocol.Interface, logger log.Logger, store storage.StateStorer, chequebook chequebook.Service, chequeStore chequebook.ChequeStore, addressbook Addressbook, networkID uint64, cashout chequebook.CashoutService, accounting settlement.Accounting, cashoutAddress common.Address) *Service {
	_ = "STUB: not implemented"
	return nil
}

// ReceiveCheque is called by the swap protocol if a cheque is received.
func (s *Service) ReceiveCheque(ctx context.Context, peer swarm.Address, cheque *chequebook.SignedCheque, exchangeRate, deduction *big.Int) (err error) {
	_ = "STUB: not implemented"
	// check this is the same chequebook for this peer as previously
	return nil
}

// Pay initiates a payment to the given peer
func (s *Service) Pay(ctx context.Context, peer swarm.Address, amount *big.Int) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) SetAccounting(accounting settlement.Accounting) {
	_ = "STUB: not implemented"
	return
}

// TotalSent returns the total amount sent to a peer
func (s *Service) TotalSent(peer swarm.Address) (totalSent *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalReceived returns the total amount received from a peer
func (s *Service) TotalReceived(peer swarm.Address) (totalReceived *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SettlementsSent returns sent settlements for each individual known peer
func (s *Service) SettlementsSent() (map[string]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SettlementsReceived returns received settlements for each individual known peer.
func (s *Service) SettlementsReceived() (map[string]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handshake is called by the swap protocol when a handshake is received.
func (s *Service) Handshake(peer swarm.Address, beneficiary common.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// LastSentCheque returns the last sent cheque for the peer
func (s *Service) LastSentCheque(peer swarm.Address) (*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastReceivedCheque returns the last received cheque for the peer
func (s *Service) LastReceivedCheque(peer swarm.Address) (*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastSentCheques returns the list of last sent cheques for all peers
func (s *Service) LastSentCheques() (map[string]*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastReceivedCheques returns the list of last received cheques for all peers
func (s *Service) LastReceivedCheques() (map[string]*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CashCheque sends a cashing transaction for the last cheque of the peer
func (s *Service) CashCheque(ctx context.Context, peer swarm.Address) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// CashoutStatus gets the status of the latest cashout transaction for the peers chequebook
func (s *Service) CashoutStatus(ctx context.Context, peer swarm.Address) (*chequebook.CashoutStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
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

type NoOpSwap struct{}

func (*NoOpSwap) TotalSent(peer swarm.Address) (totalSent *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TotalReceived returns the total amount received from a peer
func (*NoOpSwap) TotalReceived(peer swarm.Address) (totalSent *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SettlementsSent returns sent settlements for each individual known peer
func (*NoOpSwap) SettlementsSent() (map[string]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SettlementsReceived returns received settlements for each individual known peer
func (*NoOpSwap) SettlementsReceived() (map[string]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*NoOpSwap) LastSentCheque(peer swarm.Address) (*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastSentCheques returns the list of last sent cheques for all peers
func (*NoOpSwap) LastSentCheques() (map[string]*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastReceivedCheque returns the last received cheque for the peer
func (*NoOpSwap) LastReceivedCheque(peer swarm.Address) (*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LastReceivedCheques returns the list of last received cheques for all peers
func (*NoOpSwap) LastReceivedCheques() (map[string]*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CashCheque sends a cashing transaction for the last cheque of the peer
func (*NoOpSwap) CashCheque(ctx context.Context, peer swarm.Address) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// CashoutStatus gets the status of the latest cashout transaction for the peers chequebook
func (*NoOpSwap) CashoutStatus(ctx context.Context, peer swarm.Address) (*chequebook.CashoutStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
