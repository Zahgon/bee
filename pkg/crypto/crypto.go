// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// RecoverFunc is a function to recover the public key from a signature
type RecoverFunc func(signature, data []byte) (*ecdsa.PublicKey, error)

var ErrBadHashLength = errors.New("wrong block hash length")

const (
	AddressSize = 20
)

// NewOverlayAddress constructs a Swarm Address from ECDSA public key.
func NewOverlayAddress(p ecdsa.PublicKey, networkID uint64, nonce []byte) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// NewOverlayFromEthereumAddress constructs a Swarm Address for an Ethereum address.
func NewOverlayFromEthereumAddress(ethAddr []byte, networkID uint64, nonce []byte) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// GenerateSecp256k1Key generates an ECDSA private key using
// secp256k1 elliptic curve.
func GenerateSecp256k1Key() (*ecdsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

// EncodeSecp256k1PrivateKey encodes raw ECDSA private key.
func EncodeSecp256k1PrivateKey(k *ecdsa.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	//nolint:staticcheck // SA1019: ecdsa fields are deprecated, but secp256k1 is not supported by crypto/ecdh
	return nil, nil
}

// EncodeSecp256k1PublicKey encodes raw ECDSA public key in a 33-byte compressed format.
func EncodeSecp256k1PublicKey(k *ecdsa.PublicKey) []byte { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck // SA1019: ecdsa fields are deprecated, but secp256k1 is not supported by crypto/ecdh

//nolint:staticcheck // SA1019: ecdsa fields are deprecated, but secp256k1 is not supported by crypto/ecdh

// DecodeSecp256k1PrivateKey decodes raw ECDSA private key.
func DecodeSecp256k1PrivateKey(data []byte) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateSecp256r1Key generates an ECDSA private key using
// secp256k1 elliptic curve.
func GenerateSecp256r1Key() (*ecdsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

// EncodeSecp256r1PrivateKey encodes raw ECDSA private key.
func EncodeSecp256r1PrivateKey(k *ecdsa.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecodeSecp256r1PrivateKey decodes raw ECDSA private key.
func DecodeSecp256r1PrivateKey(data []byte) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Secp256k1PrivateKeyFromBytes returns an ECDSA private key based on
// the byte slice.
func Secp256k1PrivateKeyFromBytes(data []byte) *ecdsa.PrivateKey {
	_ = "STUB: not implemented"
	return nil
}

// NewEthereumAddress returns a binary representation of ethereum blockchain address.
// This function is based on github.com/ethereum/go-ethereum/crypto.PubkeyToAddress.
func NewEthereumAddress(p ecdsa.PublicKey) ([]byte, error) {
	_ = "STUB: not implemented"
	//nolint:staticcheck // SA1019: ecdsa fields are deprecated, but secp256k1 is not supported by crypto/ecdh
	return nil, nil
}

//nolint:staticcheck // SA1019: ecdsa fields are deprecated, but secp256k1 is not supported by crypto/ecdh

func LegacyKeccak256(data []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
