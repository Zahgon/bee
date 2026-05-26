// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package node

import (
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/metrics"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// InitStateStore will initialize the stateStore with the given path to the
// data directory. When given an empty directory path, the function will instead
// initialize an in-memory state store that will not be persisted.
func InitStateStore(logger log.Logger, dataDir string, cacheCapacity uint64) (storage.StateStorerManager, metrics.Collector, error) {
	_ = "STUB: not implemented"
	return *new(storage.StateStorerManager), *new(metrics.Collector), nil
}

// InitStamperStore will create new stamper store with the given path to the
// data directory. When given an empty directory path, the function will instead
// initialize an in-memory state store that will not be persisted.
// The returned bool indicates whether the previous shutdown was unclean (dirty).
func InitStamperStore(logger log.Logger, dataDir string, stateStore storage.StateStorer) (storage.Store, bool, error) {
	_ = "STUB: not implemented"
	return *new(storage.Store), false, nil
}

const (
	overlayNonce     = "overlayV2_nonce"
	noncedOverlayKey = "nonce-overlay"
)

// checkOverlay checks the overlay is the same as stored in the statestore
func checkOverlay(storer storage.StateStorer, overlay swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func overlayNonceExists(s storage.StateStorer) ([]byte, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func setOverlay(s storage.StateStorer, overlay swarm.Address, nonce []byte) error {
	_ = "STUB: not implemented"
	return nil
}
