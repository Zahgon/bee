// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"testing"

	"github.com/ethersphere/bee/v2/pkg/storage"
)

const (
	key1 = "key1" // stores the serialized type
	key2 = "key2" // stores a json array
)

var (
	value1 = &Serializing{value: "value1"}
	value2 = []string{"a", "b", "c"}
)

type Serializing struct {
	value           string
	marshalCalled   bool
	unmarshalCalled bool
}

func (st *Serializing) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (st *Serializing) UnmarshalBinary(data []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// RunPersist is a specific test case for the persistent state store.
// It tests that values persist across sessions.
func RunPersist(t *testing.T, f func(t *testing.T, dir string) storage.StateStorer) {
	_ = "STUB: not implemented"
	return
}

// insert some values

// test that the iterator works

// close the store

// bootstrap with the same old dir

// test that the iterator works

// insert some more random entries

// check again

func Run(t *testing.T, f func(t *testing.T) storage.StateStorer) { _ = "STUB: not implemented"; return }

func testDelete(t *testing.T, f func(t *testing.T) storage.StateStorer) {
	_ = "STUB: not implemented"

	// create a store
	return
}

// insert some values

// check that the persisted values match

// check that the store is empty

func testPutGet(t *testing.T, f func(t *testing.T) storage.StateStorer) {
	_ = "STUB: not implemented"

	// create a store
	return
}

// insert some values

// check that the persisted values match

func testIterator(t *testing.T, f func(t *testing.T) storage.StateStorer) {
	_ = "STUB: not implemented"

	// create a store
	return
}

// insert some values

// test that the iterator works

func insertValues(t *testing.T, store storage.StateStorer, key1, key2 string, value1 *Serializing, value2 []string) {
	_ = "STUB: not implemented"
	return
}

func insert(t *testing.T, store storage.StateStorer, prefix string, count int) {
	_ = "STUB: not implemented"
	return
}

func testPersistedValues(t *testing.T, store storage.StateStorer, key1, key2 string, value1 *Serializing, value2 []string) {
	_ = "STUB: not implemented"
	return
}

func testStoreIterator(t *testing.T, store storage.StateStorer, prefix string, size int) {
	_ = "STUB: not implemented"
	return
}

func testEmpty(t *testing.T, store storage.StateStorer) { _ = "STUB: not implemented"; return }
