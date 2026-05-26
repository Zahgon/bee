// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package multiresolver

import (
	"errors"
	"fmt"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/resolver"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "multiresolver"

// Ensure MultiResolver implements Resolver interface.
var _ resolver.Interface = (*MultiResolver)(nil)

var (
	// ErrTLDTooLong denotes when a TLD in a name exceeds maximum length.
	ErrTLDTooLong = fmt.Errorf("TLD exceeds maximum length of %d characters", maxTLDLength)
	// ErrInvalidTLD denotes passing an invalid TLD to the MultiResolver.
	ErrInvalidTLD = errors.New("invalid TLD")
	// ErrResolverChainEmpty denotes trying to pop an empty resolver chain.
	ErrResolverChainEmpty = errors.New("resolver chain empty")
	// ErrResolverChainFailed denotes that an entire name resolution chain
	// for a given TLD failed.
	ErrResolverChainFailed = errors.New("resolver chain failed")
	// ErrCloseFailed denotes that closing the multiresolver failed.
	ErrCloseFailed = errors.New("close failed")
	// ErrResolverService denotes that no resolver service is configured for the requested name or the resolver service is not available.
	ErrResolverService = errors.New("cannot communicate with the resolver or no resolver service configured")
)

type resolverMap map[string][]resolver.Interface

// MultiResolver performs name resolutions based on the TLD label in the name.
type MultiResolver struct {
	resolvers resolverMap
	logger    log.Logger
	cfgs      []ConnectionConfig
	// ForceDefault will force all names to be resolved by the default
	// resolution chain, regadless of their TLD.
	ForceDefault bool
}

// Option is a function that applies an option to a MultiResolver.
type Option func(*MultiResolver)

// NewMultiResolver will return a new MultiResolver instance.
func NewMultiResolver(opts ...Option) *MultiResolver { _ = "STUB: not implemented"; return nil }

// Apply all options.

// Discard log output by default.

// Attempt to connect to each resolver using the connection string.

// NOTE: if we want to create a specific client based on the TLD
// we can do it here.

// WithConnectionConfigs will set the initial connection configuration.
func WithConnectionConfigs(cfgs []ConnectionConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLogger will set the logger used by the MultiResolver.
func WithLogger(logger log.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithForceDefault will force resolution using the default resolver chain.
func WithForceDefault() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDefaultCIDResolver() Option { _ = "STUB: not implemented"; return *new(Option) }

// PushResolver will push a new Resolver to the name resolution chain for the
// given TLD. An empty TLD will push to the default resolver chain.
func (mr *MultiResolver) PushResolver(tld string, r resolver.Interface) {
	_ = "STUB: not implemented"
	return
}

// PopResolver will pop the last resolver from the name resolution chain for the
// given TLD. An empty TLD will pop from the default resolver chain.
func (mr *MultiResolver) PopResolver(tld string) error { _ = "STUB: not implemented"; return nil }

// ChainCount returns the number of resolvers in a resolver chain for the given
// tld.
// TLD names should be prepended with a dot (eg ".tld"). An empty TLD will
// return the number of resolvers in the default resolver chain.
func (mr *MultiResolver) ChainCount(tld string) int { _ = "STUB: not implemented"; return 0 }

// GetChain will return the resolution chain for a given TLD.
// TLD names should be prepended with a dot (eg ".tld"). An empty TLD will
// return all resolvers in the default resolver chain.
func (mr *MultiResolver) GetChain(tld string) []resolver.Interface {
	_ = "STUB: not implemented"
	return nil

	// Resolve will attempt to resolve a name to an address.
	// The resolution chain is selected based on the TLD of the name. If the name
	// does not end in a TLD, the default resolution chain is selected.
	// The resolution will be performed iteratively on the resolution chain,
	// returning the result of the first Resolver that succeeds. If all resolvers
	// in the chain return an error, the function will return an ErrResolveFailed.
}

func (mr *MultiResolver) Resolve(name string) (addr resolver.Address, err error) {
	_ = "STUB: not implemented"
	return *new(resolver.Address), nil
}

// If no resolver chain is found, switch to the default chain.

// only the CID resolver is defined

// Close all will call Close on all resolvers in all resolver chains.
func (mr *MultiResolver) Close() error { _ = "STUB: not implemented"; return nil }

func getTLD(name string) string { _ = "STUB: not implemented"; return "" }

func (mr *MultiResolver) connectENSClient(tld, address, endpoint string) {
	_ = "STUB: not implemented"
	return
}
