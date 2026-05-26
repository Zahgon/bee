// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto

import (
	"crypto/ecdsa"
)

// edgSecp256_k1 aggregates private key cryptography functions that employ secp256k1
type edgSecp256_k1 struct{}

var EDGSecp256_K1 = new(edgSecp256_k1)

func (s *edgSecp256_k1) Generate() (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *edgSecp256_k1) Encode(k *ecdsa.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *edgSecp256_k1) Decode(data []byte) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// edgSecp256_r1 aggregates private key cryptography functions that employ secp256r1
type edgSecp256_r1 struct{}

var EDGSecp256_R1 = new(edgSecp256_r1)

func (s *edgSecp256_r1) Generate() (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *edgSecp256_r1) Encode(k *ecdsa.PrivateKey) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *edgSecp256_r1) Decode(data []byte) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
