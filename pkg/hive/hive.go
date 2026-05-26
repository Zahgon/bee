// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hive exposes the hive protocol implementation
// which is the discovery protocol used to inform and be
// informed about other peers in the network. It gossips
// about all peers by default and performs no specific
// prioritization about which peers are gossipped to
// others.
package hive

import (
	"context"
	"errors"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"

	"github.com/ethersphere/bee/v2/pkg/addressbook"
	"github.com/ethersphere/bee/v2/pkg/hive/pb"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/ratelimit"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	ma "github.com/multiformats/go-multiaddr"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "hive"

const (
	protocolName    = "hive"
	protocolVersion = "1.1.0"
	peersStreamName = "peers"
	messageTimeout  = 1 * time.Minute // maximum allowed time for a message to be read or written.
	maxBatchSize    = 30
)

var (
	limitBurst = 4 * int(swarm.MaxBins)
	limitRate  = time.Minute

	ErrRateLimitExceeded = errors.New("rate limit exceeded")
)

type Service struct {
	streamer          p2p.Bee260CompatibilityStreamer
	addressBook       addressbook.GetPutter
	addPeersHandler   func(...swarm.Address)
	networkID         uint64
	logger            log.Logger
	metrics           metrics
	inLimiter         *ratelimit.Limiter
	outLimiter        *ratelimit.Limiter
	quit              chan struct{}
	wg                sync.WaitGroup
	peersChan         chan pb.Peers
	sem               *semaphore.Weighted
	bootnode          bool
	allowPrivateCIDRs bool
	overlay           swarm.Address
}

func New(streamer p2p.Bee260CompatibilityStreamer, addressbook addressbook.GetPutter, networkID uint64, bootnode bool, allowPrivateCIDRs bool, overlay swarm.Address, logger log.Logger) *Service {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) Protocol() p2p.ProtocolSpec {
	_ = "STUB: not implemented"
	return *new(p2p.ProtocolSpec)
}

var ErrShutdownInProgress = errors.New("shutdown in progress")

func (s *Service) BroadcastPeers(ctx context.Context, addressee swarm.Address, peers ...swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// If broadcasting limit is exceeded, return early

func (s *Service) SetAddPeersHandler(h func(addr ...swarm.Address)) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Service) sendPeers(ctx context.Context, peer swarm.Address, peers []swarm.Address) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) peersHandler(ctx context.Context, peer p2p.Peer, stream p2p.Stream) error {
	_ = "STUB: not implemented"
	return nil
}

// close the stream before processing in order to unblock the sending side
// fullclose is called async because there is no need to wait for confirmation,
// but we still want to handle not closed stream from the other side to avoid zombie stream

func (s *Service) disconnect(peer p2p.Peer) error { _ = "STUB: not implemented"; return nil }

func (s *Service) startCheckPeersHandler() { _ = "STUB: not implemented"; return }

func (s *Service) checkAndAddPeers(peers pb.Peers) { _ = "STUB: not implemented"; return }

// no underlays sent

// if peer exists already in the addressBook
// and if the underlays match, skip

// filterAdvertisableUnderlays returns underlay addresses that can be advertised
// based on the allowPrivateCIDRs setting. If allowPrivateCIDRs is false,
// only public addresses are returned.
func (s *Service) filterAdvertisableUnderlays(underlays []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}
