// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"crypto/ecdsa"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type GranteeListMock interface {
	Add(publicKeys []*ecdsa.PublicKey) error
	Remove(removeList []*ecdsa.PublicKey) error
	Get() []*ecdsa.PublicKey
	Save() (swarm.Address, error)
}

type GranteeListStructMock struct {
	grantees []*ecdsa.PublicKey
}

func (g *GranteeListStructMock) Get() []*ecdsa.PublicKey { _ = "STUB: not implemented"; return nil }

func (g *GranteeListStructMock) Add(addList []*ecdsa.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GranteeListStructMock) Remove(removeList []*ecdsa.PublicKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *GranteeListStructMock) Save() (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func NewGranteeList() *GranteeListStructMock { _ = "STUB: not implemented"; return nil }
