// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/bigint"
)

const (
	errCantBalances   = "Cannot get balances"
	errCantBalance    = "Cannot get balance"
	errNoBalance      = "No balance for peer"
	errInvalidAddress = "Invalid address"
)

type balanceResponse struct {
	Peer              string         `json:"peer"`
	Balance           *bigint.BigInt `json:"balance"`
	ThresholdReceived *bigint.BigInt `json:"thresholdreceived"`
	ThresholdGiven    *bigint.BigInt `json:"thresholdgiven"`
}

type balancesResponse struct {
	Balances []balanceResponse `json:"balances"`
}

func (s *Service) balancesHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) peerBalanceHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) compensatedBalancesHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) compensatedPeerBalanceHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
