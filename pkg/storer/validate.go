// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/sharky"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// Validate ensures that all retrievalIndex chunks are correctly stored in sharky.
func ValidateReserve(ctx context.Context, basePath string, opts *Options) error {
	_ = "STUB: not implemented"
	return nil
}

// ValidateRetrievalIndex ensures that all retrievalIndex chunks are correctly stored in sharky.
func ValidateRetrievalIndex(ctx context.Context, basePath string, opts *Options) error {
	_ = "STUB: not implemented"
	return nil
}

func validateWork(logger log.Logger, store storage.Store, readFn func(context.Context, sharky.Location, []byte) error) {
	_ = "STUB: not implemented"
	return
}

// ValidatePinCollectionChunks collects all chunk addresses that are present in a pin collection but
// are either invalid or missing altogether.
func ValidatePinCollectionChunks(ctx context.Context, basePath, pin, location string, opts *Options) error {
	_ = "STUB: not implemented"
	return nil
}

type PinIntegrity struct {
	Store  storage.Store
	Sharky *sharky.Store
}

type PinStat struct {
	Ref                     swarm.Address
	Total, Missing, Invalid int
}

func (p *PinIntegrity) Check(ctx context.Context, logger log.Logger, pin string, out chan PinStat) {
	_ = "STUB: not implemented"
	return
}
