// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nbhdutil

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

func MineOverlay(ctx context.Context, p ecdsa.PublicKey, networkID uint64, targetNeighborhood string) (swarm.Address, []byte, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil, nil
}
