// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simple

// WalkEntryFunc is the type of the function called for each entry visited
// by WalkEntry.
type WalkEntryFunc func(path string, entry Entry, err error) error

func (m *manifest) WalkEntry(root string, walkFn WalkEntryFunc) (err error) {
	_ = "STUB: not implemented"
	return nil
}
