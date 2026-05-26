// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"crypto/ecdsa"

	"github.com/ethersphere/bee/v2/pkg/keystore"
)

var _ keystore.Service = (*Service)(nil)

const (
	keyHeaderKDF = "scrypt"
	keyVersion   = 3

	scryptN     = 1 << 15
	scryptR     = 8
	scryptP     = 1
	scryptDKLen = 32
)

// This format is compatible with Ethereum JSON v3 key file format.
type encryptedKey struct {
	Address string    `json:"address"`
	Crypto  keyCripto `json:"crypto"`
	Version int       `json:"version"`
	Id      string    `json:"id"`
}

type keyCripto struct {
	Cipher       string       `json:"cipher"`
	CipherText   string       `json:"ciphertext"`
	CipherParams cipherParams `json:"cipherparams"`
	KDF          string       `json:"kdf"`
	KDFParams    kdfParams    `json:"kdfparams"`
	MAC          string       `json:"mac"`
}

type cipherParams struct {
	IV string `json:"iv"`
}

type kdfParams struct {
	N     int    `json:"n"`
	R     int    `json:"r"`
	P     int    `json:"p"`
	DKLen int    `json:"dklen"`
	Salt  string `json:"salt"`
}

func encryptKey(k *ecdsa.PrivateKey, password string, edg keystore.EDG) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptKey(data []byte, password string, edg keystore.EDG) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encryptData(data, password []byte) (*keyCripto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptData(v keyCripto, password string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if this fails we might be trying to load an ethereum V3 keyfile

func aesCTRXOR(key, inText, iv []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func getKDFKey(v keyCripto, password []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
