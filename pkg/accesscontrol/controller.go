// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package accesscontrol provides functionalities needed
// for managing access control on Swarm
package accesscontrol

import (
	"context"
	"crypto/ecdsa"
	"io"

	"github.com/ethersphere/bee/v2/pkg/accesscontrol/kvs"
	"github.com/ethersphere/bee/v2/pkg/file"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// Grantees represents an interface for managing and retrieving grantees for a publisher.
type Grantees interface {
	// UpdateHandler manages the grantees for the given publisher, updating the list based on provided public keys to add or remove.
	// Only the publisher can make changes to the grantee list.
	UpdateHandler(ctx context.Context, ls file.LoadSaver, gls file.LoadSaver, granteeRef swarm.Address, historyRef swarm.Address, publisher *ecdsa.PublicKey, addList, removeList []*ecdsa.PublicKey) (swarm.Address, swarm.Address, swarm.Address, swarm.Address, error)
	// Get returns the list of grantees for the given publisher.
	// The list is accessible only by the publisher.
	Get(ctx context.Context, ls file.LoadSaver, publisher *ecdsa.PublicKey, encryptedglRef swarm.Address) ([]*ecdsa.PublicKey, error)
}

// Controller represents an interface for managing access control on Swarm.
// It provides methods for handling downloads, uploads and updates for grantee lists and references.
type Controller interface {
	Grantees
	// DownloadHandler decrypts the encryptedRef using the lookupkey based on the history and timestamp.
	DownloadHandler(ctx context.Context, ls file.LoadSaver, encryptedRef swarm.Address, publisher *ecdsa.PublicKey, historyRef swarm.Address, timestamp int64) (swarm.Address, error)
	// UploadHandler encrypts the reference and stores it in the history as the latest update.
	UploadHandler(ctx context.Context, ls file.LoadSaver, reference swarm.Address, publisher *ecdsa.PublicKey, historyRef swarm.Address) (swarm.Address, swarm.Address, swarm.Address, error)
	io.Closer
}

// ControllerStruct represents a controller for access control logic.
type ControllerStruct struct {
	access ActLogic
}

var _ Controller = (*ControllerStruct)(nil)

// NewController creates a new access controller with the given access logic.
func NewController(access ActLogic) *ControllerStruct { _ = "STUB: not implemented"; return nil }

// DownloadHandler decrypts the encryptedRef using the lookupkey based on the history and timestamp.
func (c *ControllerStruct) DownloadHandler(
	ctx context.Context,
	ls file.LoadSaver,
	encryptedRef swarm.Address,
	publisher *ecdsa.PublicKey,
	historyRef swarm.Address,
	timestamp int64,
) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// UploadHandler encrypts the reference and stores it in the history as the latest update.
func (c *ControllerStruct) UploadHandler(
	ctx context.Context,
	ls file.LoadSaver,
	reference swarm.Address,
	publisher *ecdsa.PublicKey,
	historyRef swarm.Address,
) (swarm.Address, swarm.Address, swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), *new(swarm.Address), *new(swarm.Address), nil
}

// UpdateHandler manages the grantees for the given publisher, updating the list based on provided public keys to add or remove.
// Only the publisher can make changes to the grantee list.
// Limitation: If an update is called again within a second from the latest upload/update then mantaray save fails with ErrInvalidInput,
// because the key (timestamp) is already present, hence a new fork is not created.
func (c *ControllerStruct) UpdateHandler(
	ctx context.Context,
	ls file.LoadSaver,
	gls file.LoadSaver,
	encryptedglRef swarm.Address,
	historyRef swarm.Address,
	publisher *ecdsa.PublicKey,
	addList []*ecdsa.PublicKey,
	removeList []*ecdsa.PublicKey,
) (swarm.Address, swarm.Address, swarm.Address, swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), *new(swarm.Address), *new(swarm.Address), *new(swarm.Address), nil
}

// generate new access key and new act, only if history was not newly created

// need to re-initialize history, because Lookup loads the forks causing the manifest save to skip the root node

// Get returns the list of grantees for the given publisher.
// The list is accessible only by the publisher.
func (c *ControllerStruct) Get(ctx context.Context, ls file.LoadSaver, publisher *ecdsa.PublicKey, encryptedglRef swarm.Address) ([]*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ControllerStruct) newActWithPublisher(ctx context.Context, ls file.LoadSaver, publisher *ecdsa.PublicKey) (kvs.KeyValueStore, error) {
	_ = "STUB: not implemented"
	return *new(kvs.KeyValueStore), nil
}

func (c *ControllerStruct) getHistoryAndAct(ctx context.Context, ls file.LoadSaver, historyRef swarm.Address, publisher *ecdsa.PublicKey, timestamp int64) (history History, act kvs.KeyValueStore, err error) {
	_ = "STUB: not implemented"
	return *new(History), *new(kvs.KeyValueStore), nil
}

func (c *ControllerStruct) saveHistoryAndAct(ctx context.Context, history History, mtdt *map[string]string, act kvs.KeyValueStore) (swarm.Address, swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), *new(swarm.Address), nil
}

func (c *ControllerStruct) getGranteeList(ctx context.Context, ls file.LoadSaver, encryptedglRef swarm.Address, publisher *ecdsa.PublicKey) (gl GranteeList, err error) {
	_ = "STUB: not implemented"
	return *new(GranteeList), nil
}

func (c *ControllerStruct) encryptRefForPublisher(publisherPubKey *ecdsa.PublicKey, ref swarm.Address) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func (c *ControllerStruct) decryptRefForPublisher(publisherPubKey *ecdsa.PublicKey, encryptedRef swarm.Address) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// Close simply returns nil
func (c *ControllerStruct) Close() error { _ = "STUB: not implemented"; return nil }
