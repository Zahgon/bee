// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
)

const welcomeMessageMaxRequestSize = 512

type welcomeMessageRequest struct {
	WelcomeMesssage string `json:"welcomeMessage"`
}

type welcomeMessageResponse struct {
	WelcomeMesssage string `json:"welcomeMessage"`
}

func (s *Service) getWelcomeMessageHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) setWelcomeMessageHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
