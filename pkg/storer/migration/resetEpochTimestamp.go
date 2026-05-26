// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package migration

import (
	"github.com/ethersphere/bee/v2/pkg/storer/internal/transaction"
)

// resetReserveEpochTimestamp is a migration that resets the epoch timestamp of the reserve
// so that peers in the network can resync chunks.
func resetReserveEpochTimestamp(st transaction.Storage) func() error {
	_ = "STUB: not implemented"
	return nil
}
