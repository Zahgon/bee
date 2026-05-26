// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package headerutils

import (
	"errors"

	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const (
	priceFieldName  = "price"
	targetFieldName = "target"
	indexFieldName  = "index"
)

var (
	// ErrFieldLength denotes p2p.Header having malformed field length in bytes
	ErrFieldLength = errors.New("field length error")
	// ErrNoIndexHeader denotes p2p.Header lacking specified field
	ErrNoIndexHeader = errors.New("no index header")
	// ErrNoTargetHeader denotes p2p.Header lacking specified field
	ErrNoTargetHeader = errors.New("no target header")
	// ErrNoPriceHeader denotes p2p.Header lacking specified field
	ErrNoPriceHeader = errors.New("no price header")
)

// Headers, utility functions

func MakePricingHeaders(chunkPrice uint64, addr swarm.Address) (p2p.Headers, error) {
	_ = "STUB: not implemented"
	return *new(p2p.Headers), nil
}

func MakePricingResponseHeaders(chunkPrice uint64, addr swarm.Address, index uint8) (p2p.Headers, error) {
	_ = "STUB: not implemented"
	return *new(p2p.Headers), nil
}

// ParsePricingHeaders used by responder to read address and price from stream headers
// Returns an error if no target field attached or the contents of it are not readable
func ParsePricingHeaders(receivedHeaders p2p.Headers) (swarm.Address, uint64, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), 0, nil
}

// ParsePricingResponseHeaders used by requester to read address, price and index from response headers
// Returns an error if any fields are missing or target is unreadable
func ParsePricingResponseHeaders(receivedHeaders p2p.Headers) (swarm.Address, uint64, uint8, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), 0, 0, nil
}

func ParseIndexHeader(receivedHeaders p2p.Headers) (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ParseTargetHeader(receivedHeaders p2p.Headers) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func ParsePriceHeader(receivedHeaders p2p.Headers) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
