// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
)

type postEnvelopeResponse struct {
	Issuer    string `json:"issuer"`    // Ethereum address of the postage batch owner
	Index     string `json:"index"`     // used index of the Postage Batch
	Timestamp string `json:"timestamp"` // timestamp of the postage stamp
	Signature string `json:"signature"` // postage stamp signature
}

// envelopePostHandler generates new postage stamp for requested chunk address
func (s *Service) envelopePostHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
