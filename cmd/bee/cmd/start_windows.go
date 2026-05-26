// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package cmd

import (
	"golang.org/x/sys/windows/svc/debug"

	"github.com/ethersphere/bee/v2/pkg/log"
)

func isWindowsService() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func createWindowsEventLogger(svcName string, logger log.Logger) (log.Logger, error) {
	_ = "STUB: not implemented"
	return *new(log.Logger), nil
}

type windowsEventLogger struct {
	log.Logger
	winlog debug.Log
}

func (l windowsEventLogger) Debug(_ string, _ ...interface{}) { _ = "STUB: not implemented"; return }

func (l windowsEventLogger) Info(msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l windowsEventLogger) Warning(msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (l windowsEventLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}
