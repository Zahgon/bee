// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package node

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p/libp2p"
	"github.com/ethersphere/bee/v2/pkg/settlement"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/chequebook"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/erc20"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/priceoracle"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/transaction"
)

const (
	maxDelay                = 1 * time.Minute
	cancellationDepth       = 12
	additionalConfirmations = 2
)

// BlockchainRPCConfig holds the configuration parameters for the blockchain RPC client transport.
type BlockchainRPCConfig struct {
	Endpoint    string
	DialTimeout time.Duration
	TLSTimeout  time.Duration
	IdleTimeout time.Duration
	Keepalive   time.Duration
}

// InitChain will initialize the Ethereum backend at the given endpoint and
// set up the Transaction Service to interact with it using the provided signer.
func InitChain(
	ctx context.Context,
	logger log.Logger,
	stateStore storage.StateStorer,
	chainID int64,
	signer crypto.Signer,
	pollingInterval time.Duration,
	chainEnabled bool,
	minimumGasTipCap uint64,
	fallbackGasLimit uint64,
	rpcCfg BlockchainRPCConfig,
	blockSyncInterval uint64,
) (transaction.Backend, common.Address, int64, transaction.Monitor, transaction.Service, error) {
	_ = "STUB: not implemented"
	return *new(transaction.Backend), *new(common.Address), 0, *new(transaction.Monitor), *new(transaction.Service), nil
}

// InitChequebookFactory will initialize the chequebook factory with the given
// chain backend.
func InitChequebookFactory(logger log.Logger, backend transaction.Backend, chainID int64, transactionService transaction.Service, factoryAddress string) (chequebook.Factory, error) {
	_ = "STUB: not implemented"
	return *new(chequebook.Factory), nil
}

// InitChequebookService will initialize the chequebook service with the given
// chequebook factory and chain backend.
func InitChequebookService(
	ctx context.Context,
	logger log.Logger,
	stateStore storage.StateStorer,
	signer crypto.Signer,
	chainID int64,
	backend transaction.Backend,
	overlayEthAddress common.Address,
	transactionService transaction.Service,
	chequebookFactory chequebook.Factory,
	initialDeposit string,
	erc20Service erc20.Service,
) (chequebook.Service, error) {
	_ = "STUB: not implemented"
	return *new(chequebook.Service), nil
}

func initChequeStoreCashout(
	stateStore storage.StateStorer,
	swapBackend transaction.Backend,
	chequebookFactory chequebook.Factory,
	chainID int64,
	overlayEthAddress common.Address,
	transactionService transaction.Service,
) (chequebook.ChequeStore, chequebook.CashoutService) {
	_ = "STUB: not implemented"
	return *new(chequebook.ChequeStore), *new(chequebook.CashoutService)
}

// InitSwap will initialize and register the swap service.
func InitSwap(
	p2ps *libp2p.Service,
	logger log.Logger,
	stateStore storage.StateStorer,
	networkID uint64,
	overlayEthAddress common.Address,
	chequebookService chequebook.Service,
	chequeStore chequebook.ChequeStore,
	cashoutService chequebook.CashoutService,
	accounting settlement.Accounting,
	priceOracleAddress string,
	chainID int64,
	transactionService transaction.Service,
) (*swap.Service, priceoracle.Service, error) {
	_ = "STUB: not implemented"
	return nil, *new(priceoracle.Service), nil
}

func GetTxHash(stateStore storage.StateStorer, logger log.Logger, trxString string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetTxNextBlock(ctx context.Context, logger log.Logger, backend transaction.Backend, monitor transaction.Monitor, duration time.Duration, trx []byte, blockHash string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// noOpChequebookService is a noOp implementation for chequebook.Service interface.
type noOpChequebookService struct{}

func (m *noOpChequebookService) Deposit(context.Context, *big.Int) (hash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (m *noOpChequebookService) Withdraw(context.Context, *big.Int) (hash common.Hash, err error) {
	_ = "STUB: not implemented"
	return *new(common.Hash), nil
}

func (m *noOpChequebookService) WaitForDeposit(context.Context, common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *noOpChequebookService) Balance(context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *noOpChequebookService) AvailableBalance(context.Context) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *noOpChequebookService) Address() common.Address {
	_ = "STUB: not implemented"
	return *new(common.Address)
}

func (m *noOpChequebookService) Issue(context.Context, common.Address, *big.Int, chequebook.SendChequeFunc) (*big.Int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *noOpChequebookService) LastCheque(common.Address) (*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *noOpChequebookService) LastCheques() (map[common.Address]*chequebook.SignedCheque, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
