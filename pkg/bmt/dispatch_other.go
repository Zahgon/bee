// Copyright 2026 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux || !amd64 || purego

package bmt

// NewPool returns a BMT pool. On this platform only the goroutine implementation
// is compiled in, so SIMDOptIn() is ignored.
func NewPool(c *Conf) Pool { _ = "STUB: not implemented"; return *new(Pool) }

// NewHasher returns a standalone (non-pooled) BMT hasher.
func NewHasher() Hasher { _ = "STUB: not implemented"; return *new(Hasher) }

// NewPrefixHasher returns a standalone BMT hasher with the given prefix.
func NewPrefixHasher(prefix []byte) Hasher { _ = "STUB: not implemented"; return *new(Hasher) }
