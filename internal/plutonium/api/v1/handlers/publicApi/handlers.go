package publicapi

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/ole-larsen/plutonium/models"
	"github.com/ole-larsen/plutonium/restapi/operations/public"
)

type PublicAPI interface {
	GetPingHandler(_ public.GetPingParams) middleware.Responder
}

func GetPingHandler(_ public.GetPingParams) middleware.Responder {
	return public.NewGetPingOK().
		WithPayload(&models.PingResponse{Message: "pong", Timestamp: strfmt.DateTime(time.Now())})
}
