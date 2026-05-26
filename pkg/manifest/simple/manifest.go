// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simple

import (
	"encoding"
	"errors"
	"sync"
)

// Error used when lookup path does not match
var (
	ErrNotFound  = errors.New("not found")
	ErrEmptyPath = errors.New("empty path")
)

// Manifest is a representation of a manifest.
type Manifest interface {
	// Add adds a manifest entry to the specified path.
	Add(string, string, map[string]string) error
	// Remove removes a manifest entry on the specified path.
	Remove(string) error
	// Lookup returns a manifest node entry if one is found in the specified path.
	Lookup(string) (Entry, error)
	// HasPrefix tests whether the specified prefix path exists.
	HasPrefix(string) bool
	// Length returns an implementation-specific count of elements in the manifest.
	// For Manifest, this means the number of all the existing entries.
	Length() int

	// WalkEntry walks all entries, calling walkFn for each entry in the map.
	// All errors that arise visiting entries are filtered by walkFn.
	WalkEntry(string, WalkEntryFunc) error

	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

// manifest is a JSON representation of a manifest.
// It stores manifest entries in a map based on string keys.
type manifest struct {
	Entries map[string]*entry `json:"entries,omitempty"`

	mu sync.RWMutex // mutex for accessing the entries map
}

// NewManifest creates a new Manifest struct and returns a pointer to it.
func NewManifest() Manifest { _ = "STUB: not implemented"; return *new(Manifest) }

func notFound(path string) error { _ = "STUB: not implemented"; return nil }

func (m *manifest) Add(path, entry string, metadata map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manifest) Remove(path string) error { _ = "STUB: not implemented"; return nil }

func (m *manifest) Lookup(path string) (Entry, error) {
	_ = "STUB: not implemented"
	return *new(Entry), nil
}

// return a copy to prevent external modification

func (m *manifest) HasPrefix(path string) bool { _ = "STUB: not implemented"; return false }

func (m *manifest) Length() int { _ = "STUB: not implemented"; return 0 }

// MarshalBinary implements encoding.BinaryMarshaler.
func (m *manifest) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (m *manifest) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }
