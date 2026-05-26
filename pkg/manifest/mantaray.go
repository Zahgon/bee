// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package manifest

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/file"
	"github.com/ethersphere/bee/v2/pkg/manifest/mantaray"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const (
	// ManifestMantarayContentType represents content type used for noting that
	// specific file should be processed as mantaray manifest.
	ManifestMantarayContentType = "application/bzz-manifest-mantaray+octet-stream"
)

type mantarayManifest struct {
	trie *mantaray.Node

	ls file.LoadSaver
}

// NewMantarayManifest creates a new mantaray-based manifest.
func NewMantarayManifest(
	ls file.LoadSaver,
	encrypted bool,
) (Interface, error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

// use empty obfuscation key if not encrypting

// NOTE: it will be copied to all trie nodes

// NewMantarayManifestReference loads existing mantaray-based manifest.
func NewMantarayManifestReference(
	reference swarm.Address,
	ls file.LoadSaver,
) (Interface, error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

func (m *mantarayManifest) Root() *mantaray.Node { _ = "STUB: not implemented"; return nil }

func (m *mantarayManifest) Type() string { _ = "STUB: not implemented"; return "" }

func (m *mantarayManifest) Add(ctx context.Context, path string, entry Entry) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mantarayManifest) Remove(ctx context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mantarayManifest) Lookup(ctx context.Context, path string) (Entry, error) {
	_ = "STUB: not implemented"
	return *new(Entry), nil
}

func (m *mantarayManifest) HasPrefix(ctx context.Context, prefix string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *mantarayManifest) Store(ctx context.Context, storeSizeFn ...StoreSizeFunc) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func (m *mantarayManifest) IterateAddresses(ctx context.Context, fn swarm.AddressIterFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// The following comparison to the emptyAddr is
// a dirty hack which prevents the walker to
// fail when it encounters an empty address
// (e.g.: during the unpin traversal operation
// for manifest). This workaround should be
// removed after the manifest serialization bug
// is fixed.

type mantarayLoadSaver struct {
	ls          file.LoadSaver
	storeSizeFn []StoreSizeFunc
}

func (ls *mantarayLoadSaver) Load(ctx context.Context, ref []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ls *mantarayLoadSaver) Save(ctx context.Context, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
