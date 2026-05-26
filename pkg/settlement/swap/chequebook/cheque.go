// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chequebook

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/crypto/eip712"
)

// Cheque represents a cheque for a SimpleSwap chequebook
type Cheque struct {
	Chequebook       common.Address
	Beneficiary      common.Address
	CumulativePayout *big.Int
}

// SignedCheque represents a cheque together with its signature
type SignedCheque struct {
	Cheque
	Signature []byte
}

// chequebookDomain computes chainId-dependent EIP712 domain
func chequebookDomain(chainID int64) eip712.TypedDataDomain {
	_ = "STUB: not implemented"
	return *new(eip712.TypedDataDomain)
}

// ChequeTypes are the needed type descriptions for cheque signing
var ChequeTypes = eip712.Types{
	"EIP712Domain": eip712.EIP712DomainType,
	"Cheque": []eip712.Type{
		{
			Name: "chequebook",
			Type: "address",
		},
		{
			Name: "beneficiary",
			Type: "address",
		},
		{
			Name: "cumulativePayout",
			Type: "uint256",
		},
	},
}

// ChequeSigner signs cheque
type ChequeSigner interface {
	// Sign signs a cheque
	Sign(cheque *Cheque) ([]byte, error)
}

type chequeSigner struct {
	signer  crypto.Signer // the underlying signer used
	chainID int64         // the chainID used for EIP712
}

// NewChequeSigner creates a new cheque signer for the given chainID.
func NewChequeSigner(signer crypto.Signer, chainID int64) ChequeSigner {
	_ = "STUB: not implemented"
	return *new(ChequeSigner)
}

// eip712DataForCheque converts a cheque into the correct TypedData structure.
func eip712DataForCheque(cheque *Cheque, chainID int64) *eip712.TypedData {
	_ = "STUB: not implemented"
	return nil
}

// Sign signs a cheque.
func (s *chequeSigner) Sign(cheque *Cheque) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cheque *Cheque) String() string { _ = "STUB: not implemented"; return "" }

func (cheque *Cheque) Equal(other *Cheque) bool { _ = "STUB: not implemented"; return false }

func (cheque *SignedCheque) Equal(other *SignedCheque) bool {
	_ = "STUB: not implemented"
	return false
}
