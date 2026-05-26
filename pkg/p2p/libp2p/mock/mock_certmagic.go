// Copyright 2025 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package mock

import (
	"context"
	"crypto/tls"
	"sync"

	_ "embed"

	"github.com/libp2p/go-libp2p/config"
	"github.com/libp2p/go-libp2p/core/host"
)

//go:embed testdata/cert.pem
var certPEM []byte

//go:embed testdata/key.pem
var keyPEM []byte

// MockFileStorage is a minimal implementation of certmagic.FileStorage for testing.
type MockFileStorage struct {
	path string
}

func NewMockFileStorage(path string) *MockFileStorage { _ = "STUB: not implemented"; return nil }

func (m *MockFileStorage) Store(_ context.Context, _ string, _ []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockFileStorage) Load(_ context.Context, _ string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockFileStorage) Delete(_ context.Context, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockFileStorage) Exists(_ context.Context, _ string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *MockFileStorage) List(_ context.Context, _ string, _ bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockFileStorage) Lock(_ context.Context, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockFileStorage) Unlock(_ context.Context, _ string) error {
	_ = "STUB: not implemented"

	// MockCache is a minimal implementation of certmagic.Cache for testing.
	return nil
}

type MockCache struct {
	stopped bool
	mu      sync.Mutex
}

func NewMockCache() *MockCache { _ = "STUB: not implemented"; return nil }

func (m *MockCache) Stop() { _ = "STUB: not implemented"; return }

// MockConfig is a minimal implementation of certmagic.Config for testing.
type MockConfig struct {
	cache *MockCache
}

func NewMockConfig() *MockConfig { _ = "STUB: not implemented"; return nil }

func (m *MockConfig) TLSConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

// MockP2PForgeCertMgr is a mock implementation of p2pforge.P2PForgeCertMgr.
type MockP2PForgeCertMgr struct {
	cache        *MockCache
	onCertLoaded func()
	started      bool
	mu           sync.Mutex
	ProvideHost  func(host.Host) error
}

func NewMockP2PForgeCertMgr(onCertLoaded func()) *MockP2PForgeCertMgr {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockP2PForgeCertMgr) Start() error { _ = "STUB: not implemented"; return nil }

func (m *MockP2PForgeCertMgr) Stop() { _ = "STUB: not implemented"; return }

func (m *MockP2PForgeCertMgr) TLSConfig() *tls.Config {
	_ = "STUB: not implemented"
	// Use tls.X509KeyPair to create a certificate from the hardcoded strings
	return nil
}

// This should not fail if the strings are pasted correctly

func (m *MockP2PForgeCertMgr) SetOnCertLoaded(cb func()) { _ = "STUB: not implemented"; return }

func (m *MockP2PForgeCertMgr) AddressFactory() config.AddrsFactory {
	_ = "STUB: not implemented"
	return *new(config.AddrsFactory)
}

func (m *MockP2PForgeCertMgr) GetCache() *MockCache { _ = "STUB: not implemented"; return nil }
