// Code generated by go-swagger; DO NOT EDIT.

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ole-larsen/plutonium/models"
)

// GetFrontendUsersOKCode is the HTTP code returned for type GetFrontendUsersOK
const GetFrontendUsersOKCode int = 200

/*
GetFrontendUsersOK ok

swagger:response getFrontendUsersOK
*/
type GetFrontendUsersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.PublicUser `json:"body,omitempty"`
}

// NewGetFrontendUsersOK creates GetFrontendUsersOK with default headers values
func NewGetFrontendUsersOK() *GetFrontendUsersOK {

	return &GetFrontendUsersOK{}
}

// WithPayload adds the payload to the get frontend users o k response
func (o *GetFrontendUsersOK) WithPayload(payload []*models.PublicUser) *GetFrontendUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend users o k response
func (o *GetFrontendUsersOK) SetPayload(payload []*models.PublicUser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.PublicUser, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetFrontendUsersBadRequestCode is the HTTP code returned for type GetFrontendUsersBadRequest
const GetFrontendUsersBadRequestCode int = 400

/*
GetFrontendUsersBadRequest Bad request due to missing or invalid parameters.

swagger:response getFrontendUsersBadRequest
*/
type GetFrontendUsersBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendUsersBadRequest creates GetFrontendUsersBadRequest with default headers values
func NewGetFrontendUsersBadRequest() *GetFrontendUsersBadRequest {

	return &GetFrontendUsersBadRequest{}
}

// WithPayload adds the payload to the get frontend users bad request response
func (o *GetFrontendUsersBadRequest) WithPayload(payload *models.ErrorResponse) *GetFrontendUsersBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend users bad request response
func (o *GetFrontendUsersBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendUsersBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendUsersUnauthorizedCode is the HTTP code returned for type GetFrontendUsersUnauthorized
const GetFrontendUsersUnauthorizedCode int = 401

/*
GetFrontendUsersUnauthorized Unauthorized. The request is missing valid authentication.

swagger:response getFrontendUsersUnauthorized
*/
type GetFrontendUsersUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendUsersUnauthorized creates GetFrontendUsersUnauthorized with default headers values
func NewGetFrontendUsersUnauthorized() *GetFrontendUsersUnauthorized {

	return &GetFrontendUsersUnauthorized{}
}

// WithPayload adds the payload to the get frontend users unauthorized response
func (o *GetFrontendUsersUnauthorized) WithPayload(payload *models.ErrorResponse) *GetFrontendUsersUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend users unauthorized response
func (o *GetFrontendUsersUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendUsersUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendUsersNotFoundCode is the HTTP code returned for type GetFrontendUsersNotFound
const GetFrontendUsersNotFoundCode int = 404

/*
GetFrontendUsersNotFound Not found. The requested resource could not be found.

swagger:response getFrontendUsersNotFound
*/
type GetFrontendUsersNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendUsersNotFound creates GetFrontendUsersNotFound with default headers values
func NewGetFrontendUsersNotFound() *GetFrontendUsersNotFound {

	return &GetFrontendUsersNotFound{}
}

// WithPayload adds the payload to the get frontend users not found response
func (o *GetFrontendUsersNotFound) WithPayload(payload *models.ErrorResponse) *GetFrontendUsersNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend users not found response
func (o *GetFrontendUsersNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendUsersNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendUsersInternalServerErrorCode is the HTTP code returned for type GetFrontendUsersInternalServerError
const GetFrontendUsersInternalServerErrorCode int = 500

/*
GetFrontendUsersInternalServerError Internal server error.

swagger:response getFrontendUsersInternalServerError
*/
type GetFrontendUsersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendUsersInternalServerError creates GetFrontendUsersInternalServerError with default headers values
func NewGetFrontendUsersInternalServerError() *GetFrontendUsersInternalServerError {

	return &GetFrontendUsersInternalServerError{}
}

// WithPayload adds the payload to the get frontend users internal server error response
func (o *GetFrontendUsersInternalServerError) WithPayload(payload *models.ErrorResponse) *GetFrontendUsersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend users internal server error response
func (o *GetFrontendUsersInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
