// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"github.com/ethersphere/bee/v2/pkg/feeds"
	storage "github.com/ethersphere/bee/v2/pkg/storage"
)

type factory struct {
	storage.Getter
}

func New(getter storage.Getter) feeds.Factory {
	_ = "STUB: not implemented"
	return *new(feeds.Factory)
}

func (f *factory) NewLookup(t feeds.Type, feed *feeds.Feed) (feeds.Lookup, error) {
	_ = "STUB: not implemented"
	return *new(feeds.Lookup), nil
}
