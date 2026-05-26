// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const subscribePushEventKey = "subscribe-push"

func (db *DB) SubscribePush(ctx context.Context) (<-chan swarm.Chunk, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

// close the returned chunkInfo channel at the end to
// signal that the subscription is done

// gracefully stop the iteration
// on stop

// if we get storage.ErrNotFound, it could happen that the previous
// iteration happened on a snapshot that was not fully updated yet.
// in this case, we wait for the next event to trigger the iteration
// again. This trigger ensures that we perform the iteration on the
// latest snapshot.

// wait for the next event
