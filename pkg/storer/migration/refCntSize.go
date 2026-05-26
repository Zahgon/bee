// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package migration

import (
	"errors"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/sharky"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const oldRretrievalIndexItemSize = swarm.HashSize + 8 + sharky.LocationSize + 1

var _ storage.Item = (*OldRetrievalIndexItem)(nil)

var (
	// errMarshalInvalidRetrievalIndexAddress is returned if the RetrievalIndexItem address is zero during marshaling.
	errMarshalInvalidRetrievalIndexAddress = errors.New("marshal RetrievalIndexItem: address is zero")
	// errMarshalInvalidRetrievalIndexLocation is returned if the RetrievalIndexItem location is invalid during marshaling.
	errMarshalInvalidRetrievalIndexLocation = errors.New("marshal RetrievalIndexItem: location is invalid")
	// errUnmarshalInvalidRetrievalIndexSize is returned during unmarshaling if the passed buffer is not the expected size.
	errUnmarshalInvalidRetrievalIndexSize = errors.New("unmarshal RetrievalIndexItem: invalid size")
	// errUnmarshalInvalidRetrievalIndexLocationBytes is returned during unmarshaling if the location buffer is invalid.
	errUnmarshalInvalidRetrievalIndexLocationBytes = errors.New("unmarshal RetrievalIndexItem: invalid location bytes")
)

// OldRetrievalIndexItem is the index which gives us the sharky location from the swarm.Address.
// The RefCnt stores the reference of each time a Put operation is issued on this Address.
type OldRetrievalIndexItem struct {
	Address   swarm.Address
	Timestamp uint64
	Location  sharky.Location
	RefCnt    uint8
}

func (r *OldRetrievalIndexItem) ID() string { _ = "STUB: not implemented"; return "" }

func (OldRetrievalIndexItem) Namespace() string { _ = "STUB: not implemented"; return "" }

// Stored in bytes as:
// |--Address(32)--|--Timestamp(8)--|--Location(7)--|--RefCnt(1)--|
func (r *OldRetrievalIndexItem) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *OldRetrievalIndexItem) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }

func (r *OldRetrievalIndexItem) Clone() storage.Item {
	_ = "STUB: not implemented"
	return *new(storage.Item)
}

func (r OldRetrievalIndexItem) String() string { _ = "STUB: not implemented"; return "" }

func RefCountSizeInc(s storage.BatchStore, logger log.Logger) func() error {
	_ = "STUB: not implemented"
	return nil
}

// create new
