// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kademlia

import (
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// generateCommonBinPrefixes generates the common bin prefixes
// used by the bin balancer.
func generateCommonBinPrefixes(base swarm.Address, suffixLength int) [][]swarm.Address {
	_ = "STUB: not implemented"
	return nil
}

// copy base address

// flip first bit for bin

// set pseudo suffix

// clear rest of the bits

// Clears the bit at pos in n.
func clearBit(n, pos uint8) uint8 { _ = "STUB: not implemented"; return 0 }

// Sets the bit at pos in the integer n.
func setBit(n, pos uint8) uint8 { _ = "STUB: not implemented"; return 0 }

func hasBit(n, pos uint8) bool { _ = "STUB: not implemented"; return false }
