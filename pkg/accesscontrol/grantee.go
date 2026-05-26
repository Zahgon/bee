// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package accesscontrol

import (
	"context"
	"crypto/ecdsa"
	"errors"

	"github.com/ethersphere/bee/v2/pkg/file"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const (
	publicKeyLen = 65
)

var (
	// ErrNothingToRemove indicates that the remove list is empty.
	ErrNothingToRemove = errors.New("nothing to remove")
	// ErrNoGranteeFound indicates that the grantee list is empty.
	ErrNoGranteeFound = errors.New("no grantee found")
	// ErrNothingToAdd indicates that the add list is empty.
	ErrNothingToAdd = errors.New("nothing to add")
)

// GranteeList manages a list of public keys.
type GranteeList interface {
	// Add adds a list of public keys to the grantee list. It filters out duplicates.
	Add(addList []*ecdsa.PublicKey) error
	// Remove removes a list of public keys from the grantee list, if there is any.
	Remove(removeList []*ecdsa.PublicKey) error
	// Get simply returns the list of public keys.
	Get() []*ecdsa.PublicKey
	// Save saves the grantee list to the underlying storage and returns the reference.
	Save(ctx context.Context) (swarm.Address, error)
}

// GranteeListStruct represents a list of grantee public keys.
type GranteeListStruct struct {
	grantees []*ecdsa.PublicKey
	loadSave file.LoadSaver
}

var _ GranteeList = (*GranteeListStruct)(nil)

// Get simply returns the list of public keys.
func (g *GranteeListStruct) Get() []*ecdsa.PublicKey {
	_ = "STUB: not implemented"

	// Add adds a list of public keys to the grantee list. It filters out duplicates.
	return nil
}

func (g *GranteeListStruct) Add(addList []*ecdsa.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

// Save saves the grantee list to the underlying storage and returns the reference.
func (g *GranteeListStruct) Save(ctx context.Context) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// Remove removes a list of public keys from the grantee list, if there is any.
func (g *GranteeListStruct) Remove(keysToRemove []*ecdsa.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

// NewGranteeList creates a new (and empty) grantee list.
func NewGranteeList(ls file.LoadSaver) *GranteeListStruct { _ = "STUB: not implemented"; return nil }

// NewGranteeListReference loads an existing grantee list.
func NewGranteeListReference(ctx context.Context, ls file.LoadSaver, reference swarm.Address) (*GranteeListStruct, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func serialize(publicKeys []*ecdsa.PublicKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: check if this is the correct way to serialize the public key
// Is this the only curve we support?
// Should we have switch case for different curves?
//nolint:staticcheck // SA1019: ecdsa fields are deprecated, but secp256k1 is not supported by crypto/ecdh

func deserialize(data []byte) []*ecdsa.PublicKey { _ = "STUB: not implemented"; return nil }

func deserializeBytes(data []byte) *ecdsa.PublicKey { _ = "STUB: not implemented"; return nil }
