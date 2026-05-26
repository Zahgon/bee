// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// mapStructureTagName represents the name of the tag used to map values.
const mapStructureTagName = "map"

// errHexLength reports an attempt to decode an odd-length input.
// It's a drop-in replacement for hex.ErrLength.
var errHexLength = errors.New("odd length hex string")

// hexInvalidByteError values describe errors resulting
// from an invalid byte in a hex string.
// It's a drop-in replacement for hex.InvalidByteError.
type hexInvalidByteError byte

// Error implements the error interface.
func (e hexInvalidByteError) Error() string { _ = "STUB: not implemented"; return "" }

// parseError is returned when an entry cannot be parsed.
type parseError struct {
	Entry string
	Value string
	Cause error
}

// Error implements the error interface.
func (e *parseError) Error() string { _ = "STUB: not implemented"; return "" }

// Unwrap implements the interface required by errors.Unwrap function.
func (e *parseError) Unwrap() error {
	_ = "STUB: not implemented"

	// Equal returns true if the given error
	// type and fields are equal to this error.
	// It is used to compare errors in tests.
	return nil
}

func (e *parseError) Equal(err error) bool { _ = "STUB: not implemented"; return false }

// newParseError returns a new mapStructure error.
// If the cause is strconv.NumError, its
// underlying error is unwrapped and
// used as a cause. The hex.InvalidByteError
// and hex.ErrLength errors are replaced in
// order to hide unnecessary information.
func newParseError(entry, value string, cause error) error { _ = "STUB: not implemented"; return nil }

// flattenErrorsFormat flattens the errors in
// the multierror.Error as a one-line string.
var flattenErrorsFormat = func(es []error) string {
	messages := make([]string, len(es))
	for i, err := range es {
		messages[i] = err.Error()
	}
	return fmt.Sprintf(
		"%d error(s) occurred: %v",
		len(es),
		strings.Join(messages, "; "),
	)
}

// mapStructure maps the input to the output values.
// The input is one of the following:
//   - map[string]string
//   - map[string][]string
//
// In the second case, the first value of
// the string array is taken as a value.
//
// The output struct fields can contain the
// `map` tag that refers to the map input key.
// For example:
//
//	type Output struct {
//		BoolVal bool `map:"boolVal,omitempty"`
//	}
//
// If the `map` tag is not present, the field name is used.
// If the field name or the `map` tag is not present in
// the input map, the field is skipped. If the map value
// is empty and the` omitempty` tag is present then the
// field is skipped.
//
// In case of parsing error, a new parseError is returned to the caller.
// The caller can use the Unwrap method to get the original error.
func mapStructure(input, output any, hooks map[string]func(v string) (string, error)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Do input sanity checks.

// Do output sanity checks.

// set is the workhorse here, parsing and setting the values.

// Clear the field on error.

// Nil slice.

// parseFieldTags parses the given field tags into name, hook, and omitempty.

// Map input into output.

// numberSize returns the size of the number in bits.
func numberSize(k reflect.Kind) int { _ = "STUB: not implemented"; return 0 }

// flattenValue returns the first element of the value if it is a slice.
func flattenValue(val reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string { _ = "STUB: not implemented"; return "" }
