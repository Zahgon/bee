// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type chunkAddressResponse struct {
	Reference swarm.Address `json:"reference"`
}

func (s *Service) chunkUploadHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Currently the localstore supports session based uploads. We don't want to
// create new session for single chunk uploads. So if the chunk upload is not
// part of a session already, then we directly push the chunk. This way we dont
// need to go through the UploadStore.

// not a valid cac chunk. Check if it's a replica soc chunk.

// FromChunk only uses the chunk data to recreate the soc chunk. So the address is irrelevant.

func (s *Service) chunkGetHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
