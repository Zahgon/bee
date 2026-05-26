// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"github.com/ethersphere/bee/v2/pkg/resolver/client"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// Ensure mock Client implements the Client interface.
var _ client.Interface = (*Client)(nil)

// Client is the mock resolver client implementation.
type Client struct {
	isConnected    bool
	endpoint       string
	defaultAddress swarm.Address
	resolveFn      func(string) (swarm.Address, error)
}

// Option is a function that applies an option to a Client.
type Option func(*Client)

// NewClient construct a new mock Client.
func NewClient(opts ...Option) *Client { _ = "STUB: not implemented"; return nil }

// WithEndpoint will set the endpoint.
func WithEndpoint(endpoint string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WitResolveAddress will set the address returned by Resolve.
func WitResolveAddress(addr swarm.Address) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithResolveFunc will set the Resolve function implementation.
func WithResolveFunc(fn func(string) (swarm.Address, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// IsConnected is the mock IsConnected implementation.
func (cl *Client) IsConnected() bool { _ = "STUB: not implemented"; return false }

// Endpoint is the mock Endpoint implementation.
func (cl *Client) Endpoint() string {
	_ = "STUB: not implemented"

	// Resolve is the mock Resolve implementation
	return ""
}

func (cl *Client) Resolve(name string) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// Close is the mock Close implementation.
func (cl *Client) Close() error { _ = "STUB: not implemented"; return nil }
