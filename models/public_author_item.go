// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PublicAuthorItem public author item
//
// swagger:model PublicAuthorItem
type PublicAuthorItem struct {
	Image       *PublicFile     `json:"image,omitempty"`
	BtnLink     string          `json:"btnLink,omitempty"`
	BtnText     string          `json:"btnText,omitempty"`
	Description string          `json:"description,omitempty"`
	Link        string          `json:"link,omitempty"`
	Name        string          `json:"name,omitempty"`
	Title       string          `json:"title,omitempty"`
	Socials     []*PublicSocial `json:"socials"`
	Wallets     []*PublicWallet `json:"wallets"`
	ID          int64           `json:"id,omitempty"`
}

// Validate validates this public author item
func (m *PublicAuthorItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSocials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWallets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicAuthorItem) validateImage(formats strfmt.Registry) error {
	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *PublicAuthorItem) validateSocials(formats strfmt.Registry) error {
	if swag.IsZero(m.Socials) { // not required
		return nil
	}

	for i := 0; i < len(m.Socials); i++ {
		if swag.IsZero(m.Socials[i]) { // not required
			continue
		}

		if m.Socials[i] != nil {
			if err := m.Socials[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("socials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("socials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PublicAuthorItem) validateWallets(formats strfmt.Registry) error {
	if swag.IsZero(m.Wallets) { // not required
		return nil
	}

	for i := 0; i < len(m.Wallets); i++ {
		if swag.IsZero(m.Wallets[i]) { // not required
			continue
		}

		if m.Wallets[i] != nil {
			if err := m.Wallets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wallets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wallets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this public author item based on the context it is used
func (m *PublicAuthorItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSocials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWallets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicAuthorItem) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

	if m.Image != nil {

		if swag.IsZero(m.Image) { // not required
			return nil
		}

		if err := m.Image.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *PublicAuthorItem) contextValidateSocials(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Socials); i++ {

		if m.Socials[i] != nil {

			if swag.IsZero(m.Socials[i]) { // not required
				return nil
			}

			if err := m.Socials[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("socials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("socials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PublicAuthorItem) contextValidateWallets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Wallets); i++ {

		if m.Wallets[i] != nil {

			if swag.IsZero(m.Wallets[i]) { // not required
				return nil
			}

			if err := m.Wallets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wallets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wallets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicAuthorItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicAuthorItem) UnmarshalBinary(b []byte) error {
	var res PublicAuthorItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
