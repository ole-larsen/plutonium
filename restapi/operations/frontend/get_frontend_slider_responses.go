// Code generated by go-swagger; DO NOT EDIT.

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ole-larsen/plutonium/models"
)

// GetFrontendSliderOKCode is the HTTP code returned for type GetFrontendSliderOK
const GetFrontendSliderOKCode int = 200

/*
GetFrontendSliderOK ok

swagger:response getFrontendSliderOK
*/
type GetFrontendSliderOK struct {

	/*
	  In: Body
	*/
	Payload *models.PublicSlider `json:"body,omitempty"`
}

// NewGetFrontendSliderOK creates GetFrontendSliderOK with default headers values
func NewGetFrontendSliderOK() *GetFrontendSliderOK {

	return &GetFrontendSliderOK{}
}

// WithPayload adds the payload to the get frontend slider o k response
func (o *GetFrontendSliderOK) WithPayload(payload *models.PublicSlider) *GetFrontendSliderOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend slider o k response
func (o *GetFrontendSliderOK) SetPayload(payload *models.PublicSlider) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendSliderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendSliderBadRequestCode is the HTTP code returned for type GetFrontendSliderBadRequest
const GetFrontendSliderBadRequestCode int = 400

/*
GetFrontendSliderBadRequest Bad request due to missing or invalid parameters.

swagger:response getFrontendSliderBadRequest
*/
type GetFrontendSliderBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendSliderBadRequest creates GetFrontendSliderBadRequest with default headers values
func NewGetFrontendSliderBadRequest() *GetFrontendSliderBadRequest {

	return &GetFrontendSliderBadRequest{}
}

// WithPayload adds the payload to the get frontend slider bad request response
func (o *GetFrontendSliderBadRequest) WithPayload(payload *models.ErrorResponse) *GetFrontendSliderBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend slider bad request response
func (o *GetFrontendSliderBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendSliderBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendSliderUnauthorizedCode is the HTTP code returned for type GetFrontendSliderUnauthorized
const GetFrontendSliderUnauthorizedCode int = 401

/*
GetFrontendSliderUnauthorized Unauthorized. The request is missing valid authentication.

swagger:response getFrontendSliderUnauthorized
*/
type GetFrontendSliderUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendSliderUnauthorized creates GetFrontendSliderUnauthorized with default headers values
func NewGetFrontendSliderUnauthorized() *GetFrontendSliderUnauthorized {

	return &GetFrontendSliderUnauthorized{}
}

// WithPayload adds the payload to the get frontend slider unauthorized response
func (o *GetFrontendSliderUnauthorized) WithPayload(payload *models.ErrorResponse) *GetFrontendSliderUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend slider unauthorized response
func (o *GetFrontendSliderUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendSliderUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendSliderNotFoundCode is the HTTP code returned for type GetFrontendSliderNotFound
const GetFrontendSliderNotFoundCode int = 404

/*
GetFrontendSliderNotFound Not found. The requested resource could not be found.

swagger:response getFrontendSliderNotFound
*/
type GetFrontendSliderNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendSliderNotFound creates GetFrontendSliderNotFound with default headers values
func NewGetFrontendSliderNotFound() *GetFrontendSliderNotFound {

	return &GetFrontendSliderNotFound{}
}

// WithPayload adds the payload to the get frontend slider not found response
func (o *GetFrontendSliderNotFound) WithPayload(payload *models.ErrorResponse) *GetFrontendSliderNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend slider not found response
func (o *GetFrontendSliderNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendSliderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendSliderInternalServerErrorCode is the HTTP code returned for type GetFrontendSliderInternalServerError
const GetFrontendSliderInternalServerErrorCode int = 500

/*
GetFrontendSliderInternalServerError Internal server error.

swagger:response getFrontendSliderInternalServerError
*/
type GetFrontendSliderInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendSliderInternalServerError creates GetFrontendSliderInternalServerError with default headers values
func NewGetFrontendSliderInternalServerError() *GetFrontendSliderInternalServerError {

	return &GetFrontendSliderInternalServerError{}
}

// WithPayload adds the payload to the get frontend slider internal server error response
func (o *GetFrontendSliderInternalServerError) WithPayload(payload *models.ErrorResponse) *GetFrontendSliderInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend slider internal server error response
func (o *GetFrontendSliderInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendSliderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}