// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chequebook

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/transaction"
)

type chequebookContract struct {
	address            common.Address
	transactionService transaction.Service
}

func newChequebookContract(address common.Address, transactionService transaction.Service) *chequebookContract {
	_ = "STUB: not implemented"
	return nil
}

func (c *chequebookContract) Issuer(ctx context.Context) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

// Balance returns the token balance of the chequebook.
func (c *chequebookContract) Balance(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *chequebookContract) PaidOut(ctx context.Context, address common.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *chequebookContract) TotalPaidOut(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
