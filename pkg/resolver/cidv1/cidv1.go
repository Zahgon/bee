// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cidv1

import (
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// https://github.com/multiformats/multicodec/blob/master/table.csv
const (
	SwarmNsCodec       uint64 = 0xe4
	SwarmManifestCodec uint64 = 0xfa
	SwarmFeedCodec     uint64 = 0xfb
)

type Resolver struct{}

func (Resolver) Resolve(name string) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func (Resolver) Close() error { _ = "STUB: not implemented"; return nil }
