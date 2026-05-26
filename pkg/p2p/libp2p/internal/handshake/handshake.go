// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handshake

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethersphere/bee/v2/pkg/bzz"
	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/p2p/libp2p/internal/handshake/pb"
	"github.com/ethersphere/bee/v2/pkg/swarm"

	libp2ppeer "github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "handshake"

const (
	// ProtocolName is the text of the name of the handshake protocol.
	ProtocolName = "handshake"
	// ProtocolVersion is the current handshake protocol version.
	ProtocolVersion = "14.0.0"
	// StreamName is the name of the stream used for handshake purposes.
	StreamName = "handshake"
	// MaxWelcomeMessageLength is maximum number of characters allowed in the welcome message.
	MaxWelcomeMessageLength = 140
	handshakeTimeout        = 15 * time.Second
)

var (
	// ErrNetworkIDIncompatible is returned if response from the other peer does not have valid networkID.
	ErrNetworkIDIncompatible = errors.New("incompatible network ID")

	// ErrInvalidAck is returned if data in received in ack is not valid (invalid signature for example).
	ErrInvalidAck = errors.New("invalid ack")

	// ErrInvalidSyn is returned if observable address in ack is not a valid..
	ErrInvalidSyn = errors.New("invalid syn")

	// ErrWelcomeMessageLength is returned if the welcome message is longer than the maximum length
	ErrWelcomeMessageLength = fmt.Errorf("handshake welcome message longer than maximum of %d characters", MaxWelcomeMessageLength)

	// ErrPicker is returned if the picker (kademlia) rejects the peer
	ErrPicker = errors.New("picker rejection")
)

// AdvertisableAddressResolver can Resolve a Multiaddress.
type AdvertisableAddressResolver interface {
	Resolve(observedAddress ma.Multiaddr) (ma.Multiaddr, error)
}

type Addresser interface {
	AdvertizableAddrs() ([]ma.Multiaddr, error)
}

type Option struct {
	bee260compatibility bool
}

// WithBee260Compatibility option ensures that only one underlay address is
// passed to the peer in p2p protocol messages, so that nodes with version 2.6.0
// and older can deserialize it. This option can be safely removed when bee
// version 2.6.0 is deprecated.
func WithBee260Compatibility(yes bool) func(*Option) { _ = "STUB: not implemented"; return nil }

// Service can perform initiate or handle a handshake between peers.
type Service struct {
	signer                crypto.Signer
	advertisableAddresser AdvertisableAddressResolver
	overlay               swarm.Address
	fullNode              bool
	nonce                 []byte
	networkID             uint64
	validateOverlay       bool
	welcomeMessage        atomic.Value
	logger                log.Logger
	libp2pID              libp2ppeer.ID
	metrics               metrics
	picker                p2p.Picker
	mu                    sync.RWMutex
	hostAddresser         Addresser
}

// Info contains the information received from the handshake.
type Info struct {
	BzzAddress *bzz.Address
	FullNode   bool
}

func (i *Info) LightString() string { _ = "STUB: not implemented"; return "" }

// New creates a new handshake Service.
func New(signer crypto.Signer, advertisableAddresser AdvertisableAddressResolver, overlay swarm.Address, networkID uint64, fullNode bool, nonce []byte, hostAddresser Addresser, welcomeMessage string, validateOverlay bool, ownPeerID libp2ppeer.ID, logger log.Logger) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) SetPicker(n p2p.Picker) { _ = "STUB: not implemented"; return }

// Handshake initiates a handshake with a peer.
func (s *Service) Handshake(ctx context.Context, stream p2p.Stream, peerMultiaddrs []ma.Multiaddr, opts ...func(*Option)) (i *Info, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sort to remove potential duplicates

// remove duplicates

// Synced read:

// Handle handles an incoming handshake from a peer.
func (s *Service) Handle(ctx context.Context, stream p2p.Stream, peerMultiaddrs []ma.Multiaddr, opts ...func(*Option)) (i *Info, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sort to remove potential duplicates

// remove duplicates

// SetWelcomeMessage sets the new handshake welcome message.
func (s *Service) SetWelcomeMessage(msg string) (err error) { _ = "STUB: not implemented"; return nil }

// GetWelcomeMessage returns the current handshake welcome message.
func (s *Service) GetWelcomeMessage() string { _ = "STUB: not implemented"; return "" }

func (s *Service) parseCheckAck(ack *pb.Ack) (*bzz.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
