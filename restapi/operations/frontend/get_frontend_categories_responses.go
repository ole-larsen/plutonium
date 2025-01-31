// Code generated by go-swagger; DO NOT EDIT.

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ole-larsen/plutonium/models"
)

// GetFrontendCategoriesOKCode is the HTTP code returned for type GetFrontendCategoriesOK
const GetFrontendCategoriesOKCode int = 200

/*
GetFrontendCategoriesOK ok

swagger:response getFrontendCategoriesOK
*/
type GetFrontendCategoriesOK struct {

	/*
	  In: Body
	*/
	Payload models.PublicCategories `json:"body,omitempty"`
}

// NewGetFrontendCategoriesOK creates GetFrontendCategoriesOK with default headers values
func NewGetFrontendCategoriesOK() *GetFrontendCategoriesOK {

	return &GetFrontendCategoriesOK{}
}

// WithPayload adds the payload to the get frontend categories o k response
func (o *GetFrontendCategoriesOK) WithPayload(payload models.PublicCategories) *GetFrontendCategoriesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend categories o k response
func (o *GetFrontendCategoriesOK) SetPayload(payload models.PublicCategories) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendCategoriesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.PublicCategories{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetFrontendCategoriesBadRequestCode is the HTTP code returned for type GetFrontendCategoriesBadRequest
const GetFrontendCategoriesBadRequestCode int = 400

/*
GetFrontendCategoriesBadRequest Bad request due to missing or invalid parameters.

swagger:response getFrontendCategoriesBadRequest
*/
type GetFrontendCategoriesBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendCategoriesBadRequest creates GetFrontendCategoriesBadRequest with default headers values
func NewGetFrontendCategoriesBadRequest() *GetFrontendCategoriesBadRequest {

	return &GetFrontendCategoriesBadRequest{}
}

// WithPayload adds the payload to the get frontend categories bad request response
func (o *GetFrontendCategoriesBadRequest) WithPayload(payload *models.ErrorResponse) *GetFrontendCategoriesBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend categories bad request response
func (o *GetFrontendCategoriesBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendCategoriesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendCategoriesUnauthorizedCode is the HTTP code returned for type GetFrontendCategoriesUnauthorized
const GetFrontendCategoriesUnauthorizedCode int = 401

/*
GetFrontendCategoriesUnauthorized Unauthorized. The request is missing valid authentication.

swagger:response getFrontendCategoriesUnauthorized
*/
type GetFrontendCategoriesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendCategoriesUnauthorized creates GetFrontendCategoriesUnauthorized with default headers values
func NewGetFrontendCategoriesUnauthorized() *GetFrontendCategoriesUnauthorized {

	return &GetFrontendCategoriesUnauthorized{}
}

// WithPayload adds the payload to the get frontend categories unauthorized response
func (o *GetFrontendCategoriesUnauthorized) WithPayload(payload *models.ErrorResponse) *GetFrontendCategoriesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend categories unauthorized response
func (o *GetFrontendCategoriesUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendCategoriesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendCategoriesNotFoundCode is the HTTP code returned for type GetFrontendCategoriesNotFound
const GetFrontendCategoriesNotFoundCode int = 404

/*
GetFrontendCategoriesNotFound Not found. The requested resource could not be found.

swagger:response getFrontendCategoriesNotFound
*/
type GetFrontendCategoriesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendCategoriesNotFound creates GetFrontendCategoriesNotFound with default headers values
func NewGetFrontendCategoriesNotFound() *GetFrontendCategoriesNotFound {

	return &GetFrontendCategoriesNotFound{}
}

// WithPayload adds the payload to the get frontend categories not found response
func (o *GetFrontendCategoriesNotFound) WithPayload(payload *models.ErrorResponse) *GetFrontendCategoriesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend categories not found response
func (o *GetFrontendCategoriesNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendCategoriesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendCategoriesInternalServerErrorCode is the HTTP code returned for type GetFrontendCategoriesInternalServerError
const GetFrontendCategoriesInternalServerErrorCode int = 500

/*
GetFrontendCategoriesInternalServerError Internal server error.

swagger:response getFrontendCategoriesInternalServerError
*/
type GetFrontendCategoriesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendCategoriesInternalServerError creates GetFrontendCategoriesInternalServerError with default headers values
func NewGetFrontendCategoriesInternalServerError() *GetFrontendCategoriesInternalServerError {

	return &GetFrontendCategoriesInternalServerError{}
}

// WithPayload adds the payload to the get frontend categories internal server error response
func (o *GetFrontendCategoriesInternalServerError) WithPayload(payload *models.ErrorResponse) *GetFrontendCategoriesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend categories internal server error response
func (o *GetFrontendCategoriesInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendCategoriesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
