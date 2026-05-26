// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type socPostResponse struct {
	Reference swarm.Address `json:"reference"`
}

func (s *Service) socUploadHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Resolve the deferred upload mode. Historically /soc always pushed
// directly to the network; preserve that default when neither header is
// provided. An explicit Swarm-Deferred-Upload header wins, otherwise the
// presence of a Swarm-Tag opts the caller into deferred mode (matching
// /chunks' auto-defer semantics).

func (s *Service) socGetHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// include additional headers
