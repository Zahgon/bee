// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
)

type BeeNodeMode uint

const (
	UnknownMode BeeNodeMode = iota
	LightMode
	FullMode
	UltraLightMode
)

type nodeResponse struct {
	BeeMode           string `json:"beeMode"`
	ChequebookEnabled bool   `json:"chequebookEnabled"`
	SwapEnabled       bool   `json:"swapEnabled"`
}

func (b BeeNodeMode) String() string { _ = "STUB: not implemented"; return "" }

// nodeGetHandler gives back information about the Bee node configuration.
func (s *Service) nodeGetHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}
