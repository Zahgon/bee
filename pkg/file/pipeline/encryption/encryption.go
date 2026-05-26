// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encryption

import (
	"github.com/ethersphere/bee/v2/pkg/encryption"
	"github.com/ethersphere/bee/v2/pkg/file/pipeline"
)

type encryptionWriter struct {
	next pipeline.ChainWriter
	enc  encryption.ChunkEncrypter
}

func NewEncryptionWriter(encrypter encryption.ChunkEncrypter, next pipeline.ChainWriter) pipeline.ChainWriter {
	_ = "STUB: not implemented"
	return *new(pipeline.ChainWriter)
}

// Write assumes that the span is prepended to the actual data before the write !
func (e *encryptionWriter) ChainWrite(p *pipeline.PipeWriteArgs) error {
	_ = "STUB: not implemented"
	return nil
}

// replace the verbatim data with the encrypted data

func (e *encryptionWriter) Sum() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
