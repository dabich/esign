// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package roomfolders implements the DocuSign SDK
// category RoomFolders.
//
//
//
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/rooms-api/reference/RoomFolders
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/rooms"
//   )
//   ...
//   roomfoldersService := roomfolders.New(esignCredential)
package roomfolders // import "github.com/jfcote87/esignrooms//roomfolders"

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/rooms"
)

// Service implements DocuSign RoomFolders API operations
type Service struct {
	credential esign.Credential
}

// New initializes a roomfolders service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// GetRoomFolders gets a list of room folders accessible to the current user.
//
// https://developers.docusign.com/docs/rooms-api/reference/roomfolders/roomfolders/getroomfolders
//
// SDK Method RoomFolders::GetRoomFolders
func (s *Service) GetRoomFolders(roomID string) *GetRoomFoldersOp {
	return &GetRoomFoldersOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"rooms", roomID, "room_folders"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// GetRoomFoldersOp implements DocuSign API SDK RoomFolders::GetRoomFolders
type GetRoomFoldersOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetRoomFoldersOp) Do(ctx context.Context) (*rooms.RoomFolderList, error) {
	var res *rooms.RoomFolderList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// StartPosition position of the first item in the total results. Defaults to 0.
func (op *GetRoomFoldersOp) StartPosition(val int) *GetRoomFoldersOp {
	if op != nil {
		op.QueryOpts.Set("startPosition", fmt.Sprintf("%d", val))
	}
	return op
}

// Count number of room folders to return. Defaults to the maximum which is 100.
func (op *GetRoomFoldersOp) Count(val int) *GetRoomFoldersOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}
