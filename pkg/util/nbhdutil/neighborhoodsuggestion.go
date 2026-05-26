// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nbhdutil

import (
	"net/http"
)

type httpClient interface {
	Get(url string) (*http.Response, error)
}

func FetchNeighborhood(client httpClient, suggester string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
