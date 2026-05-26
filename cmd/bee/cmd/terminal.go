// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/spf13/cobra"
)

type passwordReader interface {
	ReadPassword() (password string, err error)
}

type stdInPasswordReader struct{}

func (stdInPasswordReader) ReadPassword() (password string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func terminalPromptPassword(cmd *cobra.Command, r passwordReader, title string) (password string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func terminalPromptCreatePassword(cmd *cobra.Command, r passwordReader) (password string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}
