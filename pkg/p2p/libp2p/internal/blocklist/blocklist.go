// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blocklist

import (
	"time"

	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var keyPrefix = "blocklist-"

type currentTimeFn = func() time.Time

type Blocklist struct {
	store         storage.StateStorer
	currentTimeFn currentTimeFn
}

func NewBlocklist(store storage.StateStorer) *Blocklist { _ = "STUB: not implemented"; return nil }

type entry struct {
	Timestamp time.Time `json:"timestamp"`
	Duration  string    `json:"duration"` // Duration is string because the time.Duration does not implement MarshalJSON/UnmarshalJSON methods.
	Reason    string    `json:"reason"`
	Full      bool      `json:"full"`
}

func (b *Blocklist) Exists(overlay swarm.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (b *Blocklist) Add(overlay swarm.Address, duration time.Duration, reason string, full bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// if peer is already blacklisted, blacklist it for the maximum amount of time

// Peers returns all currently blocklisted peers.
func (b *Blocklist) Peers() ([]p2p.BlockListedPeer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip to the next item

func (b *Blocklist) get(key string) (entry, time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(entry), *new(time.Duration), nil
}

func generateKey(overlay swarm.Address) string { _ = "STUB: not implemented"; return "" }

func unmarshalKey(s string) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *
	// trim prefix
	new(swarm.Address), nil
}
