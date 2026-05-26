// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p2p

import (
	"context"

	ma "github.com/multiformats/go-multiaddr"
)

func sortAddrsByTCPPreference(addrs []ma.Multiaddr) { _ = "STUB: not implemented"; return }

func isDNSProtocol(protoCode int) bool { _ = "STUB: not implemented"; return false }

func Discover(ctx context.Context, addr ma.Multiaddr, f func(ma.Multiaddr) (bool, error)) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// If the resolved addresses are real (non-DNS) multiaddrs, order them
// with TCP first. Otherwise (still DNS), shuffle randomly as before.
