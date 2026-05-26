// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package accesscontrol

import (
	"crypto/ecdsa"
	"errors"
)

var (
	// ErrInvalidPublicKey is an error that is returned when a public key is nil.
	ErrInvalidPublicKey = errors.New("invalid public key")
	// ErrSecretKeyInfinity is an error that is returned when the shared secret is a point at infinity.
	ErrSecretKeyInfinity = errors.New("shared secret is point at infinity")
)

// Session represents an interface for a Diffie-Hellmann key derivation
type Session interface {
	// Key returns a derived key for each nonce.
	Key(publicKey *ecdsa.PublicKey, nonces [][]byte) ([][]byte, error)
}

var _ Session = (*SessionStruct)(nil)

// SessionStruct represents a session with an access control key.
type SessionStruct struct {
	key *ecdsa.PrivateKey
}

// Key returns a derived key for each nonce.
func (s *SessionStruct) Key(publicKey *ecdsa.PublicKey, nonces [][]byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck // SA1019: ecdsa fields are deprecated, but secp256k1 is not supported by crypto/ecdh

// NewDefaultSession creates a new session from a private key.
func NewDefaultSession(key *ecdsa.PrivateKey) *SessionStruct { _ = "STUB: not implemented"; return nil }
