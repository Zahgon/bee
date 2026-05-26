// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mantaray

import (
	"context"
	"errors"
)

const (
	PathSeparator = '/' // path separator
)

var ZeroObfuscationKey = make([]byte, 32)

// Error used when lookup path does not match
var (
	ErrNotFound         = errors.New("not found")
	ErrEmptyPath        = errors.New("empty path")
	ErrMetadataTooLarge = errors.New("metadata too large")
)

// Node represents a mantaray Node
type Node struct {
	nodeType       uint8
	refBytesSize   int
	obfuscationKey []byte
	ref            []byte // reference to uninstantiated Node persisted serialised
	entry          []byte
	metadata       map[string]string
	forks          map[byte]*fork
}

type fork struct {
	prefix []byte // the non-branching part of the subpath
	*Node         // in memory structure that represents the Node
}

const (
	nodeTypeValue             = uint8(2)
	nodeTypeEdge              = uint8(4)
	nodeTypeWithPathSeparator = uint8(8)
	nodeTypeWithMetadata      = uint8(16)

	nodeTypeMask = uint8(255)
)

func nodeTypeIsWithMetadataType(nodeType uint8) bool { _ = "STUB: not implemented"; return false }

// NewNodeRef is the exported Node constructor used to represent manifests by reference
func NewNodeRef(ref []byte) *Node { _ = "STUB: not implemented"; return nil }

// New is the constructor for in-memory Node structure
func New() *Node { _ = "STUB: not implemented"; return nil }

func notFound(path []byte) error { _ = "STUB: not implemented"; return nil }

// IsValueType returns true if the node contains entry.
func (n *Node) IsValueType() bool { _ = "STUB: not implemented"; return false }

// IsEdgeType returns true if the node forks into other nodes.
func (n *Node) IsEdgeType() bool { _ = "STUB: not implemented"; return false }

// IsWithPathSeparatorType returns true if the node path contains separator character.
func (n *Node) IsWithPathSeparatorType() bool { _ = "STUB: not implemented"; return false }

// IsWithMetadataType returns true if the node contains metadata.
func (n *Node) IsWithMetadataType() bool { _ = "STUB: not implemented"; return false }

func (n *Node) makeValue() { _ = "STUB: not implemented"; return }

func (n *Node) makeEdge() { _ = "STUB: not implemented"; return }

func (n *Node) makeWithPathSeparator() { _ = "STUB: not implemented"; return }

func (n *Node) makeWithMetadata() { _ = "STUB: not implemented"; return }

func (n *Node) makeNotWithPathSeparator() { _ = "STUB: not implemented"; return }

func (n *Node) SetObfuscationKey(obfuscationKey []byte) { _ = "STUB: not implemented"; return }

// Reference returns the address of the mantaray node if saved.
func (n *Node) Reference() []byte {
	_ = "STUB: not implemented"

	// Entry returns the value stored on the specific path.
	return nil
}

func (n *Node) Entry() []byte {
	_ = "STUB: not implemented"

	// Metadata returns the metadata stored on the specific path.
	return nil
}

func (n *Node) Metadata() map[string]string {
	_ = "STUB: not implemented"

	// LookupNode finds the node for a path or returns error if not found
	return nil
}

func (n *Node) LookupNode(ctx context.Context, path []byte, l Loader) (*Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lookup finds the entry for a path or returns error if not found
func (n *Node) Lookup(ctx context.Context, path []byte, l Loader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add adds an entry to the path
func (n *Node) Add(ctx context.Context, path, entry []byte, metadata map[string]string, ls LoadSaver) error {
	_ = "STUB: not implemented"
	return nil
}

// empty entry for directories

// check for prefix size limit

// move current common prefix node

// if common path is full path new node is value type

// NOTE: special case on edge split

// add new for shared prefix

func (n *Node) updateIsWithPathSeparator(path []byte) { _ = "STUB: not implemented"; return }

// Remove removes a path from the node
func (n *Node) Remove(ctx context.Context, path []byte, ls LoadSaver) error {
	_ = "STUB: not implemented"
	return nil
}

func common(a, b []byte) (c []byte) { _ = "STUB: not implemented"; return nil }

// HasPrefix tests whether the node contains prefix path.
func (n *Node) HasPrefix(ctx context.Context, path []byte, l Loader) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
