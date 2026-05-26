// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package accesscontrol

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethersphere/bee/v2/pkg/accesscontrol/kvs"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"golang.org/x/crypto/sha3"
)

//nolint:gochecknoglobals
var (
	hashFunc      = sha3.NewLegacyKeccak256
	oneByteArray  = []byte{1}
	zeroByteArray = []byte{0}
)

// Decryptor is a read-only interface for the ACT.
type Decryptor interface {
	// DecryptRef will return a decrypted reference, for given encrypted reference and grantee.
	DecryptRef(ctx context.Context, storage kvs.KeyValueStore, encryptedRef swarm.Address, publisher *ecdsa.PublicKey) (swarm.Address, error)
	Session
}

// Control interface for the ACT (does write operations).
type Control interface {
	Decryptor
	// AddGrantee adds a new grantee to the ACT.
	AddGrantee(ctx context.Context, storage kvs.KeyValueStore, publisherPubKey, granteePubKey *ecdsa.PublicKey) error
	// EncryptRef encrypts a Swarm reference for a given grantee.
	EncryptRef(ctx context.Context, storage kvs.KeyValueStore, grantee *ecdsa.PublicKey, ref swarm.Address) (swarm.Address, error)
}

// ActLogic represents the access control logic.
type ActLogic struct {
	Session
}

var _ Control = (*ActLogic)(nil)

// EncryptRef encrypts a Swarm reference for a publisher.
func (al ActLogic) EncryptRef(ctx context.Context, storage kvs.KeyValueStore, publisherPubKey *ecdsa.PublicKey, ref swarm.Address) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// AddGrantee adds a new grantee to the ACT.
func (al ActLogic) AddGrantee(ctx context.Context, storage kvs.KeyValueStore, publisherPubKey, granteePubKey *ecdsa.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

// Create new access key because grantee is the publisher.

// Get previously generated access key.

// Encrypt the access key for the new Grantee.

// Add the new encrypted access key to the Act.

// Will return the access key for a publisher (public key).
func (al *ActLogic) getAccessKey(ctx context.Context, storage kvs.KeyValueStore, publisherPubKey *ecdsa.PublicKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no need for constructor call if value not found in act.

// Generate lookup key and access key decryption key for a given public key.
func (al *ActLogic) getKeys(publicKey *ecdsa.PublicKey) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// DecryptRef will return a decrypted reference, for given encrypted reference and publisher.
func (al ActLogic) DecryptRef(ctx context.Context, storage kvs.KeyValueStore, encryptedRef swarm.Address, publisher *ecdsa.PublicKey) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// NewLogic creates a new ACT Logic from a session.
func NewLogic(s Session) ActLogic { _ = "STUB: not implemented"; return *new(ActLogic) }
