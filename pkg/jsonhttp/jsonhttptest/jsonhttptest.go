// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonhttptest

import (
	"context"
	"io"
	"net/http"
	"testing"
)

// Request is a testing helper function that makes an HTTP request using
// provided client with provided method and url. It performs a validation on
// expected response code and additional options. It returns response headers if
// the request and all validation are successful. In case of any error, testing
// Errorf or Fatal functions will be called.
func Request(tb testing.TB, client *http.Client, method, url string, responseCode int, opts ...Option) http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

// When "Content-Length" header is set additionally assert
// that resp.ContentLength has the same value.

// WithContext sets a context to the request made by the Request function.
func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRequestBody writes a request body to the request made by the Request
// function.
func WithRequestBody(body io.Reader) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithJSONRequestBody writes a request JSON-encoded body to the request made by
// the Request function.
func WithJSONRequestBody(r any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMultipartRequest writes a multipart request with a single file in it to
// the request made by the Request function.
func WithMultipartRequest(body io.Reader, length int, filename, contentType string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRequestHeader adds a single header to the request made by the Request
// function. To add multiple headers call multiple times this option when as
// arguments to the Request function.
func WithRequestHeader(key, value string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExpectedResponse validates that the response from the request in the
// Request function matches completely bytes provided here.
func WithExpectedResponse(response []byte) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExpectedResponseHeader validates that the response from the request
// has header with specified value
func WithExpectedResponseHeader(key, value string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithExpectedContentLength is shorthand for creating "Content-Length" header check.
func WithExpectedContentLength(value int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNonEmptyResponseHeader validates that the response from the request
// has header with non empty value.
func WithNonEmptyResponseHeader(key string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExpectedJSONResponse validates that the response from the request in the
// Request function matches JSON-encoded body provided here.
func WithExpectedJSONResponse(response any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUnmarshalJSONResponse unmarshals response body from the request in the
// Request function to the provided response. Response must be a pointer.
func WithUnmarshalJSONResponse(response any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPutResponseBody replaces the data in the provided byte slice with the
// data from the response body of the request in the Request function.
//
// Example:
//
//	var respBytes []byte
//	options := []jsonhttptest.Option{
//		jsonhttptest.WithPutResponseBody(&respBytes),
//	}
func WithPutResponseBody(b *[]byte) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNoResponseBody ensures that there is no data sent by the response of the
// request in the Request function.
func WithNoResponseBody() Option { _ = "STUB: not implemented"; return *new(Option) }

type options struct {
	ctx                     context.Context
	requestBody             io.Reader
	requestHeaders          http.Header
	expectedResponseHeaders http.Header
	nonEmptyResponseHeaders []string
	expectedResponse        []byte
	expectedJSONResponse    any
	unmarshalResponse       any
	responseBody            *[]byte
	noResponseBody          bool
}

type Option interface {
	apply(*options) error
}
type optionFunc func(*options) error

func (f optionFunc) apply(r *options) error { _ = "STUB: not implemented"; return nil }
