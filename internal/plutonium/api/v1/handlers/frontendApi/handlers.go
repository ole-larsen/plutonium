package frontendapi

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ole-larsen/plutonium/models"
	"github.com/ole-larsen/plutonium/restapi/operations/frontend"
)

func GetHeaderHandler(_ frontend.GetFrontendHeaderParams, _ *models.Principal) middleware.Responder {
	return frontend.NewGetFrontendHeaderOK().WithPayload(nil)
}

func GetFooterHandler(_ frontend.GetFrontendFooterParams, _ *models.Principal) middleware.Responder {
	return frontend.NewGetFrontendFooterOK().WithPayload(nil)
}
