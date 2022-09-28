// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package powerforms implements the DocuSign SDK
// category PowerForms.
//
// The PowerForms category enables PowerForms to be created and managed.
//
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/PowerForms
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/v2/model"
//   )
//   ...
//   powerformsService := powerforms.New(esignCredential)
package powerforms // import "github.com/jfcote87/esignv2/powerforms"

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/v2/model"
)

// Service implements DocuSign PowerForms API operations
type Service struct {
	credential esign.Credential
}

// New initializes a powerforms service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// DataList returns the form data associated with the usage of a PowerForm.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/powerforms/powerformdata/list
//
// SDK Method PowerForms::getPowerFormData
func (s *Service) DataList(powerFormID string) *DataListOp {
	return &DataListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"powerforms", powerFormID, "form_data"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// DataListOp implements DocuSign API SDK PowerForms::getPowerFormData
type DataListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DataListOp) Do(ctx context.Context) (*model.PowerFormsFormDataResponse, error) {
	var res *model.PowerFormsFormDataResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// DataLayout set the call query parameter data_layout
func (op *DataListOp) DataLayout(val string) *DataListOp {
	if op != nil {
		op.QueryOpts.Set("data_layout", val)
	}
	return op
}

// FromDate start of the search date range. Only returns templates created on or after this date/time. If no value is specified, there is no limit on the earliest date created.
func (op *DataListOp) FromDate(val time.Time) *DataListOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
	}
	return op
}

// ToDate end of the search date range. Only returns templates created up to this date/time. If no value is provided, this defaults to the current date.
func (op *DataListOp) ToDate(val time.Time) *DataListOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
	}
	return op
}

// Create creates a new PowerForm.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/powerforms/powerforms/create
//
// SDK Method PowerForms::createPowerForm
func (s *Service) Create(powerForm *model.PowerForm) *CreateOp {
	return &CreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "powerforms",
		Payload:    powerForm,
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// CreateOp implements DocuSign API SDK PowerForms::createPowerForm
type CreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateOp) Do(ctx context.Context) (*model.PowerForm, error) {
	var res *model.PowerForm
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Delete delete a PowerForm.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/powerforms/powerforms/delete
//
// SDK Method PowerForms::deletePowerForm
func (s *Service) Delete(powerFormID string) *DeleteOp {
	return &DeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"powerforms", powerFormID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// DeleteOp implements DocuSign API SDK PowerForms::deletePowerForm
type DeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// DeleteList deletes one or more PowerForms
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/powerforms/powerforms/deletelist
//
// SDK Method PowerForms::deletePowerForms
func (s *Service) DeleteList(powerFormsRequest *model.PowerFormsRequest) *DeleteListOp {
	return &DeleteListOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       "powerforms",
		Payload:    powerFormsRequest,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// DeleteListOp implements DocuSign API SDK PowerForms::deletePowerForms
type DeleteListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteListOp) Do(ctx context.Context) (*model.PowerFormsResponse, error) {
	var res *model.PowerFormsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// Get returns a single PowerForm.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/powerforms/powerforms/get
//
// SDK Method PowerForms::getPowerForm
func (s *Service) Get(powerFormID string) *GetOp {
	return &GetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"powerforms", powerFormID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// GetOp implements DocuSign API SDK PowerForms::getPowerForm
type GetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetOp) Do(ctx context.Context) (*model.PowerForm, error) {
	var res *model.PowerForm
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// List returns the list of PowerForms available to the user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/powerforms/powerforms/list
//
// SDK Method PowerForms::listPowerForms
func (s *Service) List() *ListOp {
	return &ListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "powerforms",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ListOp implements DocuSign API SDK PowerForms::listPowerForms
type ListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListOp) Do(ctx context.Context) (*model.PowerFormsResponse, error) {
	var res *model.PowerFormsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// FromDate start of the search date range. Only returns templates created on or after this date/time. If no value is specified, there is no limit on the earliest date created.
func (op *ListOp) FromDate(val time.Time) *ListOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
	}
	return op
}

// Order is an optional value that sets the direction order used to sort the item list.
//
// Valid values are:
//
// * asc = ascending sort order
// * desc = descending sort order
func (op *ListOp) Order(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("order", val)
	}
	return op
}

// OrderBy is an optional value that sets the file attribute used to sort the item list.
//
// Valid values are:
//
// * modified
// * name
func (op *ListOp) OrderBy(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("order_by", val)
	}
	return op
}

// ToDate end of the search date range. Only returns templates created up to this date/time. If no value is provided, this defaults to the current date.
func (op *ListOp) ToDate(val time.Time) *ListOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
	}
	return op
}

// ListSenders returns the list of PowerForms available to the user.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/powerforms/powerforms/listsenders
//
// SDK Method PowerForms::listPowerFormSenders
func (s *Service) ListSenders() *ListSendersOp {
	return &ListSendersOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "powerforms/senders",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// ListSendersOp implements DocuSign API SDK PowerForms::listPowerFormSenders
type ListSendersOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListSendersOp) Do(ctx context.Context) (*model.PowerFormSendersResponse, error) {
	var res *model.PowerFormSendersResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// StartPosition is the position within the total result set from which to start returning values. The value **thumbnail** may be used to return the page image.
func (op *ListSendersOp) StartPosition(val int) *ListSendersOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// Update creates a new PowerForm.
//
// https://developers.docusign.com/docs/esign-rest-api/v2/reference/powerforms/powerforms/update
//
// SDK Method PowerForms::updatePowerForm
func (s *Service) Update(powerFormID string, powerForm *model.PowerForm) *UpdateOp {
	return &UpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"powerforms", powerFormID}, "/"),
		Payload:    powerForm,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.APIv2,
	}
}

// UpdateOp implements DocuSign API SDK PowerForms::updatePowerForm
type UpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateOp) Do(ctx context.Context) (*model.PowerForm, error) {
	var res *model.PowerForm
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
