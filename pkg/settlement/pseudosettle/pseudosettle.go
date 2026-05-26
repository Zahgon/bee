// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pseudosettle

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/settlement"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "pseudosettle"

const (
	protocolName    = "pseudosettle"
	protocolVersion = "1.0.0"
	streamName      = "pseudosettle"
)

var (
	SettlementReceivedPrefix = "pseudosettle_total_received_"
	SettlementSentPrefix     = "pseudosettle_total_sent_"

	ErrSettlementTooSoon              = errors.New("settlement too soon")
	ErrNoPseudoSettlePeer             = errors.New("settlement peer not found")
	ErrDisconnectAllowanceCheckFailed = errors.New("settlement allowance below enforced amount")
	ErrTimeOutOfSyncAlleged           = errors.New("settlement allowance timestamps from peer were decreasing")
	ErrTimeOutOfSyncRecent            = errors.New("settlement allowance timestamps from peer differed from our measurement by more than 2 seconds")
	ErrTimeOutOfSyncInterval          = errors.New("settlement allowance interval from peer differed from local interval by more than 3 seconds")
	ErrRefreshmentBelowExpected       = errors.New("refreshment below expected")
	ErrRefreshmentAboveExpected       = errors.New("refreshment above expected")
)

type Service struct {
	streamer         p2p.Streamer
	logger           log.Logger
	store            storage.StateStorer
	accounting       settlement.Accounting
	metrics          metrics
	refreshRate      *big.Int
	lightRefreshRate *big.Int
	p2pService       p2p.Service
	timeNow          func() time.Time
	peersMu          sync.Mutex
	peers            map[string]*pseudoSettlePeer
}

type pseudoSettlePeer struct {
	lock     sync.Mutex // lock to be held during receiving a payment from this peer
	fullNode bool
}

type lastPayment struct {
	Timestamp      int64
	CheckTimestamp int64
	Total          *big.Int
}

func New(streamer p2p.Streamer, logger log.Logger, store storage.StateStorer, accounting settlement.Accounting, refreshRate, lightRefreshRate *big.Int, p2pService p2p.Service) *Service {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) Protocol() p2p.ProtocolSpec {
	_ = "STUB: not implemented"
	return *new(p2p.ProtocolSpec)
}

func (s *Service) init(ctx context.Context, p p2p.Peer) error {
	s.peersMu.Lock()
	defer s.peersMu.Unlock()

	_, ok := s.peers[p.Address.String()]
	if !ok {
		peerData := &pseudoSettlePeer{fullNode: p.FullNode}
		s.peers[p.Address.String()] = peerData
	}

	go s.accounting.Connect(p.Address, p.FullNode)
	return nil
}

func (s *Service) terminate(p p2p.Peer) error { _ = "STUB: not implemented"; return nil }

func totalKey(peer swarm.Address, prefix string) string { _ = "STUB: not implemented"; return "" }

func totalKeyPeer(key []byte, prefix string) (peer swarm.Address, err error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// peerAllowance computes the maximum incoming payment value we accept
// this is the time based allowance or the peers actual debt, whichever is less
func (s *Service) peerAllowance(peer swarm.Address, fullNode bool) (limit *big.Int, stamp int64, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (s *Service) handler(ctx context.Context, p p2p.Peer, stream p2p.Stream) (err error) {
	_ = "STUB: not implemented"
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

// SettlementsSent returns all stored sent settlement values for a given type of prefix
func (s *Service) SettlementsSent() (map[string]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SettlementsReceived returns all stored received settlement values for a given type of prefix
func (s *Service) SettlementsReceived() (map[string]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
