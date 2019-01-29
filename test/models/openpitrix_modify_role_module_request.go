// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyRoleModuleRequest openpitrix modify role module request
// swagger:model openpitrixModifyRoleModuleRequest
type OpenpitrixModifyRoleModuleRequest struct {

	// role module
	RoleModule *OpenpitrixRoleModule `json:"role_module,omitempty"`
}

// Validate validates this openpitrix modify role module request
func (m *OpenpitrixModifyRoleModuleRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoleModule(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixModifyRoleModuleRequest) validateRoleModule(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleModule) { // not required
		return nil
	}

	if m.RoleModule != nil {

		if err := m.RoleModule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role_module")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyRoleModuleRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyRoleModuleRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyRoleModuleRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}