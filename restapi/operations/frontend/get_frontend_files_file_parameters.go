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
	"github.com/go-openapi/swag"
)

// NewGetFrontendFilesFileParams creates a new GetFrontendFilesFileParams object
//
// There are no default values defined in the spec.
func NewGetFrontendFilesFileParams() GetFrontendFilesFileParams {

	return GetFrontendFilesFileParams{}
}

// GetFrontendFilesFileParams contains all the bound params for the get frontend files file operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetFrontendFilesFile
type GetFrontendFilesFileParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Dpr *float64
	/*
	  In: query
	*/
	Format *string
	/*
	  In: query
	*/
	W *float64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFrontendFilesFileParams() beforehand.
func (o *GetFrontendFilesFileParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDpr, qhkDpr, _ := qs.GetOK("dpr")
	if err := o.bindDpr(qDpr, qhkDpr, route.Formats); err != nil {
		res = append(res, err)
	}

	qFormat, qhkFormat, _ := qs.GetOK("format")
	if err := o.bindFormat(qFormat, qhkFormat, route.Formats); err != nil {
		res = append(res, err)
	}

	qW, qhkW, _ := qs.GetOK("w")
	if err := o.bindW(qW, qhkW, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDpr binds and validates parameter Dpr from query.
func (o *GetFrontendFilesFileParams) bindDpr(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("dpr", "query", "float64", raw)
	}
	o.Dpr = &value

	return nil
}

// bindFormat binds and validates parameter Format from query.
func (o *GetFrontendFilesFileParams) bindFormat(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Format = &raw

	return nil
}

// bindW binds and validates parameter W from query.
func (o *GetFrontendFilesFileParams) bindW(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("w", "query", "float64", raw)
	}
	o.W = &value

	return nil
}
