// Copyright 2022 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package roles implements the DocuSign SDK
// category Roles.
//
// Each role is associated with specific permissions. Each new member is assigned a role when you create them, automatically granting them the permissions associated with that role.
//
// Roles use the following permission types to control the actions that users can perform:
//
// - Room
// - Room Details
// - Documents
// - Envelopes
// - Member Management
// - Company Settings
//
// ## Roles in Rooms v5
//
// Rooms v5 has three predefined roles:
//
// - **Admin**: This role can manage all members and company settings.
// - **Manager**: This role is designed to manage rooms in one or more offices or regions. Depending on the permissions that you set for the individual member in the `ClassicManagerPermissions`, managers can also perform additional tasks, such as managing member accounts and company account information.
// - **Agent**: This role is for contributors who only have access to rooms that they create or that are created for them. Agents cannot manage company settings or other members.
//
// However, you can reconfigure the `permissions` for these roles in the `Members` object.
//
// [members]: ./Members. h t m l Fix this when we can link to objects.
// [ClassicMemberPermissions]: ./definition/ClassicManagerPermissions. h t m l
//
// ## Roles in Rooms v6
//
// Rooms v6 enables you to fully configure custom roles containing permissions that make sense for your company. Because each new member is assigned a role, you must set up these roles before you can invite members to join your account.
//
// In Rooms v6, permissions for roles are tied to the `roleId` property and not yet exposed. You can learn more about these permission types and configure them in the console.
//
// ### Internal and External Roles
//
// In Rooms v6, a role can be either internal or external. You assign internal roles to people inside your company. You assign external roles to people outside your company when you invite them to a room.
//
// Each member inside your company has a default company role. However, they can also be assigned additional roles with different permissions on a per-room basis. Regardless of the member's default company role, what they can do in a room is entirely controlled by their role in that particular room.
//
//
// Service Api documentation may be found at:
// https://developers.docusign.com/docs/rooms-api/reference/Roles
// Usage example:
//
//   import (
//       "github.com/jfcote87/esign"
//       "github.com/jfcote87/esign/rooms"
//   )
//   ...
//   rolesService := roles.New(esignCredential)
package roles // import "github.com/jfcote87/esignrooms//roles"

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/rooms"
)

// Service implements DocuSign Roles API operations
type Service struct {
	credential esign.Credential
}

// New initializes a roles service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// CreateRole creates a role.
//
// https://developers.docusign.com/docs/rooms-api/reference/roles/roles/createrole
//
// SDK Method Roles::CreateRole
func (s *Service) CreateRole(body *rooms.RoleForCreate) *CreateRoleOp {
	return &CreateRoleOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "roles",
		Payload:    body,
		Accept:     "application/json-patch+json, application/json, text/json, application/*+json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// CreateRoleOp implements DocuSign API SDK Roles::CreateRole
type CreateRoleOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateRoleOp) Do(ctx context.Context) (*rooms.Role, error) {
	var res *rooms.Role
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// DeleteRole deletes a role.
//
// https://developers.docusign.com/docs/rooms-api/reference/roles/roles/deleterole
//
// SDK Method Roles::DeleteRole
func (s *Service) DeleteRole(roleID string) *DeleteRoleOp {
	return &DeleteRoleOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"roles", roleID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// DeleteRoleOp implements DocuSign API SDK Roles::DeleteRole
type DeleteRoleOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteRoleOp) Do(ctx context.Context) error {
	return ((*esign.Op)(op)).Do(ctx, nil)
}

// GetRole gets a role.
//
// https://developers.docusign.com/docs/rooms-api/reference/roles/roles/getrole
//
// SDK Method Roles::GetRole
func (s *Service) GetRole(roleID string) *GetRoleOp {
	return &GetRoleOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"roles", roleID}, "/"),
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// GetRoleOp implements DocuSign API SDK Roles::GetRole
type GetRoleOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetRoleOp) Do(ctx context.Context) (*rooms.Role, error) {
	var res *rooms.Role
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// IncludeIsAssigned (Optional) When set to **true**, the response includes the `isAssigned` property, which specifies whether the role is currently assigned to any users. The default is **false**.
func (op *GetRoleOp) IncludeIsAssigned() *GetRoleOp {
	if op != nil {
		op.QueryOpts.Set("includeIsAssigned", "true")
	}
	return op
}

// GetRoles gets roles.
//
// https://developers.docusign.com/docs/rooms-api/reference/roles/roles/getroles
//
// SDK Method Roles::GetRoles
func (s *Service) GetRoles() *GetRolesOp {
	return &GetRolesOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "roles",
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// GetRolesOp implements DocuSign API SDK Roles::GetRoles
type GetRolesOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetRolesOp) Do(ctx context.Context) (*rooms.RoleSummaryList, error) {
	var res *rooms.RoleSummaryList
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}

// OnlyAssignable (Optional) When set to **true**, returns only the roles that the current user can assign to someone else. The default value is **false**.
func (op *GetRolesOp) OnlyAssignable() *GetRolesOp {
	if op != nil {
		op.QueryOpts.Set("onlyAssignable", "true")
	}
	return op
}

// Filter (Optional) A search filter that returns roles by the beginning of the role name. You can enter the beginning of the role name only to return all of the roles that begin with the text that you entered.
//
// For example, if your company has set up roles such as Manager Beginner, Manager Pro, Agent Expert, and Agent Superstar, you could enter `Manager` to return all of the Manager roles (Manager Beginner and Manager Pro).
//
// **Note**: You do not enter a wildcard (*) at the end of the name fragment.
func (op *GetRolesOp) Filter(val string) *GetRolesOp {
	if op != nil {
		op.QueryOpts.Set("filter", val)
	}
	return op
}

// StartPosition (Optional) The starting zero-based index position of the result set. The default value is 0.
func (op *GetRolesOp) StartPosition(val int) *GetRolesOp {
	if op != nil {
		op.QueryOpts.Set("startPosition", fmt.Sprintf("%d", val))
	}
	return op
}

// Count (Optional) The number of results to return. This value must be a number between `1` and `100` (default).
func (op *GetRolesOp) Count(val int) *GetRolesOp {
	if op != nil {
		op.QueryOpts.Set("count", fmt.Sprintf("%d", val))
	}
	return op
}

// UpdateRole updates a role.
//
// https://developers.docusign.com/docs/rooms-api/reference/roles/roles/updaterole
//
// SDK Method Roles::UpdateRole
func (s *Service) UpdateRole(roleID string, body *rooms.RoleForUpdate) *UpdateRoleOp {
	return &UpdateRoleOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"roles", roleID}, "/"),
		Payload:    body,
		Accept:     "application/json",
		QueryOpts:  make(url.Values),
		Version:    esign.RoomsV2,
	}
}

// UpdateRoleOp implements DocuSign API SDK Roles::UpdateRole
type UpdateRoleOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateRoleOp) Do(ctx context.Context) (*rooms.Role, error) {
	var res *rooms.Role
	return res, ((*esign.Op)(op)).Do(ctx, &res)
}
