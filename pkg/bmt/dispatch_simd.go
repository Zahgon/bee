// Copyright 2026 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && amd64 && !purego

package bmt

// NewPool returns a BMT pool. If SIMDOptIn() is true and the CPU exposes AVX2
// or AVX-512, a SIMD-batched pool is returned. Otherwise the goroutine-based
// pool is returned (silent fallback).
func NewPool(c *Conf) Pool { _ = "STUB: not implemented"; return *new(Pool) }

// NewHasher returns a standalone (non-pooled) BMT hasher.
func NewHasher() Hasher { _ = "STUB: not implemented"; return *new(Hasher) }

// NewPrefixHasher returns a standalone BMT hasher with the given prefix.
func NewPrefixHasher(prefix []byte) Hasher { _ = "STUB: not implemented"; return *new(Hasher) }
