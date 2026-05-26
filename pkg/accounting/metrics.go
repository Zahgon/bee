// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package accounting

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	// all metrics fields must be exported
	// to be able to return them by Metrics()
	// using reflection
	TotalDebitedAmount                       prometheus.Counter
	TotalCreditedAmount                      prometheus.Counter
	DebitEventsCount                         prometheus.Counter
	CreditEventsCount                        prometheus.Counter
	AccountingDisconnectsEnforceRefreshCount prometheus.Counter
	AccountingRefreshAttemptCount            prometheus.Counter
	AccountingNonFatalRefreshFailCount       prometheus.Counter
	AccountingDisconnectsOverdrawCount       prometheus.Counter
	AccountingDisconnectsGhostOverdrawCount  prometheus.Counter
	AccountingDisconnectsReconnectCount      prometheus.Counter
	AccountingBlocksCount                    prometheus.Counter
	AccountingReserveCount                   prometheus.Counter
	TotalOriginatedCreditedAmount            prometheus.Counter
	OriginatedCreditEventsCount              prometheus.Counter
	SettleErrorCount                         prometheus.Counter
	PaymentAttemptCount                      prometheus.Counter
	PaymentErrorCount                        prometheus.Counter
	ErrTimeOutOfSyncAlleged                  prometheus.Counter
	ErrTimeOutOfSyncRecent                   prometheus.Counter
	ErrTimeOutOfSyncInterval                 prometheus.Counter
	ErrRefreshmentBelowExpected              prometheus.Counter
	ErrRefreshmentAboveExpected              prometheus.Counter
}

func newMetrics() metrics { _ = "STUB: not implemented"; return *new(metrics) }

// Metrics returns the prometheus Collector for the accounting service.
func (a *Accounting) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }
