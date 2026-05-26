// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package store

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/encryption"
	storage "github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type decryptingStore struct {
	storage.Getter
}

func New(s storage.Getter) storage.Getter { _ = "STUB: not implemented"; return *new(storage.Getter) }

func (s *decryptingStore) Get(ctx context.Context, addr swarm.Address) (ch swarm.Chunk, err error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

// normal, unencrypted content

// encrypted reference

func DecryptChunkData(chunkData []byte, encryptionKey encryption.Key) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// removing extra bytes which were just added for padding

func decrypt(chunkData []byte, key encryption.Key) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
