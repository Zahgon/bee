// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mantaray

import (
	"context"
)

// WalkNodeFunc is the type of the function called for each node visited
// by WalkNode.
type WalkNodeFunc func(path []byte, node *Node, err error) error

func walkNodeFnCopyBytes(path []byte, node *Node, walkFn WalkNodeFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// walkNode recursively descends path, calling walkFn.
func walkNode(ctx context.Context, path []byte, l Loader, n *Node, walkFn WalkNodeFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// WalkNode walks the node tree structure rooted at root, calling walkFn for
// each node in the tree, including root. All errors that arise visiting nodes
// are filtered by walkFn.
func (n *Node) WalkNode(ctx context.Context, root []byte, l Loader, walkFn WalkNodeFunc) error {
	_ = "STUB: not implemented"
	return nil
}
