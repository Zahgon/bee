// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bzz exposes the data structure and operations
// necessary on the bzz.Address type which used in the handshake
// protocol, address-book and hive protocol.
package bzz

import (
	"errors"

	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/swarm"

	ma "github.com/multiformats/go-multiaddr"
)

var ErrInvalidAddress = errors.New("invalid address")

// Address represents the bzz address in swarm.
// It consists of a peers underlay (physical) address, overlay (topology) address and signature.
// Signature is used to verify the `Overlay/Underlay` pair, as it is based on `underlay|networkID`, signed with the public key of Overlay address
type Address struct {
	Underlays       []ma.Multiaddr
	Overlay         swarm.Address
	Signature       []byte
	Nonce           []byte
	EthereumAddress []byte
}

type addressJSON struct {
	Overlay   string   `json:"overlay"`
	Underlay  string   `json:"underlay"` // For backward compatibility
	Underlays []string `json:"underlays"`
	Signature string   `json:"signature"`
	Nonce     string   `json:"transaction"`
}

func NewAddress(signer crypto.Signer, underlays []ma.Multiaddr, overlay swarm.Address, networkID uint64, nonce []byte) (*Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseAddress(underlay, overlay, signature, nonce []byte, validateOverlay bool, networkID uint64) (*Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// no underlays sent

func generateSignData(underlay, overlay []byte, networkID uint64) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (a *Address) Equal(b *Address) bool { _ = "STUB: not implemented"; return false }

func AreUnderlaysEqual(a, b []ma.Multiaddr) bool { _ = "STUB: not implemented"; return false }

func (a *Address) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// select the underlay address for backward compatibility

func (a *Address) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// append the underlay for backward compatibility

func (a *Address) String() string { _ = "STUB: not implemented"; return "" }

// ShortString returns shortened versions of bzz address in a format: [Overlay, Underlay]
// It can be used for logging
func (a *Address) ShortString() string { _ = "STUB: not implemented"; return "" }

func (a *Address) underlaysAsStrings() []string { _ = "STUB: not implemented"; return nil }

func parseMultiaddrs(addrs []string) ([]ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
