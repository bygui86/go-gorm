package rest

import (
	"fmt"
	"net/http"
	"time"

	"gopkg.in/logex.v1"

	"github.com/gorilla/mux"
)

const (
	// endpoints
	rootEndpoint           = "/api/v1/products"
	productIdKey           = "id"
	productIdValue         = "{" + productIdKey + ":[0-9]+}"
	getProductByIdEndpoint = "/" + productIdValue

	// header keys
	contentTypeHeaderKey = "Content-Type"
	acceptHeaderKey      = "Accept"

	// header values
	applicationJsonValue = "application/json"

	// server parameters
	httpServerHostFormat          = "%s:%d"
	httpServerWriteTimeoutDefault = time.Second * 15
	httpServerReadTimeoutDefault  = time.Second * 15
	httpServerIdleTimeoutDefault  = time.Second * 60
)

func (r *RestInterfaceImpl) setupRouter() {
	logex.Debug("Create new router")

	r.router = mux.NewRouter().StrictSlash(true)

	r.router.
		HandleFunc(rootEndpoint, r.getProducts).
		Methods(http.MethodGet).
		Headers(acceptHeaderKey, applicationJsonValue).
		Headers(contentTypeHeaderKey, applicationJsonValue)
	r.router.
		HandleFunc(rootEndpoint+getProductByIdEndpoint, r.getProductById).
		Methods(http.MethodGet).
		Headers(acceptHeaderKey, applicationJsonValue).
		Headers(contentTypeHeaderKey, applicationJsonValue).
		Queries(productIdKey, productIdValue)
	r.router.
		HandleFunc(rootEndpoint+getProductByIdEndpoint, r.createProduct).
		Methods(http.MethodPost).
		Headers(acceptHeaderKey, applicationJsonValue).
		Headers(contentTypeHeaderKey, applicationJsonValue).
		Queries(productIdKey, productIdValue)
	r.router.
		HandleFunc(rootEndpoint+getProductByIdEndpoint, r.updateProduct).
		Methods(http.MethodPut).
		Headers(acceptHeaderKey, applicationJsonValue).
		Headers(contentTypeHeaderKey, applicationJsonValue).
		Queries(productIdKey, productIdValue)
	r.router.
		HandleFunc(rootEndpoint+getProductByIdEndpoint, r.deleteProductById).
		Methods(http.MethodDelete).
		Headers(acceptHeaderKey, applicationJsonValue).
		Queries(productIdKey, productIdValue)
}

func (r *RestInterfaceImpl) setupHTTPServer() {
	logex.Debugf("Create new HTTP server on port %d", r.config.RestPort)

	if r.config != nil {
		r.httpServer = &http.Server{
			Addr:    fmt.Sprintf(httpServerHostFormat, r.config.RestHost, r.config.RestPort),
			Handler: r.router,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: httpServerWriteTimeoutDefault,
			ReadTimeout:  httpServerReadTimeoutDefault,
			IdleTimeout:  httpServerIdleTimeoutDefault,
		}
		return
	}

	logex.Error("HTTP server creation failed: REST server configurations not initialized")
}
