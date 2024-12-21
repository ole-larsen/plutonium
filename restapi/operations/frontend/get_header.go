// Code generated by go-swagger; DO NOT EDIT.

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/ole-larsen/plutonium/models"
)

// GetHeaderHandlerFunc turns a function with the right signature into a get header handler
type GetHeaderHandlerFunc func(GetHeaderParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHeaderHandlerFunc) Handle(params GetHeaderParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetHeaderHandler interface for that can handle valid get header params
type GetHeaderHandler interface {
	Handle(GetHeaderParams, *models.Principal) middleware.Responder
}

// NewGetHeader creates a new http.Handler for the get header operation
func NewGetHeader(ctx *middleware.Context, handler GetHeaderHandler) *GetHeader {
	return &GetHeader{Context: ctx, Handler: handler}
}

/*
	GetHeader swagger:route GET /header Frontend getHeader

Fetches the public header for the frontend.
*/
type GetHeader struct {
	Context *middleware.Context
	Handler GetHeaderHandler
}

func (o *GetHeader) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetHeaderParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}