// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package streamtest

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	ma "github.com/multiformats/go-multiaddr"
)

var (
	ErrRecordsNotFound    = errors.New("records not found")
	ErrStreamNotSupported = errors.New("stream not supported")
	ErrStreamClosed       = errors.New("stream closed")

	noopMiddleware = func(f p2p.HandlerFunc) p2p.HandlerFunc {
		return f
	}
)

type Recorder struct {
	base               swarm.Address
	fullNode           bool
	records            map[string][]*Record
	recordsMu          sync.Mutex
	protocols          []p2p.ProtocolSpec
	middlewares        []p2p.HandlerMiddleware
	streamErr          func(swarm.Address, string, string, string) error
	pingErr            func(ma.Multiaddr) (time.Duration, error)
	protocolsWithPeers map[string]p2p.ProtocolSpec
	messageLatency     time.Duration
}

func WithProtocols(protocols ...p2p.ProtocolSpec) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPeerProtocols(protocolsWithPeers map[string]p2p.ProtocolSpec) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMiddlewares(middlewares ...p2p.HandlerMiddleware) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithBaseAddr(a swarm.Address) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLightNode() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithStreamError(streamErr func(swarm.Address, string, string, string) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithPingErr(pingErr func(ma.Multiaddr) (time.Duration, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMessageLatency(latency time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func New(opts ...Option) *Recorder { _ = "STUB: not implemented"; return nil }

func (r *Recorder) Reset() { _ = "STUB: not implemented"; return }

func (r *Recorder) SetProtocols(protocols ...p2p.ProtocolSpec) { _ = "STUB: not implemented"; return }

func (r *Recorder) NewStream(ctx context.Context, addr swarm.Address, h p2p.Headers, protocolName, protocolVersion, streamName string) (p2p.Stream, error) {
	_ = "STUB: not implemented"
	return *new(p2p.Stream), nil
}

// pass a new context to handler,

// do not cancel it with the client stream context

func (r *Recorder) Ping(ctx context.Context, addr ma.Multiaddr) (rtt time.Duration, err error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (r *Recorder) Records(addr swarm.Address, protocolName, protocolVersio, streamName string) ([]*Record, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// wait for all records goroutines to terminate

// WaitRecords waits for some time for records to come into the recorder. If msgs is 0, the timeoutSec period is waited to verify
// that _no_ messages arrive during this time period.
func (r *Recorder) WaitRecords(t *testing.T, addr swarm.Address, proto, version, stream string, msgs, timeoutSec int) []*Record {
	_ = "STUB: not implemented"
	return nil
}

// we can be here if msgs == 0 && l == 0
// or msgs = x && l < x, both cases are fine
// and we should continue waiting

// IsBee260 implements p2p.Bee260CompatibilityStreamer interface.
// It always returns false.
func (r *Recorder) IsBee260(overlay swarm.Address) bool { _ = "STUB: not implemented"; return false }

type Record struct {
	in    *record
	out   *record
	err   error
	errMu sync.Mutex
	done  chan struct{}
}

func (r *Record) In() []byte { _ = "STUB: not implemented"; return nil }

func (r *Record) Out() []byte { _ = "STUB: not implemented"; return nil }

func (r *Record) Err() error { _ = "STUB: not implemented"; return nil }

func (r *Record) setErr(err error) { _ = "STUB: not implemented"; return }

type stream struct {
	in              *record
	out             *record
	headers         p2p.Headers
	responseHeaders p2p.Headers
	closed          bool
	lock            sync.Mutex
}

func newStream(in, out *record) *stream { _ = "STUB: not implemented"; return nil }

func (s *stream) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *stream) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *stream) Headers() p2p.Headers { _ = "STUB: not implemented"; return *new(p2p.Headers) }

func (s *stream) ResponseHeaders() p2p.Headers { _ = "STUB: not implemented"; return *new(p2p.Headers) }

func (s *stream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *stream) Closed() bool { _ = "STUB: not implemented"; return false }

func (s *stream) FullClose() error { _ = "STUB: not implemented"; return nil }

func (s *stream) Reset() (err error) { _ = "STUB: not implemented"; return nil }

type record struct {
	b        []byte
	c        int
	lock     sync.Mutex
	dataSigC chan struct{}
	latency  time.Duration
	closed   bool
}

func newRecord(latency time.Duration) *record { _ = "STUB: not implemented"; return nil }

func (r *record) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (r *record) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *record) close() { _ = "STUB: not implemented"; return }

func (r *record) bytes() []byte { _ = "STUB: not implemented"; return nil }

func (r *record) bytesSize() int { _ = "STUB: not implemented"; return 0 }

type Option interface {
	apply(*Recorder)
}
type optionFunc func(*Recorder)

func (f optionFunc) apply(r *Recorder) { _ = "STUB: not implemented"; return }

var _ p2p.StreamerDisconnecter = (*RecorderDisconnecter)(nil)

type RecorderDisconnecter struct {
	*Recorder
	disconnected map[string]struct{}
	blocklisted  map[string]time.Duration
	mu           sync.RWMutex
}

func NewRecorderDisconnecter(r *Recorder) *RecorderDisconnecter {
	_ = "STUB: not implemented"
	return nil
}

func (r *RecorderDisconnecter) Disconnect(overlay swarm.Address, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RecorderDisconnecter) Blocklist(overlay swarm.Address, d time.Duration, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RecorderDisconnecter) IsDisconnected(overlay swarm.Address) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RecorderDisconnecter) IsBlocklisted(overlay swarm.Address) (bool, time.Duration) {
	_ = "STUB: not implemented"
	return false, *new(time.Duration)
}

// NetworkStatus implements p2p.NetworkStatuser interface.
// It always returns p2p.NetworkStatusAvailable.
func (r *RecorderDisconnecter) NetworkStatus() p2p.NetworkStatus {
	_ = "STUB: not implemented"
	return *new(p2p.NetworkStatus)
}
