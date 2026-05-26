// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libp2p

import (
	"context"
	"crypto/ecdsa"
	"sync"
	"time"

	"github.com/coreos/go-semver/semver"
	"github.com/ethersphere/bee/v2/pkg/addressbook"
	"github.com/ethersphere/bee/v2/pkg/bzz"
	beecrypto "github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/p2p/libp2p/internal/blocklist"
	"github.com/ethersphere/bee/v2/pkg/p2p/libp2p/internal/breaker"
	"github.com/ethersphere/bee/v2/pkg/p2p/libp2p/internal/handshake"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/topology"
	"github.com/ethersphere/bee/v2/pkg/topology/lightnode"
	"github.com/ethersphere/bee/v2/pkg/tracing"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	libp2ppeer "github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/p2p/host/autonat"
	basichost "github.com/libp2p/go-libp2p/p2p/host/basic"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "libp2p"

var (
	_ p2p.Service      = (*Service)(nil)
	_ p2p.DebugService = (*Service)(nil)

	// reachabilityOverridePublic overrides autonat to simply report
	// public reachability status, it is set in the makefile.
	reachabilityOverridePublic = "false"
)

const (
	defaultLightNodeLimit = 100
	peerUserAgentTimeout  = time.Second

	peerstoreWaitAddrsTimeout = 10 * time.Second

	defaultHeadersRWTimeout = 10 * time.Second

	IncomingStreamCountLimit = 5_000
	OutgoingStreamCountLimit = 10_000
)

type Service struct {
	ctx                context.Context
	host               host.Host
	natManager         basichost.NATManager
	autonatDialer      host.Host
	pingDialer         host.Host
	libp2pPeerstore    peerstore.Peerstore
	metrics            metrics
	networkID          uint64
	handshakeService   *handshake.Service
	addressbook        addressbook.Putter
	peers              *peerRegistry
	connectionBreaker  breaker.Interface
	blocklist          *blocklist.Blocklist
	protocols          []p2p.ProtocolSpec
	notifier           p2p.PickyNotifier
	logger             log.Logger
	tracer             *tracing.Tracer
	ready              chan struct{}
	halt               chan struct{}
	lightNodes         lightnodes
	lightNodeLimit     int
	protocolsmu        sync.RWMutex
	reacher            p2p.Reacher
	networkStatus      atomic.Int32
	HeadersRWTimeout   time.Duration
	autoNAT            autonat.AutoNAT
	autoTLSCertManager autoTLSCertManager
	zapLogger          *zap.Logger
	enabledTransports  map[bzz.TransportType]bool
}

type lightnodes interface {
	Connected(context.Context, p2p.Peer)
	Disconnected(p2p.Peer)
	Count() int
	RandomPeer(swarm.Address) (swarm.Address, error)
	EachPeer(pf topology.EachPeerFunc) error
}

type Options struct {
	PrivateKey                  *ecdsa.PrivateKey
	NATAddr                     string
	NATWSSAddr                  string
	EnableWS                    bool
	EnableWSS                   bool
	WSSAddr                     string
	AutoTLSStorageDir           string
	AutoTLSCAEndpoint           string
	AutoTLSDomain               string
	AutoTLSRegistrationEndpoint string
	FullNode                    bool
	LightNodeLimit              int
	WelcomeMessage              string
	Nonce                       []byte
	ValidateOverlay             bool
	hostFactory                 func(...libp2p.Option) (host.Host, error)
	HeadersRWTimeout            time.Duration
	Registry                    *prometheus.Registry
	autoTLSCertManager          autoTLSCertManager
}

func New(ctx context.Context, signer beecrypto.Signer, networkID uint64, overlay swarm.Address, addr string, ab addressbook.Putter, storer storage.StateStorer, lightNodes *lightnode.Container, logger log.Logger, tracer *tracing.Tracer, o Options) (s *Service, returnErr error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tweak certain settings

// Create our limits by using our cfg and replacing the default values with values from `scaledDefaultLimits`

// The resource manager expects a limiter, se we create one from our limits.

// IPv4 /32 (Single IP) -> 200 conns
// IPv6 /56 subnet -> 200 conns

// Custom rate limiter for connection attempts
// 20 peers cluster adaptation:
// Allow bursts of connection attempts (e.g. restart) but prevent DDOS.

// Allow unlimited local connections (same as default)

// Unlimited global

// Apply limits per individual IPv4 address (/32)
// Allow 10 connection attempts per second per IP, burst up to 40

// Apply limits per /56 IPv6 subnet
// Allow 10 connection attempts per second per IP, burst up to 40
// Subnet-level limiting prevents flooding from multiple addresses in the same block.

// Duration to retain state for an IP or subnet after it becomes inactive.

// AutoTLS is only needed for WSS

// call if service is not constructed

// Use dedicated peerstore instead the global DefaultPeerstore

// Use the default libp2p host creation

// Support same non default security and transport options as
// original host.

// If you want to help other peers to figure out if they are behind
// NATs, you can launch the server-side of AutoNAT too (AutoRelay
// already runs the client)

// Create a new dialer for libp2p ping protocol. This ensures that the protocol
// uses a different set of keys to do ping. It prevents inconsistencies in peerstore as
// the addresses used are not dialable and hence should be cleaned up. We should create
// this host with the same transports and security options to be able to dial to other
// peers.

// use default options

// TCP transport is always included

// Construct protocols.

// update peer registry on network events

type parsedAddress struct {
	IP4  string
	IP6  string
	Port string
}

func parseAddress(addr string) (*parsedAddress, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Service) reachabilityWorker() error { _ = "STUB: not implemented"; return nil }

func (s *Service) handleIncoming(stream network.Stream) { _ = "STUB: not implemented"; return }

// For the handshake we always need an observed address (ObservedUnderlay).
// If the peerstore had no addresses, fall back to RemoteMultiaddr for the
// handshake only. This typically means the peer is behind NAT and its
// address is not reachable from the outside.

// Only persist in addressbook when we have real peerstore addresses.

// light node announces explicitly

// kick another node to fit this one in

// note: this cannot be unit tested since the node
// waiting on handshakeStream.FullClose() on the other side
// might actually get a stream reset when we disconnect here
// resulting in a flaky response from the Connect method on
// the other side.
// that is why the Pick method has been added to the notifier
// interface, in addition to the possibility of deciding whether
// a peer connection is wanted prior to adding the peer to the
// peer registry and starting the protocols.

// when a full node connects, we gossip about it to the
// light nodes so that they can also have a chance at building
// a solid topology.

// isTransportSupported checks if the given transport type is supported by this service.
func (s *Service) isTransportSupported(t bzz.TransportType) bool {
	_ = "STUB: not implemented"
	return false
}

// filterSupportedAddresses filters multiaddresses to only include those
// that are supported by the available transports (TCP, WS, WSS).
func (s *Service) filterSupportedAddresses(addrs []ma.Multiaddr) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) notifyReacherConnected(overlay swarm.Address, underlays []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) SetPickyNotifier(n p2p.PickyNotifier) { _ = "STUB: not implemented"; return }

func (s *Service) AddProtocol(p p2p.ProtocolSpec) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// exchange headers

// tracing: get span tracing context and add it to the context
// silently ignore if the peer is not providing tracing

// count unexpected requests

func (s *Service) Addresses() (addresses []ma.Multiaddr, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func includeNatResolvedAddresses(addrs []ma.Multiaddr, advertisableAddresser handshake.AdvertisableAddressResolver, logger log.Logger) (addresses []ma.Multiaddr) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) NATManager() basichost.NATManager {
	_ = "STUB: not implemented"
	return *new(basichost.NATManager)
}

func (s *Service) Blocklist(overlay swarm.Address, duration time.Duration, reason string) error {
	_ = "STUB: not implemented"
	return nil
}

func buildHostAddress(peerID libp2ppeer.ID) (ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), nil
}

func (s *Service) Connect(ctx context.Context, addrs []ma.Multiaddr) (address *bzz.Address, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to connect to each underlay address one by one.

// Extract the peer ID from the multiaddr.

// Check if attempting to connect to self

// If we skipped all addresses due to self-connection, return an error

func (s *Service) Disconnect(overlay swarm.Address, reason string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// found is checked at the bottom of the function

// disconnected is a registered peer registry event
func (s *Service) disconnected(address swarm.Address) { _ = "STUB: not implemented"; return }

// peerID might not always be found on shutdown

func (s *Service) Peers() []p2p.Peer { _ = "STUB: not implemented"; return nil }

func (s *Service) Blocklisted(overlay swarm.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Service) BlocklistedPeers() ([]p2p.BlockListedPeer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) NewStream(ctx context.Context, overlay swarm.Address, headers p2p.Headers, protocolName, protocolVersion, streamName string) (p2p.Stream, error) {
	_ = "STUB: not implemented"
	return *new(p2p.Stream), nil
}

// Verify if we really have an active connection

// tracing: add span context header

// exchange headers

func (s *Service) newStreamForPeerID(ctx context.Context, peerID libp2ppeer.ID, protocolName, protocolVersion, streamName string) (network.Stream, error) {
	_ = "STUB: not implemented"
	return *new(network.Stream), nil
}

func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }

// SetWelcomeMessage sets the welcome message for the handshake protocol.
func (s *Service) SetWelcomeMessage(val string) error { _ = "STUB: not implemented"; return nil }

// GetWelcomeMessage returns the value of the welcome message.
func (s *Service) GetWelcomeMessage() string { _ = "STUB: not implemented"; return "" }

func (s *Service) Ready() error { _ = "STUB: not implemented"; return nil }

func (s *Service) Halt() { _ = "STUB: not implemented"; return }

func (s *Service) Ping(ctx context.Context, addr ma.Multiaddr) (rtt time.Duration, err error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// Add the address to libp2p peerstore for it to be dialable

// Cleanup connection after ping is done

// peerUserAgent returns User Agent string of the connected peer if the peer
// provides it. It ignores the default libp2p user agent string
// "github.com/libp2p/go-libp2p" and returns empty string in that case.
func (s *Service) peerUserAgent(ctx context.Context, peerID libp2ppeer.ID) string {
	_ = "STUB: not implemented"
	return ""
}

// Peerstore may not contain all keys and values right after the connections is created.
// This retry mechanism ensures more reliable user agent propagation.

// error is ignored as user agent is informative only

// Ignore the default user agent.

// NetworkStatus implements the p2p.NetworkStatuser interface.
func (s *Service) NetworkStatus() p2p.NetworkStatus {
	_ = "STUB: not implemented"
	return *new(p2p.NetworkStatus)
}

// determineCurrentNetworkStatus determines if the network
// is available/unavailable based on the given error, and
// returns ErrNetworkUnavailable if unavailable.
// The result of this operation is stored and can be reflected
// in the results of future NetworkStatus method calls.
func (s *Service) determineCurrentNetworkStatus(err error) error {
	_ = "STUB: not implemented"
	return nil
}

// peerMultiaddrs builds full multiaddresses for a peer using the peerstore.
func (s *Service) peerMultiaddrs(ctx context.Context, peerID libp2ppeer.ID) ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IsBee260 implements p2p.Bee260CompatibilityStreamer interface.
// It checks if a peer is running Bee version older than 2.7.0.
func (s *Service) IsBee260(overlay swarm.Address) bool { _ = "STUB: not implemented"; return false }

var version270 = *semver.Must(semver.NewVersion("2.7.0"))

func (s *Service) bee260BackwardCompatibility(peerID libp2ppeer.ID) bool {
	_ = "STUB: not implemented"
	return false
}

// Compare major.minor.patch only (ignore pre-release)
// This way 2.7.0-rc12 is treated as >= 2.7.0

// appendSpace adds a leading space character if the string is not empty.
// It is useful for constructing log messages with conditional substrings.
func appendSpace(s string) string { _ = "STUB: not implemented"; return "" }

// userAgent returns a User Agent string passed to the libp2p host to identify peer node.
func userAgent() string { _ = "STUB: not implemented"; return "" }

func newConnMetricNotify(m metrics) *connectionNotifier { _ = "STUB: not implemented"; return nil }

type connectionNotifier struct {
	metrics metrics
	network.Notifiee
}

func (c *connectionNotifier) Connected(_ network.Network, _ network.Conn) {
	_ = "STUB: not implemented"
	return
}

// isNetworkOrHostUnreachableError determines based on the
// given error whether the host or network is reachable.
func isNetworkOrHostUnreachableError(err error) bool { _ = "STUB: not implemented"; return false }

// Since TransportError doesn't implement the Unwrap
// method we need to inspect the errors manually.

type compositeAddressResolver struct {
	tcpResolver handshake.AdvertisableAddressResolver
	wssResolver handshake.AdvertisableAddressResolver
}

func newCompositeAddressResolver(tcpResolver, wssResolver handshake.AdvertisableAddressResolver) handshake.AdvertisableAddressResolver {
	_ = "STUB: not implemented"
	return *new(handshake.AdvertisableAddressResolver)
}

func (c *compositeAddressResolver) Resolve(observedAddress ma.Multiaddr) (ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), nil
}

// ma.P_WSS protocol is deprecated, multiaddrs should comtain WS and TLS protocols for WSS

type hostAddresser struct {
	host host.Host
}

func newHostAddresser(host host.Host) *hostAddresser { _ = "STUB: not implemented"; return nil }

func (h *hostAddresser) AdvertizableAddrs() ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildFullMAs(addrs []ma.Multiaddr, peerID libp2ppeer.ID) ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildFullMA(addr ma.Multiaddr, peerID libp2ppeer.ID) (ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), nil
}

// waitPeerAddrs is used to reliably get remote addresses from libp2p peerstore
// as sometimes addresses are not available soon enough from its Addrs() method.
func waitPeerAddrs(ctx context.Context, s peerstore.Peerstore, peerID libp2ppeer.ID) []ma.Multiaddr {
	_ = "STUB: not implemented"
	return nil
}

// cancel the addrStream when this function exits

// ensure that the AddrStream will receive addresses by creating it before Addrs() is called
// this may happen just after the connection is established and peerstore is not updated

// return the first address as it arrives
