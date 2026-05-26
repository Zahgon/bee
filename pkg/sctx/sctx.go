// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sctx provides convenience methods for context
// value injection and extraction.
package sctx

import (
	"context"
	"errors"
	"math/big"
)

// ErrTargetPrefix is returned when target prefix decoding fails.
var ErrTargetPrefix = errors.New("error decoding prefix string")

type (
	HTTPRequestIDKey struct{}
	requestHostKey   struct{}
	gasPriceKey      struct{}
	gasLimitKey      struct{}
)

// SetHost sets the http request host in the context
func SetHost(ctx context.Context, domain string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// GetHost gets the request host from the context
func GetHost(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func SetGasLimit(ctx context.Context, limit uint64) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetGasLimit(ctx context.Context) uint64 { _ = "STUB: not implemented"; return 0 }

func GetGasLimitWithDefault(ctx context.Context, defaultLimit uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func SetGasPrice(ctx context.Context, price *big.Int) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetGasPrice(ctx context.Context) *big.Int { _ = "STUB: not implemented"; return nil }
