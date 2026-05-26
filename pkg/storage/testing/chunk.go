// nolint:goheader
// Copyright 2019 The Swarm Authors
// This file is part of the Swarm library.
//
// The Swarm library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Swarm library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Swarm library. If not, see <http://www.gnu.org/licenses/>.

package testing

import (
	"testing"

	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// GenerateTestRandomChunk generates a valid content addressed chunk.
func GenerateTestRandomChunk() swarm.Chunk { _ = "STUB: not implemented"; return *new(swarm.Chunk) }

// GenerateTestRandomSoChunk generates a valid single owner chunk
// using supplied content addressed chunk.
func GenerateTestRandomSoChunk(tb testing.TB, cac swarm.Chunk) swarm.Chunk {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk)
}

// GenerateTestRandomInvalidChunk generates a random, however invalid, content
// addressed chunk. The chunk is also guaranteed to be invalid as a SOC: with
// purely random bytes, the embedded signature recovery byte falls inside
// btcec's valid range often enough (~0.7%) that soc.FromChunk would succeed
// and consumers expecting an "invalid" chunk see flaky behavior. Forcing the
// recovery byte (data[swarm.HashSize+swarm.SocSignatureSize-1]) outside the
// valid 27..34 range makes signature recovery deterministically fail.
func GenerateTestRandomInvalidChunk() swarm.Chunk {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk)
}

// GenerateTestRandomChunks generates a slice of random
// Chunks by using GenerateTestRandomChunk function.
func GenerateTestRandomChunks(count int) []swarm.Chunk { _ = "STUB: not implemented"; return nil }

// GenerateTestRandomChunkAt generates an invalid (!) chunk with address of proximity order po wrt target.
func GenerateTestRandomChunkAt(tb testing.TB, target swarm.Address, po int) swarm.Chunk {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk)
}

// GenerateValidRandomChunkAt generates an valid chunk with address of proximity order po wrt target.
func GenerateValidRandomChunkAt(tb testing.TB, target swarm.Address, po int) swarm.Chunk {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk)
}

// FixtureChunk gets a pregenerated content-addressed chunk and
// panics if one is not found.
func FixtureChunk(prefix string) swarm.Chunk { _ = "STUB: not implemented"; return *new(swarm.Chunk) }

// fixtureChunks returns pregenerated content-addressed chunks necessary for explicit
// test scenarios where random generated chunks are not good enough.
func fixtureChunks() map[string]swarm.Chunk { _ = "STUB: not implemented"; return nil }
