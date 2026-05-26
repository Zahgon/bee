// Copyright 2025 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libp2ptest

import (
	"testing"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p/libp2p"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// NewLibp2pService creates a new libp2p service for testing purposes.
func NewLibp2pService(t *testing.T, networkID uint64, logger log.Logger) (*libp2p.Service, swarm.Address) {
	_ = "STUB: not implemented"
	return nil, *new(swarm.Address)
}

// Disable default NAT manager
