// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chequebook

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/transaction"
	"github.com/ethersphere/bee/v2/pkg/util/abiutil"
	"github.com/ethersphere/go-sw3-abi/sw3abi"
)

var (
	ErrInvalidFactory       = errors.New("not a valid factory contract")
	ErrNotDeployedByFactory = errors.New("chequebook not deployed by factory")
	errDecodeABI            = errors.New("could not decode abi data")

	factoryABI                  = abiutil.MustParseABI(sw3abi.SimpleSwapFactoryABIv0_6_9)
	simpleSwapDeployedEventType = factoryABI.Events["SimpleSwapDeployed"]
)

// Factory is the main interface for interacting with the chequebook factory.
type Factory interface {
	// ERC20Address returns the token for which this factory deploys chequebooks.
	ERC20Address(ctx context.Context) (common.Address, error)
	// Deploy deploys a new chequebook and returns once the transaction has been submitted.
	Deploy(ctx context.Context, issuer common.Address, defaultHardDepositTimeoutDuration *big.Int, nonce common.Hash) (common.Hash, error)
	// WaitDeployed waits for the deployment transaction to confirm and returns the chequebook address
	WaitDeployed(ctx context.Context, txHash common.Hash) (common.Address, error)
	// VerifyChequebook checks that the supplied chequebook has been deployed by this factory.
	VerifyChequebook(ctx context.Context, chequebook common.Address) error
}

type factory struct {
	backend            transaction.Backend
	transactionService transaction.Service
	address            common.Address // address of the factory to use for deployments
}

type simpleSwapDeployedEvent struct {
	ContractAddress common.Address
}

// NewFactory creates a new factory service for the provided factory contract.
func NewFactory(backend transaction.Backend, transactionService transaction.Service, address common.Address) Factory {
	_ = "STUB: not implemented"
	return *new(Factory)
}

// Deploy deploys a new chequebook and returns once the transaction has been submitted.
func (c *factory) Deploy(ctx context.Context, issuer common.Address, defaultHardDepositTimeoutDuration *big.Int, nonce common.Hash) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// WaitDeployed waits for the deployment transaction to confirm and returns the chequebook address
func (c *factory) WaitDeployed(ctx context.Context, txHash common.Hash) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}

func (c *factory) verifyChequebookAgainstFactory(ctx context.Context, factory, chequebook common.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// VerifyChequebook checks that the supplied chequebook has been deployed by a supported factory.
func (c *factory) VerifyChequebook(ctx context.Context, chequebook common.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// ERC20Address returns the token for which this factory deploys chequebooks.
func (c *factory) ERC20Address(ctx context.Context) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}
