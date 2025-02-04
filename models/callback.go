// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	timeext "time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Callback callback
//
// swagger:model Callback
type Callback struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// code
	Code string `json:"code,omitempty"`

	// expiry
	Expiry timeext.Time `json:"expiry,omitempty"`

	// original Url
	OriginalURL string `json:"originalUrl,omitempty"`

	// refresh token
	RefreshToken string `json:"refresh_token,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// token type
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this callback
func (m *Callback) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this callback based on context it is used
func (m *Callback) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Callback) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Callback) UnmarshalBinary(b []byte) error {
	var res Callback
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
