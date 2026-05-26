// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package epochs

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/feeds"
	storage "github.com/ethersphere/bee/v2/pkg/storage"
)

var _ feeds.Updater = (*updater)(nil)

// Updater encapsulates a feeds putter to generate successive updates for epoch based feeds
// it persists the last update
type updater struct {
	*feeds.Putter
	last  int64
	epoch feeds.Index
}

// NewUpdater constructs a feed updater
func NewUpdater(putter storage.Putter, signer crypto.Signer, topic []byte) (feeds.Updater, error) {
	_ = "STUB: not implemented"
	return *new(feeds.Updater), nil
}

// Update pushes an update to the feed through the chunk stores
func (u *updater) Update(ctx context.Context, at int64, payload []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *updater) Feed() *feeds.Feed { _ = "STUB: not implemented"; return nil }
