// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/soc"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// MockSOC defines a mocked SOC with exported fields for easy testing.
type MockSOC struct {
	ID           soc.ID
	Owner        []byte
	Signature    []byte
	WrappedChunk swarm.Chunk
}

// Address returns the SOC address of the mocked SOC.
func (ms MockSOC) Address() swarm.Address { _ = "STUB: not implemented"; return *new(swarm.Address) }

// Chunk returns the SOC chunk of the mocked SOC.
func (ms MockSOC) Chunk() swarm.Chunk { _ = "STUB: not implemented"; return *new(swarm.Chunk) }

// GenerateMockSocWithSigner generates a valid mocked SOC from given data and signer.
func GenerateMockSocWithSigner(t *testing.T, data []byte, signer crypto.Signer) *MockSOC {
	_ = "STUB: not implemented"
	return nil
}

// GenerateMockSOC generates a valid mocked SOC from given data.
func GenerateMockSOC(t *testing.T, data []byte) *MockSOC { _ = "STUB: not implemented"; return nil }

// GenerateMockSOCWithSpan generates a valid mocked SOC from given chunk data (span + payload).
func GenerateMockSOCWithSpan(t *testing.T, data []byte) *MockSOC {
	_ = "STUB: not implemented"
	return nil
}

func generateMockSOC(t *testing.T, ch swarm.Chunk) *MockSOC { _ = "STUB: not implemented"; return nil }

// GenerateMockSOCWithKey generates a valid mocked SOC from given data and key.
func GenerateMockSOCWithKey(t *testing.T, data []byte, privKey *ecdsa.PrivateKey) *MockSOC {
	_ = "STUB: not implemented"
	return nil
}
