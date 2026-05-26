// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package migration

import (
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/transaction"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// ReserveRepairer is a migration step that removes all BinItem entries and migrates
// ChunkBinItem and BatchRadiusItem entries to use a new BinID field.
func ReserveRepairer(
	st transaction.Storage,
	chunkTypeFunc func(swarm.Chunk) swarm.ChunkType,
	logger log.Logger,
) func() error {
	_ = "STUB: not implemented"
	return nil

	/*
		STEP 0:	remove epoch item
		STEP 1:	remove all of the BinItem entries
		STEP 2:	remove all of the ChunkBinItem entries
		STEP 3:	iterate BatchRadiusItem, get new binID
				create new ChunkBinItem and BatchRadiusItem if the chunk exists in the chunkstore
				if the chunk is invalid, it is removed from the chunkstore
		STEP 4: save the latest binID to disk
	*/
}

// extra test that ensure that a unique binID has been issed to each item.

// STEP 0

// STEP 1

// STEP 2

// STEP 3

// STEP 4
