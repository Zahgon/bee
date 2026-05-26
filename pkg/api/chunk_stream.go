// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"
	"time"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/storer"
	"github.com/gorilla/websocket"
)

const streamReadTimeout = 15 * time.Minute

var successWsMsg = []byte{}

func (s *Service) chunkUploadStreamHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Optional: omit if caller provides pre-signed stamps per chunk

// Fallback: read tag from query parameter (browser WebSocket can't set headers)

// Create connection-level putter only if BatchID is provided.
// If BatchID is not provided, the API caller is expected to provide
// pre-signed stamps with each chunk (and is also expected to keep
// track of stamp state over time).

// if tag not specified use direct upload
// Using context.Background here because the putter's lifetime extends beyond that of the HTTP request.

// chunkDecoder extracts chunk data and optionally a stamp from a websocket message.
// When BatchID is provided in headers, decodeChunkWithoutStamp is used (no stamp in message).
// When BatchID is not provided, decodeChunkWithStamp is used (stamp prepended to chunk data).
type chunkDecoder func(msg []byte) (chunkData []byte, stamp *postage.Stamp, err error)

// decodeChunkWithoutStamp returns the message as-is (used when BatchID provided in headers).
func decodeChunkWithoutStamp(msg []byte) ([]byte, *postage.Stamp, error) {
	_ = "STUB: not implemented"
	return nil,

		// decodeChunkWithStamp extracts a stamp from the first 113 bytes of the message.
		// Returns an error if the message is too small or the stamp is invalid.
		nil, nil
}

func decodeChunkWithStamp(msg []byte) ([]byte, *postage.Stamp, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (s *Service) handleUploadStream(
	logger log.Logger,
	conn *websocket.Conn,
	putter storer.PutterSession,
	tag uint64,
	decode chunkDecoder,
) {
	_ = "STUB: not implemented"
	return
}

// Cache for batch validation to avoid database lookups for every chunk
// Key: batch ID, Value: stored batch info
// This avoids the expensive batchStore.Get() call for each chunk

// No cleanup needed for batch cache - it's just metadata

// Only call Done on connection-level putter if it exists

// shutdown

// client gone

// if there is no indication to stop, go ahead and read the next message

// Decode the message using the appropriate decoder

// Determine the putter to use

// If stamp was extracted, create a per-chunk putter

// Clean up per-chunk putter
