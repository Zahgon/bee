// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	optionNameDataDir                      = "data-dir"
	optionNameCacheCapacity                = "cache-capacity"
	optionNameDBOpenFilesLimit             = "db-open-files-limit"
	optionNameDBBlockCacheCapacity         = "db-block-cache-capacity"
	optionNameDBWriteBufferSize            = "db-write-buffer-size"
	optionNameDBDisableSeeksCompaction     = "db-disable-seeks-compaction"
	optionNamePassword                     = "password"
	optionNamePasswordFile                 = "password-file"
	optionNameAPIAddr                      = "api-addr"
	optionNameP2PAddr                      = "p2p-addr"
	optionNameNATAddr                      = "nat-addr"
	optionNameP2PWSEnable                  = "p2p-ws-enable"
	optionNameBootnodes                    = "bootnode"
	optionNameNetworkID                    = "network-id"
	optionWelcomeMessage                   = "welcome-message"
	optionCORSAllowedOrigins               = "cors-allowed-origins"
	optionNameTracingEnabled               = "tracing-enable"
	optionNameTracingEndpoint              = "tracing-endpoint"
	optionNameTracingHost                  = "tracing-host"
	optionNameTracingPort                  = "tracing-port"
	optionNameTracingServiceName           = "tracing-service-name"
	optionNameVerbosity                    = "verbosity"
	optionNamePaymentThreshold             = "payment-threshold"
	optionNamePaymentTolerance             = "payment-tolerance-percent"
	optionNamePaymentEarly                 = "payment-early-percent"
	optionNameResolverEndpoints            = "resolver-options"
	optionNameBootnodeMode                 = "bootnode-mode"
	optionNameSwapFactoryAddress           = "swap-factory-address"
	optionNameSwapInitialDeposit           = "swap-initial-deposit"
	optionNameSwapEnable                   = "swap-enable"
	optionNameChequebookEnable             = "chequebook-enable"
	optionNameFullNode                     = "full-node"
	optionNamePostageContractAddress       = "postage-stamp-address"
	optionNamePostageContractStartBlock    = "postage-stamp-start-block"
	optionNamePriceOracleAddress           = "price-oracle-address"
	optionNameRedistributionAddress        = "redistribution-address"
	optionNameStakingAddress               = "staking-address"
	optionNameBlockTime                    = "block-time"
	optionNameBlockSyncInterval            = "block-sync-interval"
	optionWarmUpTime                       = "warmup-time"
	optionNameMainNet                      = "mainnet"
	optionNameRetrievalCaching             = "cache-retrieval"
	optionNameResync                       = "resync"
	optionNamePProfBlock                   = "pprof-profile"
	optionNamePProfMutex                   = "pprof-mutex"
	optionNameStaticNodes                  = "static-nodes"
	optionNameAllowPrivateCIDRs            = "allow-private-cidrs"
	optionNameSleepAfter                   = "sleep-after"
	optionNameStorageIncentivesEnable      = "storage-incentives-enable"
	optionNameStateStoreCacheCapacity      = "statestore-cache-capacity"
	optionNameTargetNeighborhood           = "target-neighborhood"
	optionNameNeighborhoodSuggester        = "neighborhood-suggester"
	optionNameWhitelistedWithdrawalAddress = "withdrawal-addresses-whitelist"
	optionNameTransactionDebugMode         = "transaction-debug-mode"
	optionMinimumStorageRadius             = "minimum-storage-radius"
	optionReserveCapacityDoubling          = "reserve-capacity-doubling"
	optionSkipPostageSnapshot              = "skip-postage-snapshot"
	optionNameMinimumGasTipCap             = "minimum-gas-tip-cap"
	optionNameGasLimitFallback             = "gas-limit-fallback"
	optionNameP2PWSSEnable                 = "p2p-wss-enable"
	optionP2PWSSAddr                       = "p2p-wss-addr"
	optionNATWSSAddr                       = "nat-wss-addr"
	optionAutoTLSDomain                    = "autotls-domain"
	optionAutoTLSRegistrationEndpoint      = "autotls-registration-endpoint"
	optionAutoTLSCAEndpoint                = "autotls-ca-endpoint"
	optionUseSIMD                          = "use-simd-hashing"

	// blockchain-rpc
	optionNameBlockchainRpcEndpoint    = "blockchain-rpc-endpoint"
	optionNameBlockchainRpcDialTimeout = "blockchain-rpc-dial-timeout"
	optionNameBlockchainRpcTLSTimeout  = "blockchain-rpc-tls-timeout"
	optionNameBlockchainRpcIdleTimeout = "blockchain-rpc-idle-timeout"
	optionNameBlockchainRpcKeepalive   = "blockchain-rpc-keepalive"
	configKeyBlockchainRpcEndpoint     = "blockchain-rpc.endpoint"
	configKeyBlockchainRpcDialTimeout  = "blockchain-rpc.dial-timeout"
	configKeyBlockchainRpcTLSTimeout   = "blockchain-rpc.tls-timeout"
	configKeyBlockchainRpcIdleTimeout  = "blockchain-rpc.idle-timeout"
	configKeyBlockchainRpcKeepalive    = "blockchain-rpc.keepalive"
)

var blockchainRpcConfigPairs = []struct{ flat, dotted string }{
	{optionNameBlockchainRpcEndpoint, configKeyBlockchainRpcEndpoint},
	{optionNameBlockchainRpcDialTimeout, configKeyBlockchainRpcDialTimeout},
	{optionNameBlockchainRpcTLSTimeout, configKeyBlockchainRpcTLSTimeout},
	{optionNameBlockchainRpcIdleTimeout, configKeyBlockchainRpcIdleTimeout},
	{optionNameBlockchainRpcKeepalive, configKeyBlockchainRpcKeepalive},
}

var knownNestedKeys = func() map[string]bool {
	m := make(map[string]bool, len(blockchainRpcConfigPairs))
	for _, p := range blockchainRpcConfigPairs {
		m[p.dotted] = true
	}
	return m
}()

// nolint:gochecknoinits
func init() {
	cobra.EnableCommandSorting = false
}

type command struct {
	root             *cobra.Command
	config           *viper.Viper
	logger           log.Logger
	passwordReader   passwordReader
	cfgFile          string
	homeDir          string
	isWindowsService bool
}

type option func(*command)

func newCommand(opts ...option) (c *command, err error) { _ = "STUB: not implemented"; return nil, nil }

// Find home directory.

func (c *command) Execute() (err error) { _ = "STUB: not implemented"; return nil }

// Execute parses command line arguments and runs appropriate functions.
func Execute() (err error) { _ = "STUB: not implemented"; return nil }

func (c *command) initGlobalFlags() { _ = "STUB: not implemented"; return }

func (c *command) initCommandVariables() error { _ = "STUB: not implemented"; return nil }

func (c *command) initConfig() (err error) { _ = "STUB: not implemented"; return nil }

// Use config file from the flag.

// Search config in home directory with name ".bee" (without extension).

// Environment

// read in environment variables that match

// If a config file is found, read it in.

func (c *command) setHomeDir() (err error) { _ = "STUB: not implemented"; return nil }

func (c *command) setAllFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

// preRun must be called from every command's PreRunE, after which c.logger is
// ready for use in RunE. It binds CLI flags to viper and initialises the logger.
func (c *command) preRun(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

func (c *command) initLogger(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

// bindBlockchainRpcConfig supports both flat (blockchain-rpc-endpoint) and
// nested (blockchain-rpc.endpoint) YAML forms, with nested taking precedence.
func (c *command) bindBlockchainRpcConfig(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

// Check before registering the alias; afterwards the flat value is unreachable.

func newLogger(cmd *cobra.Command, verbosity string) (log.Logger, error) {
	_ = "STUB: not implemented"
	return *new(log.Logger), nil
}

// For backwards compatibility, just enable v1 debugging as trace.

func (c *command) CheckUnknownParams(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Only accept the dotted→hyphenated form for explicitly registered
// nested config keys.
