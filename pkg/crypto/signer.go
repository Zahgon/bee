// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/crypto/eip712"
)

var ErrInvalidLength = errors.New("invalid signature length")

type Signer interface {
	// Sign signs data with ethereum prefix (eip191 type 0x45).
	Sign(data []byte) ([]byte, error)
	// SignTx signs an ethereum transaction.
	SignTx(transaction *types.Transaction, chainID *big.Int) (*types.Transaction, error)
	// SignTypedData signs data according to eip712.
	SignTypedData(typedData *eip712.TypedData) ([]byte, error)
	// PublicKey returns the public key this signer uses.
	PublicKey() (*ecdsa.PublicKey, error)
	// EthereumAddress returns the ethereum address this signer uses.
	EthereumAddress() (common.Address, error)
}

// addEthereumPrefix adds the ethereum prefix to the data.
func addEthereumPrefix(data []byte) []byte { _ = "STUB: not implemented"; return nil }

// hashWithEthereumPrefix returns the hash that should be signed for the given data.
func hashWithEthereumPrefix(data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Recover verifies signature with the data base provided.
// It is using `btcec.RecoverCompact` function.
func Recover(signature, data []byte) (*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert to btcec input format with 'recovery id' v at the beginning.

type defaultSigner struct {
	key *ecdsa.PrivateKey
}

func NewDefaultSigner(key *ecdsa.PrivateKey) Signer { _ = "STUB: not implemented"; return *new(Signer) }

// PublicKey returns the public key this signer uses.
func (d *defaultSigner) PublicKey() (*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Sign signs data with ethereum prefix (eip191 type 0x45).
}

func (d *defaultSigner) Sign(data []byte) (signature []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignTx signs an ethereum transaction.
func (d *defaultSigner) SignTx(transaction *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isCompressedKey is false here so we get the expected v value (27 or 28)

// v value needs to be adjusted by 27 as transaction.WithSignature expects it to be 0 or 1

// EthereumAddress returns the ethereum address this signer uses.
func (d *defaultSigner) EthereumAddress() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// SignTypedData signs data according to eip712.
func (d *defaultSigner) SignTypedData(typedData *eip712.TypedData) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sign the provided hash and convert it to the ethereum (r,s,v) format.
func (d *defaultSigner) sign(sighash []byte, isCompressedKey bool) ([]byte, error) {
	_ = "STUB: not implemented"
	//nolint:staticcheck // SA1019: ecdsa fields are deprecated, but secp256k1 is not supported by crypto/ecdh
	return nil, nil
}

// Convert to Ethereum signature format with 'recovery id' v at the end.

// RecoverEIP712 recovers the public key for eip712 signed data.
func RecoverEIP712(signature []byte, data *eip712.TypedData) (*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert to btcec input format with 'recovery id' v at the beginning.
