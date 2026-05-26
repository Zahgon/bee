// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package shed

import (
	"github.com/syndtr/goleveldb/leveldb"
)

// StructField is a helper to store complex structure by
// encoding it in RLP format.
type StructField struct {
	db  *DB
	key []byte
}

// NewStructField returns a new StructField.
// It validates its name and type against the database schema.
func (db *DB) NewStructField(name string) (f StructField, err error) {
	_ = "STUB: not implemented"
	return *new(StructField), nil
}

// Get unmarshals data from the database to a provided val.
// If the data is not found leveldb.ErrNotFound is returned.
func (f StructField) Get(val any) (err error) { _ = "STUB: not implemented"; return nil }

// Put marshals provided val and saves it to the database.
func (f StructField) Put(val any) (err error) { _ = "STUB: not implemented"; return nil }

// PutInBatch marshals provided val and puts it into the batch.
func (f StructField) PutInBatch(batch *leveldb.Batch, val any) (err error) {
	_ = "STUB: not implemented"
	return nil
}
