// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package libp2p

import (
	protocol "github.com/libp2p/go-libp2p/core/protocol"
)

// protocolSemverMatcher returns a matcher function for a given base protocol.
// Protocol ID must be constructed according to the Swarm protocol ID
// specification, where the second to last part is the version is semver format.
// The matcher function will return a boolean indicating whether a protocol ID
// matches the base protocol. A given protocol ID matches the base protocol if
// the IDs are the same and if the semantic version of the base protocol is the
// same or higher than that of the protocol ID provided.
func (s *Service) protocolSemverMatcher(base protocol.ID) (func(protocol.ID) bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
