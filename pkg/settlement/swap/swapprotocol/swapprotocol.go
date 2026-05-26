// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package swapprotocol

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/chequebook"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/priceoracle"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "swapprotocol"

const (
	protocolName    = "swap"
	protocolVersion = "1.0.0"
	streamName      = "swap" // stream for cheques
)

var (
	ErrNegotiateRate      = errors.New("exchange rates mismatch")
	ErrNegotiateDeduction = errors.New("deduction values mismatch")
	ErrHaveDeduction      = errors.New("received deduction not zero")
)

type SendChequeFunc chequebook.SendChequeFunc

type IssueFunc func(ctx context.Context, beneficiary common.Address, amount *big.Int, sendChequeFunc chequebook.SendChequeFunc) (*big.Int, error)

// (context.Context, common.Address, *big.Int, chequebook.SendChequeFunc) (*big.Int, error)

// Interface is the main interface to send messages over swap protocol.
type Interface interface {
	// EmitCheque sends a signed cheque to a peer.
	EmitCheque(ctx context.Context, peer swarm.Address, beneficiary common.Address, amount *big.Int, issue IssueFunc) (balance *big.Int, err error)
}

// Swap is the interface the settlement layer should implement to receive cheques.
type Swap interface {
	// ReceiveCheque is called by the swap protocol if a cheque is received.
	ReceiveCheque(ctx context.Context, peer swarm.Address, cheque *chequebook.SignedCheque, exchangeRate, deduction *big.Int) error
	// Handshake is called by the swap protocol when a handshake is received.
	Handshake(peer swarm.Address, beneficiary common.Address) error
	GetDeductionForPeer(peer swarm.Address) (bool, error)
	GetDeductionByPeer(peer swarm.Address) (bool, error)
	AddDeductionByPeer(peer swarm.Address) error
}

// Service is the main implementation of the swap protocol.
type Service struct {
	streamer    p2p.Streamer
	logger      log.Logger
	swap        Swap
	priceOracle priceoracle.Service
	beneficiary common.Address
}

// New creates a new swap protocol Service.
func New(streamer p2p.Streamer, logger log.Logger, beneficiary common.Address, priceOracle priceoracle.Service) *Service {
	_ = "STUB: not implemented"
	return nil
}

// SetSwap sets the swap to notify.
func (s *Service) SetSwap(swap Swap) { _ = "STUB: not implemented"; return }

func (s *Service) Protocol() p2p.ProtocolSpec {
	_ = "STUB: not implemented"
	return *new(p2p.ProtocolSpec)
}

// init is called on outgoing connections and triggers handshake exchange
func (s *Service) init(ctx context.Context, p p2p.Peer) error {
	beneficiary := common.BytesToAddress(p.EthereumAddress)
	return s.swap.Handshake(p.Address, beneficiary)
}

func (s *Service) handler(ctx context.Context, p p2p.Peer, stream p2p.Stream) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// signature validation

func (s *Service) headler(receivedHeaders p2p.Headers, peerAddress swarm.Address) (returnHeaders p2p.Headers) {
	_ = "STUB: not implemented"
	return *new(p2p.Headers)
}

// InitiateCheque attempts to send a cheque to a peer.
func (s *Service) EmitCheque(ctx context.Context, peer swarm.Address, beneficiary common.Address, amount *big.Int, issue IssueFunc) (balance *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// reading exchangeRated headers

// comparing received headers to known truth

// get whether peer have deducted in the past

// if peer is not entitled for deduction but sent non zero deduction value, return with error

// get current global exchangeRate rate and deduction

// exchangeRate rates should match

// deduction values should match or be zero

// issue cheque call with provided callback for sending cheque to finish transaction

// for simplicity we use json marshaller. can be replaced by a binary encoding in the future.

// sending cheque
