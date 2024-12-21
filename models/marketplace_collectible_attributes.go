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

// MarketplaceCollectibleAttributes marketplace collectible attributes
//
// swagger:model MarketplaceCollectibleAttributes
type MarketplaceCollectibleAttributes struct {
	Creator      *PublicUser                     `json:"creator,omitempty"`
	Details      *MarketplaceCollectibleDetails  `json:"details,omitempty"`
	Metadata     *MarketplaceCollectibleMetadata `json:"metadata,omitempty"`
	Owner        *PublicUser                     `json:"owner,omitempty"`
	URI          string                          `json:"uri,omitempty"`
	TokenIds     []int64                         `json:"tokenIds"`
	CollectionID int64                           `json:"collectionId,omitempty"`
	ItemID       int64                           `json:"itemId,omitempty"`
}

// Validate validates this marketplace collectible attributes
func (m *MarketplaceCollectibleAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketplaceCollectibleAttributes) validateCreator(formats strfmt.Registry) error {
	if swag.IsZero(m.Creator) { // not required
		return nil
	}

	if m.Creator != nil {
		if err := m.Creator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creator")
			}
			return err
		}
	}

	return nil
}

func (m *MarketplaceCollectibleAttributes) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if m.Details != nil {
		if err := m.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *MarketplaceCollectibleAttributes) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *MarketplaceCollectibleAttributes) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this marketplace collectible attributes based on the context it is used
func (m *MarketplaceCollectibleAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketplaceCollectibleAttributes) contextValidateCreator(ctx context.Context, formats strfmt.Registry) error {

	if m.Creator != nil {

		if swag.IsZero(m.Creator) { // not required
			return nil
		}

		if err := m.Creator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creator")
			}
			return err
		}
	}

	return nil
}

func (m *MarketplaceCollectibleAttributes) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.Details != nil {

		if swag.IsZero(m.Details) { // not required
			return nil
		}

		if err := m.Details.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *MarketplaceCollectibleAttributes) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *MarketplaceCollectibleAttributes) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {

		if swag.IsZero(m.Owner) { // not required
			return nil
		}

		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MarketplaceCollectibleAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketplaceCollectibleAttributes) UnmarshalBinary(b []byte) error {
	var res MarketplaceCollectibleAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}