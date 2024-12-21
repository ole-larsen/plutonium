// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/ole-larsen/plutonium/internal/blockchain"
	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/ole-larsen/plutonium/internal/plutonium"
	v1monitoringApi "github.com/ole-larsen/plutonium/internal/plutonium/api/v1/handlers/monitoringApi"
	v1publicApi "github.com/ole-larsen/plutonium/internal/plutonium/api/v1/handlers/publicApi"
	v1middleware "github.com/ole-larsen/plutonium/internal/plutonium/api/v1/middleware"
	"github.com/ole-larsen/plutonium/internal/plutonium/settings"
	"github.com/ole-larsen/plutonium/internal/storage"
	"github.com/ole-larsen/plutonium/restapi/operations"
	"github.com/ole-larsen/plutonium/restapi/operations/monitoring"
	"github.com/ole-larsen/plutonium/restapi/operations/public"
)

//go:generate swagger generate server --target ../../plutonium --name Service --spec ../schema/swagger.yml --principal models.Principal

func configureFlags(api *operations.ServiceAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{}
}

func configureAPI(api *operations.ServiceAPI) http.Handler {
	logger := log.NewLogger("info", log.DefaultBuildLogger)
	cfg := settings.LoadConfig(".env")
	fmt.Println(cfg.DSN)

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	api.Logger = logger.Infof

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	service := plutonium.NewServer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize storage and set it on the service
	store, err := storage.SetupStorage(ctx, cfg.DSN)
	if err != nil {
		panic(err)
	}

	// initialize web3 dialer
	dialer, err := blockchain.NewWeb3Dialer(logger, cfg.Network, store.GetContractsRepository())
	if err != nil {
		panic(err)
	}

	err = dialer.Load(ctx)
	if err != nil {
		panic(err)
	}

	service.
		SetSettings(cfg).
		SetLogger(logger).
		SetStorage(store)

	api.PublicGetPingHandler = public.GetPingHandlerFunc(v1publicApi.GetPingHandler)

	api.MonitoringGetMetricsHandler = monitoring.GetMetricsHandlerFunc(v1monitoringApi.GetMetricsHandler)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(_ *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	if s != nil {
		fmt.Println(scheme, addr)
	}
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	handler = v1middleware.CorsMiddleware(handler)
	handler = v1middleware.CsrfMiddleware(handler)

	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handler = v1middleware.PrometheusMiddleware(handler)
	return handler
}
