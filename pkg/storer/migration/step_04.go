// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package migration

import (
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/transaction"
)

// step_04 is the fourth step of the migration. It forces a sharky recovery to
// be run on the localstore.
func step_04(
	sharkyBasePath string,
	sharkyNoOfShards int,
	st transaction.Storage,
	logger log.Logger,
) func() error {
	_ = "STUB: not implemented"
	return nil

	// for in-mem store, skip this step
}
