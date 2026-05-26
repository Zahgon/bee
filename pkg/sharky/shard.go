// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sharky

import (
	"context"
	"io"
)

// LocationSize is the size of the byte representation of Location
const LocationSize int = 7

// Location models the location <shard, slot, length> of a chunk
type Location struct {
	Shard  uint8
	Slot   uint32
	Length uint16
}

func (l Location) String() string { _ = "STUB: not implemented"; return "" }

// MarshalBinary returns byte representation of location
func (l *Location) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalBinary constructs the location from byte representation
func (l *Location) UnmarshalBinary(buf []byte) error { _ = "STUB: not implemented"; return nil }

// LocationFromBinary is a helper to construct a Location object from byte representation
func LocationFromBinary(buf []byte) (Location, error) {
	_ = "STUB: not implemented"
	return *new(Location), nil
}

// sharkyFile defines the minimal interface that is required for a file type for it to
// be usable in sharky. This allows us to have different implementations of file types
// that can continue using the sharky logic
type sharkyFile interface {
	io.ReadWriteCloser
	io.ReaderAt
	io.Seeker
	io.WriterAt
	Truncate(int64) error
	Sync() error
}

// write models the input to a write operation
type write struct {
	buf []byte     // variable size read buffer
	res chan entry // to put the result through
}

// entry models the output result of a write operation
type entry struct {
	loc Location // shard, slot, length combo
	err error    // signal for end of operation
}

// read models the input to read operation (the output is an error)
type read struct {
	ctx  context.Context
	buf  []byte // variable size read buffer
	slot uint32 // slot to read from
}

// shard models a shard writing to a file with periodic offsets due to fixed maxDataSize
type shard struct {
	reads       chan read     // channel for reads
	errc        chan error    // result for reads
	writes      chan write    // channel for writes
	index       uint8         // index of the shard
	maxDataSize int           // max size of blobs
	file        sharkyFile    // the file handle the shard is writing data to
	slots       *slots        // component keeping track of freed slots
	quit        chan struct{} // channel to signal quitting
}

// forever loop processing
func (sh *shard) process() { _ = "STUB: not implemented"; return }

// this condition checks if an slot is in limbo (popped but not used for write op)

// since the goroutine in the Read method can quit
// on shutdown, we need to make sure that we can actually
// write to the channel, since a shutdown is possible in
// theory between after the point that the context is cancelled

// since the Read method respects the quit channel
// we can safely quit here without writing to the channel

// only enabled if there is a free slot previously popped

// re-enable popping a free slot next time we can write
// disable popping a write operation until there is a free slot

// pop a free slot

// only if there is one can we pop a chunk to write otherwise keep back pressure on writes
// effectively enforcing another shard to be chosen
// enable popping a write operation
// disabling getting a new slot until a write is actually done

// close closes the shard:
// wait for pending operations to finish then saves free slots and blobs on disk
func (sh *shard) close() error { _ = "STUB: not implemented"; return nil }

// offset calculates the offset from the slot
// this is possible since all blobs are of fixed size
func (sh *shard) offset(slot uint32) int64 { _ = "STUB: not implemented"; return 0 }

// read reads loc.Length bytes to the buffer from the blob slot loc.Slot
func (sh *shard) read(r read) error { _ = "STUB: not implemented"; return nil }

// write writes loc.Length bytes to the buffer from the blob slot loc.Slot
func (sh *shard) write(buf []byte, slot uint32) entry {
	_ = "STUB: not implemented"
	return *new(entry)
}

// release frees the slot allowing new entry to overwrite
func (sh *shard) release(ctx context.Context, slot uint32) error {
	_ = "STUB: not implemented"
	return nil
}
