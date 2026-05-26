// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package bitvector provides functionality of a
// simple bit vector implementation.
package bitvector

import (
	"errors"
)

var errInvalidLength = errors.New("invalid length")

// BitVector is a convenience object for manipulating and representing bit vectors
type BitVector struct {
	len int
	b   []byte
}

// New creates a new bit vector with the given length
func New(l int) (*BitVector, error) { _ = "STUB: not implemented"; return nil, nil }

// NewFromBytes creates a bit vector from the passed byte slice.
//
// Leftmost bit in byte slice becomes leftmost bit in bit vector
func NewFromBytes(b []byte, l int) (*BitVector, error) { _ = "STUB: not implemented"; return nil, nil }

// Get gets the corresponding bit, counted from left to right
func (bv *BitVector) Get(i int) bool { _ = "STUB: not implemented"; return false }

// Set sets the bit corresponding to the index in the bitvector, counted from left to right
func (bv *BitVector) Set(i int) { _ = "STUB: not implemented"; return }

// Bytes retrieves the underlying bytes of the bitvector
func (bv *BitVector) Bytes() []byte { _ = "STUB: not implemented"; return nil }
