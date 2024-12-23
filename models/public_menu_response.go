// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PublicMenuResponse public menu response
//
// swagger:model PublicMenuResponse
type PublicMenuResponse struct {

	// menu
	Menu *PublicMenu `json:"menu,omitempty"`
}

// Validate validates this public menu response
func (m *PublicMenuResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMenu(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicMenuResponse) validateMenu(formats strfmt.Registry) error {
	if swag.IsZero(m.Menu) { // not required
		return nil
	}

	if m.Menu != nil {
		if err := m.Menu.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("menu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("menu")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this public menu response based on the context it is used
func (m *PublicMenuResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMenu(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicMenuResponse) contextValidateMenu(ctx context.Context, formats strfmt.Registry) error {

	if m.Menu != nil {

		if swag.IsZero(m.Menu) { // not required
			return nil
		}

		if err := m.Menu.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("menu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("menu")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicMenuResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicMenuResponse) UnmarshalBinary(b []byte) error {
	var res PublicMenuResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
