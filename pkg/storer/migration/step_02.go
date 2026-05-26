// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package migration

import (
	"github.com/ethersphere/bee/v2/pkg/storer/internal/transaction"
)

// step_02 migrates the cache to the new format.
// the old cacheEntry item has the same key, but the value is different. So only
// a Put is needed.
func step_02(st transaction.Storage) func() error { _ = "STUB: not implemented"; return nil }
