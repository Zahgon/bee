// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
)

type ReadyStatusResponse healthStatusResponse

func (s *Service) readinessHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}
