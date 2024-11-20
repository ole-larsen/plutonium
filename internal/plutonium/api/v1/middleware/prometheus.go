package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/ole-larsen/plutonium/internal/plutonium/api/v1/prometheus"
)

func PrometheusMiddleware(handler http.Handler) http.Handler {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := prometheus.NewResponseWriter(w)
		if lrw != nil && r != nil {
			handler.ServeHTTP(lrw, r)
			statusCode := lrw.StatusCode
			duration := time.Since(start)
			logger := log.NewLogger("info", log.DefaultBuildLogger)
			logger.Infoln(r.URL.String(), r.Method, fmt.Sprintf("%d", statusCode), duration.Seconds())
		}
	})

	return h
}
