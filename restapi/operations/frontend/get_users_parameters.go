// Code generated by go-swagger; DO NOT EDIT.

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetUsersParams creates a new GetUsersParams object
//
// There are no default values defined in the spec.
func NewGetUsersParams() GetUsersParams {

	return GetUsersParams{}
}

// GetUsersParams contains all the bound params for the get users operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetUsers
type GetUsersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Address *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetUsersParams() beforehand.
func (o *GetUsersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAddress, qhkAddress, _ := qs.GetOK("address")
	if err := o.bindAddress(qAddress, qhkAddress, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAddress binds and validates parameter Address from query.
func (o *GetUsersParams) bindAddress(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Address = &raw

	return nil
}
