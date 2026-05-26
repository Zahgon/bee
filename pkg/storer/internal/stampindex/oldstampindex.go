// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stampindex

import (
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// ItemV1 is an store.Item that represents data relevant to stamp.
type ItemV1 struct {
	// Keys.
	namespace  []byte // The namespace of other related item.
	BatchID    []byte
	StampIndex []byte

	// Values.
	StampTimestamp   []byte
	ChunkAddress     swarm.Address
	ChunkIsImmutable bool
}

// ID implements the storage.Item interface.
func (i ItemV1) ID() string { _ = "STUB: not implemented"; return "" }

// Namespace implements the storage.Item interface.
func (i ItemV1) Namespace() string { _ = "STUB: not implemented"; return "" }

// Marshal implements the storage.Item interface.
func (i ItemV1) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal implements the storage.Item interface.
func (i *ItemV1) Unmarshal(bytes []byte) error { _ = "STUB: not implemented"; return nil }

// Clone  implements the storage.Item interface.
func (i *ItemV1) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

// String implements the fmt.Stringer interface.
func (i ItemV1) String() string { _ = "STUB: not implemented"; return "" }

func (i *ItemV1) SetNamespace(ns []byte) { _ = "STUB: not implemented"; return }
