// Code generated by go-swagger; DO NOT EDIT.

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetFrontendFilesFileHandlerFunc turns a function with the right signature into a get frontend files file handler
type GetFrontendFilesFileHandlerFunc func(GetFrontendFilesFileParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFrontendFilesFileHandlerFunc) Handle(params GetFrontendFilesFileParams) middleware.Responder {
	return fn(params)
}

// GetFrontendFilesFileHandler interface for that can handle valid get frontend files file params
type GetFrontendFilesFileHandler interface {
	Handle(GetFrontendFilesFileParams) middleware.Responder
}

// NewGetFrontendFilesFile creates a new http.Handler for the get frontend files file operation
func NewGetFrontendFilesFile(ctx *middleware.Context, handler GetFrontendFilesFileHandler) *GetFrontendFilesFile {
	return &GetFrontendFilesFile{Context: ctx, Handler: handler}
}

/*
	GetFrontendFilesFile swagger:route GET /frontend/files/:file/ Frontend getFrontendFilesFile

Serve Static Images
*/
type GetFrontendFilesFile struct {
	Context *middleware.Context
	Handler GetFrontendFilesFileHandler
}

func (o *GetFrontendFilesFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetFrontendFilesFileParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
