// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PublicSubscribeForm public subscribe form
//
// swagger:model PublicSubscribeForm
type PublicSubscribeForm struct {

	// csrf
	Csrf string `json:"csrf,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`
}

// Validate validates this public subscribe form
func (m *PublicSubscribeForm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this public subscribe form based on context it is used
func (m *PublicSubscribeForm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PublicSubscribeForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicSubscribeForm) UnmarshalBinary(b []byte) error {
	var res PublicSubscribeForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
