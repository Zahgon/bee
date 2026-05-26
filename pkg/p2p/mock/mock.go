// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"time"

	"github.com/ethersphere/bee/v2/pkg/bzz"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	ma "github.com/multiformats/go-multiaddr"
)

// Service is the mock of a P2P Service
type Service struct {
	addProtocolFunc       func(p2p.ProtocolSpec) error
	connectFunc           func(ctx context.Context, addr []ma.Multiaddr) (address *bzz.Address, err error)
	disconnectFunc        func(overlay swarm.Address, reason string) error
	peersFunc             func() []p2p.Peer
	blocklistedPeersFunc  func() ([]p2p.BlockListedPeer, error)
	addressesFunc         func() ([]ma.Multiaddr, error)
	notifierFunc          p2p.PickyNotifier
	setWelcomeMessageFunc func(string) error
	getWelcomeMessageFunc func() string
	blocklistFunc         func(swarm.Address, time.Duration, string) error
	welcomeMessage        string
}

// WithAddProtocolFunc sets the mock implementation of the AddProtocol function
func WithAddProtocolFunc(f func(p2p.ProtocolSpec) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithConnectFunc sets the mock implementation of the Connect function
func WithConnectFunc(f func(ctx context.Context, addr []ma.Multiaddr) (address *bzz.Address, err error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDisconnectFunc sets the mock implementation of the Disconnect function
func WithDisconnectFunc(f func(overlay swarm.Address, reason string) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPeersFunc sets the mock implementation of the Peers function
func WithPeersFunc(f func() []p2p.Peer) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBlocklistedPeersFunc sets the mock implementation of the BlocklistedPeers function
func WithBlocklistedPeersFunc(f func() ([]p2p.BlockListedPeer, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAddressesFunc sets the mock implementation of the Addresses function
func WithAddressesFunc(f func() ([]ma.Multiaddr, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGetWelcomeMessageFunc sets the mock implementation of the GetWelcomeMessage function
func WithGetWelcomeMessageFunc(f func() string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSetWelcomeMessageFunc sets the mock implementation of the SetWelcomeMessage function
func WithSetWelcomeMessageFunc(f func(string) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithBlocklistFunc(f func(swarm.Address, time.Duration, string) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// New will create a new mock P2P Service with the given options
func New(opts ...Option) *Service { _ = "STUB: not implemented"; return nil }

func (s *Service) AddProtocol(spec p2p.ProtocolSpec) error { _ = "STUB: not implemented"; return nil }

func (s *Service) Connect(ctx context.Context, addr []ma.Multiaddr) (address *bzz.Address, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) Disconnect(overlay swarm.Address, reason string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) Addresses() ([]ma.Multiaddr, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Service) Peers() []p2p.Peer { _ = "STUB: not implemented"; return nil }

func (s *Service) Blocklisted(overlay swarm.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Service) BlocklistedPeers() ([]p2p.BlockListedPeer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) SetWelcomeMessage(val string) error { _ = "STUB: not implemented"; return nil }

func (s *Service) GetWelcomeMessage() string { _ = "STUB: not implemented"; return "" }

func (s *Service) Halt() { _ = "STUB: not implemented"; return }

func (s *Service) Blocklist(overlay swarm.Address, duration time.Duration, reason string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) SetPickyNotifier(f p2p.PickyNotifier) {
	_ = "STUB: not implemented"

	// NetworkStatus implements p2p.NetworkStatuser interface.
	// It always returns p2p.NetworkStatusAvailable.
	return
}

func (s *Service) NetworkStatus() p2p.NetworkStatus {
	_ = "STUB: not implemented"
	return *new(p2p.NetworkStatus)
}

type Option interface {
	apply(*Service)
}
type optionFunc func(*Service)

func (f optionFunc) apply(r *Service) { _ = "STUB: not implemented"; return }
