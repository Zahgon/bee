// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethersphere/bee/v2/pkg/postage"
)

type optionFunc func(*mockPostage)

// Option is an option passed to a mock postage Service.
type Option interface {
	apply(*mockPostage)
}

func (f optionFunc) apply(r *mockPostage) {
	_ = "STUB: not implemented"

	// New creates a new mock postage service.
	return
}

func New(o ...Option) postage.Service { _ = "STUB: not implemented"; return *new(postage.Service) }

// WithAcceptAll sets the mock to return a new BatchIssuer on every
// call to GetStampIssuer.
func WithAcceptAll() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithIssuer(s *postage.StampIssuer) Option { _ = "STUB: not implemented"; return *new(Option) }

type mockPostage struct {
	issuersMap map[string]*postage.StampIssuer
	issuerLock sync.Mutex
	acceptAll  bool
}

func (m *mockPostage) HandleStampExpiry(ctx context.Context, id []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mockPostage) Add(s *postage.StampIssuer) error { _ = "STUB: not implemented"; return nil }

func (m *mockPostage) StampIssuers() []*postage.StampIssuer { _ = "STUB: not implemented"; return nil }

func (m *mockPostage) GetStampIssuer(id []byte) (*postage.StampIssuer, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (m *mockPostage) IssuerUsable(_ *postage.StampIssuer) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *mockPostage) HandleCreate(_ *postage.Batch, _ *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *mockPostage) HandleTopUp(_ []byte, _ *big.Int) { _ = "STUB: not implemented"; return }

func (m *mockPostage) HandleDepthIncrease(_ []byte, _ uint8) { _ = "STUB: not implemented"; return }

func (m *mockPostage) Close() error { _ = "STUB: not implemented"; return nil }

var _ postage.BatchExpiryHandler = (*mockPostage)(nil)
