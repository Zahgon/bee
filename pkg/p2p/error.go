// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p2p

import (
	"errors"
	"time"
)

var (
	// ErrPeerNotFound should be returned by p2p service methods when the requested
	// peer is not found.
	ErrPeerNotFound = errors.New("peer not found")
	// ErrAlreadyConnected is returned if connect was called for already connected node.
	ErrAlreadyConnected = errors.New("already connected")
	// ErrDialLightNode is returned if connect was attempted to a light node.
	ErrDialLightNode = errors.New("target peer is a light node")
	// ErrPeerBlocklisted is returned if peer is on blocklist
	ErrPeerBlocklisted = errors.New("peer blocklisted")
	// ErrUnsupportedAddresses is returned when all peer addresses use unsupported transports
	ErrUnsupportedAddresses = errors.New("no supported addresses")
)

const (
	DefaultBlocklistTime = 1 * time.Minute
)

// ConnectionBackoffError indicates that connection calls will not be executed until `tryAfter` timestamp.
// The reason is provided in the wrapped error.
type ConnectionBackoffError struct {
	tryAfter time.Time
	err      error
}

// NewConnectionBackoffError creates new `ConnectionBackoffError` with provided underlying error and `tryAfter` timestamp.
func NewConnectionBackoffError(err error, tryAfter time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// TryAfter returns a tryAfter timestamp.
func (e *ConnectionBackoffError) TryAfter() time.Time {
	_ = "STUB: not implemented"

	// Unwrap returns an underlying error.
	return *new(time.Time)
}

func (e *ConnectionBackoffError) Unwrap() error {
	_ = "STUB: not implemented"

	// Error implements function of the standard go error interface.
	return nil
}

func (e *ConnectionBackoffError) Error() string { _ = "STUB: not implemented"; return "" }

// DisconnectError is an error that is specifically handled inside p2p. If returned by specific protocol
// handler it causes peer disconnect.
type DisconnectError struct {
	err error
}

// NewDisconnectError wraps error and creates a special error that is treated specially
// by p2p. It causes peer to disconnect.
func NewDisconnectError(err error) error { _ = "STUB: not implemented"; return nil }

// Unwrap returns an underlying error.
func (e *DisconnectError) Unwrap() error {
	_ = "STUB: not implemented"

	// Error implements function of the standard go error interface.
	return nil
}

func (e *DisconnectError) Error() string { _ = "STUB: not implemented"; return "" }

type BlockPeerError struct {
	duration time.Duration
	err      error
}

// NewBlockPeerError wraps error and creates a special error that is treated specially
// by p2p. It causes peer to be disconnected and blocks any new connection for this peer for the provided duration.
func NewBlockPeerError(duration time.Duration, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// Unwrap returns an underlying error.
func (e *BlockPeerError) Unwrap() error {
	_ = "STUB: not implemented"

	// Error implements function of the standard go error interface.
	return nil
}

func (e *BlockPeerError) Error() string { _ = "STUB: not implemented"; return "" }

// Duration represents the period for which the peer will be blocked.
// 0 duration is treated as infinity
func (e *BlockPeerError) Duration() time.Duration {
	_ = "STUB: not implemented"

	// IncompatibleStreamError is the error that should be returned by p2p service
	// NewStream method when the stream or its version is not supported.
	return *new(time.Duration)
}

type IncompatibleStreamError struct {
	err error
}

// NewIncompatibleStreamError wraps the error that is the cause of stream
// incompatibility with IncompatibleStreamError that it can be detected and
// returns it.
func NewIncompatibleStreamError(err error) *IncompatibleStreamError {
	_ = "STUB: not implemented"
	return nil
}

// Unwrap returns an underlying error.
func (e *IncompatibleStreamError) Unwrap() error {
	_ = "STUB: not implemented"

	// Error implements function of the standard go error interface.
	return nil
}

func (e *IncompatibleStreamError) Error() string { _ = "STUB: not implemented"; return "" }
