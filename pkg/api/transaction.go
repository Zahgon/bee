// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/bigint"
)

const (
	errCantGetTransaction    = "cannot get transaction"
	errUnknownTransaction    = "unknown transaction"
	errAlreadyImported       = "already imported"
	errCantResendTransaction = "can't resend transaction"
)

type transactionInfo struct {
	TransactionHash common.Hash     `json:"transactionHash"`
	To              *common.Address `json:"to"`
	Nonce           uint64          `json:"nonce"`
	GasPrice        *bigint.BigInt  `json:"gasPrice"`
	GasLimit        uint64          `json:"gasLimit"`
	GasTipBoost     int             `json:"gasTipBoost"`
	GasTipCap       *bigint.BigInt  `json:"gasTipCap"`
	GasFeeCap       *bigint.BigInt  `json:"gasFeeCap"`
	Data            string          `json:"data"`
	Created         time.Time       `json:"created"`
	Description     string          `json:"description"`
	Value           *bigint.BigInt  `json:"value"`
}

type transactionPendingList struct {
	PendingTransactions []transactionInfo `json:"pendingTransactions"`
}

func (s *Service) transactionListHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) transactionDetailHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type transactionHashResponse struct {
	TransactionHash common.Hash `json:"transactionHash"`
}

func (s *Service) transactionResendHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) transactionCancelHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
