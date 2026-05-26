// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package accesscontrol

import (
	"context"
	"errors"

	"github.com/ethersphere/bee/v2/pkg/file"
	"github.com/ethersphere/bee/v2/pkg/manifest"
	"github.com/ethersphere/bee/v2/pkg/manifest/mantaray"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var (
	// ErrEndIteration indicates that the iteration terminated.
	ErrEndIteration = errors.New("end iteration")
	// ErrUnexpectedType indicates that an error occurred during the mantary-manifest creation.
	ErrUnexpectedType = errors.New("unexpected type")
	// ErrInvalidTimestamp indicates that the timestamp given to Lookup is invalid.
	ErrInvalidTimestamp = errors.New("invalid timestamp")
	// ErrNotFound is returned when an Entry is not found in the history.
	ErrNotFound = errors.New("access control: not found")
)

// History represents the interface for managing access control history.
type History interface {
	// Add adds a new entry to the access control history with the given timestamp and metadata.
	Add(ctx context.Context, ref swarm.Address, timestamp *int64, metadata *map[string]string) error
	// Lookup retrieves the entry from the history based on the given timestamp or returns error if not found.
	Lookup(ctx context.Context, timestamp int64) (manifest.Entry, error)
	// Store stores the history to the underlying storage and returns the reference.
	Store(ctx context.Context) (swarm.Address, error)
}

var _ History = (*HistoryStruct)(nil)

// manifestInterface extends the `manifest.Interface` interface and adds a `Root` method.
type manifestInterface interface {
	manifest.Interface
	Root() *mantaray.Node
}

// HistoryStruct represents an access control history with a mantaray-based manifest.
type HistoryStruct struct {
	manifest manifestInterface
	ls       file.LoadSaver
}

// NewHistory creates a new history with a mantaray-based manifest.
func NewHistory(ls file.LoadSaver) (*HistoryStruct, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewHistoryReference loads a history with a mantaray-based manifest.
func NewHistoryReference(ls file.LoadSaver, ref swarm.Address) (*HistoryStruct, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add adds a new entry to the access control history with the given timestamp and metadata.
func (h *HistoryStruct) Add(ctx context.Context, ref swarm.Address, timestamp *int64, metadata *map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// add timestamps transformed so that the latests timestamp becomes the smallest key.

// Lookup retrieves the entry from the history based on the given timestamp or returns error if not found.
func (h *HistoryStruct) Lookup(ctx context.Context, timestamp int64) (manifest.Entry, error) {
	_ = "STUB: not implemented"
	return *new(manifest.Entry), nil
}

func (h *HistoryStruct) lookupNode(ctx context.Context, searchedTimestamp int64) (*mantaray.Node, error) {
	_ = "STUB: not implemented"
	// before node's timestamp is the closest one that is less than or equal to the searched timestamp
	// for instance: 2030, 2020, 1994 -> search for 2021 -> before is 2020
	return nil, nil
}

// after node's timestamp is after the latest
// for instance: 2030, 2020, 1994 -> search for 1980 -> after is 1994

// return error to stop the walk, this is how WalkNode works...

// Store stores the history to the underlying storage and returns the reference.
func (h *HistoryStruct) Store(ctx context.Context) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func bytesToInt64(b []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func isBeforeMatch(pathTimestamp []byte, searchedTimestamp int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
