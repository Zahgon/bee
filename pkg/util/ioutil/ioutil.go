// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ioutil

// DB folders paths from bee datadir
const (
	DataPathLocalstore = "localstore"
	DataPathKademlia   = "kademlia-metrics"
)

// The WriterFunc type is an adapter to allow the use of
// ordinary functions as io.Writer Write method. If f is
// a function with the appropriate signature, WriterFunc(f)
// is an io.Writer that calls f.
type WriterFunc func([]byte) (int, error)

// WriterFunc calls f(p).
func (f WriterFunc) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"

	// RemoveContent removes all files in path. Copied function from cmd/db.go
	return 0, nil
}

func RemoveContent(path string) error { _ = "STUB: not implemented"; return nil }
