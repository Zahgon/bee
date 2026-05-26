// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonhttp

import (
	"net/http"
)

type MethodHandler map[string]http.Handler

func (h MethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// HandleMethods uses a corresponding Handler based on HTTP request method.
// If Handler is not found, a method not allowed HTTP response is returned
// with specified body and Content-Type header.
func HandleMethods(methods map[string]http.Handler, body string, contentType string, w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// true if not COR request

func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"

	// NewMaxBodyBytesHandler is an http middleware constructor that limits the
	// maximal number of bytes that can be read from the request body. When a body
	// is read, the error can be handled with a helper function HandleBodyReadError
	// in order to respond with Request Entity Too Large response.
	// See TestNewMaxBodyBytesHandler as an example.
	return
}

func NewMaxBodyBytesHandler(limit int64) func(http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

// HandleBodyReadError checks for particular errors and writes appropriate
// response accordingly. If no known error is found, no response is written and
// the function returns false.
func HandleBodyReadError(err error, w http.ResponseWriter) (responded bool) {
	_ = "STUB: not implemented"
	return false
}

// http.MaxBytesReader returns an unexported error,
// this is the only way to detect it
