// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// pinRootHash pins root hash of given reference. This method is idempotent.
func (s *Service) pinRootHash(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// unpinRootHash unpin's an already pinned root hash. This method is idempotent.
func (s *Service) unpinRootHash(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// getPinnedRootHash returns back the given reference if its root hash is pinned.
func (s *Service) getPinnedRootHash(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// listPinnedRootHashes lists all the references of the pinned root hashes.
func (s *Service) listPinnedRootHashes(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type PinIntegrityResponse struct {
	Reference swarm.Address `json:"reference"`
	Total     int           `json:"total"`
	Missing   int           `json:"missing"`
	Invalid   int           `json:"invalid"`
}

func (s *Service) pinIntegrityHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
