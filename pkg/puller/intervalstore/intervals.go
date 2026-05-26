// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package intervalstore provides a persistence layer
// for intervals relating to a peer. The store
// provides basic operation such as adding intervals to
// existing ones and persisting the results, as well as
// getting next interval for a peer.
package intervalstore

import (
	"sync"
)

// Intervals store a list of intervals. Its purpose is to provide
// methods to add new intervals and retrieve missing intervals that
// need to be added.
// It may be used in synchronization of streaming data to persist
// retrieved data ranges between sessions.
type Intervals struct {
	start  uint64
	ranges [][2]uint64
	mu     sync.RWMutex
}

// New creates a new instance of Intervals.
// Start argument limits the lower bound of intervals.
// No range below start bound will be added by Add method or
// returned by Next method. This limit may be used for
// tracking "live" synchronization, where the sync session
// starts from a specific value, and if "live" sync intervals
// need to be merged with historical ones, it can be safely done.
func NewIntervals(start uint64) *Intervals { _ = "STUB: not implemented"; return nil }

// Add adds a new range to intervals. Range start and end are values
// are both inclusive.
func (i *Intervals) Add(start, end uint64) { _ = "STUB: not implemented"; return }

func (i *Intervals) add(start, end uint64) { _ = "STUB: not implemented"; return }

// Merge adds all the intervals from the m Interval to current one.
func (i *Intervals) Merge(m *Intervals) { _ = "STUB: not implemented"; return }

// Next returns the first range interval that is not fulfilled. Returned
// start and end values are both inclusive, meaning that the whole range
// including start and end need to be added in order to fill the gap
// in intervals.
// Returned value for end is 0 if the next interval is after the whole
// range that is stored in Intervals. Zero end value represents no limit
// on the next interval length.
// Argument ceiling is the upper bound for the returned range.
// Returned empty boolean indicates if both start and end values have
// reached the ceiling value which means that the returned range is empty,
// not containing a single element.
func (i *Intervals) Next(ceiling uint64) (start, end uint64, empty bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

// Last returns the value that is at the end of the last interval.
func (i *Intervals) Last() (end uint64) { _ = "STUB: not implemented"; return 0 }

// String returns a descriptive representation of range intervals
// in [] notation, as a list of two element vectors.
func (i *Intervals) String() string { _ = "STUB: not implemented"; return "" }

// MarshalBinary encodes Intervals parameters into a semicolon separated list.
// The first element in the list is base36-encoded start value. The following
// elements are two base36-encoded value ranges separated by comma.
func (i *Intervals) MarshalBinary() (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalBinary decodes data according to the Intervals.MarshalBinary format.
func (i *Intervals) UnmarshalBinary(data []byte) (err error) { _ = "STUB: not implemented"; return nil }
