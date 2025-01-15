// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetFrontendAuthCallbackParams creates a new GetFrontendAuthCallbackParams object
//
// There are no default values defined in the spec.
func NewGetFrontendAuthCallbackParams() GetFrontendAuthCallbackParams {

	return GetFrontendAuthCallbackParams{}
}

// GetFrontendAuthCallbackParams contains all the bound params for the get frontend auth callback operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetFrontendAuthCallback
type GetFrontendAuthCallbackParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*code
	  In: query
	*/
	Code *string
	/*state
	  In: query
	*/
	State *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFrontendAuthCallbackParams() beforehand.
func (o *GetFrontendAuthCallbackParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCode, qhkCode, _ := qs.GetOK("code")
	if err := o.bindCode(qCode, qhkCode, route.Formats); err != nil {
		res = append(res, err)
	}

	qState, qhkState, _ := qs.GetOK("state")
	if err := o.bindState(qState, qhkState, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCode binds and validates parameter Code from query.
func (o *GetFrontendAuthCallbackParams) bindCode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Code = &raw

	return nil
}

// bindState binds and validates parameter State from query.
func (o *GetFrontendAuthCallbackParams) bindState(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.State = &raw

	return nil
}
