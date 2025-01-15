// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ole-larsen/plutonium/models"
)

// GetFrontendAuthWalletConnectOKCode is the HTTP code returned for type GetFrontendAuthWalletConnectOK
const GetFrontendAuthWalletConnectOKCode int = 200

/*
GetFrontendAuthWalletConnectOK check public address in db

swagger:response getFrontendAuthWalletConnectOK
*/
type GetFrontendAuthWalletConnectOK struct {

	/*
	  In: Body
	*/
	Payload *models.Nonce `json:"body,omitempty"`
}

// NewGetFrontendAuthWalletConnectOK creates GetFrontendAuthWalletConnectOK with default headers values
func NewGetFrontendAuthWalletConnectOK() *GetFrontendAuthWalletConnectOK {

	return &GetFrontendAuthWalletConnectOK{}
}

// WithPayload adds the payload to the get frontend auth wallet connect o k response
func (o *GetFrontendAuthWalletConnectOK) WithPayload(payload *models.Nonce) *GetFrontendAuthWalletConnectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend auth wallet connect o k response
func (o *GetFrontendAuthWalletConnectOK) SetPayload(payload *models.Nonce) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendAuthWalletConnectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendAuthWalletConnectBadRequestCode is the HTTP code returned for type GetFrontendAuthWalletConnectBadRequest
const GetFrontendAuthWalletConnectBadRequestCode int = 400

/*
GetFrontendAuthWalletConnectBadRequest Bad request due to missing or invalid parameters.

swagger:response getFrontendAuthWalletConnectBadRequest
*/
type GetFrontendAuthWalletConnectBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendAuthWalletConnectBadRequest creates GetFrontendAuthWalletConnectBadRequest with default headers values
func NewGetFrontendAuthWalletConnectBadRequest() *GetFrontendAuthWalletConnectBadRequest {

	return &GetFrontendAuthWalletConnectBadRequest{}
}

// WithPayload adds the payload to the get frontend auth wallet connect bad request response
func (o *GetFrontendAuthWalletConnectBadRequest) WithPayload(payload *models.ErrorResponse) *GetFrontendAuthWalletConnectBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend auth wallet connect bad request response
func (o *GetFrontendAuthWalletConnectBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendAuthWalletConnectBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendAuthWalletConnectUnauthorizedCode is the HTTP code returned for type GetFrontendAuthWalletConnectUnauthorized
const GetFrontendAuthWalletConnectUnauthorizedCode int = 401

/*
GetFrontendAuthWalletConnectUnauthorized Unauthorized. The request is missing valid authentication.

swagger:response getFrontendAuthWalletConnectUnauthorized
*/
type GetFrontendAuthWalletConnectUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendAuthWalletConnectUnauthorized creates GetFrontendAuthWalletConnectUnauthorized with default headers values
func NewGetFrontendAuthWalletConnectUnauthorized() *GetFrontendAuthWalletConnectUnauthorized {

	return &GetFrontendAuthWalletConnectUnauthorized{}
}

// WithPayload adds the payload to the get frontend auth wallet connect unauthorized response
func (o *GetFrontendAuthWalletConnectUnauthorized) WithPayload(payload *models.ErrorResponse) *GetFrontendAuthWalletConnectUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend auth wallet connect unauthorized response
func (o *GetFrontendAuthWalletConnectUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendAuthWalletConnectUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendAuthWalletConnectNotFoundCode is the HTTP code returned for type GetFrontendAuthWalletConnectNotFound
const GetFrontendAuthWalletConnectNotFoundCode int = 404

/*
GetFrontendAuthWalletConnectNotFound Not found. The requested resource could not be found.

swagger:response getFrontendAuthWalletConnectNotFound
*/
type GetFrontendAuthWalletConnectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendAuthWalletConnectNotFound creates GetFrontendAuthWalletConnectNotFound with default headers values
func NewGetFrontendAuthWalletConnectNotFound() *GetFrontendAuthWalletConnectNotFound {

	return &GetFrontendAuthWalletConnectNotFound{}
}

// WithPayload adds the payload to the get frontend auth wallet connect not found response
func (o *GetFrontendAuthWalletConnectNotFound) WithPayload(payload *models.ErrorResponse) *GetFrontendAuthWalletConnectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend auth wallet connect not found response
func (o *GetFrontendAuthWalletConnectNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendAuthWalletConnectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendAuthWalletConnectInternalServerErrorCode is the HTTP code returned for type GetFrontendAuthWalletConnectInternalServerError
const GetFrontendAuthWalletConnectInternalServerErrorCode int = 500

/*
GetFrontendAuthWalletConnectInternalServerError Internal server error.

swagger:response getFrontendAuthWalletConnectInternalServerError
*/
type GetFrontendAuthWalletConnectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendAuthWalletConnectInternalServerError creates GetFrontendAuthWalletConnectInternalServerError with default headers values
func NewGetFrontendAuthWalletConnectInternalServerError() *GetFrontendAuthWalletConnectInternalServerError {

	return &GetFrontendAuthWalletConnectInternalServerError{}
}

// WithPayload adds the payload to the get frontend auth wallet connect internal server error response
func (o *GetFrontendAuthWalletConnectInternalServerError) WithPayload(payload *models.ErrorResponse) *GetFrontendAuthWalletConnectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend auth wallet connect internal server error response
func (o *GetFrontendAuthWalletConnectInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendAuthWalletConnectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
