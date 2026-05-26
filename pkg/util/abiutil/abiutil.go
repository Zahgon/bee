// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abiutil

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

// MustParseABI parses is the same as calling abi.JSON
// but panics on error (if the given ABI is invalid).
func MustParseABI(json string) abi.ABI { _ = "STUB: not implemented"; return *new(abi.ABI) }
