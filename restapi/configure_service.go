// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"runtime/pprof"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"golang.org/x/oauth2"

	genericRuntime "runtime"

	"github.com/ole-larsen/plutonium/internal/blockchain"
	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/ole-larsen/plutonium/internal/plutonium"
	v1authApi "github.com/ole-larsen/plutonium/internal/plutonium/api/v1/handlers/authApi"
	v1frontendApi "github.com/ole-larsen/plutonium/internal/plutonium/api/v1/handlers/frontendApi"
	v1monitoringApi "github.com/ole-larsen/plutonium/internal/plutonium/api/v1/handlers/monitoringApi"
	v1publicApi "github.com/ole-larsen/plutonium/internal/plutonium/api/v1/handlers/publicApi"
	v1middleware "github.com/ole-larsen/plutonium/internal/plutonium/api/v1/middleware"
	"github.com/ole-larsen/plutonium/internal/plutonium/grpcserver"
	"github.com/ole-larsen/plutonium/internal/plutonium/httpclient"
	"github.com/ole-larsen/plutonium/internal/plutonium/oauth2client"
	"github.com/ole-larsen/plutonium/internal/plutonium/settings"
	"github.com/ole-larsen/plutonium/internal/storage"
	"github.com/ole-larsen/plutonium/restapi/operations"
	"github.com/ole-larsen/plutonium/restapi/operations/auth"
	"github.com/ole-larsen/plutonium/restapi/operations/frontend"
	"github.com/ole-larsen/plutonium/restapi/operations/monitoring"
	"github.com/ole-larsen/plutonium/restapi/operations/public"
)

//go:generate swagger generate server --target ../../plutonium --name Service --spec ../schema/swagger.yml --principal models.Principal

const defaultTimeout = 30

func configureFlags(api *operations.ServiceAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{}
}

func configureAPI(api *operations.ServiceAPI) http.Handler {
	logger := log.NewLogger("info", log.DefaultBuildLogger)
	cfg := settings.LoadConfig(".env")
	fmt.Println(cfg.DSN)

	// profileling cpu
	fcpu, err := os.Create(`cpu.profile`)
	if err != nil {
		panic(err)
	}

	defer fcpu.Close()

	if err := pprof.StartCPUProfile(fcpu); err != nil {
		panic(err)
	}

	defer pprof.StopCPUProfile()

	// memory profile
	fmem, err := os.Create(`mem.profile`)
	if err != nil {
		panic(err)
	}

	defer fmem.Close()
	genericRuntime.GC() // memory usage stat

	if err := pprof.WriteHeapProfile(fmem); err != nil {
		panic(err)
	}

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

	// initialize web3 web3Dialer
	web3Dialer, err := blockchain.NewWeb3Dialer(logger, cfg.Network, store.GetContractsRepository())
	if err != nil {
		panic(err)
	}

	err = web3Dialer.Load(ctx)
	if err != nil {
		panic(err)
	}

	httpDialer := httpclient.NewHTTPClient().
		SetTimeout(defaultTimeout * time.Second).
		SetSettings(cfg).
		SetTransport(httpclient.SetDefaultTransport())

	oauth2Client := &oauth2client.Oauth2{
		Client: oauth2client.NewClient().SetSettings(cfg),
		Config: make(map[string]oauth2.Config),
	}

	// initialize grpc server
	grpc := grpcserver.SetupGRPC(
		cfg.GRPC.Host,
		cfg.GRPC.Port,
	).
		SetStorage(store).
		SetWeb3Dialer(web3Dialer).
		SetHTTPDialer(httpDialer).
		SetOauth2(oauth2Client).
		SetSettings(cfg)

	service.
		SetSettings(cfg).
		SetLogger(logger).
		SetStorage(store).
		SetGRPC(grpc).
		SetWeb3Dialer(web3Dialer).
		SetHTTPDialer(httpDialer).
		SetOauth2(oauth2Client)

	go func() {
		logger.Infoln("starting grpc server")

		err := grpc.ListenAndServe(cfg)
		if err != nil {
			logger.Errorln(err)
		}
	}()

	// setup file server to serve images and files
	http.Handle("/", http.FileServer(http.Dir("./uploads")))

	api.PublicGetPingHandler = public.GetPingHandlerFunc(v1publicApi.GetPingHandler)

	api.MonitoringGetMetricsHandler = monitoring.GetMetricsHandlerFunc(v1monitoringApi.GetMetricsHandler)

	// frontend handlers
	frontendAPI := v1frontendApi.NewFrontendAPI(service)

	// applies when the "x-token" header is set
	api.XTokenAuth = frontendAPI.XTokenAuth

	api.FrontendGetFrontendMenuHandler = frontend.GetFrontendMenuHandlerFunc(frontendAPI.GetMenuHandler)
	api.FrontendGetFrontendSliderHandler = frontend.GetFrontendSliderHandlerFunc(frontendAPI.GetSliderHandler)
	api.FrontendGetFrontendFilesFileHandler = frontend.GetFrontendFilesFileHandlerFunc(frontendAPI.GetFileHandler)
	api.FrontendGetFrontendCategoriesHandler = frontend.GetFrontendCategoriesHandlerFunc(frontendAPI.GetCategoriesHandler)
	api.FrontendGetFrontendContractsHandler = frontend.GetFrontendContractsHandlerFunc(frontendAPI.GetContractsHandler)
	api.FrontendGetFrontendUsersHandler = frontend.GetFrontendUsersHandlerFunc(frontendAPI.GetUsersHandler)
	api.FrontendGetFrontendPageSlugHandler = frontend.GetFrontendPageSlugHandlerFunc(frontendAPI.GetPagesHandler)
	api.FrontendGetFrontendContactHandler = frontend.GetFrontendContactHandlerFunc(frontendAPI.GetContactsHandler)
	api.FrontendPostFrontendContactFormHandler = frontend.PostFrontendContactFormHandlerFunc(frontendAPI.PostContactsHandler)
	api.FrontendGetFrontendFaqHandler = frontend.GetFrontendFaqHandlerFunc(frontendAPI.GetFaqHandler)
	api.FrontendGetFrontendHelpCenterHandler = frontend.GetFrontendHelpCenterHandlerFunc(frontendAPI.GetHelpCenterHandler)
	api.FrontendGetFrontendBlogHandler = frontend.GetFrontendBlogHandlerFunc(frontendAPI.GetBlogsHandler)
	api.FrontendGetFrontendBlogSlugHandler = frontend.GetFrontendBlogSlugHandlerFunc(frontendAPI.GetBlogsSlugHandler)
	api.FrontendGetFrontendWalletConnectHandler = frontend.GetFrontendWalletConnectHandlerFunc(frontendAPI.GetWalletConnectHandler)
	api.FrontendGetFrontendCreateAndSellHandler = frontend.GetFrontendCreateAndSellHandlerFunc(frontendAPI.GetCreateAndSellHandler)

	authAPI := v1authApi.NewAuthAPI(service)

	api.AuthGetFrontendAuthWalletConnectHandler = auth.GetFrontendAuthWalletConnectHandlerFunc(authAPI.GetWalletConnect)
	api.AuthPostFrontendAuthWalletConnectHandler = auth.PostFrontendAuthWalletConnectHandlerFunc(authAPI.PostWalletConnect)
	api.AuthGetFrontendAuthCallbackHandler = auth.GetFrontendAuthCallbackHandlerFunc(authAPI.GetOauth2Callback)

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
	fmt.Println("----------------------------------------------------------------")
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handler = v1middleware.CorsMiddleware(handler)
	handler = v1middleware.CsrfMiddleware(handler)
	handler = v1middleware.PrometheusMiddleware(handler)

	return handler
}
