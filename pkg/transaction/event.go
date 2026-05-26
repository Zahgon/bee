// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package transaction

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	ErrEventNotFound = errors.New("event not found")
	ErrNoTopic       = errors.New("no topic")
)

// ParseEvent will parse the specified abi event from the given log
func ParseEvent(a *abi.ABI, eventName string, c any, e types.Log) error {
	_ = "STUB: not implemented"
	return nil
}

// FindSingleEvent will find the first event of the given kind.
func FindSingleEvent(abi *abi.ABI, receipt *types.Receipt, contractAddress common.Address, event abi.Event, out any) error {
	_ = "STUB: not implemented"
	return nil
}
