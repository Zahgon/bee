// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package api

import (
	"archive/tar"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/ethersphere/bee/v2/pkg/file/redundancy"
	"github.com/ethersphere/bee/v2/pkg/log"
	"github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/storer"
	"github.com/ethersphere/bee/v2/pkg/swarm"
	"github.com/opentracing/opentracing-go"
)

var errEmptyDir = errors.New("no files in root directory")

// dirUploadHandler uploads a directory supplied as a tar in an HTTP request
func (s *Service) dirUploadHandler(
	ctx context.Context,
	logger log.Logger,
	span opentracing.Span,
	w http.ResponseWriter,
	r *http.Request,
	putter storer.PutterSession,
	encrypt bool,
	tag uint64,
	rLevel redundancy.Level,
	act bool,
	historyAddress swarm.Address,
) {
	_ = "STUB: not implemented"
	return
}

// Parse error is ignored; unsupported media types are caught by the default case below.

// storeDir stores all files recursively contained in the directory given as a tar/multipart
// it returns the hash for the uploaded manifest corresponding to the uploaded dir
func storeDir(
	ctx context.Context,
	encrypt bool,
	reader dirReader,
	log log.Logger,
	putter storage.Putter,
	getter storage.Getter,
	indexFilename,
	errorFilename string,
	rLevel redundancy.Level,
) (swarm.Address, error) {
	_ = "STUB: not implemented"
	return *new(swarm.Address), nil
}

// iterate through the files in the supplied tar

// add file entry to dir manifest

// check if files were uploaded through the manifest

// store website information

// save manifest

type FileInfo struct {
	Path        string
	Name        string
	ContentType string
	Size        int64
	Reader      io.Reader
}

type dirReader interface {
	Next() (*FileInfo, error)
}

type tarReader struct {
	r      *tar.Reader
	logger log.Logger
}

func (t *tarReader) Next() (*FileInfo, error) { _ = "STUB: not implemented"; return nil, nil }

// always use Unix path separator

// only store regular files

// multipart reader returns files added as a multipart form. We will ensure all the
// part headers are passed correctly
type multipartReader struct {
	r *multipart.Reader
}

func (m *multipartReader) Next() (*FileInfo, error) { _ = "STUB: not implemented"; return nil, nil }
