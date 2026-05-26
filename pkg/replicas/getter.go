// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// the code below implements the integration of dispersed replicas in chunk fetching.
// using storage.Getter interface.
package replicas

import (
	"context"
	"errors"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/file/redundancy"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// ErrSwarmageddon is returned in case of a vis mayor called Swarmageddon.
// Swarmageddon is the situation when none of the replicas can be retrieved.
// If 2^{depth} replicas were uploaded and they all have valid postage stamps
// then the probability of Swarmageddon is less than 0.000001
// assuming the error rate of chunk retrievals stays below the level expressed
// as depth by the publisher.
var ErrSwarmageddon = errors.New("swarmageddon has begun")

// getter is the private implementation of storage.Getter, an interface for
// retrieving chunks. This getter embeds the original simple chunk getter and extends it
// to a multiplexed variant that fetches chunks with replicas.
//
// the strategy to retrieve a chunk that has replicas can be configured with a few parameters:
//   - RetryInterval: the delay before a new batch of replicas is fetched.
//   - depth: 2^{depth} is the total number of additional replicas that have been uploaded
//     (by default, it is assumed to be 4, ie. total of 16)
//   - (not implemented) pivot: replicas with address in the proximity of pivot will be tried first
type getter struct {
	wg sync.WaitGroup
	storage.Getter
	level redundancy.Level
}

// NewGetter is the getter constructor
func NewGetter(g storage.Getter, level redundancy.Level) storage.Getter {
	_ = "STUB: not implemented"
	return *new(storage.Getter)
}

// Get makes the getter satisfy the storage.Getter interface
func (g *getter) Get(ctx context.Context, addr swarm.Address) (ch swarm.Chunk, err error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// channel that the results (retrieved chunks) are gathered to from concurrent
// workers each fetching a replica

// errc collects the errors

// concurrently call to retrieve chunk using original CAC address

// counters
// counts the replica addresses tried
// the number of replicas attempted to download in this batch

//

// nil channel to disable case
// addresses used are doubling each period of search expansion
// (at intervals of RetryInterval)

// at least one chunk is retrieved, cancel the rest and return early

// ticker switches on the address channel

// getting the addresses in order
