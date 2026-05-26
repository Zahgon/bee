// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeDeadline = 4 * time.Second // write deadline. should be smaller than the shutdown timeout on api close
)

func (s *Service) pssPostHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) pssWsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) pumpWs(conn *websocket.Conn, t string) { _ = "STUB: not implemented"; return }

// shutdown

// client gone

// error encountered while pinging client. client probably gone
