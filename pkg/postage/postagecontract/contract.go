// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package postagecontract

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/transaction"
	"github.com/ethersphere/bee/v2/pkg/util/abiutil"
	"github.com/ethersphere/go-sw3-abi/sw3abi"
)

var (
	BucketDepth = uint8(16)

	erc20ABI = abiutil.MustParseABI(sw3abi.ERC20ABIv0_6_9)

	ErrBatchCreate          = errors.New("batch creation failed")
	ErrInsufficientFunds    = errors.New("insufficient token balance")
	ErrInvalidDepth         = errors.New("invalid depth")
	ErrBatchTopUp           = errors.New("batch topUp failed")
	ErrBatchDilute          = errors.New("batch dilute failed")
	ErrChainDisabled        = errors.New("chain disabled")
	ErrNotImplemented       = errors.New("not implemented")
	ErrInsufficientValidity = errors.New("insufficient validity")

	approveDescription     = "Approve tokens for postage operations"
	createBatchDescription = "Postage batch creation"
	topUpBatchDescription  = "Postage batch top up"
	diluteBatchDescription = "Postage batch dilute"
)

type Interface interface {
	CreateBatch(ctx context.Context, initialBalance *big.Int, depth uint8, immutable bool, label string) (common.Hash, []byte, error)
	TopUpBatch(ctx context.Context, batchID []byte, topupBalance *big.Int) (common.Hash, error)
	DiluteBatch(ctx context.Context, batchID []byte, newDepth uint8) (common.Hash, error)
	Paused(ctx context.Context) (bool, error)
	MinimumValidityBlocks(ctx context.Context) (uint64, error)
	PostageBatchExpirer
}

type PostageBatchExpirer interface {
	ExpireBatches(ctx context.Context) error
}

type postageContract struct {
	owner                       common.Address
	postageStampContractAddress common.Address
	postageStampContractABI     abi.ABI
	bzzTokenAddress             common.Address
	transactionService          transaction.Service
	postageService              postage.Service
	postageStorer               postage.Storer

	// Cached postage stamp contract event topics.
	batchCreatedTopic       common.Hash
	batchTopUpTopic         common.Hash
	batchDepthIncreaseTopic common.Hash

	gasLimit uint64
}

func New(
	owner common.Address,
	postageStampContractAddress common.Address,
	postageStampContractABI abi.ABI,
	bzzTokenAddress common.Address,
	transactionService transaction.Service,
	postageService postage.Service,
	postageStorer postage.Storer,
	chainEnabled bool,
	gasLimit uint64,
) Interface {
	_ = "STUB: not implemented"
	return *new(Interface)
}

func (c *postageContract) ExpireBatches(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *postageContract) expiredBatchesExists(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *postageContract) expireLimitedBatches(ctx context.Context, count *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *postageContract) sendApproveTransaction(ctx context.Context, amount *big.Int) (receipt *types.Receipt, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *postageContract) sendTransaction(ctx context.Context, callData []byte, desc string) (receipt *types.Receipt, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *postageContract) sendCreateBatchTransaction(ctx context.Context, owner common.Address, initialBalance *big.Int, depth uint8, nonce common.Hash, immutable bool) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *postageContract) sendTopUpBatchTransaction(ctx context.Context, batchID []byte, topUpAmount *big.Int) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *postageContract) sendDiluteTransaction(ctx context.Context, batchID []byte, newDepth uint8) (*types.Receipt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *postageContract) getBalance(ctx context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *postageContract) getProperty(ctx context.Context, propertyName string, out any) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *postageContract) getMinInitialBalance(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *postageContract) CreateBatch(ctx context.Context, initialBalance *big.Int, depth uint8, immutable bool, label string) (txHash common.Hash, batchID []byte, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil, nil
}

func (c *postageContract) TopUpBatch(ctx context.Context, batchID []byte, topupBalance *big.Int) (txHash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (c *postageContract) DiluteBatch(ctx context.Context, batchID []byte, newDepth uint8) (txHash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (c *postageContract) Paused(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *postageContract) MinimumValidityBlocks(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type batchCreatedEvent struct {
	BatchId           [32]byte
	TotalAmount       *big.Int
	NormalisedBalance *big.Int
	Owner             common.Address
	Depth             uint8
	BucketDepth       uint8
	ImmutableFlag     bool
}

type noOpPostageContract struct{}

func (m *noOpPostageContract) CreateBatch(context.Context, *big.Int, uint8, bool, string) (common.Hash, []byte, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil, nil
}

func (m *noOpPostageContract) TopUpBatch(context.Context, []byte, *big.Int) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (m *noOpPostageContract) DiluteBatch(context.Context, []byte, uint8) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (m *noOpPostageContract) Paused(context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *noOpPostageContract) MinimumValidityBlocks(context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *noOpPostageContract) ExpireBatches(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func LookupERC20Address(ctx context.Context, transactionService transaction.Service, postageStampContractAddress common.Address, postageStampContractABI abi.ABI, chainEnabled bool) (common.Address, error) {
	_ = "STUB: not implemented"
	return *new(common.Address), nil
}
