// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ens

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	goens "github.com/wealdtech/go-ens/v3"

	"github.com/ethersphere/bee/v2/pkg/resolver/client"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

const (
	defaultENSContractAddress = "00000000000C2E074eC69A0dFb2997BA6C7d2e1e"
	swarmContentHashPrefix    = "bzz://"
)

// Address is the swarm bzz address.
type Address = swarm.Address

// Make sure Client implements the resolver.Client interface.
var _ client.Interface = (*Client)(nil)

var (
	// ErrFailedToConnect denotes that the resolver failed to connect to the
	// provided endpoint.
	ErrFailedToConnect = errors.New("failed to connect")
	// ErrResolveFailed denotes that a name could not be resolved.
	ErrResolveFailed = errors.New("resolve failed")
	// ErrNotImplemented denotes that the function has not been implemented.
	ErrNotImplemented = errors.New("function not implemented")
	// errNameNotRegistered denotes that the name is not registered.
	errNameNotRegistered = errors.New("name is not registered")
)

// Client is a name resolution client that can connect to ENS via an
// Ethereum endpoint.
type Client struct {
	endpoint     string
	contractAddr string
	ethCl        *ethclient.Client
	connectFn    func(string, string) (*ethclient.Client, *goens.Registry, error)
	resolveFn    func(*goens.Registry, common.Address, string) (string, error)
	registry     *goens.Registry
}

// Option is a function that applies an option to a Client.
type Option func(*Client)

// NewClient will return a new Client.
func NewClient(endpoint string, opts ...Option) (client.Interface, error) {
	_ = "STUB: not implemented"
	return *new(client.Interface), nil
}

// Apply all options to the Client.

// Set the default ENS contract address.

// Establish a connection to the ENS.

// WithContractAddress will set the ENS contract address.
func WithContractAddress(addr string) Option { _ = "STUB: not implemented"; return *new(Option) }

// IsConnected returns true if there is an active RPC connection with an
// Ethereum node at the configured endpoint.
func (c *Client) IsConnected() bool { _ = "STUB: not implemented"; return false }

// Endpoint returns the endpoint the client was connected to.
func (c *Client) Endpoint() string {
	_ = "STUB: not implemented"

	// Resolve implements the resolver.Client interface.
	return ""
}

func (c *Client) Resolve(name string) (Address, error) {
	_ = "STUB: not implemented"
	return *new(Address), nil
}

// Ensure that the content hash string is in a valid format, eg.
// "bzz://<address>".

// Trim the prefix and try to parse the result as a bzz address.

// Close closes the RPC connection with the client, terminating all unfinished
// requests. If the connection is already closed, this call is a noop.
func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }

func wrapDial(endpoint, contractAddr string) (*ethclient.Client, *goens.Registry, error) {
	_ = "STUB: not implemented"
	// Dial the eth client.
	return nil, nil, nil
}

// Obtain the ENS registry.

// Ensure that the ENS registry client is deployed to the given contract address.

func wrapResolve(registry *goens.Registry, _ common.Address, name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// it returns error only if the service is not available

// If the name is not registered, return an error.

// Obtain the resolver for this domain name.

// Try and read out the content hash record.

// Check if it's a service error (rate limiting, network issues)
