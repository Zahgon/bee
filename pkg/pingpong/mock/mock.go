// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"time"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type Service struct {
	pingFunc func(ctx context.Context, address swarm.Address, msgs ...string) (rtt time.Duration, err error)
}

func New(pingFunc func(ctx context.Context, address swarm.Address, msgs ...string) (rtt time.Duration, err error)) *Service {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) Ping(ctx context.Context, address swarm.Address, msgs ...string) (rtt time.Duration, err error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
