// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// this is a pseudorandom reader that generates a deterministic
// sequence of bytes based on the seed. It is used in tests to
// enable large volumes of pseudorandom data to be generated
// and compared without having to store the data in memory.
package pseudorand

import (
	"io"
)

const bufSize = 4096

// Reader is a pseudorandom reader that generates a deterministic
// sequence of bytes based on the seed.
type Reader struct {
	cur int
	len int
	seg [40]byte
	buf [bufSize]byte
}

// NewSeed creates a new seed.
func NewSeed() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// New creates a new pseudorandom reader seeded with the given seed.
func NewReader(seed []byte, l int) *Reader { _ = "STUB: not implemented"; return nil }

// Size returns the size of the reader.
func (r *Reader) Size() int {
	_ = "STUB: not implemented"

	// Read reads len(buf) bytes into buf.	It returns the number of bytes	 read (0 <= n <= len(buf))	and any error encountered.	Even if Read returns n < len(buf), it may use all of buf as scratch space during the call. If some data is available but not len(buf) bytes, Read conventionally returns what is available instead of waiting for more.
	return 0
}

func (r *Reader) Read(buf []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Equal compares the contents of the reader with the contents of
// the given reader. It returns true if the contents are equal upto n bytes
func (r1 *Reader) Equal(r2 io.Reader) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Match compares the contents of the reader with the contents of
// the given reader. It returns true if the contents are equal upto n bytes
func (r1 *Reader) Match(r2 io.Reader, l int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Seek sets the offset for the next Read to offset, interpreted
// according to whence: 0 means relative to the start of the file,
// 1 means relative to the current offset, and 2 means relative to
// the end. It returns the new offset and an error, if any.
func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Offset returns the current offset of the reader.
func (r *Reader) Offset() int64 { _ = "STUB: not implemented"; return 0 }

// ReadAt reads len(buf) bytes into buf starting at offset off.
func (r *Reader) ReadAt(buf []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// fill fills the buffer with the hash of the current segment.
func (r *Reader) fill() { _ = "STUB: not implemented"; return }
