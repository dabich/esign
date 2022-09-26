// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package esignpermissionprofiles implements the DocuSign SDK
// category ESignPermissionProfiles.
// 
// When you create or invite a new member in Rooms, the system creates an eSignature account for the member at the same time. This section shows you how to retrieve a list of eSignature permission profiles that the current user can assign to a new member.
// 
// You can use the API only to retrieve and assign eSignature permission profiles. You create and modify permission profiles in the DocuSign console.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/rooms-api/reference/ESignPermissionProfiles
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/rooms"
//   ) 
//   ...
//   esignpermissionprofilesService := esignpermissionprofiles.New(esignCredential)
package esignpermissionprofiles // import "github.com/jfcote87/esignrooms//esignpermissionprofiles"

import (
    "context"
    "net/url"
    
    "github.com/jfcote87/esign"
    "github.com/jfcote87/esign/rooms"
)

// Service implements DocuSign ESignPermissionProfiles Category API operations
type Service struct {
    credential esign.Credential 
}

// New initializes a esignpermissionprofiles service using cred to authorize ops.
func New(cred esign.Credential) *Service {
    return &Service{credential: cred}
}

// GetESignPermissionProfiles gets eSignature Permission Profiles.
// 
// https://developers.docusign.com/docs/rooms-api/reference/esignpermissionprofiles/esignpermissionprofiles/getesignpermissionprofiles
// 
// SDK Method ESignPermissionProfiles::GetESignPermissionProfiles
func (s *Service) GetESignPermissionProfiles() *GetESignPermissionProfilesOp {
    return &GetESignPermissionProfilesOp{
        Credential: s.credential,
        Method:  "GET",
        Path: "esign_permission_profiles",
        Accept: "application/json",
        QueryOpts: make(url.Values),
        Version: esign.RoomsV2,
    }
}

// GetESignPermissionProfilesOp implements DocuSign API SDK ESignPermissionProfiles::GetESignPermissionProfiles
type GetESignPermissionProfilesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetESignPermissionProfilesOp) Do(ctx context.Context)  (*rooms.ESignPermissionProfileList, error) {
    var res *rooms.ESignPermissionProfileList
    return res, ((*esign.Op)(op)).Do(ctx, &res)
}
