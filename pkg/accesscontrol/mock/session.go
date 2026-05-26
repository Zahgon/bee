// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"crypto/ecdsa"

	"github.com/ethersphere/bee/v2/pkg/keystore"
)

type SessionMock struct {
	KeyFunc func(publicKey *ecdsa.PublicKey, nonces [][]byte) ([][]byte, error)
	key     *ecdsa.PrivateKey
}

func (s *SessionMock) Key(publicKey *ecdsa.PublicKey, nonces [][]byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSessionMock(key *ecdsa.PrivateKey) *SessionMock { _ = "STUB: not implemented"; return nil }

func NewFromKeystore(
	ks keystore.Service,
	tag,
	password string,
	keyFunc func(publicKey *ecdsa.PublicKey, nonces [][]byte) ([][]byte, error),
) *SessionMock {
	_ = "STUB: not implemented"
	return nil
}
