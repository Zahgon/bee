// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"crypto/ecdsa"

	"github.com/ethersphere/bee/v2/pkg/keystore"
)

// Service is the file-based keystore.Service implementation.
//
// Keys are stored in directory where each private key is stored in a file,
// which is encrypted with symmetric key using some password.
type Service struct {
	dir string
}

// New creates new file-based keystore.Service implementation.
func New(dir string) *Service { _ = "STUB: not implemented"; return nil }

func (s *Service) Exists(name string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (s *Service) SetKey(name, password string, edg keystore.EDG) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) Key(name, password string, edg keystore.EDG) (pk *ecdsa.PrivateKey, created bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (s *Service) keyFilename(name string) string { _ = "STUB: not implemented"; return "" }
