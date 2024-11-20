// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ole-larsen/plutonium/models"
)

// GetPingOKCode is the HTTP code returned for type GetPingOK
const GetPingOKCode int = 200

/*
GetPingOK Successful response indicating server availability.

swagger:response getPingOK
*/
type GetPingOK struct {

	/*
	  In: Body
	*/
	Payload *models.PingResponse `json:"body,omitempty"`
}

// NewGetPingOK creates GetPingOK with default headers values
func NewGetPingOK() *GetPingOK {

	return &GetPingOK{}
}

// WithPayload adds the payload to the get ping o k response
func (o *GetPingOK) WithPayload(payload *models.PingResponse) *GetPingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ping o k response
func (o *GetPingOK) SetPayload(payload *models.PingResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPingInternalServerErrorCode is the HTTP code returned for type GetPingInternalServerError
const GetPingInternalServerErrorCode int = 500

/*
GetPingInternalServerError Internal Server Error. This typically indicates a server-side issue or
an unexpected runtime error preventing proper functionality.

swagger:response getPingInternalServerError
*/
type GetPingInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetPingInternalServerError creates GetPingInternalServerError with default headers values
func NewGetPingInternalServerError() *GetPingInternalServerError {

	return &GetPingInternalServerError{}
}

// WithPayload adds the payload to the get ping internal server error response
func (o *GetPingInternalServerError) WithPayload(payload *models.ErrorResponse) *GetPingInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get ping internal server error response
func (o *GetPingInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPingInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}