// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var Spans []int64

// nolint:gochecknoinits
func init() {
	Spans = GenerateSpanSizes(9, swarm.Branches)
}

// GenerateSpanSizes generates a dictionary of maximum span lengths per level represented by one SectionSize() of data
func GenerateSpanSizes(levels, branches int) []int64 { _ = "STUB: not implemented"; return nil }

// Levels calculates the last level index which a particular data section count will result in.
// The returned level will be the level of the root hash.
func Levels(length int64, sectionSize, branches int) int { _ = "STUB: not implemented"; return 0 }
