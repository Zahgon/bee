// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package swap

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

var (
	peerPrefix            = "swap_chequebook_peer_"
	peerChequebookPrefix  = "swap_peer_chequebook_"
	beneficiaryPeerPrefix = "swap_beneficiary_peer_"
	peerBeneficiaryPrefix = "swap_peer_beneficiary_"
	deductedForPeerPrefix = "swap_deducted_for_peer_"
	deductedByPeerPrefix  = "swap_deducted_by_peer_"
)

// Addressbook maps peers to beneficaries, chequebooks and in reverse.
type Addressbook interface {
	// Beneficiary returns the beneficiary for the given peer.
	Beneficiary(peer swarm.Address) (beneficiary common.Address, known bool, err error)
	// Chequebook returns the chequebook for the given peer.
	Chequebook(peer swarm.Address) (chequebookAddress common.Address, known bool, err error)
	// BeneficiaryPeer returns the peer for a beneficiary.
	BeneficiaryPeer(beneficiary common.Address) (peer swarm.Address, known bool, err error)
	// ChequebookPeer returns the peer for a beneficiary.
	ChequebookPeer(chequebook common.Address) (peer swarm.Address, known bool, err error)
	// PutBeneficiary stores the beneficiary for the given peer.
	PutBeneficiary(peer swarm.Address, beneficiary common.Address) error
	// PutChequebook stores the chequebook for the given peer.
	PutChequebook(peer swarm.Address, chequebook common.Address) error
	// AddDeductionFor peer stores the flag indicating the peer have already issued a cheque that has been deducted
	AddDeductionFor(peer swarm.Address) error
	// AddDeductionBy peer stores the flag indicating the peer have already issued a cheque that has been deducted
	AddDeductionBy(peer swarm.Address) error
	// GetDeductionFor returns whether a peer have already issued a cheque that has been deducted
	GetDeductionFor(peer swarm.Address) (bool, error)
	// GetDeductionBy returns whether a peer have already received a cheque that has been deducted
	GetDeductionBy(peer swarm.Address) (bool, error)
	// MigratePeer returns whether a peer have already received a cheque that has been deducted
	MigratePeer(oldPeer, newPeer swarm.Address) error
}

type addressbook struct {
	store storage.StateStorer
}

// NewAddressbook creates a new addressbook using the store.
func NewAddressbook(store storage.StateStorer) Addressbook {
	_ = "STUB: not implemented"
	return *new(Addressbook)
}

func (a *addressbook) MigratePeer(oldPeer, newPeer swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// Beneficiary returns the beneficiary for the given peer.
func (a *addressbook) Beneficiary(peer swarm.Address) (beneficiary common.Address, known bool, err error) {
	_ = "STUB: not implemented"
	return *new(common.Address), false, nil
}

// BeneficiaryPeer returns the peer for a beneficiary.
func (a *addressbook) BeneficiaryPeer(beneficiary common.Address) (peer swarm.Address, known bool, err error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), false, nil
}

// Chequebook returns the chequebook for the given peer.
func (a *addressbook) Chequebook(peer swarm.Address) (chequebookAddress common.Address, known bool, err error) {
	_ = "STUB: not implemented"
	return *new(common.Address), false, nil
}

// ChequebookPeer returns the peer for a beneficiary.
func (a *addressbook) ChequebookPeer(chequebook common.Address) (peer swarm.Address, known bool, err error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), false, nil
}

// PutBeneficiary stores the beneficiary for the given peer.
func (a *addressbook) PutBeneficiary(peer swarm.Address, beneficiary common.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// PutChequebook stores the chequebook for the given peer.
func (a *addressbook) PutChequebook(peer swarm.Address, chequebook common.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *addressbook) AddDeductionFor(peer swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *addressbook) AddDeductionBy(peer swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *addressbook) GetDeductionFor(peer swarm.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (a *addressbook) GetDeductionBy(peer swarm.Address) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// peerKey computes the key where to store the chequebook from a peer.
func peerKey(peer swarm.Address) string { _ = "STUB: not implemented"; return "" }

// chequebookPeerKey computes the key where to store the peer for a chequebook.
func chequebookPeerKey(chequebook common.Address) string { _ = "STUB: not implemented"; return "" }

// peerBeneficiaryKey computes the key where to store the beneficiary for a peer.
func peerBeneficiaryKey(peer swarm.Address) string { _ = "STUB: not implemented"; return "" }

// beneficiaryPeerKey computes the key where to store the peer for a beneficiary.
func beneficiaryPeerKey(peer common.Address) string { _ = "STUB: not implemented"; return "" }

func peerDeductedByKey(peer swarm.Address) string { _ = "STUB: not implemented"; return "" }

func peerDeductedForKey(peer swarm.Address) string { _ = "STUB: not implemented"; return "" }
