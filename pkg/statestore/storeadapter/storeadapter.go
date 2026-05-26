// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storeadapter

import (
	"github.com/ethersphere/bee/v2/pkg/storage"
)

// stateStoreNamespace is the namespace used for state storage.
const stateStoreNamespace = "ss"

var _ storage.Item = (*proxyItem)(nil)

// proxyItem is a proxy object that implements the Item interface.
// It is an intermediary object between StateStorer and Store interfaces calls.
type proxyItem struct {
	ns  string
	key string
	obj any
}

// ID implements Item interface.
func (pi *proxyItem) ID() string {
	_ = "STUB: not implemented"

	// Namespace implements Item interface.
	return ""
}

func (pi *proxyItem) Namespace() string {
	_ = "STUB: not implemented"

	// Marshal implements Item interface.
	return ""
}

func (pi *proxyItem) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal implements Item interface.
func (pi *proxyItem) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

// Clone implements Item interface.
func (pi *proxyItem) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

// String implements Item interface.
func (pi proxyItem) String() string { _ = "STUB: not implemented"; return "" }

// newProxyItem creates a new proxyItem.
func newProxyItem(key string, obj any) *proxyItem { _ = "STUB: not implemented"; return nil }

var _ storage.Item = (*rawItem)(nil)

// rawItem is a proxy object that implements the Item interface.
// It is an intermediary object between StateStorer and Store iterator calls.
type rawItem struct {
	*proxyItem
}

// Marshal implements Item interface.
func (ri *rawItem) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal implements Item interface.
func (ri *rawItem) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

var (
	_ storage.StateStorer        = (*StateStorerAdapter)(nil)
	_ storage.StateStorerCleaner = (*StateStorerAdapter)(nil)
)

// StateStorerAdapter is an adapter from Store to the StateStorer.
type StateStorerAdapter struct {
	storage storage.Store
}

// Close implements StateStorer interface.
func (s *StateStorerAdapter) Close() error { _ = "STUB: not implemented"; return nil }

// Get implements StateStorer interface.
func (s *StateStorerAdapter) Get(key string, obj any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Put implements StateStorer interface.
func (s *StateStorerAdapter) Put(key string, obj any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Delete implements StateStorer interface.
func (s *StateStorerAdapter) Delete(key string) (err error) { _ = "STUB: not implemented"; return nil }

// Iterate implements StateStorer interface.
func (s *StateStorerAdapter) Iterate(prefix string, iterFunc storage.StateIterFunc) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *StateStorerAdapter) Nuke() error { _ = "STUB: not implemented"; return nil }

func (s *StateStorerAdapter) ClearForHopping() error { _ = "STUB: not implemented"; return nil }

// to not redeploy chequebook contract
// avoid unnecessary syncing
// to not resync blockchain transactions

func (s *StateStorerAdapter) collectKeysExcept(prefixesToPreserve []string) (keys []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StateStorerAdapter) deleteKeys(keys []string) error { _ = "STUB: not implemented"; return nil }

// NewStateStorerAdapter creates a new StateStorerAdapter.
func NewStateStorerAdapter(storage storage.Store) (*StateStorerAdapter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
