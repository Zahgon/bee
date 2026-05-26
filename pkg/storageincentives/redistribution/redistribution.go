// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redistribution

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/transaction"
)

const (
	loggerName      = "redistributionContract"
	BoostTipPercent = 50
)

type Contract interface {
	ReserveSalt(context.Context) ([]byte, error)
	IsPlaying(context.Context, uint8) (bool, error)
	IsWinner(context.Context) (bool, error)
	Claim(context.Context, ChunkInclusionProofs) (common.Hash, error)
	Commit(context.Context, []byte, uint64) (common.Hash, error)
	Reveal(context.Context, uint8, []byte, []byte) (common.Hash, error)
}

type contract struct {
	overlay                   swarm.Address
	owner                     common.Address
	logger                    log.Logger
	txService                 transaction.Service
	incentivesContractAddress common.Address
	incentivesContractABI     abi.ABI
	gasLimit                  uint64
}

func New(
	overlay swarm.Address,
	owner common.Address,
	logger log.Logger,
	txService transaction.Service,
	incentivesContractAddress common.Address,
	incentivesContractABI abi.ABI,
	gasLimit uint64,
) Contract {
	_ = "STUB: not implemented"
	return *new(Contract)
}

// IsPlaying checks if the overlay is participating in the upcoming round.
func (c *contract) IsPlaying(ctx context.Context, depth uint8) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// IsWinner checks if the overlay is winner by sending a transaction to blockchain.
func (c *contract) IsWinner(ctx context.Context) (isWinner bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Claim sends a transaction to blockchain if a win is claimed.
func (c *contract) Claim(ctx context.Context, proofs ChunkInclusionProofs) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// Commit submits the obfusHash hash by sending a transaction to the blockchain.
func (c *contract) Commit(ctx context.Context, obfusHash []byte, round uint64) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// Reveal submits the storageDepth, reserveCommitmentHash and RandomNonce in a transaction to blockchain.
func (c *contract) Reveal(ctx context.Context, storageDepth uint8, reserveCommitmentHash []byte, RandomNonce []byte) (common.Hash, error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// ReserveSalt provides the current round anchor by transacting on the blockchain.
func (c *contract) ReserveSalt(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *contract) sendAndWait(ctx context.Context, request *transaction.TxRequest, boostPercent int) (txHash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

// callTx simulates a transaction based on tx request.
func (c *contract) callTx(ctx context.Context, callData []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
