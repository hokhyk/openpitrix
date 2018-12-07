// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixGetAttachmentsResponse openpitrix get attachments response
// swagger:model openpitrixGetAttachmentsResponse
type OpenpitrixGetAttachmentsResponse struct {

	// attachments
	Attachments OpenpitrixGetAttachmentsResponseAttachments `json:"attachments,omitempty"`
}

// Validate validates this openpitrix get attachments response
func (m *OpenpitrixGetAttachmentsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixGetAttachmentsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixGetAttachmentsResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixGetAttachmentsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}