// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapped

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	TotalRPCCalls  prometheus.Counter
	TotalRPCErrors prometheus.Counter

	TransactionReceiptCalls       prometheus.Counter
	TransactionCalls              prometheus.Counter
	BlockHeaderAsBlockNumberCalls prometheus.Counter
	BlockHeaderCalls              prometheus.Counter
	BalanceCalls                  prometheus.Counter
	NonceAtCalls                  prometheus.Counter
	PendingNonceCalls             prometheus.Counter
	CallContractCalls             prometheus.Counter
	SuggestGasTipCapCalls         prometheus.Counter
	EstimateGasCalls              prometheus.Counter
	SendTransactionCalls          prometheus.Counter
	FilterLogsCalls               prometheus.Counter
	ChainIDCalls                  prometheus.Counter
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

func (b *wrappedBackend) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
