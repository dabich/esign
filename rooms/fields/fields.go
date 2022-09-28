// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package fields implements the DocuSign SDK
// category Fields.
//
// Information about field sets.
//
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/rooms-api/reference/Fields
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/rooms"
//   )
//   ...
//   fieldsService := fields.New(esignCredential)
package fields // import "github.com/jfcote87/esignrooms//fields"

import (
	"context"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/rooms"
)

// Service implements DocuSign Fields API operations
type Service struct {
	credential esign.Credential
}

// New initializes a fields service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// GetFieldSet gets a field set.
//
// https://developers.docusign.com/docs/rooms-api/reference/fields/fields/getfieldset
//
// SDK Method Fields::GetFieldSet
func (s *Service) GetFieldSet(fieldSetID string) *GetFieldSetOp {
	return &GetFieldSetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"field_sets", fieldSetID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// GetFieldSetOp implements DocuSign API SDK Fields::GetFieldSet
type GetFieldSetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetFieldSetOp) Do(ctx context.Context) (*rooms.FieldSet, error) {
	var res *rooms.FieldSet
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// FieldsCustomDataFilters (Optional) An comma-separated list that limits the fields to return:
//
// - `IsRequiredOnCreate`: include fields that are required in room creation.
// - `IsRequiredOnSubmit`: include fields that are required when submitting a room for review.
func (op *GetFieldSetOp) FieldsCustomDataFilters(val ...string) *GetFieldSetOp {
	if op != nil {
		op.QueryOpts.Set("fieldsCustomDataFilters", strings.Join(val, ","))
	}
	return op
}
