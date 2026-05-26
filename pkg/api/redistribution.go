// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/bigint"
)

type redistributionStatusResponse struct {
	MinimumGasFunds           *bigint.BigInt `json:"minimumGasFunds"`
	HasSufficientFunds        bool           `json:"hasSufficientFunds"`
	IsFrozen                  bool           `json:"isFrozen"`
	IsFullySynced             bool           `json:"isFullySynced"`
	Phase                     string         `json:"phase"`
	Round                     uint64         `json:"round"`
	LastWonRound              uint64         `json:"lastWonRound"`
	LastPlayedRound           uint64         `json:"lastPlayedRound"`
	LastFrozenRound           uint64         `json:"lastFrozenRound"`
	LastSelectedRound         uint64         `json:"lastSelectedRound"`
	LastSampleDurationSeconds float64        `json:"lastSampleDurationSeconds"`
	Block                     uint64         `json:"block"`
	Reward                    *bigint.BigInt `json:"reward"`
	Fees                      *bigint.BigInt `json:"fees"`
	IsHealthy                 bool           `json:"isHealthy"`
}

func (s *Service) redistributionStatusHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
