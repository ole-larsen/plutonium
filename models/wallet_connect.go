// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WalletConnect wallet connect
//
// swagger:model WalletConnect
type WalletConnect struct {
	Created     strfmt.Date `json:"created,omitempty"`
	Deleted     strfmt.Date `json:"deleted,omitempty"`
	Updated     strfmt.Date `json:"updated,omitempty"`
	Description string      `json:"description,omitempty"`
	Title       string      `json:"title,omitempty"`
	CreatedByID int64       `json:"created_by_id,omitempty"`
	ID          int64       `json:"id,omitempty"`
	ImageID     int64       `json:"image_id,omitempty"`
	OrderBy     int64       `json:"order_by,omitempty"`
	UpdatedByID int64       `json:"updated_by_id,omitempty"`
	Enabled     bool        `json:"enabled,omitempty"`
}

// Validate validates this wallet connect
func (m *WalletConnect) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WalletConnect) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WalletConnect) validateDeleted(formats strfmt.Registry) error {
	if swag.IsZero(m.Deleted) { // not required
		return nil
	}

	if err := validate.FormatOf("deleted", "body", "date", m.Deleted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WalletConnect) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wallet connect based on context it is used
func (m *WalletConnect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WalletConnect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WalletConnect) UnmarshalBinary(b []byte) error {
	var res WalletConnect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
