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

// Package encryption exposes functionalities needed for
// encryption and decryption operations in Bee.
package encryption

import (
	"hash"
)

const (
	KeyLength     = 32
	ReferenceSize = 64
)

type Key []byte

type Encrypter interface {
	Key() Key
	Encrypt(data []byte) ([]byte, error)
}

type Decrypter interface {
	Key() Key
	Decrypt(data []byte) ([]byte, error)
}

type Interface interface {
	Encrypter
	Decrypter
	Reset()
}

type Encryption struct {
	key      Key              // the encryption key (hashSize bytes long)
	keyLen   int              // length of the key = length of blockcipher block
	padding  int              // encryption will pad the data upto this if > 0
	index    int              // counter index
	initCtr  uint32           // initial counter used for counter mode blockcipher
	hashFunc func() hash.Hash // hasher constructor function
}

// New constructs a new encrypter/decrypter
func New(key Key, padding int, initCtr uint32, hashFunc func() hash.Hash) Interface {
	_ = "STUB: not implemented"
	return *new(Interface)
}

// Key returns the base key
func (e *Encryption) Key() Key {
	_ = "STUB: not implemented"

	// Encrypt encrypts the data and does padding if specified
	return *new(Key)
}

func (e *Encryption) Encrypt(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decrypt decrypts the data, if padding was used caller must know original length and truncate
func (e *Encryption) Decrypt(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset resets the counter. It is only safe to call after an encryption operation is completed
// After Reset is called, the Encryption object can be reused for other data
func (e *Encryption) Reset() {
	_ = "STUB: not implemented"

	// split up input into keylength segments and encrypt sequentially
	return
}

func (e *Encryption) transform(in, out []byte) error { _ = "STUB: not implemented"; return nil }

// pad the rest if out is longer

// used for segmentwise transformation
// if in is shorter than out, padding is used
func (e *Encryption) Transcrypt(i int, in, out []byte) error {
	_ = "STUB: not implemented"
	// first hash key with counter (initial counter + i)
	return nil
}

// second round of hashing for selective disclosure

// XOR bytes uptil length of in (out must be at least as long)

// insert padding if out is longer

func pad(b []byte) { _ = "STUB: not implemented"; return }

// GenerateRandomKey generates a random key of length l
func GenerateRandomKey(l int) Key { _ = "STUB: not implemented"; return *new(Key) }
