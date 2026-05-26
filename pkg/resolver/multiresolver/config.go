// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package multiresolver

// Defined as per RFC 1034. For reference, see:
// https://en.wikipedia.org/wiki/Domain_Name_System#cite_note-rfc1034-1
const maxTLDLength = 63

// ConnectionConfig contains the TLD, endpoint and contract address used to
// establish to a resolver.
type ConnectionConfig struct {
	TLD      string
	Address  string
	Endpoint string
}

// ParseConnectionString will try to parse a connection string used to connect
// the Resolver to a name resolution service. The resulting config can be
// used to initialize a resolver Service.
func parseConnectionString(cs string) (ConnectionConfig, error) {
	_ = "STUB: not implemented"
	return *new(ConnectionConfig), nil
}

// Split TLD and Endpoint strings.

// Make sure not to grab the protocol, as it contains "://"!
// Eg. in http://... the "http" is NOT a tld.

// Split the address string.

// ParseConnectionStrings will apply ParseConnectionString to each connection
// string. Returns first error found.
func ParseConnectionStrings(cstrs []string) ([]ConnectionConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
