// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package migration

import (
	"errors"
	"fmt"

	storage "github.com/ethersphere/bee/v2/pkg/storage"
)

var ErrItemIDShouldntChange = errors.New("item.ID shouldn't be changing after update")

type (
	// ItemDeleteFn is callback function called in migration step
	// to check if Item should be removed in this step.
	ItemDeleteFn func(storage.Item) (deleted bool)

	// ItemUpdateFn is callback function called in migration step
	// to check if Item should be updated in this step.
	ItemUpdateFn func(storage.Item) (updatedItem storage.Item, hasChanged bool)
)

// WithItemDeleteFn return option with ItemDeleteFn set.
func WithItemDeleteFn(fn ItemDeleteFn) option { _ = "STUB: not implemented"; return *new(option) }

// WithItemUpdaterFn return option with ItemUpdateFn set.
func WithItemUpdaterFn(fn ItemUpdateFn) option { _ = "STUB: not implemented"; return *new(option) }

type option func(*options)

type options struct {
	deleteFn   ItemDeleteFn
	updateFn   ItemUpdateFn
	opPerBatch int
}

func defaultOptions() *options { _ = "STUB: not implemented"; return nil }

func (o *options) applyAll(opts []option) { _ = "STUB: not implemented"; return }

// NewStepOnIndex creates new migration step with update and/or delete operation.
// Migration will iterate on all elements selected by query and delete or update items
// based on supplied callback functions.
func NewStepOnIndex(s storage.BatchStore, query storage.Query, opts ...option) StepFn {
	_ = "STUB: not implemented"
	return *new(StepFn)
}

func stepOnIndex(s storage.Store, query storage.Query, o *options) error {
	_ = "STUB: not implemented"
	return nil
}

func deleteAll(s storage.Store, items []storage.Item) error { _ = "STUB: not implemented"; return nil }

func putAll(s storage.Store, items []storage.Item) error { _ = "STUB: not implemented"; return nil }

type key struct {
	storage.Marshaler
	storage.Unmarshaler
	storage.Cloner
	fmt.Stringer

	id        string
	namespace string
}

func newKey(k storage.Key) *key { _ = "STUB: not implemented"; return nil }

func (k *key) ID() string        { _ = "STUB: not implemented"; return "" }
func (k *key) Namespace() string { _ = "STUB: not implemented"; return "" }
