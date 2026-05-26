// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package accounting provides functionalities needed
// to do per-peer accounting.
package accounting

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/pricing"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "accounting"

const (
	linearCheckpointNumber = 1800
	linearCheckpointStep   = 100
)

var (
	_                        Interface = (*Accounting)(nil)
	balancesPrefix                     = "accounting_balance_"
	balancesSurplusPrefix              = "accounting_surplusbalance_"
	balancesOriginatedPrefix           = "accounting_originatedbalance_"
	// fraction of the refresh rate that is the minimum for monetary settlement
	// this value is chosen so that tiny payments are prevented while still allowing small payments in environments with lower payment thresholds
	minimumPaymentDivisor    = int64(5)
	failedSettlementInterval = int64(10) // seconds
)

// Interface is the Accounting interface.
type Interface interface {
	// PrepareCredit action to prevent overspending in case of concurrent requests.
	PrepareCredit(ctx context.Context, peer swarm.Address, price uint64, originated bool) (Action, error)
	// PrepareDebit returns an accounting Action for the later debit to be executed on and to implement shadowing a possibly credited part of reserve on the other side.
	PrepareDebit(ctx context.Context, peer swarm.Address, price uint64) (Action, error)
	// Balance returns the current balance for the given peer.
	Balance(peer swarm.Address) (*big.Int, error)
	// SurplusBalance returns the current surplus balance for the given peer.
	SurplusBalance(peer swarm.Address) (*big.Int, error)
	// Balances returns balances for all known peers.
	Balances() (map[string]*big.Int, error)
	// CompensatedBalance returns the current balance deducted by current surplus balance for the given peer.
	CompensatedBalance(peer swarm.Address) (*big.Int, error)
	// CompensatedBalances returns the compensated balances for all known peers.
	CompensatedBalances() (map[string]*big.Int, error)
	// PeerAccounting returns the associated values for all known peers
	PeerAccounting() (map[string]PeerInfo, error)
}

// Action represents an accounting action that can be applied
type Action interface {
	// Cleanup cleans up an action. Must be called whether it was applied or not.
	Cleanup()
	// Apply applies an action
	Apply() error
}

// debitAction represents a future debit
type debitAction struct {
	accounting     *Accounting
	price          *big.Int
	peer           swarm.Address
	accountingPeer *accountingPeer
	applied        bool
}

// creditAction represents a future debit
type creditAction struct {
	accounting     *Accounting
	price          *big.Int
	peer           swarm.Address
	accountingPeer *accountingPeer
	originated     bool
	applied        bool
}

// PayFunc is the function used for async monetary settlement
type PayFunc func(context.Context, swarm.Address, *big.Int)

// RefreshFunc is the function used for sync time-based settlement
type RefreshFunc func(context.Context, swarm.Address, *big.Int)

// Mutex is a drop in replacement for the sync.Mutex
// it will not lock if the context is expired
type Mutex struct {
	mu chan struct{}
}

func NewMutex() *Mutex { _ = "STUB: not implemented"; return nil }

// unlocked by default

var ErrFailToLock = errors.New("failed to lock")

func (m *Mutex) TryLock(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// locked

func (m *Mutex) Lock() { _ = "STUB: not implemented"; return }

func (m *Mutex) Unlock() {
	_ = "STUB: not implemented"

	// accountingPeer holds all in-memory accounting information for one peer.
	return
}

type accountingPeer struct {
	lock                           *Mutex   // lock to be held during any accounting action for this peer
	reservedBalance                *big.Int // amount currently reserved for active peer interaction
	shadowReservedBalance          *big.Int // amount potentially to be debited for active peer interaction
	refreshReservedBalance         *big.Int // amount debt potentially decreased during an ongoing refreshment
	ghostBalance                   *big.Int // amount potentially could have been debited for but was not
	paymentThreshold               *big.Int // the threshold at which the peer expects us to pay
	earlyPayment                   *big.Int // individual early payment threshold calculated from payment threshold and early payment percentage
	paymentThresholdForPeer        *big.Int // individual payment threshold at which the peer is expected to pay
	disconnectLimit                *big.Int // individual disconnect threshold calculated from tolerance and payment threshold for peer
	refreshTimestampMilliseconds   int64    // last time we attempted and succeeded time-based settlement
	refreshReceivedTimestamp       int64    // last time we accepted time-based settlement
	paymentOngoing                 bool     // indicate if we are currently settling with the peer
	refreshOngoing                 bool     // indicates if we are currently refreshing with the peer
	lastSettlementFailureTimestamp int64    // time of last unsuccessful attempt to issue a cheque
	connected                      bool     // indicates whether the peer is currently connected
	fullNode                       bool     // the peer connected as full node or light node
	totalDebtRepay                 *big.Int // since being connected, amount of cumulative debt settled by the peer
	thresholdGrowAt                *big.Int // cumulative debt to be settled by the peer in order to give threshold upgrade
}

// Accounting is the main implementation of the accounting interface.
type Accounting struct {
	// Mutex for accessing the accountingPeers map.
	accountingPeersMu sync.Mutex
	accountingPeers   map[string]*accountingPeer
	logger            log.Logger
	store             storage.StateStorer
	// The payment threshold in BZZ we communicate to our peers.
	paymentThreshold *big.Int
	// The amount in percent we let peers exceed the payment threshold before we
	// disconnect them.
	paymentTolerance int64
	// Start settling when reserve plus debt reaches this close to threshold in percent.
	earlyPayment int64
	// Limit to disconnect peer after going in debt over
	disconnectLimit *big.Int
	// function used for monetary settlement
	payFunction PayFunc
	// function used for time settlement
	refreshFunction RefreshFunc
	// allowance based on time used in pseudo settle
	refreshRate      *big.Int
	lightRefreshRate *big.Int
	// lower bound for the value of issued cheques
	minimumPayment      *big.Int
	pricing             pricing.Interface
	metrics             metrics
	wg                  sync.WaitGroup
	p2p                 p2p.Service
	timeNow             func() time.Time
	thresholdGrowStep   *big.Int
	thresholdGrowChange *big.Int
	// light node counterparts
	lightPaymentThreshold    *big.Int
	lightDisconnectLimit     *big.Int
	lightThresholdGrowStep   *big.Int
	lightThresholdGrowChange *big.Int
}

var (
	// ErrOverdraft denotes the expected debt in Reserve would exceed the payment thresholds.
	ErrOverdraft = errors.New("attempted overdraft")
	// ErrDisconnectThresholdExceeded denotes a peer has exceeded the disconnect threshold.
	ErrDisconnectThresholdExceeded = errors.New("disconnect threshold exceeded")
	// ErrPeerNoBalance is the error returned if no balance in store exists for a peer
	ErrPeerNoBalance = errors.New("no balance for peer")
	// ErrInvalidValue denotes an invalid value read from store
	ErrInvalidValue = errors.New("invalid value")
	// ErrOverRelease
	ErrOverRelease = errors.New("attempting to release more balance than was reserved for peer")
	// ErrEnforceRefresh
	ErrEnforceRefresh = errors.New("allowance expectation refused")
)

// NewAccounting creates a new Accounting instance with the provided options.
func NewAccounting(
	PaymentThreshold *big.Int,
	PaymentTolerance,
	EarlyPayment int64,
	logger log.Logger,
	Store storage.StateStorer,
	Pricing pricing.Interface,
	refreshRate *big.Int,
	lightFactor int64,
	p2pService p2p.Service,
) (*Accounting, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *Accounting) getIncreasedExpectedDebt(peer swarm.Address, accountingPeer *accountingPeer, bigPrice *big.Int) (*big.Int, *big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// debt if all reserved operations are successfully credited excluding debt created by surplus balance

// additionalDebt is debt created by incoming payments which we don't consider debt for monetary settlement purposes

// debt if all reserved operations are successfully credited including debt created by surplus balance

func (a *Accounting) PrepareCredit(ctx context.Context, peer swarm.Address, price uint64, originated bool) (Action, error) {
	_ = "STUB: not implemented"
	return *new(Action), nil
}

// debt if all reserved operations are successfully credited including debt created by surplus balance

// debt if all reserved operations are successfully credited and all shadow reserved operations are debited including debt created by surplus balance
// in other words this the debt the other node sees if everything pending is successful

// If our expected debt reduced by what could have been credited on the other side already is less than earlyPayment away from our payment threshold
// and we are actually in debt, trigger settlement.
// we pay early to avoid needlessly blocking request later when concurrent requests occur and we are already close to the payment threshold.

// if expectedDebt would still exceed the paymentThreshold at this point block this request
// this can happen if there is a large number of concurrent requests to the same peer

func (c *creditAction) Apply() error { _ = "STUB: not implemented"; return nil }

// debt if all reserved operations are successfully credited including debt created by surplus balance

// Calculate next balance by decreasing current balance with the price we credit

// debt if all reserved operations are successfully credited and all shadow reserved operations are debited including debt created by surplus balance
// in other words this the debt the other node sees if everything pending is successful

// Calculate next balance by decreasing current balance with the price we credit

// only consider negative balance for limiting originated balance

// If originated balance is more into the negative domain, set it to balance

// debt if all reserved operations are successfully credited and all shadow reserved operations are debited including debt created by surplus balance
// in other words this the debt the other node sees if everything pending is successful

func (c *creditAction) Cleanup() { _ = "STUB: not implemented"; return }

// Settle all debt with a peer. The lock on the accountingPeer must be held when
// called.
func (a *Accounting) settle(peer swarm.Address, balance *accountingPeer) error {
	_ = "STUB: not implemented"
	return nil
}

// get debt towards peer decreased by any amount that is to be debited soon

// Don't do anything if there is not enough actual debt
// This might be the case if the peer owes us and the total reserve for a peer exceeds the payment threshold.
// Minimum amount to trigger settlement for is 1 * refresh rate to avoid ineffective use of refreshments

// Only trigger refreshment if last refreshment finished at least 1000 milliseconds ago
// This is to avoid a peer refusing refreshment because not enough time passed since last refreshment

// if a settlement failed recently, wait until failedSettlementInterval before trying again

// if there is no monetary settlement happening, check if there is something to settle
// compute debt excluding debt created by incoming payments

// if the remaining debt is still larger than some minimum amount, trigger monetary settlement

// add settled amount to shadow reserve before sending it

// if a refreshment is ongoing, add this amount sent to cumulative potential debt decrease during refreshment

// Balance returns the current balance for the given peer.
func (a *Accounting) Balance(peer swarm.Address) (balance *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OriginatedBalance returns the current balance for the given peer.
func (a *Accounting) OriginatedBalance(peer swarm.Address) (balance *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SurplusBalance returns the current balance for the given peer.
func (a *Accounting) SurplusBalance(peer swarm.Address) (balance *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CompensatedBalance returns balance decreased by surplus balance
func (a *Accounting) CompensatedBalance(peer swarm.Address) (compensated *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if surplus is 0 and peer has no balance, propagate ErrPeerNoBalance

// Compensated balance is balance decreased by surplus balance

// peerBalanceKey returns the balance storage key for the given peer.
func peerBalanceKey(peer swarm.Address) string { _ = "STUB: not implemented"; return "" }

// peerSurplusBalanceKey returns the surplus balance storage key for the given peer
func peerSurplusBalanceKey(peer swarm.Address) string { _ = "STUB: not implemented"; return "" }

func originatedBalanceKey(peer swarm.Address) string { _ = "STUB: not implemented"; return "" }

// getAccountingPeer returns the accountingPeer for a given swarm address.
// If not found in memory it will initialize it.
func (a *Accounting) getAccountingPeer(peer swarm.Address) *accountingPeer {
	_ = "STUB: not implemented"
	return nil
}

// initially assume the peer has the same threshold as us

// notifyPaymentThresholdUpgrade is used when cumulative debt settled by peer reaches current checkpoint,
// to set the next checkpoint and increase the payment threshold given by 1 * refreshment rate
// must be called under accountingPeer lock
func (a *Accounting) notifyPaymentThresholdUpgrade(peer swarm.Address, accountingPeer *accountingPeer) {
	_ = "STUB: not implemented"
	// get appropriate linear growth limit based on whether the peer is a full node or a light node
	return
}

// if current checkpoint already passed linear growth limit, set next checkpoint exponentially

// otherwise set next linear checkpoint

// get appropriate refresh rate

// increase given threshold by refresh rate

// recalculate disconnectLimit for peer

// announce new payment threshold to peer

// Balances gets balances for all peers from store.
func (a *Accounting) Balances() (map[string]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PeerInfo struct {
	Balance                  *big.Int
	ConsumedBalance          *big.Int
	ThresholdReceived        *big.Int
	ThresholdGiven           *big.Int
	CurrentThresholdReceived *big.Int
	CurrentThresholdGiven    *big.Int
	SurplusBalance           *big.Int
	ReservedBalance          *big.Int
	ShadowReservedBalance    *big.Int
	GhostBalance             *big.Int
}

func (a *Accounting) PeerAccounting() (map[string]PeerInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get appropriate refresh rate

// get appropriate refresh rate

// CompensatedBalances gets balances for all peers from store.
func (a *Accounting) CompensatedBalances() (map[string]*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// balanceKeyPeer returns the embedded peer from the balance storage key.
func balanceKeyPeer(key []byte) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

func surplusBalanceKeyPeer(key []byte) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// PeerDebt returns the positive part of the sum of the outstanding balance and the shadow reserve
func (a *Accounting) PeerDebt(peer swarm.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// peerLatentDebt returns the sum of the positive part of the outstanding balance, shadow reserve and the ghost balance
func (a *Accounting) peerLatentDebt(peer swarm.Address) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// shadowBalance returns the current debt reduced by any potentially debitable amount stored in shadowReservedBalance
// this represents how much less our debt could potentially be seen by the other party if it's ahead with processing credits corresponding to our shadow reserve
func (a *Accounting) shadowBalance(peer swarm.Address, accountingPeer *accountingPeer) (shadowBalance *big.Int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NotifyPaymentSent is triggered by async monetary settlement to update our balance and remove it's price from the shadow reserve
func (a *Accounting) NotifyPaymentSent(peer swarm.Address, amount *big.Int, receivedError error) {
	_ = "STUB: not implemented"
	return
}

// decrease shadow reserve by payment value

// Get nextBalance by increasing current balance with price

// NotifyPaymentThreshold should be called to notify accounting of changes in the payment threshold
func (a *Accounting) NotifyPaymentThreshold(peer swarm.Address, paymentThreshold *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// NotifyPaymentReceived is called by Settlement when we receive a payment.
func (a *Accounting) NotifyPaymentReceived(peer swarm.Address, amount *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// if balance is already negative or zero, we credit full amount received to surplus balance and terminate early

// if current balance is positive, let's make a partial credit to

// Don't allow a payment to put us into debt
// This is to prevent another node tricking us into settling by settling
// first (e.g. send a bouncing cheque to trigger an honest cheque in swap).

// If payment would have put us into debt, rather, let's add to surplusBalance,
// so as that an oversettlement attempt creates balance for future forwarding services
// charges to be deducted of

// NotifyRefreshmentSent is called by pseudosettle when refreshment is done or failed
func (a *Accounting) NotifyRefreshmentSent(peer swarm.Address, attemptedAmount, amount *big.Int, timestamp int64, allegedInterval int64, receivedError error) {
	_ = "STUB: not implemented"
	return
}

// conclude ongoing refreshment

// save timestamp received in milliseconds of when the refreshment completed locally

// if specific error is received increment metrics

// if refreshment failed with connected peer, blocklist

// enforce allowance
// calculate expectation decreased by any potential debt decreases occurred during the refreshment

// reset cumulative potential debt decrease during an ongoing refreshment as refreshment just completed

// dont expect higher amount accepted than attempted (sanity check)

// calculate time based allowance

// expect minimum of time based allowance and debt / attempted amount based expectation

// compare received refreshment amount to expectation

// if expectation is not met, blocklist peer

// update balance

// update originated balance

// NotifyRefreshmentReceived is called by pseudosettle when we receive a time based settlement.
func (a *Accounting) NotifyRefreshmentReceived(peer swarm.Address, amount *big.Int, timestamp int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Get nextBalance by increasing current balance with amount

// We allow a refreshment to potentially put us into debt as it was previously negotiated and be limited to the peer's outstanding debt plus shadow reserve

// PrepareDebit prepares a debit operation by increasing the shadowReservedBalance
func (a *Accounting) PrepareDebit(ctx context.Context, peer swarm.Address, price uint64) (Action, error) {
	_ = "STUB: not implemented"
	return *new(Action), nil
}

// if a refreshment is ongoing, add this amount to the potential debt decrease during an ongoing refreshment

func (a *Accounting) increaseBalance(peer swarm.Address, _ *accountingPeer, price *big.Int) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// see if peer has surplus balance to deduct this transaction of

// get new surplus balance after deduct

// if nothing left for debiting, store new surplus balance and return from debit

// if surplus balance didn't cover full transaction, let's continue with leftover part as cost

// a sanity check

// if we still have something to debit, than have run out of surplus balance,
// let's store 0 as surplus balance

// Get nextBalance by increasing current balance with price

// Apply applies the debit operation and decreases the shadowReservedBalance
func (d *debitAction) Apply() error { _ = "STUB: not implemented"; return nil }

// get appropriate refresh rate

// peer too much in debt

// Cleanup reduces shadow reserve if and only if debitaction have not been applied
func (d *debitAction) Cleanup() { _ = "STUB: not implemented"; return }

func (a *Accounting) blocklistUntil(peer swarm.Address, multiplier int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (a *Accounting) blocklist(peer swarm.Address, multiplier int64, reason string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Accounting) Connect(peer swarm.Address, fullNode bool) { _ = "STUB: not implemented"; return }

// decreaseOriginatedBalanceTo decreases the originated balance to provided limit or 0 if limit is positive
func (a *Accounting) decreaseOriginatedBalanceTo(peer swarm.Address, limit *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// If originated balance is more into the negative domain, set it to limit

// decreaseOriginatedBalanceBy decreases the originated balance by provided amount even below 0
func (a *Accounting) decreaseOriginatedBalanceBy(peer swarm.Address, amount *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// Move originated balance into the positive domain by amount

func (a *Accounting) Disconnect(peer swarm.Address) { _ = "STUB: not implemented"; return }

func (a *Accounting) SetRefreshFunc(f RefreshFunc) { _ = "STUB: not implemented"; return }

func (a *Accounting) SetPayFunc(f PayFunc) {
	_ = "STUB: not implemented"

	// Close hangs up running websockets on shutdown.
	return
}

func (a *Accounting) Close() error { _ = "STUB: not implemented"; return nil }

func percentOf(percent int64, of *big.Int) *big.Int { _ = "STUB: not implemented"; return nil }
