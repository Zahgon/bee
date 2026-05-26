// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package retrieval provides the retrieval protocol
// implementation. The protocol is used to retrieve
// chunks over the network using forwarding-kademlia
// routing.
package retrieval

import (
	"context"
	"time"

	"github.com/ethersphere/bee/v2/pkg/accounting"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/pricer"
	"github.com/ethersphere/bee/v2/pkg/skippeers"
	storage "github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/topology"
	"github.com/ethersphere/bee/v2/pkg/tracing"
	"github.com/opentracing/opentracing-go"
	"resenje.org/singleflight"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "retrieval"

const (
	protocolName    = "retrieval"
	protocolVersion = "1.4.0"
	streamName      = "retrieval"
)

var _ Interface = (*Service)(nil)

type Interface interface {
	// RetrieveChunk retrieves a chunk from the network using the retrieval protocol.
	// it takes as parameters a context, a chunk address to retrieve (content-addressed or single-owner) and
	// a source peer address, for the case that we are requesting the chunk for another peer. In case the request
	// originates at the current node (i.e. no forwarding involved), the caller should use swarm.ZeroAddress
	// as the value for sourcePeerAddress.
	RetrieveChunk(ctx context.Context, address, sourcePeerAddr swarm.Address) (chunk swarm.Chunk, err error)
}

type retrievalResult struct {
	chunk swarm.Chunk
	peer  swarm.Address
	err   error
}

type Storer interface {
	Cache() storage.Putter
	Lookup() storage.Getter
}

type Service struct {
	addr          swarm.Address
	radiusFunc    func() (uint8, error)
	streamer      p2p.Streamer
	peerSuggester topology.ClosestPeerer
	storer        Storer
	singleflight  singleflight.Group[string, swarm.Chunk]
	logger        log.Logger
	accounting    accounting.Interface
	metrics       metrics
	pricer        pricer.Interface
	tracer        *tracing.Tracer
	caching       bool
	errSkip       *skippeers.List
}

func New(
	addr swarm.Address,
	radiusFunc func() (uint8, error),
	storer Storer,
	streamer p2p.Streamer,
	chunkPeerer topology.ClosestPeerer,
	logger log.Logger,
	accounting accounting.Interface,
	pricer pricer.Interface,
	tracer *tracing.Tracer,
	forwarderCaching bool,
) *Service {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) Protocol() p2p.ProtocolSpec {
	_ = "STUB: not implemented"
	return *new(p2p.ProtocolSpec)
}

const (
	RetrieveChunkTimeout = time.Second * 30
	preemptiveInterval   = time.Second
	overDraftRefresh     = time.Millisecond * 600
	skiplistDur          = time.Minute
	originSuffix         = "_origin"
	maxOriginErrors      = 32
	maxMultiplexForwards = 2
)

func (s *Service) RetrieveChunk(ctx context.Context, chunkAddr, sourcePeerAddr swarm.Address) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// if we are the origin node, allow many preemptive retries to speed up the retrieval of the chunk.

// no overdraft peers, we have depleted ALL peers

// there is still an inflight request, wait for it's result

// since we can reach into the neighborhood of the chunk
// act as the multiplexer and push the chunk in parallel to multiple peers.
// neighbor peers will also have multiple retries, which means almost the entire neighborhood
// will be scanned for the chunk, starting from the closest to the furthest peer in the neighborhood.

func (s *Service) retrieveChunk(ctx context.Context, quit chan struct{}, chunkAddr, peer swarm.Address, result chan retrievalResult, action accounting.Action, span opentracing.Span) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) prepareCredit(ctx context.Context, peer, chunk swarm.Address, origin bool) (accounting.Action, error) {
	_ = "STUB: not implemented"
	return *new(accounting.Action), nil
}

// closestPeer returns address of the peer that is closest to the chunk with
// provided address addr. This function will ignore peers with addresses
// provided in skipPeers and if allowUpstream is true, peers that are further of
// the chunk than this node is, could also be returned, allowing the upstream
// retrieve request.
func (s *Service) closestPeer(addr swarm.Address, skipPeers []swarm.Address, allowUpstream bool) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func (s *Service) handler(p2pctx context.Context, p p2p.Peer, stream p2p.Stream) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// forward the request

// debit price from p's balance

// cache the request last, so that putting to the localstore does not slow down the request flow

func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }
