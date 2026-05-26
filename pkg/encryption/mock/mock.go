// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"errors"

	"github.com/ethersphere/bee/v2/pkg/encryption"
)

var _ encryption.Interface = (*Encryptor)(nil)

var (
	// ErrNotImplemented is returned when a required Encryptor function is not set.
	ErrNotImplemented = errors.New("not implemented")
	// ErrInvalidXORKey is returned when the key for XOR encryption is not valid.
	ErrInvalidXORKey = errors.New("invalid xor key")
)

// Encryptor implements encryption Interface in order to mock it in tests.
type Encryptor struct {
	encryptFunc func(data []byte) ([]byte, error)
	decryptFunc func(data []byte) ([]byte, error)
	resetFunc   func()
	keyFunc     func() encryption.Key
}

// New returns a new Encryptor configured with provided options.
func New(opts ...Option) *Encryptor { _ = "STUB: not implemented"; return nil }

// Key has only bogus
func (e *Encryptor) Key() encryption.Key { _ = "STUB: not implemented"; return *new(encryption.Key) }

// Encrypt calls the configured encrypt function, or returns ErrNotImplemented
// if it is not set.
func (e *Encryptor) Encrypt(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Decrypt calls the configured decrypt function, or returns ErrNotImplemented
// if it is not set.
func (e *Encryptor) Decrypt(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset calls the configured reset function, if it is set.
func (e *Encryptor) Reset() { _ = "STUB: not implemented"; return }

// Option represents configures the Encryptor instance.
type Option func(*Encryptor)

// WithEncryptFunc sets the Encryptor Encrypt function.
func WithEncryptFunc(f func([]byte) ([]byte, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDecryptFunc sets the Encryptor Decrypt function.
func WithDecryptFunc(f func([]byte) ([]byte, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithResetFunc sets the Encryptor Reset function.
func WithResetFunc(f func()) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithKeyFunc sets the Encryptor Key function.
func WithKeyFunc(f func() encryption.Key) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithXOREncryption sets Encryptor Encrypt and Decrypt functions with XOR
// encryption function that uses the provided key for encryption.
func WithXOREncryption(key []byte) Option { _ = "STUB: not implemented"; return *new(Option) }

func newXORFunc(key []byte) func([]byte) ([]byte, error) { _ = "STUB: not implemented"; return nil }

func xor(input, key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
