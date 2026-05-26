// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
	"time"

	storer "github.com/ethersphere/bee/v2/pkg/storer"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type tagRequest struct {
	Address swarm.Address `json:"address"`
}

type tagResponse struct {
	Split     uint64        `json:"split"`
	Seen      uint64        `json:"seen"`
	Stored    uint64        `json:"stored"`
	Sent      uint64        `json:"sent"`
	Synced    uint64        `json:"synced"`
	Uid       uint64        `json:"uid"`
	Address   swarm.Address `json:"address"`
	StartedAt time.Time     `json:"startedAt"`
}

func newTagResponse(tag storer.SessionInfo) tagResponse {
	_ = "STUB: not implemented"
	return *new(tagResponse)
}

type listTagsResponse struct {
	Tags []tagResponse `json:"tags"`
}

func (s *Service) createTagHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) getTagHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) deleteTagHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) doneSplitHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) listTagsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Default limit.
