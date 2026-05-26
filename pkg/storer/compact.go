// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"context"
)

// Compact minimizes sharky disk usage by, using the current sharky locations from the storer,
// relocating chunks starting from the end of the used slots to the first available slots.
func Compact(ctx context.Context, basePath string, opts *Options, validate bool) error {
	_ = "STUB: not implemented"
	return nil
}

// we deliberately choose to iterate the whole store again for each shard
// so that we do not store all the items in memory (for operators with huge localstores)

// marks free and used slots

// start begins at the zero slot. The loop below will increment the position of start until a free slot is found.
// end points to the last slot, and the loop will decrement the position of end until a used slot is found.
// Once start and end point to free and used slots, respectively, the swap of the chunk location will occur.

// walk to the right until a free slot is found

// walk to the left until a used slot found
