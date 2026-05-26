// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pingpong exposes the simple ping-pong protocol
// which measures round-trip-time with other peers.
package pingpong

import (
	"context"
	"time"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/tracing"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "pingpong"

const (
	protocolName    = "pingpong"
	protocolVersion = "1.0.0"
	streamName      = "pingpong"
)

type Interface interface {
	Ping(ctx context.Context, address swarm.Address, msgs ...string) (rtt time.Duration, err error)
}

type Service struct {
	streamer p2p.Streamer
	logger   log.Logger
	tracer   *tracing.Tracer
	metrics  metrics
}

func New(streamer p2p.Streamer, logger log.Logger, tracer *tracing.Tracer) *Service {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) Protocol() p2p.ProtocolSpec {
	_ = "STUB: not implemented"
	return *new(p2p.ProtocolSpec)
}

func (s *Service) Ping(ctx context.Context, address swarm.Address, msgs ...string) (rtt time.Duration, err error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (s *Service) handler(ctx context.Context, p p2p.Peer, stream p2p.Stream) error {
	_ = "STUB: not implemented"
	return nil
}
