package monitoringapi

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/ole-larsen/plutonium/restapi/operations/monitoring"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MonitoringAPI interface {
	GetMetricsHandler(params monitoring.GetMetricsParams) middleware.Responder
}

func GetMetricsHandler(params monitoring.GetMetricsParams) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
		promhttp.Handler().ServeHTTP(w, params.HTTPRequest)
	})
}
