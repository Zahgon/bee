// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"io"

	"github.com/ethersphere/bee/v2/pkg/file/redundancy"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/spf13/cobra"
)

// putter is a putter that stores all the split chunk addresses of a file
type putter struct {
	cb func(chunk swarm.Chunk) error
}

func (s *putter) Put(_ context.Context, chunk swarm.Chunk) error {
	_ = "STUB: not implemented"
	return nil
}

func newPutter(cb func(ch swarm.Chunk) error) *putter { _ = "STUB: not implemented"; return nil }

var _ storage.Putter = (*putter)(nil)

type pipelineFunc func(context.Context, io.Reader) (swarm.Address, error)

func requestPipelineFn(s storage.Putter, encrypt bool, rLevel redundancy.Level) pipelineFunc {
	_ = "STUB: not implemented"
	return *new(pipelineFunc)
}

func (c *command) initSplitCmd() error { _ = "STUB: not implemented"; return nil }

func splitRefs(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func splitChunks(cmd *cobra.Command) { _ = "STUB: not implemented"; return }
