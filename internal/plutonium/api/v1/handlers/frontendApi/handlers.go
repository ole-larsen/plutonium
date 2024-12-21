package frontendapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ole-larsen/plutonium/models"
	"github.com/ole-larsen/plutonium/restapi/operations/frontend"
)

func GetHeaderHandler(params frontend.GetFrontendHeaderParams, principal *models.Principal) middleware.Responder {
	return frontend.NewGetFrontendHeaderOK().WithPayload(nil)
}

func GetFooterHandler(params frontend.GetFrontendFooterParams, principal *models.Principal) middleware.Responder {
	return frontend.NewGetFrontendFooterOK().WithPayload(nil)
}
