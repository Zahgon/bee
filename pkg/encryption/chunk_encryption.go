// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encryption

// ChunkEncrypter encrypts chunk data.
type ChunkEncrypter interface {
	EncryptChunk([]byte) (key Key, encryptedSpan, encryptedData []byte, err error)
}

type chunkEncrypter struct{}

func NewChunkEncrypter() ChunkEncrypter { _ = "STUB: not implemented"; return *new(ChunkEncrypter) }

func (c *chunkEncrypter) EncryptChunk(chunkData []byte) (Key, []byte, []byte, error) {
	_ = "STUB: not implemented"
	return *new(Key), nil, nil, nil
}

func NewSpanEncryption(key Key) Interface { _ = "STUB: not implemented"; return *new(Interface) }

func NewDataEncryption(key Key) Interface { _ = "STUB: not implemented"; return *new(Interface) }
