// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/multiformats/go-multiaddr"
)

type addressesResponse struct {
	Overlay      *swarm.Address        `json:"overlay"`
	Underlay     []multiaddr.Multiaddr `json:"underlay"`
	Ethereum     common.Address        `json:"ethereum"` // deprecated
	ChainAddress common.Address        `json:"chain_address"`
	PublicKey    string                `json:"publicKey"`
	PSSPublicKey string                `json:"pssPublicKey"`
}

func (s *Service) addressesHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

// initialize variable to json encode as [] instead null if p2p is nil

// addresses endpoint is exposed before p2p service is configured
// to provide information about other addresses.
