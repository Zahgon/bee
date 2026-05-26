// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"github.com/ethersphere/bee/v2/pkg/pusher"
	"github.com/ethersphere/bee/v2/pkg/storage"
)

// DirectUpload is the implementation of the NetStore.DirectUpload method.
func (db *DB) DirectUpload() PutterSession {
	_ = "STUB: not implemented"
	// egCtx will allow early exit of Put operations if we have
	// already encountered error.
	return *new(PutterSession)
}

// Download is the implementation of the NetStore.Download method.
func (db *DB) Download(cache bool) storage.Getter {
	_ = "STUB: not implemented"
	return *new(storage.Getter)
}

// if chunk is not found locally, retrieve it from the network

// PusherFeed is the implementation of the NetStore.PusherFeed method.
func (db *DB) PusherFeed() <-chan *pusher.Op { _ = "STUB: not implemented"; return nil }
