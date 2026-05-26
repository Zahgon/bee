// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
)

const (
	apiVersion = "v1" // Only one api version exists, this should be configurable with more.
	rootPath   = "/" + apiVersion
)

func (s *Service) Mount() { _ = "STUB: not implemented"; return }

// EnableFullAPI will enable all available endpoints, because some endpoints are not available during syncing.
func (s *Service) EnableFullAPI() { _ = "STUB: not implemented"; return }

// Skip compression for GET requests on download endpoints.
// This is done in order to preserve Content-Length header in response,
// because CompressHandler is always removing it.

func (s *Service) mountTechnicalDebug() { _ = "STUB: not implemented"; return }

func (s *Service) checkRouteAvailability(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (s *Service) checkSwapAvailability(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (s *Service) checkChequebookAvailability(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (s *Service) checkStorageIncentivesAvailability(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (s *Service) checkChainAvailability(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (s *Service) mountAPI() { _ = "STUB: not implemented"; return }

// handle is a helper closure which simplifies the router setup.

func (s *Service) mountBusinessDebug() { _ = "STUB: not implemented"; return }
