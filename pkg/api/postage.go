// Copyright 2021 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/bigint"
	"github.com/ethersphere/bee/v2/pkg/postage"
)

var defaultImmutable = true

func (s *Service) postageAccessHandler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (s *Service) postageSyncStatusCheckHandler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// hexByte takes care that a byte slice gets correctly
// marshaled by the json serializer.
type hexByte []byte

func (b hexByte) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type postageCreateResponse struct {
	BatchID hexByte `json:"batchID"`
	TxHash  string  `json:"txHash"`
}

func (s *Service) postageCreateHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Set the default.

type postageStampResponse struct {
	BatchID          hexByte        `json:"batchID"`
	Utilization      uint32         `json:"utilization"`
	UtilizationRatio float64        `json:"utilizationRatio"`
	Usable           bool           `json:"usable"`
	Label            string         `json:"label"`
	Depth            uint8          `json:"depth"`
	Amount           *bigint.BigInt `json:"amount"`
	BucketDepth      uint8          `json:"bucketDepth"`
	BlockNumber      uint64         `json:"blockNumber"`
	ImmutableFlag    bool           `json:"immutableFlag"`
	Exists           bool           `json:"exists"`
	BatchTTL         int64          `json:"batchTTL"`
}

type postageStampsResponse struct {
	Stamps []postageStampResponse `json:"stamps"`
}

type postageBatchResponse struct {
	BatchID     hexByte        `json:"batchID"`
	Value       *bigint.BigInt `json:"value"`
	Start       uint64         `json:"start"`
	Owner       hexByte        `json:"owner"`
	Depth       uint8          `json:"depth"`
	BucketDepth uint8          `json:"bucketDepth"`
	Immutable   bool           `json:"immutable"`
	BatchTTL    int64          `json:"batchTTL"`
}

type postageStampBucketsResponse struct {
	Depth            uint8        `json:"depth"`
	BucketDepth      uint8        `json:"bucketDepth"`
	BucketUpperBound uint32       `json:"bucketUpperBound"`
	Buckets          []bucketData `json:"buckets"`
}

type bucketData struct {
	BucketID   uint32 `json:"bucketID"`
	Collisions uint32 `json:"collisions"`
}

func (s *Service) postageGetStampsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) postageGetAllBatchesHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) postageGetBatchHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) postageGetStampBucketsHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) postageGetStampHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type reserveStateResponse struct {
	Radius                  uint8  `json:"radius"`
	StorageRadius           uint8  `json:"storageRadius"`
	Commitment              uint64 `json:"commitment"`
	ReserveCapacityDoubling uint8  `json:"reserveCapacityDoubling"`
}

type chainStateResponse struct {
	ChainTip              uint64         `json:"chainTip"`                        // ChainTip (block height).
	Block                 uint64         `json:"block"`                           // The block number of the last postage event.
	TotalAmount           *bigint.BigInt `json:"totalAmount"`                     // Cumulative amount paid per stamp.
	CurrentPrice          *bigint.BigInt `json:"currentPrice"`                    // Bzz/chunk/block normalised price.
	MinimumValidityBlocks uint64         `json:"minimumValidityBlocks,omitempty"` // Minimum number of blocks a new postage batch must remain valid.
}

func (s *Service) reserveStateHandler(w http.ResponseWriter, _ *http.Request) {
	_ = "STUB: not implemented"
	return
}

// chainStateHandler returns the current chain state.
func (s *Service) chainStateHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// estimateBatchTTL estimates the time remaining until the batch expires.
// The -1 signals that the batch never expires.
func (s *Service) estimateBatchTTLFromID(id []byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// estimateBatchTTL estimates the time remaining until the batch expires.
// The -1 signals that the batch never expires.
func (s *Service) estimateBatchTTL(batch *postage.Batch) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Service) postageTopUpHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) postageDiluteHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
