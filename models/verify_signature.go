// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VerifySignature verify signature
//
// swagger:model VerifySignature
type VerifySignature struct {

	// address
	Address string `json:"address,omitempty"`

	// msg
	Msg string `json:"msg,omitempty"`

	// signature
	Signature string `json:"signature,omitempty"`
}

// Validate validates this verify signature
func (m *VerifySignature) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this verify signature based on context it is used
func (m *VerifySignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VerifySignature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VerifySignature) UnmarshalBinary(b []byte) error {
	var res VerifySignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
