// Copyright 2026 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bzz

import (
	ma "github.com/multiformats/go-multiaddr"
)

// TransportType represents the transport protocol of a multiaddress.
type TransportType int

const (
	// TransportUnknown indicates an unrecognized transport.
	TransportUnknown TransportType = iota
	// TransportTCP indicates plain TCP without WebSocket.
	TransportTCP
	// TransportWS indicates WebSocket without TLS.
	TransportWS
	// TransportWSS indicates WebSocket with TLS (secure).
	TransportWSS
)

// String returns a string representation of the transport type.
func (t TransportType) String() string { _ = "STUB: not implemented"; return "" }

// Priority returns the sorting priority for the transport type.
// Lower value = higher priority: TCP (0) > WS (1) > WSS (2) > Unknown (3)
func (t TransportType) Priority() int { _ = "STUB: not implemented"; return 0 }

// ClassifyTransport returns the transport type of a multiaddress.
// It distinguishes between plain TCP, WebSocket (WS), and secure WebSocket (WSS).
func ClassifyTransport(addr ma.Multiaddr) TransportType {
	_ = "STUB: not implemented"
	return *new(TransportType)
}

func SelectBestAdvertisedAddress(addrs []ma.Multiaddr, fallback ma.Multiaddr) ma.Multiaddr {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr)
}

// Sort addresses: first by transport priority (TCP > WS > WSS), preserving relative order
