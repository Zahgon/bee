// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/sharky"
	storage "github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const (
	sharkyDirtyFileName = ".DIRTY"
)

func sharkyRecovery(ctx context.Context, sharkyBasePath string, store storage.Store, opts *Options) (closerFn, int, error) {
	_ = "STUB: not implemented"
	return *new(closerFn), 0, nil
}

// validateAndAddLocations iterates every chunk index entry, reads its data from
// Sharky, and validates the content hash. Valid chunks are registered with the
// recovery so their slots are preserved. Corrupted entries (unreadable data or
// hash mismatch) are logged, excluded from the recovery bitmap, and deleted from
// the index store — including all associated reserve metadata (BatchRadiusItem,
// ChunkBinItem, stampindex, chunkstamp) — so the node starts clean without
// serving invalid data and with correct reserve size accounting.
// If a corrupted index entry cannot be deleted, an error is returned and the
// node startup is aborted to prevent serving or operating on corrupt state.
// It returns the number of corrupted entries that were pruned.
func validateAndAddLocations(ctx context.Context, store storage.Store, sharkyRecover *sharky.Recovery, baseAddr swarm.Address, logger log.Logger) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
