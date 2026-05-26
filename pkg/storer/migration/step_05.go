// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package migration

import (
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/transaction"
)

// step_05 is a migration step that removes all upload items from the store.
func step_05(st transaction.Storage, logger log.Logger) func() error {
	_ = "STUB: not implemented"
	return nil
}
