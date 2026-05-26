// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package swarm

import (
	"hash"
)

// NewHasher returns new Keccak-256 hasher.
func NewHasher() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

type PrefixHasher struct {
	hash.Hash
	prefix []byte
}

// NewPrefixHasher returns new hasher which is Keccak-256 hasher
// with prefix value added as initial data.
func NewPrefixHasher(prefix []byte) hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

func (h *PrefixHasher) Reset() { _ = "STUB: not implemented"; return }
