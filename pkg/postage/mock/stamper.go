// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type mockStamper struct{}

// NewStamper returns anew new mock stamper.
func NewStamper() postage.Stamper {
	_ = "STUB: not implemented"
	return *

	// Stamp implements the Stamper interface. It returns an empty postage stamp.
	new(postage.Stamper)
}

func (mockStamper) Stamp(_, _ swarm.Address) (*postage.Stamp, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Stamp implements the Stamper interface. It returns an empty postage stamp.
}

func (mockStamper) BatchId() []byte { _ = "STUB: not implemented"; return nil }
