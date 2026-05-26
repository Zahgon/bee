// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/spf13/cobra"
)

const (
	optionNameValidation     = "validate"
	optionNameCollectionPin  = "pin"
	optionNameOutputLocation = "output"
)

func (c *command) initDBCmd() { _ = "STUB: not implemented"; return }

func dbInfoCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func dbCompactCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func dbValidatePinsCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func dbRepairReserve(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func dbValidateCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func dbExportCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

type noopRadiusSetter struct{}

func (noopRadiusSetter) SetStorageRadius(uint8) { _ = "STUB: not implemented"; return }

func dbExportReserveCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func dbExportPinningCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func dbImportCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func dbImportReserveCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func dbImportPinningCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func dbNukeCmd(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func removeContent(path string) error { _ = "STUB: not implemented"; return nil }

func MarshalChunkToBinary(c swarm.Chunk) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnmarshalChunkFromBinary(data []byte, address string) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// stamp not present
