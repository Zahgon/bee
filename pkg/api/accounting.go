// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/bigint"
)

const (
	httpErrGetAccountingInfo = "Cannot get accounting info"
)

type peerData struct {
	InfoResponse map[string]peerDataResponse `json:"peerData"`
}

type peerDataResponse struct {
	Balance                  *bigint.BigInt `json:"balance"`
	ConsumedBalance          *bigint.BigInt `json:"consumedBalance"`
	ThresholdReceived        *bigint.BigInt `json:"thresholdReceived"`
	ThresholdGiven           *bigint.BigInt `json:"thresholdGiven"`
	CurrentThresholdReceived *bigint.BigInt `json:"currentThresholdReceived"`
	CurrentThresholdGiven    *bigint.BigInt `json:"currentThresholdGiven"`
	SurplusBalance           *bigint.BigInt `json:"surplusBalance"`
	ReservedBalance          *bigint.BigInt `json:"reservedBalance"`
	ShadowReservedBalance    *bigint.BigInt `json:"shadowReservedBalance"`
	GhostBalance             *bigint.BigInt `json:"ghostBalance"`
}

func (s *Service) accountingInfoHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}
