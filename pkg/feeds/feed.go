// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package feeds implements generic interfaces and methods for time-based feeds
// indexing schemes are implemented in subpackages
// - epochs
// - sequence
package feeds

import (
	"encoding"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var ErrFeedTypeNotFound = errors.New("no such feed type")

// Factory creates feed lookups for different types of feeds.
type Factory interface {
	NewLookup(Type, *Feed) (Lookup, error)
}

// Type enumerates the time-based feed types
type Type int

const (
	Sequence Type = iota
	Epoch
)

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

// FromString constructs the type from a string
func (t *Type) FromString(s string) error { _ = "STUB: not implemented"; return nil }

type id struct {
	topic []byte
	index []byte
}

var _ encoding.BinaryMarshaler = (*id)(nil)

func (i *id) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Feed is representing an epoch based feed
type Feed struct {
	Topic []byte
	Owner common.Address
}

// New constructs an epoch based feed from a keccak256 digest of a plaintext
// topic and an ether address.
func New(topic []byte, owner common.Address) *Feed { _ = "STUB: not implemented"; return nil }

// Index is the interface for feed implementations.
type Index interface {
	encoding.BinaryMarshaler
	Next(last int64, at uint64) Index
	fmt.Stringer
}

// Update represents an update instance of a feed, i.e., pairing of a Feed with an Epoch
type Update struct {
	*Feed
	index Index
}

// Update called on a feed with an index and returns an Update
func (f *Feed) Update(index Index) *Update { _ = "STUB: not implemented"; return nil }

// NewUpdate creates an update from an index, timestamp, payload and signature
func NewUpdate(f *Feed, idx Index, timestamp int64, payload, sig []byte) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// Id calculates the identifier if a  feed update to be used in single owner chunks
func (u *Update) Id() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Id calculates the feed id from a topic and an index
		nil
}

func Id(topic []byte, index Index) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Address calculates the soc address of a feed update
func (u *Update) Address() (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}
