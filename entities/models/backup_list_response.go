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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BackupListResponse The definition of a backup create response body
//
// swagger:model BackupListResponse
type BackupListResponse []*BackupListResponseItems0

// Validate validates this backup list response
func (m BackupListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this backup list response based on the context it is used
func (m BackupListResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if m[i] != nil {
			if err := m[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// BackupListResponseItems0 backup list response items0
//
// swagger:model BackupListResponseItems0
type BackupListResponseItems0 struct {

	// The list of classes for which the existed backup process
	Classes []string `json:"classes"`

	// The ID of the backup. Must be URL-safe and work as a filesystem path, only lowercase, numbers, underscore, minus characters allowed.
	ID string `json:"id,omitempty"`

	// destination path of backup files proper to selected backend
	Path string `json:"path,omitempty"`

	// status of backup process
	// Enum: [STARTED TRANSFERRING TRANSFERRED SUCCESS FAILED CANCELED]
	Status string `json:"status,omitempty"`
}

// Validate validates this backup list response items0
func (m *BackupListResponseItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var backupListResponseItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STARTED","TRANSFERRING","TRANSFERRED","SUCCESS","FAILED","CANCELED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupListResponseItems0TypeStatusPropEnum = append(backupListResponseItems0TypeStatusPropEnum, v)
	}
}

const (

	// BackupListResponseItems0StatusSTARTED captures enum value "STARTED"
	BackupListResponseItems0StatusSTARTED string = "STARTED"

	// BackupListResponseItems0StatusTRANSFERRING captures enum value "TRANSFERRING"
	BackupListResponseItems0StatusTRANSFERRING string = "TRANSFERRING"

	// BackupListResponseItems0StatusTRANSFERRED captures enum value "TRANSFERRED"
	BackupListResponseItems0StatusTRANSFERRED string = "TRANSFERRED"

	// BackupListResponseItems0StatusSUCCESS captures enum value "SUCCESS"
	BackupListResponseItems0StatusSUCCESS string = "SUCCESS"

	// BackupListResponseItems0StatusFAILED captures enum value "FAILED"
	BackupListResponseItems0StatusFAILED string = "FAILED"

	// BackupListResponseItems0StatusCANCELED captures enum value "CANCELED"
	BackupListResponseItems0StatusCANCELED string = "CANCELED"
)

// prop value enum
func (m *BackupListResponseItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, backupListResponseItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BackupListResponseItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup list response items0 based on context it is used
func (m *BackupListResponseItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupListResponseItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupListResponseItems0) UnmarshalBinary(b []byte) error {
	var res BackupListResponseItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
