// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mock provides a mock implementation for the
// access control functionalities.
//
//nolint:ireturn
package mock

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethersphere/bee/v2/pkg/accesscontrol"
	"github.com/ethersphere/bee/v2/pkg/encryption"
	"github.com/ethersphere/bee/v2/pkg/file"
	"github.com/ethersphere/bee/v2/pkg/file/pipeline"
	"github.com/ethersphere/bee/v2/pkg/file/redundancy"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type mockController struct {
	historyMap map[string]accesscontrol.History
	refMap     map[string]swarm.Address
	acceptAll  bool
	publisher  string
	encrypter  encryption.Interface
	ls         file.LoadSaver
}

type optionFunc func(*mockController)

// Option is an option passed to a mock accesscontrol Controller.
type Option interface {
	apply(*mockController)
}

func (f optionFunc) apply(r *mockController) {
	_ = "STUB: not implemented"

	// New creates a new mock accesscontrol Controller.
	return
}

func New(o ...Option) accesscontrol.Controller {
	_ = "STUB: not implemented"
	return *new(accesscontrol.Controller)
}

// WithAcceptAll sets the mock to return fixed references on every call to DownloadHandler.
func WithAcceptAll() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHistory sets the mock to use the given history reference.
func WithHistory(h accesscontrol.History, ref string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPublisher sets the mock to use the given reference as the publisher address.
func WithPublisher(ref string) Option { _ = "STUB: not implemented"; return *new(Option) }

func (m *mockController) DownloadHandler(ctx context.Context, ls file.LoadSaver, encryptedRef swarm.Address, publisher *ecdsa.PublicKey, historyRootHash swarm.Address, timestamp int64) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func (m *mockController) UploadHandler(ctx context.Context, ls file.LoadSaver, reference swarm.Address, publisher *ecdsa.PublicKey, historyRootHash swarm.Address) (swarm.Address, swarm.Address, swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), *new(swarm.Address), *new(swarm.Address), nil
}

func (m *mockController) Close() error { _ = "STUB: not implemented"; return nil }

func (m *mockController) UpdateHandler(_ context.Context, ls file.LoadSaver, gls file.LoadSaver, encryptedglref swarm.Address, historyref swarm.Address, publisher *ecdsa.PublicKey, addList []*ecdsa.PublicKey, removeList []*ecdsa.PublicKey) (swarm.Address, swarm.Address, swarm.Address, swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), *new(swarm.Address), *new(swarm.Address), *new(swarm.Address), nil
}

func (m *mockController) Get(ctx context.Context, ls file.LoadSaver, publisher *ecdsa.PublicKey, encryptedglref swarm.Address) ([]*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func requestPipelineFactory(ctx context.Context, s storage.Putter, encrypt bool, rLevel redundancy.Level) func() pipeline.Interface {
	_ = "STUB: not implemented"
	return nil
}

var _ accesscontrol.Controller = (*mockController)(nil)
