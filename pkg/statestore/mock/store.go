// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mock

import (
	"sync"

	"github.com/ethersphere/bee/v2/pkg/storage"
)

var _ storage.StateStorer = (*store)(nil)

type store struct {
	store map[string][]byte
	mtx   sync.RWMutex
}

func NewStateStore() storage.StateStorer {
	_ = "STUB: not implemented"
	return *new(storage.StateStorer)
}

func (s *store) Get(key string, i any) (err error) { _ = "STUB: not implemented"; return nil }

func (s *store) Put(key string, i any) (err error) { _ = "STUB: not implemented"; return nil }

func (s *store) Delete(key string) (err error) { _ = "STUB: not implemented"; return nil }

func (s *store) Iterate(prefix string, iterFunc storage.StateIterFunc) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *store) Close() (err error) { _ = "STUB: not implemented"; return nil }
