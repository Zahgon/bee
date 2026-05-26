// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redundancy

// EncodeLevel encodes used redundancy level for uploading into span keeping the real byte count for the chunk.
// assumes span is LittleEndian
func EncodeLevel(span []byte, level Level) {
	_ = "STUB: not implemented"
	// set parity in the most signifact byte
	return
}

// p + 128

// DecodeSpan decodes the used redundancy level from span keeping the real byte count for the chunk.
// assumes span is LittleEndian
func DecodeSpan(span []byte) (Level, []byte) { _ = "STUB: not implemented"; return *new(Level), nil }

// IsLevelEncoded checks whether the redundancy level is encoded in the span
// assumes span is LittleEndian
func IsLevelEncoded(span []byte) bool { _ = "STUB: not implemented"; return false }
