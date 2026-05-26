// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package migration

import (
	"errors"

	storage "github.com/ethersphere/bee/v2/pkg/storage"
)

type (
	// StepFn is a function that migrates the storage to the next version
	StepFn func() error
	// Steps is a map of versions and their migration functions
	Steps = map[uint64]StepFn
)

// errStorageVersionItemUnmarshalInvalidSize is returned when trying
// to unmarshal buffer that is not of size storageVersionItemSize.
var errStorageVersionItemUnmarshalInvalidSize = errors.New("unmarshal StorageVersionItem: invalid size")

// Migrate migrates the storage to the latest version.
// The steps are separated by groups so different lists of steps can run individually, for example,
// two groups of migrations that run before and after the storer is initialized.
func Migrate(s storage.IndexStore, group string, sm Steps) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateVersions checks versions if they are in order n (where n min version value), n+1, n+2, n+3... (all values are increasing orders)
func ValidateVersions(sm Steps) error { _ = "STUB: not implemented"; return nil }

var _ storage.Item = (*StorageVersionItem)(nil)

// storageVersionItemSize is the size of the marshaled storage version item.
const storageVersionItemSize = 8

type StorageVersionItem struct {
	Version uint64
	Group   string
}

// ID implements the storage.Item interface.
func (s *StorageVersionItem) ID() string { _ = "STUB: not implemented"; return "" }

// Namespace implements the storage.Item interface.
func (s StorageVersionItem) Namespace() string {
	_ = "STUB: not implemented"

	// Marshal implements the storage.Item interface.
	return ""
}

func (s *StorageVersionItem) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal implements the storage.Item interface.
func (s *StorageVersionItem) Unmarshal(bytes []byte) error { _ = "STUB: not implemented"; return nil }

// Clone implements the storage.Item interface.
func (s *StorageVersionItem) Clone() storage.Item {
	_ = "STUB: not implemented"
	return *new(storage.Item)
}

// Clone implements the storage.Item interface.
func (s StorageVersionItem) String() string { _ = "STUB: not implemented"; return "" }

// Version returns the current version of the storage
func Version(s storage.Reader, group string) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// setVersion sets the current version of the storage
func setVersion(s storage.Writer, v uint64, g string) error { _ = "STUB: not implemented"; return nil }

// LatestVersion returns latest version from supplied migration steps.
func LatestVersion(sm Steps) uint64 { _ = "STUB: not implemented"; return 0 }
