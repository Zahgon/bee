// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package swarm

import (
	"math/big"
)

// Distance returns the distance between address x and address y as a (comparable) big integer using the distance metric defined in the swarm specification.
// Fails if not all addresses are of equal length.
func Distance(x, y Address) (*big.Int, error) { _ = "STUB: not implemented"; return nil, nil }

// DistanceRaw returns the distance between address x and address y in big-endian binary format using the distance metric defined in the swarm specification.
// Fails if not all addresses are of equal length.
func DistanceRaw(x, y Address) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// DistanceCmp compares x and y to a in terms of the distance metric defined in the swarm specification.
// it returns:
//   - 1 if x is closer to a than y
//   - 0 if x and y are equally far apart from a (this means that x and y are the same address)
//   - -1 if x is farther from a than y
//
// Fails if not all addresses are of equal length.
func DistanceCmp(a, x, y Address) (int, error) { _ = "STUB: not implemented"; return 0, nil }
