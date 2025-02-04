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

// PublicUser public user
//
// swagger:model PublicUser
type PublicUser struct {

	// attributes
	Attributes *PublicUserAttributes `json:"attributes,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this public user
func (m *PublicUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicUser) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this public user based on the context it is used
func (m *PublicUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicUser) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	if m.Attributes != nil {

		if swag.IsZero(m.Attributes) { // not required
			return nil
		}

		if err := m.Attributes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicUser) UnmarshalBinary(b []byte) error {
	var res PublicUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PublicUserAttributes public user attributes
//
// swagger:model PublicUserAttributes
type PublicUserAttributes struct {
	Wallpaper   *PublicFile `json:"wallpaper,omitempty"`
	Address     string      `json:"address,omitempty"`
	Created     string      `json:"created,omitempty"`
	Email       string      `json:"email,omitempty"`
	Funds       string      `json:"funds,omitempty"`
	Gravatar    string      `json:"gravatar,omitempty"`
	Nonce       string      `json:"nonce,omitempty"`
	Token       string      `json:"token,omitempty"`
	Username    string      `json:"username,omitempty"`
	UUID        string      `json:"uuid,omitempty"`
	WallpaperID int64       `json:"wallpaperId,omitempty"`
}

// Validate validates this public user attributes
func (m *PublicUserAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWallpaper(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicUserAttributes) validateWallpaper(formats strfmt.Registry) error {
	if swag.IsZero(m.Wallpaper) { // not required
		return nil
	}

	if m.Wallpaper != nil {
		if err := m.Wallpaper.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "wallpaper")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "wallpaper")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this public user attributes based on the context it is used
func (m *PublicUserAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWallpaper(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicUserAttributes) contextValidateWallpaper(ctx context.Context, formats strfmt.Registry) error {

	if m.Wallpaper != nil {

		if swag.IsZero(m.Wallpaper) { // not required
			return nil
		}

		if err := m.Wallpaper.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes" + "." + "wallpaper")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attributes" + "." + "wallpaper")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicUserAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicUserAttributes) UnmarshalBinary(b []byte) error {
	var res PublicUserAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
