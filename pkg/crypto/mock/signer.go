// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/crypto/eip712"
)

type signerMock struct {
	signTx          func(transaction *types.Transaction, chainID *big.Int) (*types.Transaction, error)
	signTypedData   func(*eip712.TypedData) ([]byte, error)
	ethereumAddress func() (common.Address, error)
	signFunc        func([]byte) ([]byte, error)
}

func (m *signerMock) EthereumAddress() (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func (m *signerMock) Sign(data []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *signerMock) SignTx(transaction *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*signerMock) PublicKey() (*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *signerMock) SignTypedData(d *eip712.TypedData) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func New(opts ...Option) crypto.Signer { _ = "STUB: not implemented"; return *new(crypto.Signer) }

// Option is the option passed to the mock Chequebook service
type Option interface {
	apply(*signerMock)
}

type optionFunc func(*signerMock)

func (f optionFunc) apply(r *signerMock) { _ = "STUB: not implemented"; return }

func WithSignFunc(f func(data []byte) ([]byte, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSignTxFunc(f func(transaction *types.Transaction, chainID *big.Int) (*types.Transaction, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSignTypedDataFunc(f func(*eip712.TypedData) ([]byte, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithEthereumAddressFunc(f func() (common.Address, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
