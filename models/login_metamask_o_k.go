// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LoginMetamaskOK login metamask o k
//
// swagger:model LoginMetamaskOK
type LoginMetamaskOK struct {

	// address
	Address string `json:"address,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// nonce
	Nonce string `json:"nonce,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// username
	Username string `json:"username,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this login metamask o k
func (m *LoginMetamaskOK) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this login metamask o k based on context it is used
func (m *LoginMetamaskOK) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoginMetamaskOK) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginMetamaskOK) UnmarshalBinary(b []byte) error {
	var res LoginMetamaskOK
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
