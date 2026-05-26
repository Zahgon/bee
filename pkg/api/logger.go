// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/log"
)

// The following variables exist only to be mocked in tests.
var (
	logRegistryIterate   = log.RegistryIterate
	logSetVerbosityByExp = log.SetVerbosityByExp
)

type (
	data struct {
		Next  node     `json:"/,omitempty"`
		Names []string `json:"+,omitempty"`
	}

	node map[string]*data

	loggerInfo struct {
		Logger    string `json:"logger"`
		Verbosity string `json:"verbosity"`
		Subsystem string `json:"subsystem"`
		ID        string `json:"id"`
	}

	loggerResult struct {
		Tree    node         `json:"tree"`
		Loggers []loggerInfo `json:"loggers"`
	}
)

// loggerGetHandler returns all available loggers that match the specified expression.
func (s *Service) loggerGetHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Tree structure.

// Flat structure.

// loggerSetVerbosityHandler sets logger(s) verbosity level based on
// the specified expression or subsystem that matches the logger(s).
func (s *Service) loggerSetVerbosityHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
