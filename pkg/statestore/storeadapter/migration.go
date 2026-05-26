// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storeadapter

import (
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/storage/migration"
)

func allSteps(st storage.Store) migration.Steps {
	_ = "STUB: not implemented"
	return *new(migration.Steps)
}

// bigIntPrefixes lists all statestore key prefixes whose values are stored as
// big.Int and need to be migrated to bigint.BigInt binary (Gob) encoding.
var bigIntPrefixes = []string{
	"accounting_balance_",
	"accounting_surplusbalance_",
	"accounting_originatedbalance_",
	"swap_chequebook_total_issued_",
}

// migrateBigIntKeys rewrites all balance values from their legacy JSON encoding
// (unquoted decimal from json.Marshal(*big.Int), or quoted string from
// json.Marshal(bigint.BigInt)) to the canonical Gob binary encoding produced
// by bigint.BigInt.MarshalBinary. After this migration UnmarshalBinary only
// needs to handle Gob.
func migrateBigIntKeys(s storage.Store) migration.StepFn {
	_ = "STUB: not implemented"
	return *new(migration.StepFn)
}

// Gob-encoded data found before migration — database is in an unexpected state

// quoted decimal string from json.Marshal(bigint.BigInt)

// unquoted decimal from json.Marshal(*big.Int)

// The StateStorerAdapter iterator gives key = prefix + res.ID,
// where res.ID already contains the full key (including prefix).
// Strip the leading prefix duplicate to get the actual statestore key.

func deletePrefix(s storage.Store, prefix string) migration.StepFn {
	_ = "STUB: not implemented"
	return *new(migration.StepFn)
}

func epochMigration(s storage.Store) migration.StepFn {
	_ = "STUB: not implemented"
	return *new(migration.StepFn)
}
