// Code generated by go-swagger; DO NOT EDIT.

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ole-larsen/plutonium/models"
)

// GetFrontendMenuOKCode is the HTTP code returned for type GetFrontendMenuOK
const GetFrontendMenuOKCode int = 200

/*
GetFrontendMenuOK Successfully fetched the menu.

swagger:response getFrontendMenuOK
*/
type GetFrontendMenuOK struct {

	/*
	  In: Body
	*/
	Payload *models.PublicMenuResponse `json:"body,omitempty"`
}

// NewGetFrontendMenuOK creates GetFrontendMenuOK with default headers values
func NewGetFrontendMenuOK() *GetFrontendMenuOK {

	return &GetFrontendMenuOK{}
}

// WithPayload adds the payload to the get frontend menu o k response
func (o *GetFrontendMenuOK) WithPayload(payload *models.PublicMenuResponse) *GetFrontendMenuOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend menu o k response
func (o *GetFrontendMenuOK) SetPayload(payload *models.PublicMenuResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendMenuOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendMenuBadRequestCode is the HTTP code returned for type GetFrontendMenuBadRequest
const GetFrontendMenuBadRequestCode int = 400

/*
GetFrontendMenuBadRequest Bad request due to missing or invalid parameters.

swagger:response getFrontendMenuBadRequest
*/
type GetFrontendMenuBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendMenuBadRequest creates GetFrontendMenuBadRequest with default headers values
func NewGetFrontendMenuBadRequest() *GetFrontendMenuBadRequest {

	return &GetFrontendMenuBadRequest{}
}

// WithPayload adds the payload to the get frontend menu bad request response
func (o *GetFrontendMenuBadRequest) WithPayload(payload *models.ErrorResponse) *GetFrontendMenuBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend menu bad request response
func (o *GetFrontendMenuBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendMenuBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendMenuUnauthorizedCode is the HTTP code returned for type GetFrontendMenuUnauthorized
const GetFrontendMenuUnauthorizedCode int = 401

/*
GetFrontendMenuUnauthorized Unauthorized. The request is missing valid authentication.

swagger:response getFrontendMenuUnauthorized
*/
type GetFrontendMenuUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendMenuUnauthorized creates GetFrontendMenuUnauthorized with default headers values
func NewGetFrontendMenuUnauthorized() *GetFrontendMenuUnauthorized {

	return &GetFrontendMenuUnauthorized{}
}

// WithPayload adds the payload to the get frontend menu unauthorized response
func (o *GetFrontendMenuUnauthorized) WithPayload(payload *models.ErrorResponse) *GetFrontendMenuUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend menu unauthorized response
func (o *GetFrontendMenuUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendMenuUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendMenuNotFoundCode is the HTTP code returned for type GetFrontendMenuNotFound
const GetFrontendMenuNotFoundCode int = 404

/*
GetFrontendMenuNotFound Not found. The requested resource could not be found.

swagger:response getFrontendMenuNotFound
*/
type GetFrontendMenuNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendMenuNotFound creates GetFrontendMenuNotFound with default headers values
func NewGetFrontendMenuNotFound() *GetFrontendMenuNotFound {

	return &GetFrontendMenuNotFound{}
}

// WithPayload adds the payload to the get frontend menu not found response
func (o *GetFrontendMenuNotFound) WithPayload(payload *models.ErrorResponse) *GetFrontendMenuNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend menu not found response
func (o *GetFrontendMenuNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendMenuNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendMenuInternalServerErrorCode is the HTTP code returned for type GetFrontendMenuInternalServerError
const GetFrontendMenuInternalServerErrorCode int = 500

/*
GetFrontendMenuInternalServerError Internal server error.

swagger:response getFrontendMenuInternalServerError
*/
type GetFrontendMenuInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendMenuInternalServerError creates GetFrontendMenuInternalServerError with default headers values
func NewGetFrontendMenuInternalServerError() *GetFrontendMenuInternalServerError {

	return &GetFrontendMenuInternalServerError{}
}

// WithPayload adds the payload to the get frontend menu internal server error response
func (o *GetFrontendMenuInternalServerError) WithPayload(payload *models.ErrorResponse) *GetFrontendMenuInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend menu internal server error response
func (o *GetFrontendMenuInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendMenuInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
