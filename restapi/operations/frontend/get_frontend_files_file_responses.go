// Code generated by go-swagger; DO NOT EDIT.

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ole-larsen/plutonium/models"
)

// GetFrontendFilesFileOKCode is the HTTP code returned for type GetFrontendFilesFileOK
const GetFrontendFilesFileOKCode int = 200

/*
GetFrontendFilesFileOK ok

swagger:response getFrontendFilesFileOK
*/
type GetFrontendFilesFileOK struct {

	/*
	  In: Body
	*/
	Payload models.FileResponse `json:"body,omitempty"`
}

// NewGetFrontendFilesFileOK creates GetFrontendFilesFileOK with default headers values
func NewGetFrontendFilesFileOK() *GetFrontendFilesFileOK {

	return &GetFrontendFilesFileOK{}
}

// WithPayload adds the payload to the get frontend files file o k response
func (o *GetFrontendFilesFileOK) WithPayload(payload models.FileResponse) *GetFrontendFilesFileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend files file o k response
func (o *GetFrontendFilesFileOK) SetPayload(payload models.FileResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendFilesFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetFrontendFilesFileBadRequestCode is the HTTP code returned for type GetFrontendFilesFileBadRequest
const GetFrontendFilesFileBadRequestCode int = 400

/*
GetFrontendFilesFileBadRequest Bad request due to missing or invalid parameters.

swagger:response getFrontendFilesFileBadRequest
*/
type GetFrontendFilesFileBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendFilesFileBadRequest creates GetFrontendFilesFileBadRequest with default headers values
func NewGetFrontendFilesFileBadRequest() *GetFrontendFilesFileBadRequest {

	return &GetFrontendFilesFileBadRequest{}
}

// WithPayload adds the payload to the get frontend files file bad request response
func (o *GetFrontendFilesFileBadRequest) WithPayload(payload *models.ErrorResponse) *GetFrontendFilesFileBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend files file bad request response
func (o *GetFrontendFilesFileBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendFilesFileBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendFilesFileUnauthorizedCode is the HTTP code returned for type GetFrontendFilesFileUnauthorized
const GetFrontendFilesFileUnauthorizedCode int = 401

/*
GetFrontendFilesFileUnauthorized Unauthorized. The request is missing valid authentication.

swagger:response getFrontendFilesFileUnauthorized
*/
type GetFrontendFilesFileUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendFilesFileUnauthorized creates GetFrontendFilesFileUnauthorized with default headers values
func NewGetFrontendFilesFileUnauthorized() *GetFrontendFilesFileUnauthorized {

	return &GetFrontendFilesFileUnauthorized{}
}

// WithPayload adds the payload to the get frontend files file unauthorized response
func (o *GetFrontendFilesFileUnauthorized) WithPayload(payload *models.ErrorResponse) *GetFrontendFilesFileUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend files file unauthorized response
func (o *GetFrontendFilesFileUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendFilesFileUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendFilesFileNotFoundCode is the HTTP code returned for type GetFrontendFilesFileNotFound
const GetFrontendFilesFileNotFoundCode int = 404

/*
GetFrontendFilesFileNotFound Not found. The requested resource could not be found.

swagger:response getFrontendFilesFileNotFound
*/
type GetFrontendFilesFileNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendFilesFileNotFound creates GetFrontendFilesFileNotFound with default headers values
func NewGetFrontendFilesFileNotFound() *GetFrontendFilesFileNotFound {

	return &GetFrontendFilesFileNotFound{}
}

// WithPayload adds the payload to the get frontend files file not found response
func (o *GetFrontendFilesFileNotFound) WithPayload(payload *models.ErrorResponse) *GetFrontendFilesFileNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend files file not found response
func (o *GetFrontendFilesFileNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendFilesFileNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFrontendFilesFileInternalServerErrorCode is the HTTP code returned for type GetFrontendFilesFileInternalServerError
const GetFrontendFilesFileInternalServerErrorCode int = 500

/*
GetFrontendFilesFileInternalServerError Internal server error.

swagger:response getFrontendFilesFileInternalServerError
*/
type GetFrontendFilesFileInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetFrontendFilesFileInternalServerError creates GetFrontendFilesFileInternalServerError with default headers values
func NewGetFrontendFilesFileInternalServerError() *GetFrontendFilesFileInternalServerError {

	return &GetFrontendFilesFileInternalServerError{}
}

// WithPayload adds the payload to the get frontend files file internal server error response
func (o *GetFrontendFilesFileInternalServerError) WithPayload(payload *models.ErrorResponse) *GetFrontendFilesFileInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend files file internal server error response
func (o *GetFrontendFilesFileInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendFilesFileInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
