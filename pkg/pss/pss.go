// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pss exposes functionalities needed to communicate
// with other peers on the network. Pss uses pushsync and
// pullsync for message delivery and mailboxing. All messages are disguised as content-addressed chunks. Sending and
// receiving of messages is exposed over the HTTP API, with
// websocket subscriptions for incoming messages.
package pss

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"io"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/pushsync"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "pss"

var (
	_            Interface = (*pss)(nil)
	ErrNoHandler           = errors.New("no handler found")
)

type Sender interface {
	// Send arbitrary byte slice with the given topic to Targets.
	Send(context.Context, Topic, []byte, postage.Stamper, *ecdsa.PublicKey, Targets) error
}

type Interface interface {
	Sender
	// Register a Handler for a given Topic.
	Register(Topic, Handler) func()
	// TryUnwrap tries to unwrap a wrapped trojan message.
	TryUnwrap(swarm.Chunk)

	SetPushSyncer(pushSyncer pushsync.PushSyncer)
	io.Closer
}

type pss struct {
	key        *ecdsa.PrivateKey
	pusher     pushsync.PushSyncer
	handlers   map[Topic][]*Handler
	handlersMu sync.Mutex
	metrics    metrics
	logger     log.Logger
	quit       chan struct{}
}

// New returns a new pss service.
func New(key *ecdsa.PrivateKey, logger log.Logger) Interface {
	_ = "STUB: not implemented"
	return *new(Interface)
}

func (ps *pss) Close() error { _ = "STUB: not implemented"; return nil }

// unset handlers on shutdown

func (ps *pss) SetPushSyncer(pushSyncer pushsync.PushSyncer) { _ = "STUB: not implemented"; return }

// Handler defines code to be executed upon reception of a trojan message.
type Handler func(context.Context, []byte)

// Send constructs a padded message with topic and payload,
// wraps it in a trojan chunk such that one of the targets is a prefix of the chunk address.
// Uses push-sync to deliver message.
func (p *pss) Send(ctx context.Context, topic Topic, payload []byte, stamper postage.Stamper, recipient *ecdsa.PublicKey, targets Targets) error {
	_ = "STUB: not implemented"
	return nil
}

// push the chunk using push sync so that it reaches it destination in network

// Register allows the definition of a Handler func for a specific topic on the pss struct.
func (p *pss) Register(topic Topic, handler Handler) (cleanup func()) {
	_ = "STUB: not implemented"
	return nil
}

func (p *pss) topics() []Topic { _ = "STUB: not implemented"; return nil }

// TryUnwrap allows unwrapping a chunk as a trojan message and calling its handlers based on the topic.
func (p *pss) TryUnwrap(c swarm.Chunk) { _ = "STUB: not implemented"; return }

// chunk not full

// cannot unwrap

// no handler

func (p *pss) getHandlers(topic Topic) []*Handler { _ = "STUB: not implemented"; return nil }
