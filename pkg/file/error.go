// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

// AbortError should be returned whenever a file operation is terminated
// before it has completed.
type AbortError struct {
	err error
}

// NewAbortError creates a new AbortError instance.
func NewAbortError(err error) error { _ = "STUB: not implemented"; return nil }

// Unwrap returns an underlying error.
func (e *AbortError) Unwrap() error {
	_ = "STUB: not implemented"

	// Error implements standard go error interface.
	return nil
}

func (e *AbortError) Error() string { _ = "STUB: not implemented"; return "" }

// HashError should be returned whenever a file operation is terminated
// before it has completed.
type HashError struct {
	err error
}

// NewHashError creates a new HashError instance.
func NewHashError(err error) error { _ = "STUB: not implemented"; return nil }

// Unwrap returns an underlying error.
func (e *HashError) Unwrap() error {
	_ = "STUB: not implemented"

	// Error implements standard go error interface.
	return nil
}

func (e *HashError) Error() string { _ = "STUB: not implemented"; return "" }
