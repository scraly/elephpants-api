// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/scraly/elephpants-api/pkg/swagger/server/restapi/operations"
	"github.com/scraly/elephpants-api/pkg/swagger/server/restapi/operations/elephpants"
)

//go:generate swagger generate server --target ../../server --name ElephpantsAPI --spec ../../swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.ElephpantsAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ElephpantsAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.BinProducer = runtime.ByteStreamProducer()
	api.JSONProducer = runtime.JSONProducer()
	api.TxtProducer = runtime.TextProducer()

	if api.ElephpantsDeleteElephpantHandler == nil {
		api.ElephpantsDeleteElephpantHandler = elephpants.DeleteElephpantHandlerFunc(func(params elephpants.DeleteElephpantParams) middleware.Responder {
			return middleware.NotImplemented("operation elephpants.DeleteElephpant has not yet been implemented")
		})
	}
	if api.ElephpantsGetElephpantHandler == nil {
		api.ElephpantsGetElephpantHandler = elephpants.GetElephpantHandlerFunc(func(params elephpants.GetElephpantParams) middleware.Responder {
			return middleware.NotImplemented("operation elephpants.GetElephpant has not yet been implemented")
		})
	}
	if api.ElephpantsGetElephpantImageHandler == nil {
		api.ElephpantsGetElephpantImageHandler = elephpants.GetElephpantImageHandlerFunc(func(params elephpants.GetElephpantImageParams) middleware.Responder {
			return middleware.NotImplemented("operation elephpants.GetElephpantImage has not yet been implemented")
		})
	}
	if api.ElephpantsGetElephpantsHandler == nil {
		api.ElephpantsGetElephpantsHandler = elephpants.GetElephpantsHandlerFunc(func(params elephpants.GetElephpantsParams) middleware.Responder {
			return middleware.NotImplemented("operation elephpants.GetElephpants has not yet been implemented")
		})
	}
	if api.ElephpantsPostElephpantHandler == nil {
		api.ElephpantsPostElephpantHandler = elephpants.PostElephpantHandlerFunc(func(params elephpants.PostElephpantParams) middleware.Responder {
			return middleware.NotImplemented("operation elephpants.PostElephpant has not yet been implemented")
		})
	}
	if api.ElephpantsPutElephpantHandler == nil {
		api.ElephpantsPutElephpantHandler = elephpants.PutElephpantHandlerFunc(func(params elephpants.PutElephpantParams) middleware.Responder {
			return middleware.NotImplemented("operation elephpants.PutElephpant has not yet been implemented")
		})
	}
	if api.ElephpantsCheckHealthHandler == nil {
		api.ElephpantsCheckHealthHandler = elephpants.CheckHealthHandlerFunc(func(params elephpants.CheckHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation elephpants.CheckHealth has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
