// Copyright 2025 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bzz

import (
	"github.com/multiformats/go-multiaddr"
)

// underlayListPrefix is a magic byte designated for identifying a serialized list of multiaddrs.
// A value of 0x99 (153) was chosen as it is not a defined multiaddr protocol code.
// This ensures that a failure is triggered by the original multiaddr.NewMultiaddrBytes function,
// which expects a valid protocol code at the start of the data.
const underlayListPrefix byte = 0x99

// SerializeUnderlays serializes a slice of multiaddrs into a single byte slice.
// If the slice contains exactly one address, the standard, backward-compatible
// multiaddr format is used. For zero or more than one address, a custom list format
// prefixed with a magic byte is utilized.
func SerializeUnderlays(addrs []multiaddr.Multiaddr) []byte {
	_ = "STUB: not implemented"
	// Backward compatibility if exactly one address is present.
	return nil
}

// For 0 or 2+ addresses, the custom list format with the prefix is used.
// The format is: [prefix_byte][varint_len_1][addr_1_bytes]...

// DeserializeUnderlays deserializes a byte slice into a slice of multiaddrs.
// The data format is automatically detected as either a single legacy multiaddr
// or a list of multiaddrs (identified by underlayListPrefix), and is parsed accordingly.
func DeserializeUnderlays(data []byte) ([]multiaddr.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the data begins with the magic prefix, it is handled as a list.

// Otherwise, the data is handled as a single, backward-compatible multiaddr.

// The result is returned as a single-element slice for a consistent return type.

// deserializeList handles the parsing of the custom list format.
// The provided data is expected to have already been stripped of the underlayListPrefix.
func deserializeList(data []byte) ([]multiaddr.Multiaddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The varint-encoded length of the next address is read.

// A sanity check is performed to ensure enough bytes remain for the declared length.

// The individual address bytes are read and parsed.
