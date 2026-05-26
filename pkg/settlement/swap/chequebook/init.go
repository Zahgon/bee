// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chequebook

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/erc20"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/transaction"
)

const (
	chequebookKey           = "swap_chequebook"
	ChequebookDeploymentKey = "swap_chequebook_transaction_deployment"

	balanceCheckBackoffDuration = 20 * time.Second
	balanceCheckMaxRetries      = 10
)

const (
	erc20SmallUnitStr = "10000000000000000"
	ethSmallUnitStr   = "1000000000000000000"
)

func checkBalance(
	ctx context.Context,
	logger log.Logger,
	swapInitialDeposit *big.Int,
	swapBackend transaction.Backend,
	chainId int64,
	overlayEthAddress common.Address,
	erc20Token erc20.Service,
) error {
	_ = "STUB: not implemented"
	return nil
}

// always round up the division as the bzzaar cannot handle decimals

// Init initialises the chequebook service.
func Init(
	ctx context.Context,
	chequebookFactory Factory,
	stateStore storage.StateStorer,
	logger log.Logger,
	swapInitialDeposit *big.Int,
	transactionService transaction.Service,
	swapBackend transaction.Backend,
	chainId int64,
	overlayEthAddress common.Address,
	chequeSigner ChequeSigner,
	erc20Service erc20.Service,
) (chequebookService Service, err error) {
	_ = "STUB: not implemented"
	return *new(Service), nil
}

// if we don't yet have a chequebook, deploy a new one

// save the address for later use

// regardless of how the chequebook service was initialised make sure that the chequebook is valid
