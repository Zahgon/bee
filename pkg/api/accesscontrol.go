// Copyright 2024 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"crypto/ecdsa"
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/storer"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

type addressKey struct{}

const granteeListEncrypt = true

// getAddressFromContext is a helper function to extract the address from the context.
func getAddressFromContext(ctx context.Context) swarm.Address {
	_ = "STUB: not implemented"
	return *new(swarm.Address)
}

// setAddressInContext sets the swarm address in the context.
func setAddressInContext(ctx context.Context, address swarm.Address) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// GranteesPatchRequest represents a request to patch the list of grantees.
type GranteesPatchRequest struct {
	// Addlist contains the list of grantees to add.
	Addlist []string `json:"add"`

	// Revokelist contains the list of grantees to revoke.
	Revokelist []string `json:"revoke"`
}

// GranteesPatchResponse represents the response structure for patching grantees.
type GranteesPatchResponse struct {
	// Reference represents the swarm address.
	Reference swarm.Address `json:"ref"`
	// HistoryReference represents the reference to the history of an access control entry.
	HistoryReference swarm.Address `json:"historyref"`
}

// GranteesPostRequest represents the request structure for adding grantees.
type GranteesPostRequest struct {
	// GranteeList represents the list of grantees to be saves on Swarm.
	GranteeList []string `json:"grantees"`
}

// GranteesPostResponse represents the response structure for adding grantees.
type GranteesPostResponse struct {
	// Reference represents the saved grantee list Swarm address.
	Reference swarm.Address `json:"ref"`
	// HistoryReference represents the reference to the history of an access control entry.
	HistoryReference swarm.Address `json:"historyref"`
}

// GranteesPatch represents a structure for modifying the list of grantees.
type GranteesPatch struct {
	// Addlist is a list of ecdsa.PublicKeys to be added to a grantee list.
	Addlist []*ecdsa.PublicKey
	// Revokelist is a list of ecdsa.PublicKeys to be removed from a grantee list
	Revokelist []*ecdsa.PublicKey
}

// actDecryptionHandler is a middleware that looks up and decrypts the given address,
// if the act headers are present.
func (s *Service) actDecryptionHandler() func(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

// Try to download the file wihthout decryption, if the act headers are not present

// actEncryptionHandler is a middleware that encrypts the given address using the publisher's public key,
// uploads the encrypted reference, history and kvs to the store.
func (s *Service) actEncryptionHandler(
	ctx context.Context,
	putter storer.PutterSession,
	reference swarm.Address,
	historyRootHash swarm.Address,
) (swarm.Address, swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), *new(swarm.Address), nil
}

// only need to upload history and kvs if a new history is created,
// meaning that the publisher uploaded to the history for the first time

// actListGranteesHandler is a middleware that decrypts the given address and returns the list of grantees,
// only the publisher is authorized to access the list.
func (s *Service) actListGranteesHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// actGrantRevokeHandler is a middleware that makes updates to the list of grantees,
// only the publisher is authorized to perform this action.
func (s *Service) actGrantRevokeHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// actCreateGranteesHandler is a middleware that creates a new list of grantees,
// only the publisher is authorized to perform this action.
func (s *Service) actCreateGranteesHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func parseKeys(list []string) ([]*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
