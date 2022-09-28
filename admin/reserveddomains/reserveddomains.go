// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package reserveddomains implements the DocuSign SDK
// category ReservedDomains.
//
// Methods to get a list of reserved domains.
//
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/admin-api/reference/ReservedDomains
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/admin"
//   )
//   ...
//   reserveddomainsService := reserveddomains.New(esignCredential)
package reserveddomains // import "github.com/jfcote87/esignadmin/reserveddomains"

import (
	"context"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/admin"
)

// Service implements DocuSign ReservedDomains API operations
type Service struct {
	credential esign.Credential
}

// New initializes a reserveddomains service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// GetReservedDomains returns the list of reserved domains for the organization.
//
// https://developers.docusign.com/docs/admin-api/reference/reserveddomains/reserveddomains/getreserveddomains
//
// SDK Method ReservedDomains::getReservedDomains
func (s *Service) GetReservedDomains(organizationID string) *GetReservedDomainsOp {
	return &GetReservedDomainsOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"", "v2", "organizations", organizationID, "reserved_domains"}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.AdminV2,
	}
}

// GetReservedDomainsOp implements DocuSign API SDK ReservedDomains::getReservedDomains
type GetReservedDomainsOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetReservedDomainsOp) Do(ctx context.Context) (*admin.DomainsResponse, error) {
	var res *admin.DomainsResponse
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
