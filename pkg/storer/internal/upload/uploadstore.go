// Copyright 2022 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upload

import (
	"context"
	"errors"
	"time"

	storage "github.com/ethersphere/bee/v2/pkg/storage"
	"github.com/ethersphere/bee/v2/pkg/storer/internal"
	"github.com/ethersphere/bee/v2/pkg/storer/internal/transaction"
	"github.com/ethersphere/bee/v2/pkg/swarm"
)

// now returns the current time.Time; used in testing.
var now = time.Now

var (
	// errPushItemMarshalAddressIsZero is returned when trying
	// to marshal a pushItem with an address that is zero.
	errPushItemMarshalAddressIsZero = errors.New("marshal pushItem: address is zero")
	// errPushItemMarshalBatchInvalid is returned when trying to
	// marshal a pushItem with invalid batch
	errPushItemMarshalBatchInvalid = errors.New("marshal pushItem: batch is invalid")
	// errPushItemUnmarshalInvalidSize is returned when trying
	// to unmarshal buffer that is not of size pushItemSize.
	errPushItemUnmarshalInvalidSize = errors.New("unmarshal pushItem: invalid size")
)

// pushItemSize is the size of a marshaled pushItem.
const pushItemSize = 8 + 2*swarm.HashSize + 8

const uploadScope = "upload"

var _ storage.Item = (*pushItem)(nil)

// pushItem is an store.Item that represents data relevant to push.
// The key is a combination of Timestamp, Address and postage stamp, where the
// Timestamp provides an order to iterate.
type pushItem struct {
	Timestamp int64
	Address   swarm.Address
	BatchID   []byte
	TagID     uint64
}

// ID implements the storage.Item interface.
func (i pushItem) ID() string { _ = "STUB: not implemented"; return "" }

// Namespace implements the storage.Item interface.
func (i pushItem) Namespace() string {
	_ = "STUB: not implemented"

	// Marshal implements the storage.Item interface.
	// If the Address is zero, an error is returned.
	return ""
}

func (i pushItem) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal implements the storage.Item interface.
// If the buffer is not of size pushItemSize, an error is returned.
func (i *pushItem) Unmarshal(bytes []byte) error { _ = "STUB: not implemented"; return nil }

// Clone implements the storage.Item interface.
func (i *pushItem) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

// String implements the fmt.Stringer interface.
func (i pushItem) String() string { _ = "STUB: not implemented"; return "" }

// errTagIDAddressItemUnmarshalInvalidSize is returned when trying
// to unmarshal buffer that is not of size tagItemSize.
var errTagItemUnmarshalInvalidSize = errors.New("unmarshal TagItem: invalid size")

// tagItemSize is the size of a marshaled TagItem.
const tagItemSize = swarm.HashSize + 7*8

var _ storage.Item = (*TagItem)(nil)

// TagItem is an store.Item that stores information about a session of upload.
type TagItem struct {
	TagID     uint64        // unique identifier for the tag
	Split     uint64        // total no of chunks processed by the splitter for hashing
	Seen      uint64        // total no of chunks already seen
	Stored    uint64        // total no of chunks stored locally on the node
	Sent      uint64        // total no of chunks sent to the neighbourhood
	Synced    uint64        // total no of chunks synced with proof
	Address   swarm.Address // swarm.Address associated with this tag
	StartedAt int64         // start timestamp
}

// ID implements the storage.Item interface.
func (i TagItem) ID() string { _ = "STUB: not implemented"; return "" }

// Namespace implements the storage.Item interface.
func (i TagItem) Namespace() string {
	_ = "STUB: not implemented"

	// Marshal implements the storage.Item interface.
	return ""
}

func (i TagItem) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// in case of encrypted reference we use the swarm hash as the address and
// avoid storing the encryption key

// Unmarshal implements the storage.Item interface.
// If the buffer is not of size tagItemSize, an error is returned.
func (i *TagItem) Unmarshal(bytes []byte) error { _ = "STUB: not implemented"; return nil }

// Clone implements the storage.Item interface.
func (i *TagItem) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

// String implements the fmt.Stringer interface.
func (i TagItem) String() string { _ = "STUB: not implemented"; return "" }

var (
	// errUploadItemMarshalAddressIsZero is returned when trying
	// to marshal a uploadItem with an address that is zero.
	errUploadItemMarshalAddressIsZero = errors.New("marshal uploadItem: address is zero")
	// errUploadItemMarshalBatchInvalid is returned when trying to
	// marshal a uploadItem with invalid batch
	errUploadItemMarshalBatchInvalid = errors.New("marshal uploadItem: batch is invalid")
	// errTagIDAddressItemUnmarshalInvalidSize is returned when trying
	// to unmarshal buffer that is not of size uploadItemSize.
	errUploadItemUnmarshalInvalidSize = errors.New("unmarshal uploadItem: invalid size")
)

// uploadItemSize is the size of a marshaled uploadItem.
const uploadItemSize = 3 * 8

var _ storage.Item = (*uploadItem)(nil)

// uploadItem is an store.Item that stores addresses of already seen chunks.
type uploadItem struct {
	Address  swarm.Address
	BatchID  []byte
	TagID    uint64
	Uploaded int64
	Synced   int64

	// IdFunc overrides the ID method.
	// This used to get the ID from the item where the address and batchID were not marshalled.
	IdFunc func() string
}

// ID implements the storage.Item interface.
func (i uploadItem) ID() string { _ = "STUB: not implemented"; return "" }

// Namespace implements the storage.Item interface.
func (i uploadItem) Namespace() string { _ = "STUB: not implemented"; return "" }

// Marshal implements the storage.Item interface.
// If the Address is zero, an error is returned.
func (i uploadItem) Marshal() ([]byte, error) {
	_ = "STUB: not implemented"
	// Address and BatchID are not part of the marshaled payload. But they are used
	// in they key and hence are required. The Marshaling is done when item is to
	// be stored, so we return errors for these cases.
	return nil, nil
}

// Unmarshal implements the storage.Item interface.
// If the buffer is not of size pushItemSize, an error is returned.
func (i *uploadItem) Unmarshal(bytes []byte) error { _ = "STUB: not implemented"; return nil }

// The Address and BatchID are required for the key, so it is assumed that
// they will be filled already. We reuse them during unmarshaling.

// Clone implements the storage.Item interface.
func (i *uploadItem) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

// String implements the fmt.Stringer interface.
func (i uploadItem) String() string { _ = "STUB: not implemented"; return "" }

// dirtyTagItemUnmarshalInvalidSize is returned when trying
// to unmarshal buffer that is not of size dirtyTagItemSize.
var errDirtyTagItemUnmarshalInvalidSize = errors.New("unmarshal dirtyTagItem: invalid size")

// dirtyTagItemSize is the size of a marshaled dirtyTagItem.
const dirtyTagItemSize = 8 + 8

type dirtyTagItem struct {
	TagID   uint64
	Started int64
}

// ID implements the storage.Item interface.
func (i dirtyTagItem) ID() string { _ = "STUB: not implemented"; return "" }

// Namespace implements the storage.Item interface.
func (i dirtyTagItem) Namespace() string { _ = "STUB: not implemented"; return "" }

// Marshal implements the storage.Item interface.
func (i dirtyTagItem) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal implements the storage.Item interface.
func (i *dirtyTagItem) Unmarshal(bytes []byte) error { _ = "STUB: not implemented"; return nil }

// Clone implements the storage.Item interface.
func (i *dirtyTagItem) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

// String implements the fmt.Stringer interface.
func (i dirtyTagItem) String() string { _ = "STUB: not implemented"; return "" }

var (
	// errPutterAlreadyClosed is returned when trying to Put a new chunk
	// after the putter has been closed.
	errPutterAlreadyClosed = errors.New("upload store: putter already closed")

	// errOverwriteOfImmutableBatch is returned when stamp index already
	// exists and the batch is immutable.
	errOverwriteOfImmutableBatch = errors.New("upload store: overwrite of existing immutable batch")

	// errOverwriteOfNewerBatch is returned if a stamp index already exists
	// and the existing chunk with the same stamp index has a newer timestamp.
	errOverwriteOfNewerBatch = errors.New("upload store: overwrite of existing batch with newer timestamp")
)

type uploadPutter struct {
	tagID  uint64
	split  uint64
	seen   uint64
	closed bool
}

// NewPutter returns a new chunk putter associated with the tagID.
// Calls to the Putter must be mutex locked to prevent concurrent upload data races.
func NewPutter(s storage.IndexStore, tagID uint64) (internal.PutterCloserWithReference, error) {
	_ = "STUB: not implemented"
	return *new(internal.PutterCloserWithReference), nil
}

// Put operation will do the following:
// 1.If upload store has already seen this chunk, it will update the tag and return
// 2.For a new chunk it will add:
// - uploadItem entry to keep track of this chunk.
// - pushItem entry to make it available for PushSubscriber
// - add chunk to the chunkstore till it is synced
// The user of the putter MUST mutex lock the call to prevent data-races across multiple upload sessions.
func (u *uploadPutter) Put(ctx context.Context, st transaction.Store, chunk swarm.Chunk) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if upload store has already seen this chunk

// Close provides the CloseWithReference interface where the session can be associated
// with a swarm reference. This can be useful while keeping track of uploads through
// the tags. It will update the tag. This will be filled with the Split and Seen count
// by the Putter.
func (u *uploadPutter) Close(s storage.IndexStore, addr swarm.Address) error {
	_ = "STUB: not implemented"
	return nil
}

// If the tag is not found, it might have been removed or never existed.
// In this case, there’s no need to update or delete it—so simply return.

func (u *uploadPutter) Cleanup(st transaction.Storage) error { _ = "STUB: not implemented"; return nil }

// CleanupDirty does a best-effort cleanup of dirty tags. This is called on startup.
func CleanupDirty(st transaction.Storage) error { _ = "STUB: not implemented"; return nil }

// Report is the implementation of the PushReporter interface.
func Report(ctx context.Context, st transaction.Store, chunk swarm.Chunk, state storage.ChunkState) error {
	_ = "STUB: not implemented"
	return nil
}

// because of the nature of the feed mechanism of the uploadstore/pusher, a chunk that is in inflight may be sent more than once to the pusher.
// this is because the chunks are removed from the queue only when they are synced, not at the start of the upload

// Once the chunk is stored/synced/failed to sync, it is deleted from the upload store as
// we no longer need to keep track of this chunk. We also need to cleanup
// the pushItem.

// tag is missing, no need update it

// update the tag

var errNextTagIDUnmarshalInvalidSize = errors.New("unmarshal nextTagID: invalid size")

// nextTagID is a storage.Item which stores a uint64 value in the store.
type nextTagID uint64

func (nextTagID) Namespace() string { _ = "STUB: not implemented"; return "" }

func (nextTagID) ID() string { _ = "STUB: not implemented"; return "" }

func (n nextTagID) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (n *nextTagID) Unmarshal(buf []byte) error { _ = "STUB: not implemented"; return nil }

func (n *nextTagID) Clone() storage.Item { _ = "STUB: not implemented"; return *new(storage.Item) }

func (n nextTagID) String() string { _ = "STUB: not implemented"; return "" }

// NextTag returns the next tag ID to be used. It reads the last used ID and
// increments it by 1. This method needs to be called under lock by user as there
// is no guarantee for parallel updates.
func NextTag(st storage.IndexStore) (TagItem, error) {
	_ = "STUB: not implemented"
	return *new(TagItem), nil
}

// TagInfo returns the TagItem for this particular tagID.
func TagInfo(st storage.Reader, tagID uint64) (TagItem, error) {
	_ = "STUB: not implemented"
	return *new(TagItem), nil
}

// ListAllTags returns all the TagItems in the store.
func ListAllTags(st storage.Reader) ([]TagItem, error) { _ = "STUB: not implemented"; return nil, nil }

func IteratePending(ctx context.Context, s transaction.ReadOnlyStore, consumerFn func(chunk swarm.Chunk) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteTag deletes TagItem associated with the given tagID.
func DeleteTag(st storage.Writer, tagID uint64) error { _ = "STUB: not implemented"; return nil }

func IterateAll(st storage.Reader, iterateFn func(item storage.Item) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

func IterateAllTagItems(st storage.Reader, cb func(ti *TagItem) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// BatchIDForChunk returns the first known batchID for the given chunk address.
func BatchIDForChunk(st storage.Reader, addr swarm.Address) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
