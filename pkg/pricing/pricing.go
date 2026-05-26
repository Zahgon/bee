// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pricing

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "pricing"

const (
	protocolName    = "pricing"
	protocolVersion = "1.0.0"
	streamName      = "pricing"
)

// ErrThresholdTooLow says that the proposed payment threshold is too low for even a single reserve.
var ErrThresholdTooLow = errors.New("threshold too low")

var _ Interface = (*Service)(nil)

// Interface is the main interface of the pricing protocol
type Interface interface {
	AnnouncePaymentThreshold(ctx context.Context, peer swarm.Address, paymentThreshold *big.Int) error
}

// PaymentThresholdObserver is used for being notified of payment threshold updates
type PaymentThresholdObserver interface {
	NotifyPaymentThreshold(peer swarm.Address, paymentThreshold *big.Int) error
}

type Service struct {
	streamer                 p2p.Streamer
	logger                   log.Logger
	paymentThreshold         *big.Int
	lightPaymentThreshold    *big.Int
	minPaymentThreshold      *big.Int
	paymentThresholdObserver PaymentThresholdObserver
}

func New(streamer p2p.Streamer, logger log.Logger, paymentThreshold, lightPaymentThreshold, minThreshold *big.Int) *Service {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) Protocol() p2p.ProtocolSpec {
	_ = "STUB: not implemented"
	return *new(p2p.ProtocolSpec)
}

func (s *Service) handler(ctx context.Context, p p2p.Peer, stream p2p.Stream) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) init(ctx context.Context, p p2p.Peer) error {
	threshold := s.paymentThreshold
	if !p.FullNode {
		threshold = s.lightPaymentThreshold
	}

	err := s.AnnouncePaymentThreshold(ctx, p.Address, threshold)
	if err != nil {
		s.logger.Warning("could not send payment threshold announcement to peer", "peer_address", p.Address)
	}
	return err
}

// AnnouncePaymentThreshold announces the payment threshold to per
func (s *Service) AnnouncePaymentThreshold(ctx context.Context, peer swarm.Address, paymentThreshold *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// SetPaymentThresholdObserver sets the PaymentThresholdObserver to be used when receiving a new payment threshold
func (s *Service) SetPaymentThresholdObserver(observer PaymentThresholdObserver) {
	_ = "STUB: not implemented"
	return
}
