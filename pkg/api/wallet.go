// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/bigint"
)

type walletResponse struct {
	BZZ                       *bigint.BigInt `json:"bzzBalance"`                // the BZZ balance of the wallet associated with the eth address of the node
	NativeToken               *bigint.BigInt `json:"nativeTokenBalance"`        // the native token balance of the wallet associated with the eth address of the node
	ChainID                   int64          `json:"chainID"`                   // the id of the blockchain
	ChequebookContractAddress common.Address `json:"chequebookContractAddress"` // the address of the chequebook contract
	WalletAddress             common.Address `json:"walletAddress"`             // the address of the bee wallet
}

func (s *Service) walletHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type walletTxResponse struct {
	TransactionHash common.Hash `json:"transactionHash"`
}

func (s *Service) walletWithdrawHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
