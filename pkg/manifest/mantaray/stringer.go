// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mantaray

import (
	"io"
)

//nolint:errcheck
func (n *Node) String() string { _ = "STUB: not implemented"; return "" }

//nolint:errcheck
func nodeStringWithPrefix(n *Node, prefix string, writer io.Writer) {
	_ = "STUB: not implemented"
	return
}

var tableCharsMap = map[string]string{
	"top":          "─",
	"top-mid":      "┬",
	"top-left":     "┌",
	"top-right":    "┐",
	"bottom":       "─",
	"bottom-mid":   "┴",
	"bottom-left":  "└",
	"bottom-right": "┘",
	"left":         "│",
	"left-mid":     "├",
	"mid":          "─",
	"mid-mid":      "┼",
	"right":        "│",
	"right-mid":    "┤",
	"middle":       "│",
}
