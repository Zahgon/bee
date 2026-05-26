// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libp2p

import (
	"net"

	ma "github.com/multiformats/go-multiaddr"
)

type staticAddressResolver struct {
	multiProto string
	port       string
}

func newStaticAddressResolver(addr string, lookupIP func(host string) ([]net.IP, error)) (*staticAddressResolver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *staticAddressResolver) Resolve(observedAddress ma.Multiaddr) (ma.Multiaddr, error) {
	_ = "STUB: not implemented"
	return *new(ma.Multiaddr), nil
}

// if address is not in a form of '/ipversion/ip/protocol/port/...` don't compare to addresses and return it

// replace sni protocol ip subdomain in domain name

// assume only one sni component

func newIPSniHost(ip string, isIPv4 bool) string { _ = "STUB: not implemented"; return "" }

// IPv6

func getMultiProto(host string, lookupIP func(host string) ([]net.IP, error)) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func ipsClassifier(ips []net.IP) (ipv4, ipv6 bool) { _ = "STUB: not implemented"; return false, false }
