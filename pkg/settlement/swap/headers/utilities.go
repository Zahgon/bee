// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package swap

import (
	"errors"
	"math/big"

	"github.com/ethersphere/bee/v2/pkg/p2p"
)

const (
	exchangeRateFieldName = "exchange"
	deductionFieldName    = "deduction"
)

var (
	// ErrFieldLength denotes p2p.Header having malformed field length in bytes
	ErrFieldLength = errors.New("field length error")
	// ErrNoExchangeHeader denotes p2p.Header lacking specified field
	ErrNoExchangeHeader = errors.New("no exchange header")
	// ErrNoDeductionHeader denotes p2p.Header lacking specified field
	ErrNoDeductionHeader = errors.New("no deduction header")
)

func MakeSettlementHeaders(exchangeRate, deduction *big.Int) p2p.Headers {
	_ = "STUB: not implemented"
	return *new(p2p.Headers)
}

func ParseSettlementResponseHeaders(receivedHeaders p2p.Headers) (exchange, deduction *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ParseExchangeHeader(receivedHeaders p2p.Headers) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseDeductionHeader(receivedHeaders p2p.Headers) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
