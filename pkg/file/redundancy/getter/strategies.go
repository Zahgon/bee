// Copyright 2023 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package getter

import (
	"context"
	"time"

	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/retrieval"
)

const (
	DefaultStrategy     = DATA                           // default prefetching strategy
	DefaultStrict       = false                          // default fallback modes
	DefaultFetchTimeout = retrieval.RetrieveChunkTimeout // timeout for each chunk retrieval
)

type (
	strategyKey     struct{}
	modeKey         struct{}
	fetchTimeoutKey struct{}
	loggerKey       struct{}
	Strategy        = int
)

// Config is the configuration for the getter - public
type Config struct {
	Strategy     Strategy
	Strict       bool
	FetchTimeout time.Duration
	Logger       log.Logger
}

const (
	NONE Strategy = iota // no prefetching and no decoding
	DATA                 // just retrieve data shards no decoding
	PROX                 // proximity driven selective fetching
	RACE                 // aggressive fetching racing all chunks
	strategyCnt
)

// DefaultConfig is the default configuration for the getter
var DefaultConfig = Config{
	Strategy:     DefaultStrategy,
	Strict:       DefaultStrict,
	FetchTimeout: DefaultFetchTimeout,
	Logger:       log.Noop,
}

// NewConfigFromContext returns a new Config based on the context
func NewConfigFromContext(ctx context.Context, def Config) (conf Config, err error) {
	_ = "STUB: not implemented"
	return *new(Config), nil
}

// SetStrategy sets the strategy for the retrieval
func SetStrategy(ctx context.Context, s Strategy) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// SetStrict sets the strict mode for the retrieval
func SetStrict(ctx context.Context, strict bool) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// SetFetchTimeout sets the timeout for each fetch
func SetFetchTimeout(ctx context.Context, timeout time.Duration) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func SetLogger(ctx context.Context, l log.Logger) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// SetConfigInContext sets the config params in the context
func SetConfigInContext(ctx context.Context, s *Strategy, fallbackmode *bool, fetchTimeout *string, logger log.Logger) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
