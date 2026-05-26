// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gsoc

import (
	"sync"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/soc"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// Handler defines code to be executed upon reception of a GSOC sub message.
// it is used as a parameter definition.
type Handler func([]byte)

type Listener interface {
	Subscribe(address swarm.Address, handler Handler) (cleanup func())
	Handle(c *soc.SOC)
	Close() error
}

type listener struct {
	handlers   map[string][]*Handler
	handlersMu sync.Mutex
	quit       chan struct{}
	logger     log.Logger
}

// New returns a new GSOC listener service.
func New(logger log.Logger) Listener { _ = "STUB: not implemented"; return *new(Listener) }

// Subscribe allows the definition of a Handler func on a specific GSOC address.
func (l *listener) Subscribe(address swarm.Address, handler Handler) (cleanup func()) {
	_ = "STUB: not implemented"
	return nil
}

// Handle is called by push/pull sync and passes the chunk its registered handler
func (l *listener) Handle(c *soc.SOC) { _ = "STUB: not implemented"; return }

// no handler

// no handler

func (p *listener) getHandlers(address swarm.Address) []*Handler {
	_ = "STUB: not implemented"
	return nil
}

func (l *listener) Close() error { _ = "STUB: not implemented"; return nil }

// unset handlers on shutdown
