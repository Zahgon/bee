// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storageincentives

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/postage/postagecontract"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/erc20"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/storageincentives/redistribution"
	"github.com/ethersphere/bee/v2/pkg/storageincentives/staking"
	"github.com/ethersphere/bee/v2/pkg/storer"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/transaction"
	"resenje.org/singleflight"
)

const loggerName = "storageincentives"

const (
	DefaultBlocksPerRound = 152
	DefaultBlocksPerPhase = DefaultBlocksPerRound / 4

	// min # of transactions our wallet should be able to cover
	minTxCountToCover = 15

	// average tx gas used by transactions issued from agent
	avgTxGas = 250_000
)

type ChainBackend interface {
	BlockNumber(context.Context) (uint64, error)
	HeaderByNumber(context.Context, *big.Int) (*types.Header, error)
	BalanceAt(ctx context.Context, address common.Address, block *big.Int) (*big.Int, error)
	SuggestedFeeAndTip(ctx context.Context, gasPrice *big.Int, boostPercent int) (*big.Int, *big.Int, error)
}

type Health interface {
	IsHealthy() bool
}

type Agent struct {
	logger                 log.Logger
	metrics                metrics
	backend                ChainBackend
	blocksPerRound         uint64
	contract               redistribution.Contract
	batchExpirer           postagecontract.PostageBatchExpirer
	redistributionStatuser staking.RedistributionStatuser
	store                  storer.Reserve
	fullSyncedFunc         func() bool
	overlay                swarm.Address
	quit                   chan struct{}
	wg                     sync.WaitGroup
	state                  *RedistributionState
	chainStateGetter       postage.ChainStateGetter
	commitLock             sync.Mutex
	health                 Health
	sampleFlight           singleflight.Group[string, sampleResult]
}

func New(overlay swarm.Address,
	ethAddress common.Address,
	backend ChainBackend,
	contract redistribution.Contract,
	batchExpirer postagecontract.PostageBatchExpirer,
	redistributionStatuser staking.RedistributionStatuser,
	store storer.Reserve,
	fullSyncedFunc func() bool,
	blockTime time.Duration,
	blocksPerRound,
	blocksPerPhase uint64,
	stateStore storage.StateStorer,
	chainStateGetter postage.ChainStateGetter,
	erc20Service erc20.Service,
	tranService transaction.Service,
	health Health,
	logger log.Logger,
) (*Agent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// start polls the current block number, calculates, and publishes only once the current phase.
// Each round is blocksPerRound long and is divided into three blocksPerPhase long phases: commit, reveal, claim.
// The sample phase is triggered upon entering the claim phase and may run until the end of the commit phase.
// If our neighborhood is selected to participate, a sample is created during the sample phase. In the commit phase,
// the sample is submitted, and in the reveal phase, the obfuscation key from the commit phase is submitted.
// Next, in the claim phase, we check if we've won, and the cycle repeats. The cycle must occur in the length of one round.
func (a *Agent) start(blockTime time.Duration, blocksPerRound, blocksPerPhase uint64) {
	_ = "STUB: not implemented"
	return
}

// Sample handled could potentially take long time, therefore it could overlap with commit
// phase of next round. When that case happens commit event needs to be triggered once more
// in order to handle commit phase with delay.

// [0, 37]
// [38, 75]

// [76, 151]

// write the current phase only once

// check if node is frozen starting from the next block

// manually invoke phaseCheck initially in order to set initial data asap

// optimization, we do not need to check the phase change at every new block

func (a *Agent) handleCommit(ctx context.Context, round uint64) error {
	_ = "STUB: not implemented"
	// commit event handler has to be guarded with lock to avoid
	// race conditions when handler is triggered again from sample phase
	return nil
}

// already committed on this round, phase is skipped

// the sample has to come from previous round to be able to commit it

// In absence of sample, phase is skipped

func (a *Agent) handleReveal(ctx context.Context, round uint64) error {
	_ = "STUB: not implemented"
	// reveal requires the commitKey from the same round
	return nil
}

// In absence of commitKey, phase is skipped

// reveal requires sample from previous round

// Sample must have been saved so far

func (a *Agent) handleClaim(ctx context.Context, round uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// When there was no reveal in same round, phase is skipped

// When there is nothing to claim (node is not a winner), phase is played

// In case when there are too many expired batches, Claim trx could runs out of gas.
// To prevent this, node should first expire batches before Claiming a reward.

// Even when error happens, proceed with claim handler
// because this should not prevent node from claiming a reward

func (a *Agent) handleSample(ctx context.Context, round uint64) (bool, error) {
	_ = "STUB: not implemented"
	// minimum proximity between the anchor and the stored chunks
	return false, nil
}

type sampleResult struct {
	Items []storer.SampleItem
	Hash  swarm.Address
}

// reserveSampleAndHash runs getPreviousRoundTime, ReserveSample, and sampleHash
// as a singleflight keyed by anchor and depth to deduplicate concurrent calls.
func (a *Agent) reserveSampleAndHash(ctx context.Context, anchor []byte, depth uint8) (sampleResult, error) {
	_ = "STUB: not implemented"
	return *new(sampleResult), nil
}

func (a *Agent) makeSample(ctx context.Context, committedDepth uint8) (SampleData, error) {
	_ = "STUB: not implemented"
	return *new(SampleData), nil
}

func (a *Agent) minBatchBalance() *big.Int { _ = "STUB: not implemented"; return nil }

func (a *Agent) getPreviousRoundTime(ctx context.Context) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (a *Agent) commit(ctx context.Context, sample SampleData, round uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Agent) Close() error { _ = "STUB: not implemented"; return nil }

func (a *Agent) wrapCommit(storageRadius uint8, sample []byte, key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Status returns the node status
func (a *Agent) Status() (*Status, error) { _ = "STUB: not implemented"; return nil, nil }

type SampleWithProofs struct {
	Hash     swarm.Address                       `json:"hash"`
	Proofs   redistribution.ChunkInclusionProofs `json:"proofs"`
	Duration time.Duration                       `json:"duration"`
}

// SampleWithProofs is called only by rchash API
func (a *Agent) SampleWithProofs(
	ctx context.Context,
	anchor1 []byte,
	anchor2 []byte,
	storageRadius uint8,
) (SampleWithProofs, error) {
	_ = "STUB: not implemented"
	return *new(SampleWithProofs), nil
}

func (a *Agent) HasEnoughFundsToPlay(ctx context.Context) (*big.Int, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}
