// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"sync"

	"github.com/ethersphere/bee/v2/pkg/file/pipeline"
)

type MockChainWriter struct {
	sync.Mutex
	chainWriteCalls int
	sumCalls        int
}

func NewChainWriter() *MockChainWriter { _ = "STUB: not implemented"; return nil }

func (c *MockChainWriter) ChainWrite(_ *pipeline.PipeWriteArgs) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *MockChainWriter) Sum() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *MockChainWriter) ChainWriteCalls() int { _ = "STUB: not implemented"; return 0 }
func (c *MockChainWriter) SumCalls() int        { _ = "STUB: not implemented"; return 0 }
