// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package manifest

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/file"
	"github.com/ethersphere/bee/v2/pkg/manifest/simple"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const (
	// ManifestSimpleContentType represents content type used for noting that
	// specific file should be processed as 'simple' manifest
	ManifestSimpleContentType = "application/bzz-manifest-simple+json"
)

type simpleManifest struct {
	manifest simple.Manifest

	reference swarm.Address
	ls        file.LoadSaver
}

// NewSimpleManifest creates a new simple manifest.
func NewSimpleManifest(ls file.LoadSaver) (Interface, error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

// NewSimpleManifestReference loads existing simple manifest.
func NewSimpleManifestReference(ref swarm.Address, l file.LoadSaver) (Interface, error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

func (m *simpleManifest) Type() string { _ = "STUB: not implemented"; return "" }

func (m *simpleManifest) Add(_ context.Context, path string, entry Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *simpleManifest) Remove(_ context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *simpleManifest) Lookup(_ context.Context, path string) (Entry, error) {
	_ = "STUB: not implemented"
	return *new(Entry), nil
}

func (m *simpleManifest) HasPrefix(_ context.Context, prefix string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *simpleManifest) Store(ctx context.Context, storeSizeFn ...StoreSizeFunc) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func (m *simpleManifest) IterateAddresses(ctx context.Context, fn swarm.AddressIterFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: making it behave same for all manifest implementation

func (m *simpleManifest) load(ctx context.Context, reference swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}
