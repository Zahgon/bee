// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package addressbook

import (
	"errors"

	"github.com/ethersphere/bee/v2/pkg/bzz"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const keyPrefix = "addressbook_entry_"

var _ Interface = (*store)(nil)

var ErrNotFound = errors.New("addressbook: not found")

// Interface is the AddressBook interface.
type Interface interface {
	GetPutter
	Remover
	// Overlays returns a list of all overlay addresses saved in addressbook.
	Overlays() ([]swarm.Address, error)
	// IterateOverlays exposes overlays in a form of an iterator.
	IterateOverlays(func(swarm.Address) (bool, error)) error
	// Addresses returns a list of all bzz.Address-es saved in addressbook.
	Addresses() ([]bzz.Address, error)
}

type GetPutter interface {
	Getter
	Putter
}

type Getter interface {
	// Get returns pointer to saved bzz.Address for requested overlay address.
	Get(overlay swarm.Address) (addr *bzz.Address, err error)
}

type Putter interface {
	// Put saves relation between peer overlay address and bzz.Address address.
	Put(overlay swarm.Address, addr bzz.Address) (err error)
}

type Remover interface {
	// Remove removes overlay address.
	Remove(overlay swarm.Address) error
}

type store struct {
	store storage.StateStorer
}

// New creates new addressbook for state storer.
func New(storer storage.StateStorer) Interface { _ = "STUB: not implemented"; return *new(Interface) }

func (s *store) Get(overlay swarm.Address) (*bzz.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *store) Put(overlay swarm.Address, addr bzz.Address) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *store) Remove(overlay swarm.Address) error { _ = "STUB: not implemented"; return nil }

func (s *store) IterateOverlays(cb func(swarm.Address) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *store) Overlays() (overlays []swarm.Address, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *store) Addresses() (addresses []bzz.Address, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
