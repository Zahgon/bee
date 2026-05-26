// Copyright 2025 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapped

import (
	"context"
	"errors"
	"math/big"
)

const (
	percentageDivisor = 100
	baseFeeMultiplier = 2
)

var ErrEIP1559NotSupported = errors.New("network does not appear to support EIP-1559 (no baseFee)")

// SuggestedFeeAndTip calculates the recommended gasFeeCap (maxFeePerGas) and gasTipCap (maxPriorityFeePerGas) for a transaction.
// If gasPrice is provided (legacy mode):
//   - On EIP-1559 networks: gasFeeCap = gasPrice; gasTipCap = max(gasPrice - baseFee, minimumTip) to respect the total cap while enforcing a tip floor where possible.
//   - On pre-EIP-1559 networks: returns (gasPrice, gasPrice) for legacy transaction compatibility.
//
// If gasPrice is nil: Uses suggested tip with optional boost, enforces minimum, and sets gasFeeCap = 2 * baseFee + gasTipCap.
func (b *wrappedBackend) SuggestedFeeAndTip(ctx context.Context, gasPrice *big.Int, boostPercent int) (*big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// nominal tip = gasPrice - baseFee

// multiplier: 100 + boostPercent (e.g., 110 for 10% boost)

// gasTipCap = gasTipCap * (100 + boostPercent) / 100

// gasFeeCap = (2 * baseFee) + gasTipCap
