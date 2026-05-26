// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/gorilla/websocket"
)

func (s *Service) gsocWsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) gsocListeningWs(conn *websocket.Conn, socAddress swarm.Address) {
	_ = "STUB: not implemented"
	return
}

// shutdown

// client gone

// error encountered while pinging client. client probably gone
