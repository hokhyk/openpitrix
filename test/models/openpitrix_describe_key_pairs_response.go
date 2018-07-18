// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDescribeKeyPairsResponse openpitrix describe key pairs response
// swagger:model openpitrixDescribeKeyPairsResponse
type OpenpitrixDescribeKeyPairsResponse struct {

	// key pair set
	KeyPairSet OpenpitrixDescribeKeyPairsResponseKeyPairSet `json:"key_pair_set"`

	// total count
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this openpitrix describe key pairs response
func (m *OpenpitrixDescribeKeyPairsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDescribeKeyPairsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDescribeKeyPairsResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDescribeKeyPairsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
