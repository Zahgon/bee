// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eip712

import (
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// type aliases to avoid importing "core" everywhere
type (
	TypedData        = apitypes.TypedData
	TypedDataDomain  = apitypes.TypedDataDomain
	Types            = apitypes.Types
	Type             = apitypes.Type
	TypedDataMessage = apitypes.TypedDataMessage
)

// EncodeForSigning encodes the hash that will be signed for the given EIP712 data
func EncodeForSigning(typedData *TypedData) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EIP712DomainType is the type description for the EIP712 Domain
var EIP712DomainType = []Type{
	{
		Name: "name",
		Type: "string",
	},
	{
		Name: "version",
		Type: "string",
	},
	{
		Name: "chainId",
		Type: "uint256",
	},
}
