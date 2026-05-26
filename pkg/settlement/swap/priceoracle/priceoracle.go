// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package priceoracle

import (
	"context"
	"errors"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/transaction"
	"github.com/ethersphere/bee/v2/pkg/util/abiutil"
	"github.com/ethersphere/go-price-oracle-abi/priceoracleabi"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "priceoracle"

var errDecodeABI = errors.New("could not decode abi data")

type service struct {
	logger             log.Logger
	priceOracleAddress common.Address
	transactionService transaction.Service
	exchangeRate       *big.Int
	deduction          *big.Int
	timeDivisor        int64
	quitC              chan struct{}
}

type Service interface {
	io.Closer
	// CurrentRates returns the current value of exchange rate and deduction
	// according to the latest information from oracle
	CurrentRates() (exchangeRate *big.Int, deduction *big.Int, err error)
	// GetPrice retrieves latest available information from oracle
	GetPrice(ctx context.Context) (*big.Int, *big.Int, error)
	Start()
}

var priceOracleABI = abiutil.MustParseABI(priceoracleabi.PriceOracleABIv0_6_9)

func New(logger log.Logger, priceOracleAddress common.Address, transactionService transaction.Service, timeDivisor int64) Service {
	_ = "STUB: not implemented"
	return *new(Service)
}

func (s *service) Start() { _ = "STUB: not implemented"; return }

// We poll the oracle in every timestamp divisible by constant 300 (timeDivisor)
// in order to get latest version approximately at the same time on all nodes
// and to minimize polling frequency
// If the node gets newer information than what was applicable at last polling point at startup
// this minimizes the negative scenario to less than 5 minutes
// during which cheques can not be sent / accepted because of the asymmetric information

func (s *service) GetPrice(ctx context.Context) (*big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *service) CurrentRates() (exchangeRate, deduction *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *service) Close() error { _ = "STUB: not implemented"; return nil }
