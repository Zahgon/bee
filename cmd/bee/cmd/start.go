// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"crypto/ecdsa"
	_ "embed"
	"time"

	"github.com/ethersphere/bee/v2/pkg/accesscontrol"
	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/node"
	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

const (
	serviceName      = "SwarmBeeSvc"
	libp2pPKFilename = "libp2p_v2"
)

//go:embed bee-welcome-message.txt
var beeWelcomeMessage string

func (c *command) initStartCmd() (err error) { _ = "STUB: not implemented"; return nil }

// ctx is global context of bee node; which is canceled after interrupt signal is received.

// Building bee node can take up some time (because node.NewBee(...) is compute have function )
// Because of this we need to do it in background so that program could be terminated when interrupt signal is received
// while bee node is being constructed.

// Wait for bee node to fully build and initialized

// Bee has fully started at this point, from now on we
// block main goroutine until it is interrupted or stopped

// Whenever program is being stopped we need to cancel main context
// beforehand so that node could be stopped via Shutdown method

// Shutdown node (if node was fully started)

// If shutdown function is blocking too long,
// allow process termination by receiving another signal.

// start blocks until some interrupt is received

type buildBeeNodeResp struct {
	bee *node.Bee
	err error
}

func buildBeeNodeAsync(ctx context.Context, c *command, cmd *cobra.Command, logger log.Logger) <-chan buildBeeNodeResp {
	_ = "STUB: not implemented"
	return nil
}

func buildBeeNode(ctx context.Context, c *command, cmd *cobra.Command, logger log.Logger) (*node.Bee, error) {
	_ = "STUB: not implemented"

	// If the resolver is specified, resolve all connection strings
	// and fail on any errors.
	return nil, nil
}

// if the user has provided a value - we use it and overwrite the default
// if mainnet is true then we only accept networkID value 1, error otherwise
// if the user has not provided a network ID but mainnet is true - just overwrite with mainnet network ID (1)
// in all the other cases we default to test network ID (10)

// Rebuild the global bmtpool instance so the new SIMDOptIn value
// is reflected in the pool created for hot-path BMT hashing.

type program struct {
	start func()
	stop  func()
}

func (p *program) Start(s service.Service) error {
	_ = "STUB: not implemented"
	// Start should not block. Do the actual work async.
	return nil
}

func (p *program) Stop(s service.Service) error { _ = "STUB: not implemented"; return nil }

type signerConfig struct {
	signer           crypto.Signer
	publicKey        *ecdsa.PublicKey
	libp2pPrivateKey *ecdsa.PrivateKey
	pssPrivateKey    *ecdsa.PrivateKey
	session          accesscontrol.Session
}

func (c *command) configureSigner(cmd *cobra.Command, logger log.Logger) (config *signerConfig, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if libp2p key exists we can assume all required keys exist
// so prompt for a password to unlock them
// otherwise prompt for new password with confirmation to create them

// postinst and post scripts inside packaging/{deb,rpm} depend and parse on this log output

type networkConfig struct {
	bootNodes []string
	blockTime time.Duration
	chainID   int64
}

func getConfigByNetworkID(networkID uint64, defaultBlockTimeInSeconds uint64) *networkConfig {
	_ = "STUB: not implemented"
	return nil
}

// Staging.

// Will use the value provided by the chain.
