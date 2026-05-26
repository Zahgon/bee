// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"context"

	"github.com/ethersphere/bee/v2/pkg/encryption"
	storage "github.com/ethersphere/bee/v2/pkg/storage"
)

// maximum amount of file tree levels this file hasher component can handle
// (128 ^ (9 - 1)) * 4096 = 295147905179352825856 bytes
const levelBufferLimit = 9

// SimpleSplitterJob encapsulated a single splitter operation, accepting blockwise
// writes of data whose length is defined in advance.
//
// After the job is constructed, Write must be called with up to ChunkSize byte slices
// until the full data length has been written. The Sum should be called which will
// return the SwarmHash of the data.
//
// Called Sum before the last Write, or Write after Sum has been called, may result in
// error and will may result in undefined result.
type SimpleSplitterJob struct {
	ctx        context.Context
	putter     storage.Putter
	spanLength int64  // target length of data
	length     int64  // number of bytes written to the data level of the hasher
	sumCounts  []int  // number of sums performed, indexed per level
	cursors    []int  // section write position, indexed per level
	buffer     []byte // keeps data and hashes, indexed by cursors
	toEncrypt  bool   // to encryrpt the chunks or not
	refSize    int64
}

// NewSimpleSplitterJob creates a new SimpleSplitterJob.
//
// The spanLength is the length of the data that will be written.
func NewSimpleSplitterJob(ctx context.Context, putter storage.Putter, spanLength int64, toEncrypt bool) *SimpleSplitterJob {
	_ = "STUB: not implemented"
	return nil
}

// double size as temp workaround for weak calculation of needed buffer space

// Write adds data to the file splitter.
func (j *SimpleSplitterJob) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Sum returns the Swarm hash of the data.
func (j *SimpleSplitterJob) Sum(b []byte) []byte {
	_ = "STUB: not implemented"

	// writeToLevel writes to the data buffer on the specified level.
	// It calls sum if chunk boundary is reached and recursively calls this function for
	// the next level with the acquired bmt hash
	//
	// It adjusts the relevant levels' cursors accordingly.
	return nil
}

func (s *SimpleSplitterJob) writeToLevel(lvl int, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// sumLevel calculates and returns the bmt sum of the last written data on the level.
//
// TODO: error handling on store write fail
func (s *SimpleSplitterJob) sumLevel(lvl int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// digest returns the calculated digest after a Sum call.
//
// The hash returned is the hash in the first section index of the work buffer
// this will be the root hash when all recursive sums have completed.
//
// The method does not check that the final hash actually has been written, so
// timing is the responsibility of the caller.
func (s *SimpleSplitterJob) digest() []byte { _ = "STUB: not implemented"; return nil }

// hashUnfinished hasher the remaining unhashed chunks at the end of each level if
// write doesn't end on a chunk boundary.
func (s *SimpleSplitterJob) hashUnfinished() error { _ = "STUB: not implemented"; return nil }

// nolint:gofmt
// moveDanglingChunk concatenates the reference to the single reference
// at the highest level of the tree in case of a balanced tree.
//
// Let F be full chunks (disregarding branching factor) and S be single references
// in the following scenario:
//
//	    S
//	  F   F
//	F   F   F
//
// # F   F   F   F S
//
// The result will be:
//
//	    SS
//	  F    F
//	F   F   F
//
// # F   F   F   F
//
// After which the SS will be hashed to obtain the final root hash
func (s *SimpleSplitterJob) moveDanglingChunk() error {
	_ = "STUB: not implemented"
	// calculate the total number of levels needed to represent the data (including the data level)
	return nil
}

// sum every intermediate level and write to the level above it

// and if there is a single reference outside a balanced tree on this level
// don't hash it again but pass it on to the next level

// TODO: simplify if possible

func (s *SimpleSplitterJob) encryptChunkData(chunkData []byte) ([]byte, encryption.Key, error) {
	_ = "STUB: not implemented"
	return nil, *new(encryption.Key), nil
}

func (s *SimpleSplitterJob) encrypt(chunkData []byte) (encryption.Key, []byte, []byte, error) {
	_ = "STUB: not implemented"
	return *new(encryption.Key), nil, nil, nil
}
