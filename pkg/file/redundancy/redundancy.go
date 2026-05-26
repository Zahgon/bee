// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redundancy

import (
	"github.com/ethersphere/bee/v2/pkg/file/pipeline"
	"github.com/klauspost/reedsolomon"
)

// ParityChunkCallback is called when a new parity chunk has been created
type ParityChunkCallback func(level int, span, address []byte) error

type RedundancyParams interface {
	MaxShards() int // returns the maximum data shard number being used in an intermediate chunk
	Level() Level
	Parities(int) int
	ChunkWrite(int, []byte, ParityChunkCallback) error
	ElevateCarrierChunk(int, ParityChunkCallback) error
	Encode(int, ParityChunkCallback) error
	GetRootData() ([]byte, error)
}

type ErasureEncoder interface {
	Encode([][]byte) error
}

var erasureEncoderFunc = func(shards, parities int) (ErasureEncoder, error) {
	return reedsolomon.New(shards, parities)
}

type Params struct {
	level      Level
	pipeLine   pipeline.PipelineFunc
	buffer     [][][]byte // keeps bytes of chunks on each level for producing erasure coded data; [levelIndex][branchIndex][byteIndex]
	cursor     []int      // index of the current buffered chunk in Buffer. this is basically the latest used branchIndex.
	maxShards  int        // number of chunks after which the parity encode function should be called
	maxParity  int        // number of parity chunks if maxShards has been reached for erasure coding
	encryption bool
}

func New(level Level, encryption bool, pipeLine pipeline.PipelineFunc) *Params {
	_ = "STUB: not implemented"
	return nil
}

// init dataBuffer for erasure coding

// 128 long always because buffer varies at encrypted chunks

func (p *Params) MaxShards() int { _ = "STUB: not implemented"; return 0 }

func (p *Params) Level() Level { _ = "STUB: not implemented"; return *new(Level) }

func (p *Params) Parities(shards int) int { _ = "STUB: not implemented"; return 0 }

// ChunkWrite caches the chunk data on the given chunk level and if it is full then it calls Encode
func (p *Params) ChunkWrite(chunkLevel int, data []byte, callback ParityChunkCallback) error {
	_ = "STUB: not implemented"
	return nil
}

// ChunkWrite caches the chunk data on the given chunk level and if it is full then it calls Encode
func (p *Params) chunkWrite(chunkLevel int, data []byte, callback ParityChunkCallback) error {
	_ = "STUB: not implemented"
	// append chunk to the buffer
	return nil
}

// add parity chunk if it is necessary

// append erasure coded data

// Encode produces and stores parity chunks that will be also passed back to the caller
func (p *Params) Encode(chunkLevel int, callback ParityChunkCallback) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Params) encode(chunkLevel int, callback ParityChunkCallback) error {
	_ = "STUB: not implemented"
	return nil
}

// realloc for parity chunks if it does not override the prev one
// calculate parity chunks

// ElevateCarrierChunk moves the last poor orphan chunk to the level above where it can fit and there are other chunks as well.
func (p *Params) ElevateCarrierChunk(chunkLevel int, callback ParityChunkCallback) error {
	_ = "STUB: not implemented"
	return nil
}

// not necessary to update current level since we will not work with it anymore

// GetRootData returns the topmost chunk in the tree.
// throws and error if the encoding has not been finished in the BMT
// OR redundancy is not used in the BMT
func (p *Params) GetRootData() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
