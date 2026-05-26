// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"os"
	"sync"
)

// defaults specifies the default global options for log
// package which every new logger will inherit on its creation.
var defaults = struct {
	pin       sync.Once // pin pins the options and formatter settings.
	options   *Options
	formatter *formatter
}{
	options: &Options{
		sink:      os.Stderr,
		verbosity: VerbosityDebug,
		fmtOptions: fmtOptions{
			timestampLayout: "2006-01-02 15:04:05.000000",
			maxLogDepth:     16,
		},
	},
}

// ModifyDefaults modifies the global default options for this log package
// that each new logger inherits when it is created. The default values can
// be modified only once, so further calls to this function will be ignored.
// This function should be called before the first call to the NewLogger
// factory constructor, otherwise it will have no effect.
func ModifyDefaults(opts ...Option) { _ = "STUB: not implemented"; return }

// loggers is the central register for Logger instances.
var loggers = new(sync.Map)

// NewLogger is a factory constructor which returns a new logger instance
// based on the given name. If such an instance already exists in the
// logger registry, then this existing instance is returned instead.
// The given options take precedence over the default options set
// by the ModifyDefaults function.
func NewLogger(name string, opts ...Option) Logger {
	_ = "STUB: not implemented"
	// Pin the default settings if
	// they are not already pinned.
	return *new(Logger)
}

// SetVerbosity sets the level
// of verbosity of the given logger.
func SetVerbosity(l Logger, v Level) error { _ = "STUB: not implemented"; return nil }

// SetVerbosityByExp sets all loggers to the given
// verbosity level v that match the given expression
// e, which can be a logger id or a regular expression.
// An error is returned if e fails to compile.
func SetVerbosityByExp(e string, v Level) error { _ = "STUB: not implemented"; return nil }

// RegistryIterate iterates through all registered loggers.
func RegistryIterate(fn func(id, path string, verbosity Level, v uint) (next bool)) {
	_ = "STUB: not implemented"
	return
}
