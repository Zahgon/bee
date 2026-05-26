// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package api provides the functionality of the Bee
// client-facing HTTP API.
package api

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethersphere/bee/v2/pkg/accesscontrol"
	"github.com/ethersphere/bee/v2/pkg/accounting"
	"github.com/ethersphere/bee/v2/pkg/crypto"
	"github.com/ethersphere/bee/v2/pkg/feeds"
	"github.com/ethersphere/bee/v2/pkg/file/pipeline"
	"github.com/ethersphere/bee/v2/pkg/file/redundancy"
	"github.com/ethersphere/bee/v2/pkg/gsoc"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/p2p"
	"github.com/ethersphere/bee/v2/pkg/pingpong"
	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/postage/postagecontract"
	"github.com/ethersphere/bee/v2/pkg/pss"
	"github.com/ethersphere/bee/v2/pkg/resolver"
	"github.com/ethersphere/bee/v2/pkg/settlement"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/chequebook"
	"github.com/ethersphere/bee/v2/pkg/settlement/swap/erc20"
	"github.com/ethersphere/bee/v2/pkg/status"
	"github.com/ethersphere/bee/v2/pkg/steward"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/storageincentives"
	"github.com/ethersphere/bee/v2/pkg/storageincentives/staking"
	"github.com/ethersphere/bee/v2/pkg/storer"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/topology"
	"github.com/ethersphere/bee/v2/pkg/topology/lightnode"
	"github.com/ethersphere/bee/v2/pkg/tracing"
	"github.com/ethersphere/bee/v2/pkg/transaction"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/semaphore"
)

// loggerName is the tree path name of the logger for this package.
const loggerName = "api"

const (
	SwarmPinHeader                    = "Swarm-Pin"
	SwarmTagHeader                    = "Swarm-Tag"
	SwarmEncryptHeader                = "Swarm-Encrypt"
	SwarmIndexDocumentHeader          = "Swarm-Index-Document"
	SwarmErrorDocumentHeader          = "Swarm-Error-Document"
	SwarmSocSignatureHeader           = "Swarm-Soc-Signature"
	SwarmFeedIndexHeader              = "Swarm-Feed-Index"
	SwarmFeedIndexNextHeader          = "Swarm-Feed-Index-Next"
	SwarmFeedResolvedVersionHeader    = "Swarm-Feed-Resolved-Version"
	SwarmOnlyRootChunk                = "Swarm-Only-Root-Chunk"
	SwarmCollectionHeader             = "Swarm-Collection"
	SwarmPostageBatchIdHeader         = "Swarm-Postage-Batch-Id"
	SwarmPostageStampHeader           = "Swarm-Postage-Stamp"
	SwarmDeferredUploadHeader         = "Swarm-Deferred-Upload"
	SwarmRedundancyLevelHeader        = "Swarm-Redundancy-Level"
	SwarmRedundancyStrategyHeader     = "Swarm-Redundancy-Strategy"
	SwarmRedundancyFallbackModeHeader = "Swarm-Redundancy-Fallback-Mode"
	SwarmChunkRetrievalTimeoutHeader  = "Swarm-Chunk-Retrieval-Timeout"
	SwarmLookAheadBufferSizeHeader    = "Swarm-Lookahead-Buffer-Size"
	SwarmActHeader                    = "Swarm-Act"
	SwarmActTimestampHeader           = "Swarm-Act-Timestamp"
	SwarmActPublisherHeader           = "Swarm-Act-Publisher"
	SwarmActHistoryAddressHeader      = "Swarm-Act-History-Address"

	ImmutableHeader = "Immutable"
	GasPriceHeader  = "Gas-Price"
	GasLimitHeader  = "Gas-Limit"
	ETagHeader      = "ETag"

	AuthorizationHeader        = "Authorization"
	AcceptEncodingHeader       = "Accept-Encoding"
	ContentTypeHeader          = "Content-Type"
	ContentDispositionHeader   = "Content-Disposition"
	ContentLengthHeader        = "Content-Length"
	RangeHeader                = "Range"
	OriginHeader               = "Origin"
	AccessControlExposeHeaders = "Access-Control-Expose-Headers"
)

const (
	multiPartFormData = "multipart/form-data"
	contentTypeTar    = "application/x-tar"
)

var (
	errInvalidNameOrAddress             = errors.New("invalid name or bzz address")
	errNoResolver                       = errors.New("no resolver connected")
	errInvalidRequest                   = errors.New("could not validate request")
	errInvalidContentType               = errors.New("invalid content-type")
	errDirectoryStore                   = errors.New("could not store directory")
	errFileStore                        = errors.New("could not store file")
	errInvalidPostageBatch              = errors.New("invalid postage batch id")
	errBatchUnusable                    = errors.New("batch not usable")
	errOperationSupportedOnlyInFullMode = errors.New("operation is supported only in full mode")
	errActDownload                      = errors.New("act download failed")
	errActUpload                        = errors.New("act upload failed")
	errActGranteeList                   = errors.New("failed to create or update grantee list")

	batchIdOrStampSig = fmt.Sprintf("Either '%s' or '%s' header must be set in the request", SwarmPostageStampHeader, SwarmPostageBatchIdHeader)
)

// Storer interface provides the functionality required from the local storage
// component of the node.
type Storer interface {
	storer.UploadStore
	storer.PinStore
	storer.CacheStore
	storer.NetStore
	storer.LocalStore
	storer.RadiusChecker
	storer.Debugger
	storer.NeighborhoodStats
}

type PinIntegrity interface {
	Check(ctx context.Context, logger log.Logger, pin string, out chan storer.PinStat)
}

type Service struct {
	storer          Storer
	resolver        resolver.Interface
	pss             pss.Interface
	gsoc            gsoc.Listener
	steward         steward.Interface
	logger          log.Logger
	loggerV1        log.Logger
	tracer          *tracing.Tracer
	feedFactory     feeds.Factory
	signer          crypto.Signer
	post            postage.Service
	accesscontrol   accesscontrol.Controller
	postageContract postagecontract.Interface
	probe           *Probe
	metricsRegistry *prometheus.Registry
	stakingContract staking.Contract
	Options

	http.Handler
	router *mux.Router

	metrics metrics

	wsWg sync.WaitGroup // wait for all websockets to close on exit
	quit chan struct{}

	overlay           *swarm.Address
	publicKey         ecdsa.PublicKey
	pssPublicKey      ecdsa.PublicKey
	ethereumAddress   common.Address
	chequebookEnabled bool
	swapEnabled       bool
	fullAPIEnabled    bool

	topologyDriver topology.Driver
	p2p            p2p.DebugService
	accounting     accounting.Interface
	chequebook     chequebook.Service
	pseudosettle   settlement.Interface
	pingpong       pingpong.Interface

	batchStore   postage.Storer
	stamperStore storage.Store
	pinIntegrity PinIntegrity

	syncStatus func() (bool, error)

	swap        swap.Interface
	transaction transaction.Service
	lightNodes  *lightnode.Container
	blockTime   time.Duration

	statusSem        *semaphore.Weighted
	postageSem       *semaphore.Weighted
	stakingSem       *semaphore.Weighted
	cashOutChequeSem *semaphore.Weighted
	beeMode          BeeNodeMode

	chainBackend transaction.Backend
	erc20Service erc20.Service
	chainID      int64

	whitelistedWithdrawalAddress []common.Address

	preMapHooks              map[string]func(v string) (string, error)
	customValidationMessages map[string]func(err validator.FieldError) error
	validate                 *validator.Validate

	redistributionAgent *storageincentives.Agent

	statusService *status.Service
	isWarmingUp   bool
}

func (s *Service) SetP2P(p2p p2p.DebugService) { _ = "STUB: not implemented"; return }

func (s *Service) SetSwarmAddress(addr *swarm.Address) { _ = "STUB: not implemented"; return }

func (s *Service) SetRedistributionAgent(redistributionAgent *storageincentives.Agent) {
	_ = "STUB: not implemented"
	return
}

type Options struct {
	CORSAllowedOrigins []string
	WsPingPeriod       time.Duration
}

type ExtraOptions struct {
	Pingpong        pingpong.Interface
	TopologyDriver  topology.Driver
	LightNodes      *lightnode.Container
	Accounting      accounting.Interface
	Pseudosettle    settlement.Interface
	Swap            swap.Interface
	Chequebook      chequebook.Service
	BlockTime       time.Duration
	Storer          Storer
	Resolver        resolver.Interface
	Pss             pss.Interface
	Gsoc            gsoc.Listener
	FeedFactory     feeds.Factory
	Post            postage.Service
	AccessControl   accesscontrol.Controller
	PostageContract postagecontract.Interface
	Staking         staking.Contract
	Steward         steward.Interface
	SyncStatus      func() (bool, error)
	NodeStatus      *status.Service
	PinIntegrity    PinIntegrity
}

func New(
	publicKey, pssPublicKey ecdsa.PublicKey,
	ethereumAddress common.Address,
	whitelistedWithdrawalAddress []string,
	logger log.Logger,
	transaction transaction.Service,
	batchStore postage.Storer,
	beeMode BeeNodeMode,
	chequebookEnabled bool,
	swapEnabled bool,
	chainBackend transaction.Backend,
	cors []string,
	stamperStore storage.Store,
) *Service {
	_ = "STUB: not implemented"
	return nil
}

// Configure will create a and initialize a new API service.
func (s *Service) Configure(signer crypto.Signer, tracer *tracing.Tracer, o Options, e ExtraOptions, chainID int64, erc20 erc20.Service) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) SetProbe(probe *Probe) { _ = "STUB: not implemented"; return }

func (s *Service) SetIsWarmingUp(v bool) {
	_ = "STUB: not implemented"

	// Close hangs up running websockets on shutdown.
	return
}

func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }

// getOrCreateSessionID attempts to get the session if an tag id is supplied, and returns an error
// if it does not exist. If no id is supplied, it will attempt to create a new session and return it.
func (s *Service) getOrCreateSessionID(tagUid uint64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// if tag ID is not supplied, create a new tag

func (s *Service) resolveNameOrAddress(str string) (swarm.Address, error) {
	_ = "STUB: not implemented"
	// Try and mapStructure the name as a bzz address.
	return *new(swarm.Address), nil
}

// If no resolver is not available, return an error.

// Try and resolve the name using the provided resolver.

func (s *Service) newTracingHandler(spanName string) func(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

// ignore

// ignore

func (s *Service) contentLengthMetricMiddleware() func(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) downloadSpeedMetricMiddleware(endpoint string) func(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

// observeUploadSpeed measures the speed of the upload and sets appropriate
// labels to the metrics. This function can be called as a deferred function in
// side of handler. This functions is not in a form of a middleware to more
// directly pass the deferred flag.
func (s *Service) observeUploadSpeed(w http.ResponseWriter, r *http.Request, start time.Time, endpoint string, deferred bool) {
	_ = "STUB: not implemented"
	return
}

// gasConfigMiddleware can be used by the APIs that allow block chain transactions to set
// gas price and gas limit through the HTTP API headers.
func (s *Service) gasConfigMiddleware(handlerName string) func(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return nil
}

// corsHandler sets CORS headers to HTTP response if allowed origins are configured.
func (s *Service) corsHandler(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// checkOrigin returns true if the origin is not set or is equal to the request host.
func (s *Service) checkOrigin(r *http.Request) bool { _ = "STUB: not implemented"; return false }

// validationError is a custom error type for validation errors.
type validationError struct {
	Entry string
	Value any
	Cause error
}

// Error implements the error interface.
func (e *validationError) Error() string { _ = "STUB: not implemented"; return "" }

// mapStructure maps the input into output struct and validates the output.
// It's a helper method for the handlers, which reduces the chattiness
// of the code.
func (s *Service) mapStructure(input, output any) func(string, log.Logger, http.ResponseWriter) {
	_ = "STUB: not implemented"
	// response unifies the response format for parsing and validation errors.
	return nil
}

// equalASCIIFold returns true if s is equal to t with ASCII case folding as
// defined in RFC 4790.
func equalASCIIFold(s, t string) bool { _ = "STUB: not implemented"; return false }

type putterOptions struct {
	BatchID  []byte
	TagID    uint64
	Deferred bool
	Pin      bool
}

type putterSessionWrapper struct {
	storer.PutterSession
	stamper postage.Stamper
	save    func() error
}

func (p *putterSessionWrapper) Put(ctx context.Context, chunk swarm.Chunk) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *putterSessionWrapper) Done(ref swarm.Address) error { _ = "STUB: not implemented"; return nil }

func (p *putterSessionWrapper) Cleanup() error { _ = "STUB: not implemented"; return nil }

func (s *Service) getStamper(batchID []byte) (postage.Stamper, func() error, error) {
	_ = "STUB: not implemented"
	return *new(postage.Stamper), nil, nil
}

func (s *Service) newStamperPutter(ctx context.Context, opts putterOptions) (storer.PutterSession, error) {
	_ = "STUB: not implemented"
	return *new(storer.PutterSession), nil
}

func (s *Service) newStampedPutter(ctx context.Context, opts putterOptions, stamp *postage.Stamp) (storer.PutterSession, error) {
	_ = "STUB: not implemented"
	return *new(storer.PutterSession), nil
}

// newStampedPutterWithBatch creates a stamped putter using a pre-fetched batch.
// This avoids the database lookup when batch info is already cached.
func (s *Service) newStampedPutterWithBatch(ctx context.Context, opts putterOptions, stamp *postage.Stamp, storedBatch *postage.Batch) (storer.PutterSession, error) {
	_ = "STUB: not implemented"
	return *new(storer.PutterSession), nil
}

type pipelineFunc func(context.Context, io.Reader) (swarm.Address, error)

func requestPipelineFn(s storage.Putter, encrypt bool, rLevel redundancy.Level) pipelineFunc {
	_ = "STUB: not implemented"
	return *new(pipelineFunc)
}

func requestPipelineFactory(ctx context.Context, s storage.Putter, encrypt bool, rLevel redundancy.Level) func() pipeline.Interface {
	_ = "STUB: not implemented"
	return nil
}

type cleanupOnErrWriter struct {
	http.ResponseWriter
	logger log.Logger
	onErr  func() error
}

func (r *cleanupOnErrWriter) WriteHeader(statusCode int) {
	_ = "STUB: not implemented"
	// if there is an error status returned, cleanup.
	return
}

// CalculateNumberOfChunks calculates the number of chunks in an arbitrary
// content length.
func CalculateNumberOfChunks(contentLength int64, isEncrypted bool) int64 {
	_ = "STUB: not implemented"
	return 0
}

// defaultUploadMethod returns true for deferred when the deferred header is not present.
func defaultUploadMethod(deferred *bool) bool { _ = "STUB: not implemented"; return false }
