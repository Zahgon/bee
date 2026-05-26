// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"github.com/ethersphere/bee/v2/pkg/encryption"
)

type chunkEncrypter struct {
	key []byte
}

func NewChunkEncrypter(key []byte) encryption.ChunkEncrypter {
	_ = "STUB: not implemented"
	return *new(encryption.ChunkEncrypter)
}

func (c *chunkEncrypter) EncryptChunk(chunkData []byte) (encryption.Key, []byte, []byte, error) {
	_ = "STUB: not implemented"
	return *new(encryption.Key), nil, nil, nil
}
