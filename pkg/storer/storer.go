// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package storer

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/stabilization"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/transaction"

	"github.com/ethersphere/bee/v2/pkg/postage"
	"github.com/ethersphere/bee/v2/pkg/pusher"
	"github.com/ethersphere/bee/v2/pkg/retrieval"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/storage/leveldbstore"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/cache"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/events"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/reserve"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/upload"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/ethersphere/bee/v2/pkg/topology"
	"github.com/ethersphere/bee/v2/pkg/tracing"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/afero"
	"resenje.org/multex"
)

// PutterSession provides a session around the storage.Putter. The session on
// successful completion commits all the operations or in case of error, rolls back
// the state.
type PutterSession interface {
	storage.Putter
	// Done is used to close the session and optionally assign a swarm.Address to
	// this session.
	Done(swarm.Address) error
	// Cleanup is used to cleanup any state related to this session in case of
	// any error.
	Cleanup() error
}

// SessionInfo is a type which exports the storer tag object. This object
// stores all the relevant information about a particular session.
type SessionInfo = upload.TagItem

// UploadStore is a logical component of the storer which deals with the upload
// of data to swarm.
type UploadStore interface {
	// Upload provides a PutterSession which is tied to the tagID. Optionally if
	// users requests to pin the data, a new pinning collection is created.
	Upload(ctx context.Context, pin bool, tagID uint64) (PutterSession, error)
	// NewSession can be used to obtain a tag ID to use for a new Upload session.
	NewSession() (SessionInfo, error)
	// Session will show the information about the session.
	Session(tagID uint64) (SessionInfo, error)
	// DeleteSession will delete the session info associated with the tag id.
	DeleteSession(tagID uint64) error
	// ListSessions will list all the Sessions currently being tracked.
	ListSessions(offset, limit int) ([]SessionInfo, error)
}

// PinStore is a logical component of the storer which deals with pinning
// functionality.
type PinStore interface {
	// NewCollection can be used to create a new PutterSession which writes a new
	// pinning collection. The address passed in during the Done of the session is
	// used as the root referencce.
	NewCollection(context.Context) (PutterSession, error)
	// DeletePin deletes all the chunks associated with the collection pointed to
	// by the swarm.Address passed in.
	DeletePin(context.Context, swarm.Address) error
	// Pins returns all the root references of pinning collections.
	Pins() ([]swarm.Address, error)
	// HasPin is a helper which checks if a collection exists with the root
	// reference passed in.
	HasPin(swarm.Address) (bool, error)
}

// PinIterator is a helper interface which can be used to iterate over all the
// chunks in a pinning collection.
type PinIterator interface {
	IteratePinCollection(root swarm.Address, iterateFn func(swarm.Address) (bool, error)) error
}

// CacheStore is a logical component of the storer that deals with cache
// content.
type CacheStore interface {
	// Lookup method provides a storage.Getter wrapped around the underlying
	// ChunkStore which will update cache related indexes if required on successful
	// lookups.
	Lookup() storage.Getter
	// Cache method provides a storage.Putter which will add the chunks to cache.
	// This will add the chunk to underlying store as well as new indexes which
	// will keep track of the chunk in the cache.
	Cache() storage.Putter
}

// NetStore is a logical component of the storer that deals with network. It will
// push/retrieve chunks from the network.
type NetStore interface {
	// DirectUpload provides a session which can be used to push chunks directly
	// to the network.
	DirectUpload() PutterSession
	// Download provides a getter which can be used to download data. If the data
	// is found locally, its returned immediately, otherwise it is retrieved from
	// the network.
	Download(cache bool) storage.Getter
	// PusherFeed is the feed for direct push chunks. This can be used by the
	// pusher component to push out the chunks.
	PusherFeed() <-chan *pusher.Op
}

var _ Reserve = (*DB)(nil)

// Reserve is a logical component of the storer that deals with reserve
// content. It will implement all the core functionality required for the protocols.
type Reserve interface {
	ReserveStore
	EvictBatch(ctx context.Context, batchID []byte) error
	ReserveSample(context.Context, []byte, uint8, uint64, *big.Int) (Sample, error)
	ReserveSize() int
}

// ReserveIterator is a helper interface which can be used to iterate over all
// the chunks in the reserve.
type ReserveIterator interface {
	ReserveIterateChunks(cb func(swarm.Chunk) (bool, error)) error
}

// ReserveStore is a logical component of the storer that deals with reserve
// content. It will implement all the core functionality required for the protocols.
type ReserveStore interface {
	ReserveGet(ctx context.Context, addr swarm.Address, batchID []byte, stampHash []byte) (swarm.Chunk, error)
	ReserveHas(addr swarm.Address, batchID []byte, stampHash []byte) (bool, error)
	ReservePutter() storage.Putter
	SubscribeBin(ctx context.Context, bin uint8, start uint64) (<-chan *BinC, func(), <-chan error)
	ReserveLastBinIDs() ([]uint64, uint64, error)
	RadiusChecker
}

// RadiusChecker provides the radius related functionality.
type RadiusChecker interface {
	IsWithinStorageRadius(addr swarm.Address) bool
	StorageRadius() uint8
	CommittedDepth() uint8
	CapacityDoubling() uint8
}

// LocalStore is a read-only ChunkStore. It can be used to check if chunk is known
// locally, but it cannot tell what is the context of the chunk (whether it is
// pinned, uploaded, etc.).
type LocalStore interface {
	ChunkStore() storage.ReadOnlyChunkStore
}

// Debugger is a helper interface which can be used to debug the storer.
type Debugger interface {
	DebugInfo(context.Context) (Info, error)
}

type NeighborhoodStats interface {
	NeighborhoodsStat(ctx context.Context) ([]*NeighborhoodStat, error)
}

type memFS struct {
	afero.Fs
}

func (m *memFS) Open(path string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

type dirFS struct {
	basedir string
}

func (d *dirFS) Open(path string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

var (
	sharkyNoOfShards = 32
	ErrDBQuit        = errors.New("db quit")
)

type closerFn func() error

func (c closerFn) Close() error { _ = "STUB: not implemented"; return nil }

func closer(closers ...io.Closer) io.Closer { _ = "STUB: not implemented"; return *new(io.Closer) }

func initInmemRepository() (transaction.Storage, io.Closer, error) {
	_ = "STUB: not implemented"
	return *new(transaction.Storage), *new(io.Closer), nil
}

// loggerName is the tree path name of the logger for this package.
const loggerName = "storer"

// Default options for levelDB.
const (
	defaultOpenFilesLimit         = uint64(256)
	defaultBlockCacheCapacity     = uint64(32 * 1024 * 1024)
	defaultWriteBufferSize        = uint64(32 * 1024 * 1024)
	defaultDisableSeeksCompaction = false
	defaultCacheCapacity          = uint64(1_000_000)
	defaultBgCacheWorkers         = 32
	DefaultReserveCapacity        = 1 << 22 // 4194304 chunks

	indexPath  = "indexstore"
	sharkyPath = "sharky"
)

func initStore(basePath string, opts *Options) (*leveldbstore.Store, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initDiskRepository(
	ctx context.Context,
	basePath string,
	opts *Options,
) (transaction.Storage, *PinIntegrity, io.Closer, int, error) {
	_ = "STUB: not implemented"
	return *new(transaction.Storage), nil, *new(io.Closer), 0, nil
}

const lockKeyNewSession string = "new_session"

// Options provides a container to configure different things in the storer.
type Options struct {
	// These are options related to levelDB. Currently, the underlying storage used is levelDB.
	LdbStats                  atomic.Pointer[prometheus.HistogramVec]
	LdbOpenFilesLimit         uint64
	LdbBlockCacheCapacity     uint64
	LdbWriteBufferSize        uint64
	LdbDisableSeeksCompaction bool
	Logger                    log.Logger
	Tracer                    *tracing.Tracer

	Address           swarm.Address
	StartupStabilizer stabilization.Subscriber
	Batchstore        postage.Storer
	ValidStamp        postage.ValidStampFn
	RadiusSetter      topology.SetStorageRadiuser
	StateStore        storage.StateStorer

	ReserveCapacity         int
	ReserveWakeUpDuration   time.Duration
	ReserveMinEvictCount    uint64
	ReserveCapacityDoubling int

	CacheCapacity      uint64
	CacheMinEvictCount uint64

	MinimumStorageRadius uint
}

func defaultOptions() *Options { _ = "STUB: not implemented"; return nil }

// cacheLimiter is used to limit the number
// of concurrent cache background workers.
type cacheLimiter struct {
	wg     sync.WaitGroup
	sem    chan struct{}
	ctx    context.Context
	cancel context.CancelFunc
}

// DB implements all the component stores described above.
type DB struct {
	logger log.Logger
	tracer *tracing.Tracer

	metrics             metrics
	storage             transaction.Storage
	multex              *multex.Multex
	cacheObj            *cache.Cache
	retrieval           retrieval.Interface
	pusherFeed          chan *pusher.Op
	quit                chan struct{}
	cacheLimiter        cacheLimiter
	dbCloser            io.Closer
	subscriptionsWG     sync.WaitGroup
	events              *events.Subscriber
	directUploadLimiter chan struct{}

	reserve          *reserve.Reserve
	inFlight         sync.WaitGroup
	reserveBinEvents *events.Subscriber
	baseAddr         swarm.Address
	batchstore       postage.Storer
	validStamp       postage.ValidStampFn
	setSyncerOnce    sync.Once
	syncer           Syncer
	reserveOptions   reserveOpts

	pinIntegrity *PinIntegrity
}

type reserveOpts struct {
	startupStabilizer  stabilization.Subscriber
	wakeupDuration     time.Duration
	minEvictCount      uint64
	cacheMinEvictCount uint64
	minimumRadius      uint8
	capacityDoubling   int
}

// New returns a newly constructed DB object which implements all the above
// component stores.
func New(ctx context.Context, dirPath string, opts *Options) (*DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cleanup any dirty state in upload and pinning stores, this could happen
// in case of dirty shutdowns

// Reset removes all entries
func (db *DB) ResetReserve(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Metrics returns set of prometheus collectors.
func (db *DB) Metrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }

// StatusMetrics exposes metrics that are exposed on the status protocol.
func (db *DB) StatusMetrics() []prometheus.Collector { _ = "STUB: not implemented"; return nil }

func (db *DB) Close() error { _ = "STUB: not implemented"; return nil }

func (db *DB) SetRetrievalService(r retrieval.Interface) {
	_ = "STUB: not implemented"

	// StartReserveWorker starts the reserve worker. It takes an optional ready channel that is closed whenever the reserve
	// worker has finished starting, as this synchronization is needed for some tests. The channel is not used for writing anywhere.
	return
}

func (db *DB) StartReserveWorker(ctx context.Context, s Syncer, radius func() (uint8, error), ready chan<- struct{}) {
	_ = "STUB: not implemented"
	return
}

type noopRetrieval struct{}

func (noopRetrieval) RetrieveChunk(_ context.Context, _ swarm.Address, _ swarm.Address) (swarm.Chunk, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Chunk), nil
}

func (db *DB) ChunkStore() storage.ReadOnlyChunkStore {
	_ = "STUB: not implemented"
	return *new(storage.ReadOnlyChunkStore)
}

func (db *DB) PinIntegrity() *PinIntegrity { _ = "STUB: not implemented"; return nil }

func (db *DB) Lock(strs ...string) func() { _ = "STUB: not implemented"; return nil }

func (db *DB) Storage() transaction.Storage {
	_ = "STUB: not implemented"
	return *new(transaction.Storage)
}

type putterSession struct {
	storage.Putter
	done    func(swarm.Address) error
	cleanup func() error
}

func (p *putterSession) Done(addr swarm.Address) error { _ = "STUB: not implemented"; return nil }

func (p *putterSession) Cleanup() error { _ = "STUB: not implemented"; return nil }
