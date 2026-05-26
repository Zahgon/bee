// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libp2p

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/p2p/libp2p/internal/headers/pb"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

func sendHeaders(ctx context.Context, headers p2p.Headers, stream *stream) error {
	_ = "STUB: not implemented"
	return nil
}

func handleHeaders(ctx context.Context, headler p2p.HeadlerFunc, stream *stream, peerAddress swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func headersPBToP2P(h *pb.Headers) p2p.Headers { _ = "STUB: not implemented"; return *new(p2p.Headers) }

func headersP2PToPB(h p2p.Headers) *pb.Headers { _ = "STUB: not implemented"; return nil }
