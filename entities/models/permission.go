//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Permission permissions attached to a role.
//
// swagger:model Permission
type Permission struct {

	// allowed actions in weaviate.
	// Required: true
	// Enum: [manage_backups read_cluster create_data read_data update_data delete_data read_nodes create_roles read_roles update_roles delete_roles create_collections read_collections update_collections delete_collections assign_and_revoke_users create_users read_users update_users delete_users create_tenants read_tenants update_tenants delete_tenants create_replicate read_replicate update_replicate delete_replicate create_aliases read_aliases update_aliases delete_aliases]
	Action *string `json:"action"`

	// aliases
	Aliases *PermissionAliases `json:"aliases,omitempty"`

	// backups
	Backups *PermissionBackups `json:"backups,omitempty"`

	// collections
	Collections *PermissionCollections `json:"collections,omitempty"`

	// data
	Data *PermissionData `json:"data,omitempty"`

	// nodes
	Nodes *PermissionNodes `json:"nodes,omitempty"`

	// replicate
	Replicate *PermissionReplicate `json:"replicate,omitempty"`

	// roles
	Roles *PermissionRoles `json:"roles,omitempty"`

	// tenants
	Tenants *PermissionTenants `json:"tenants,omitempty"`

	// users
	Users *PermissionUsers `json:"users,omitempty"`
}

// Validate validates this permission
func (m *Permission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAliases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBackups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenants(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var permissionTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["manage_backups","read_cluster","create_data","read_data","update_data","delete_data","read_nodes","create_roles","read_roles","update_roles","delete_roles","create_collections","read_collections","update_collections","delete_collections","assign_and_revoke_users","create_users","read_users","update_users","delete_users","create_tenants","read_tenants","update_tenants","delete_tenants","create_replicate","read_replicate","update_replicate","delete_replicate","create_aliases","read_aliases","update_aliases","delete_aliases"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		permissionTypeActionPropEnum = append(permissionTypeActionPropEnum, v)
	}
}

const (

	// PermissionActionManageBackups captures enum value "manage_backups"
	PermissionActionManageBackups string = "manage_backups"

	// PermissionActionReadCluster captures enum value "read_cluster"
	PermissionActionReadCluster string = "read_cluster"

	// PermissionActionCreateData captures enum value "create_data"
	PermissionActionCreateData string = "create_data"

	// PermissionActionReadData captures enum value "read_data"
	PermissionActionReadData string = "read_data"

	// PermissionActionUpdateData captures enum value "update_data"
	PermissionActionUpdateData string = "update_data"

	// PermissionActionDeleteData captures enum value "delete_data"
	PermissionActionDeleteData string = "delete_data"

	// PermissionActionReadNodes captures enum value "read_nodes"
	PermissionActionReadNodes string = "read_nodes"

	// PermissionActionCreateRoles captures enum value "create_roles"
	PermissionActionCreateRoles string = "create_roles"

	// PermissionActionReadRoles captures enum value "read_roles"
	PermissionActionReadRoles string = "read_roles"

	// PermissionActionUpdateRoles captures enum value "update_roles"
	PermissionActionUpdateRoles string = "update_roles"

	// PermissionActionDeleteRoles captures enum value "delete_roles"
	PermissionActionDeleteRoles string = "delete_roles"

	// PermissionActionCreateCollections captures enum value "create_collections"
	PermissionActionCreateCollections string = "create_collections"

	// PermissionActionReadCollections captures enum value "read_collections"
	PermissionActionReadCollections string = "read_collections"

	// PermissionActionUpdateCollections captures enum value "update_collections"
	PermissionActionUpdateCollections string = "update_collections"

	// PermissionActionDeleteCollections captures enum value "delete_collections"
	PermissionActionDeleteCollections string = "delete_collections"

	// PermissionActionAssignAndRevokeUsers captures enum value "assign_and_revoke_users"
	PermissionActionAssignAndRevokeUsers string = "assign_and_revoke_users"

	// PermissionActionCreateUsers captures enum value "create_users"
	PermissionActionCreateUsers string = "create_users"

	// PermissionActionReadUsers captures enum value "read_users"
	PermissionActionReadUsers string = "read_users"

	// PermissionActionUpdateUsers captures enum value "update_users"
	PermissionActionUpdateUsers string = "update_users"

	// PermissionActionDeleteUsers captures enum value "delete_users"
	PermissionActionDeleteUsers string = "delete_users"

	// PermissionActionCreateTenants captures enum value "create_tenants"
	PermissionActionCreateTenants string = "create_tenants"

	// PermissionActionReadTenants captures enum value "read_tenants"
	PermissionActionReadTenants string = "read_tenants"

	// PermissionActionUpdateTenants captures enum value "update_tenants"
	PermissionActionUpdateTenants string = "update_tenants"

	// PermissionActionDeleteTenants captures enum value "delete_tenants"
	PermissionActionDeleteTenants string = "delete_tenants"

	// PermissionActionCreateReplicate captures enum value "create_replicate"
	PermissionActionCreateReplicate string = "create_replicate"

	// PermissionActionReadReplicate captures enum value "read_replicate"
	PermissionActionReadReplicate string = "read_replicate"

	// PermissionActionUpdateReplicate captures enum value "update_replicate"
	PermissionActionUpdateReplicate string = "update_replicate"

	// PermissionActionDeleteReplicate captures enum value "delete_replicate"
	PermissionActionDeleteReplicate string = "delete_replicate"

	// PermissionActionCreateAliases captures enum value "create_aliases"
	PermissionActionCreateAliases string = "create_aliases"

	// PermissionActionReadAliases captures enum value "read_aliases"
	PermissionActionReadAliases string = "read_aliases"

	// PermissionActionUpdateAliases captures enum value "update_aliases"
	PermissionActionUpdateAliases string = "update_aliases"

	// PermissionActionDeleteAliases captures enum value "delete_aliases"
	PermissionActionDeleteAliases string = "delete_aliases"
)

// prop value enum
func (m *Permission) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, permissionTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Permission) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *Permission) validateAliases(formats strfmt.Registry) error {
	if swag.IsZero(m.Aliases) { // not required
		return nil
	}

	if m.Aliases != nil {
		if err := m.Aliases.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aliases")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aliases")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) validateBackups(formats strfmt.Registry) error {
	if swag.IsZero(m.Backups) { // not required
		return nil
	}

	if m.Backups != nil {
		if err := m.Backups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backups")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) validateCollections(formats strfmt.Registry) error {
	if swag.IsZero(m.Collections) { // not required
		return nil
	}

	if m.Collections != nil {
		if err := m.Collections.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collections")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collections")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) validateNodes(formats strfmt.Registry) error {
	if swag.IsZero(m.Nodes) { // not required
		return nil
	}

	if m.Nodes != nil {
		if err := m.Nodes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodes")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) validateReplicate(formats strfmt.Registry) error {
	if swag.IsZero(m.Replicate) { // not required
		return nil
	}

	if m.Replicate != nil {
		if err := m.Replicate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replicate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replicate")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	if m.Roles != nil {
		if err := m.Roles.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roles")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) validateTenants(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenants) { // not required
		return nil
	}

	if m.Tenants != nil {
		if err := m.Tenants.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenants")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenants")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) validateUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.Users) { // not required
		return nil
	}

	if m.Users != nil {
		if err := m.Users.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("users")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("users")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this permission based on the context it is used
func (m *Permission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAliases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBackups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReplicate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenants(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Permission) contextValidateAliases(ctx context.Context, formats strfmt.Registry) error {

	if m.Aliases != nil {
		if err := m.Aliases.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aliases")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aliases")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) contextValidateBackups(ctx context.Context, formats strfmt.Registry) error {

	if m.Backups != nil {
		if err := m.Backups.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("backups")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("backups")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) contextValidateCollections(ctx context.Context, formats strfmt.Registry) error {

	if m.Collections != nil {
		if err := m.Collections.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collections")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collections")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) contextValidateNodes(ctx context.Context, formats strfmt.Registry) error {

	if m.Nodes != nil {
		if err := m.Nodes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodes")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) contextValidateReplicate(ctx context.Context, formats strfmt.Registry) error {

	if m.Replicate != nil {
		if err := m.Replicate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replicate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replicate")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) contextValidateRoles(ctx context.Context, formats strfmt.Registry) error {

	if m.Roles != nil {
		if err := m.Roles.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roles")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("roles")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) contextValidateTenants(ctx context.Context, formats strfmt.Registry) error {

	if m.Tenants != nil {
		if err := m.Tenants.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenants")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenants")
			}
			return err
		}
	}

	return nil
}

func (m *Permission) contextValidateUsers(ctx context.Context, formats strfmt.Registry) error {

	if m.Users != nil {
		if err := m.Users.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("users")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("users")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Permission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Permission) UnmarshalBinary(b []byte) error {
	var res Permission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PermissionAliases Resource definition for alias-related actions and permissions. Used to specify which aliases and collections can be accessed or modified.
//
// swagger:model PermissionAliases
type PermissionAliases struct {

	// A string that specifies which aliases this permission applies to. Can be an exact alias name or a regex pattern. The default value `*` applies the permission to all aliases.
	Alias *string `json:"alias,omitempty"`

	// A string that specifies which collections this permission applies to. Can be an exact collection name or a regex pattern. The default value `*` applies the permission to all collections.
	Collection *string `json:"collection,omitempty"`
}

// Validate validates this permission aliases
func (m *PermissionAliases) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this permission aliases based on context it is used
func (m *PermissionAliases) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PermissionAliases) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionAliases) UnmarshalBinary(b []byte) error {
	var res PermissionAliases
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PermissionBackups resources applicable for backup actions
//
// swagger:model PermissionBackups
type PermissionBackups struct {

	// string or regex. if a specific collection name, if left empty it will be ALL or *
	Collection *string `json:"collection,omitempty"`
}

// Validate validates this permission backups
func (m *PermissionBackups) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this permission backups based on context it is used
func (m *PermissionBackups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PermissionBackups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionBackups) UnmarshalBinary(b []byte) error {
	var res PermissionBackups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PermissionCollections resources applicable for collection and/or tenant actions
//
// swagger:model PermissionCollections
type PermissionCollections struct {

	// string or regex. if a specific collection name, if left empty it will be ALL or *
	Collection *string `json:"collection,omitempty"`
}

// Validate validates this permission collections
func (m *PermissionCollections) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this permission collections based on context it is used
func (m *PermissionCollections) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PermissionCollections) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionCollections) UnmarshalBinary(b []byte) error {
	var res PermissionCollections
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PermissionData resources applicable for data actions
//
// swagger:model PermissionData
type PermissionData struct {

	// string or regex. if a specific collection name, if left empty it will be ALL or *
	Collection *string `json:"collection,omitempty"`

	// string or regex. if a specific object ID, if left empty it will be ALL or *
	Object *string `json:"object,omitempty"`

	// string or regex. if a specific tenant name, if left empty it will be ALL or *
	Tenant *string `json:"tenant,omitempty"`
}

// Validate validates this permission data
func (m *PermissionData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this permission data based on context it is used
func (m *PermissionData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PermissionData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionData) UnmarshalBinary(b []byte) error {
	var res PermissionData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PermissionNodes resources applicable for cluster actions
//
// swagger:model PermissionNodes
type PermissionNodes struct {

	// string or regex. if a specific collection name, if left empty it will be ALL or *
	Collection *string `json:"collection,omitempty"`

	// whether to allow (verbose) returning shards and stats data in the response
	// Enum: [verbose minimal]
	Verbosity *string `json:"verbosity,omitempty"`
}

// Validate validates this permission nodes
func (m *PermissionNodes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVerbosity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var permissionNodesTypeVerbosityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["verbose","minimal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		permissionNodesTypeVerbosityPropEnum = append(permissionNodesTypeVerbosityPropEnum, v)
	}
}

const (

	// PermissionNodesVerbosityVerbose captures enum value "verbose"
	PermissionNodesVerbosityVerbose string = "verbose"

	// PermissionNodesVerbosityMinimal captures enum value "minimal"
	PermissionNodesVerbosityMinimal string = "minimal"
)

// prop value enum
func (m *PermissionNodes) validateVerbosityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, permissionNodesTypeVerbosityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PermissionNodes) validateVerbosity(formats strfmt.Registry) error {
	if swag.IsZero(m.Verbosity) { // not required
		return nil
	}

	// value enum
	if err := m.validateVerbosityEnum("nodes"+"."+"verbosity", "body", *m.Verbosity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this permission nodes based on context it is used
func (m *PermissionNodes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PermissionNodes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionNodes) UnmarshalBinary(b []byte) error {
	var res PermissionNodes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PermissionReplicate resources applicable for replicate actions
//
// swagger:model PermissionReplicate
type PermissionReplicate struct {

	// string or regex. if a specific collection name, if left empty it will be ALL or *
	Collection *string `json:"collection,omitempty"`

	// string or regex. if a specific shard name, if left empty it will be ALL or *
	Shard *string `json:"shard,omitempty"`
}

// Validate validates this permission replicate
func (m *PermissionReplicate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this permission replicate based on context it is used
func (m *PermissionReplicate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PermissionReplicate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionReplicate) UnmarshalBinary(b []byte) error {
	var res PermissionReplicate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PermissionRoles resources applicable for role actions
//
// swagger:model PermissionRoles
type PermissionRoles struct {

	// string or regex. if a specific role name, if left empty it will be ALL or *
	Role *string `json:"role,omitempty"`

	// set the scope for the manage role permission
	// Enum: [all match]
	Scope *string `json:"scope,omitempty"`
}

// Validate validates this permission roles
func (m *PermissionRoles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var permissionRolesTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","match"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		permissionRolesTypeScopePropEnum = append(permissionRolesTypeScopePropEnum, v)
	}
}

const (

	// PermissionRolesScopeAll captures enum value "all"
	PermissionRolesScopeAll string = "all"

	// PermissionRolesScopeMatch captures enum value "match"
	PermissionRolesScopeMatch string = "match"
)

// prop value enum
func (m *PermissionRoles) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, permissionRolesTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PermissionRoles) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("roles"+"."+"scope", "body", *m.Scope); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this permission roles based on context it is used
func (m *PermissionRoles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PermissionRoles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionRoles) UnmarshalBinary(b []byte) error {
	var res PermissionRoles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PermissionTenants resources applicable for tenant actions
//
// swagger:model PermissionTenants
type PermissionTenants struct {

	// string or regex. if a specific collection name, if left empty it will be ALL or *
	Collection *string `json:"collection,omitempty"`

	// string or regex. if a specific tenant name, if left empty it will be ALL or *
	Tenant *string `json:"tenant,omitempty"`
}

// Validate validates this permission tenants
func (m *PermissionTenants) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this permission tenants based on context it is used
func (m *PermissionTenants) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PermissionTenants) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionTenants) UnmarshalBinary(b []byte) error {
	var res PermissionTenants
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PermissionUsers resources applicable for user actions
//
// swagger:model PermissionUsers
type PermissionUsers struct {

	// string or regex. if a specific name, if left empty it will be ALL or *
	Users *string `json:"users,omitempty"`
}

// Validate validates this permission users
func (m *PermissionUsers) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this permission users based on context it is used
func (m *PermissionUsers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PermissionUsers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PermissionUsers) UnmarshalBinary(b []byte) error {
	var res PermissionUsers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
