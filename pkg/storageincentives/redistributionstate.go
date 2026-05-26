// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storageincentives

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/erc20"
	"github.com/ethersphere/bee/v2/pkg/storage"
	storer "github.com/ethersphere/bee/v2/pkg/storer"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/transaction"
)

const loggerNameNode = "nodestatus"

const (
	redistributionStatusKey = "redistribution_state"
	purgeStaleDataThreshold = 10
)

type RedistributionState struct {
	mtx sync.Mutex

	stateStore     storage.StateStorer
	erc20Service   erc20.Service
	logger         log.Logger
	ethAddress     common.Address
	status         *Status
	currentBalance *big.Int
	txService      transaction.Service
}

// Status provide internal status of the nodes in the redistribution game
type Status struct {
	Phase             PhaseType
	IsFrozen          bool
	IsFullySynced     bool
	Round             uint64
	LastWonRound      uint64
	LastPlayedRound   uint64
	LastFrozenRound   uint64
	LastSelectedRound uint64
	Block             uint64
	Reward            *big.Int
	Fees              *big.Int
	RoundData         map[uint64]RoundData
	SampleDuration    time.Duration
	IsHealthy         bool
}

type RoundData struct {
	CommitKey   []byte
	SampleData  *SampleData
	HasRevealed bool
}

type SampleData struct {
	Anchor1            []byte
	ReserveSampleItems []storer.SampleItem
	ReserveSampleHash  swarm.Address
	StorageRadius      uint8
}

func NewStatus() *Status { _ = "STUB: not implemented"; return nil }

func NewRedistributionState(logger log.Logger, ethAddress common.Address, stateStore storage.StateStorer, erc20Service erc20.Service, contract transaction.Service) (*RedistributionState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Status returns the node status
func (r *RedistributionState) Status() (*Status, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *RedistributionState) save() { _ = "STUB: not implemented"; return }

func (r *RedistributionState) SetCurrentBlock(block uint64) { _ = "STUB: not implemented"; return }

func (r *RedistributionState) SetCurrentEvent(phase PhaseType, round uint64) {
	_ = "STUB: not implemented"
	return
}

func (r *RedistributionState) IsFrozen() bool { _ = "STUB: not implemented"; return false }

func (r *RedistributionState) SetFrozen(isFrozen bool, round uint64) {
	_ = "STUB: not implemented"
	return
}

// record fronzen round if not set already

func (r *RedistributionState) SetLastWonRound(round uint64) { _ = "STUB: not implemented"; return }

func (r *RedistributionState) IsFullySynced() bool { _ = "STUB: not implemented"; return false }

func (r *RedistributionState) SetFullySynced(isSynced bool) { _ = "STUB: not implemented"; return }

func (r *RedistributionState) SetLastPlayedRound(round uint64) { _ = "STUB: not implemented"; return }

func (r *RedistributionState) SetLastSelectedRound(round uint64) { _ = "STUB: not implemented"; return }

// AddFee sets the internal node status
func (r *RedistributionState) AddFee(ctx context.Context, txHash common.Hash) {
	_ = "STUB: not implemented"
	return
}

// CalculateWinnerReward calculates the reward for the winner
func (r *RedistributionState) CalculateWinnerReward(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *RedistributionState) SetBalance(ctx context.Context) error {
	_ = "STUB: not implemented"
	// get current balance
	return nil
}

func (r *RedistributionState) SampleData(round uint64) (SampleData, bool) {
	_ = "STUB: not implemented"
	return *new(SampleData), false
}

func (r *RedistributionState) SetSampleData(round uint64, sd SampleData, dur time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (r *RedistributionState) CommitKey(round uint64) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *RedistributionState) SetCommitKey(round uint64, commitKey []byte) {
	_ = "STUB: not implemented"
	return
}

func (r *RedistributionState) HasRevealed(round uint64) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *RedistributionState) SetHealthy(isHealthy bool) { _ = "STUB: not implemented"; return }

func (r *RedistributionState) IsHealthy() bool { _ = "STUB: not implemented"; return false }

func (r *RedistributionState) SetHasRevealed(round uint64) { _ = "STUB: not implemented"; return }

func (r *RedistributionState) currentRoundAndPhase() (uint64, PhaseType) {
	_ = "STUB: not implemented"
	return 0, *new(PhaseType)
}

func (r *RedistributionState) currentBlock() uint64 { _ = "STUB: not implemented"; return 0 }

func (r *RedistributionState) purgeStaleRoundData() { _ = "STUB: not implemented"; return }
