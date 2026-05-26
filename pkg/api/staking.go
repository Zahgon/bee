// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/bigint"
)

func (s *Service) stakingAccessHandler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

type getStakeResponse struct {
	StakedAmount *bigint.BigInt `json:"stakedAmount"`
}

type getWithdrawableResponse struct {
	WithdrawableAmount *bigint.BigInt `json:"withdrawableAmount"`
}
type stakeTransactionReponse struct {
	TxHash string `json:"txHash"`
}

func (s *Service) stakingDepositHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// if the deposit is successful, we should update the height of the node in the staking contract
// this is done to make sure that the node is participating in the redistribution game with the correct height
// if the node has started with insufficient stake

func (s *Service) getPotentialStake(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) getWithdrawableStakeHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) withdrawStakeHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) migrateStakeHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
