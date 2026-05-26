// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"github.com/ethersphere/bee/v2/pkg/resolver"
)

// Assure mock Resolver implements the Resolver interface.
var _ resolver.Interface = (*Resolver)(nil)

// Resolver is the mock Resolver implementation.
type Resolver struct {
	IsClosed    bool
	resolveFunc func(string) (resolver.Address, error)
}

// Option function sets the option on the mock Resolver.
type Option func(*Resolver)

// NewResolver will create a new mock Resolver.
func NewResolver(opts ...Option) resolver.Interface {
	_ = "STUB: not implemented"

	// Apply all options.
	return *new(resolver.Interface)
}

// WithResolveFunc will override the Resolve function implementation.
func WithResolveFunc(f func(string) (resolver.Address, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Resolve implements the Resolver interface.
func (r *Resolver) Resolve(name string) (resolver.Address, error) {
	_ = "STUB: not implemented"
	return *new(resolver.Address), nil
}

// Close implements the Resolver interface.
func (r *Resolver) Close() error { _ = "STUB: not implemented"; return nil }
