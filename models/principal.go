// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// Principal A unique identifier for a principal (user or entity).
//
// swagger:model principal
type Principal string

// Validate validates this principal
func (m Principal) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this principal based on context it is used
func (m Principal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
