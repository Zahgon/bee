// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package migration

import (
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/transaction"
)

// step_06 is a migration step that adds a stampHash to all BatchRadiusItems, ChunkBinItems and StampIndexItems.
func step_06(st transaction.Storage, logger log.Logger) func() error {
	_ = "STUB: not implemented"
	return nil
}

func addStampHash(logger log.Logger, st transaction.Storage) (int64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// Delete epoch timestamp

// Since the ID format has changed, we should delete the old item and put a new one with the new ID format.

// same id. Will replace.

// same id. Will replace.
