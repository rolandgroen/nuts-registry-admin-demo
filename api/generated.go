// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /web/auth)
	CreateSession(ctx echo.Context) error

	// (GET /web/customers)
	GetCustomers(ctx echo.Context) error

	// (POST /web/customers)
	ConnectCustomer(ctx echo.Context) error

	// (GET /web/customers/{id})
	GetCustomer(ctx echo.Context, id string) error

	// (PUT /web/customers/{id})
	UpdateCustomer(ctx echo.Context, id string) error

	// (GET /web/service-provider)
	GetServiceProvider(ctx echo.Context) error

	// (POST /web/service-provider)
	CreateServiceProvider(ctx echo.Context) error

	// (PUT /web/service-provider)
	UpdateServiceProvider(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateSession converts echo context to params.
func (w *ServerInterfaceWrapper) CreateSession(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateSession(ctx)
	return err
}

// GetCustomers converts echo context to params.
func (w *ServerInterfaceWrapper) GetCustomers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCustomers(ctx)
	return err
}

// ConnectCustomer converts echo context to params.
func (w *ServerInterfaceWrapper) ConnectCustomer(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ConnectCustomer(ctx)
	return err
}

// GetCustomer converts echo context to params.
func (w *ServerInterfaceWrapper) GetCustomer(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCustomer(ctx, id)
	return err
}

// UpdateCustomer converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateCustomer(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateCustomer(ctx, id)
	return err
}

// GetServiceProvider converts echo context to params.
func (w *ServerInterfaceWrapper) GetServiceProvider(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetServiceProvider(ctx)
	return err
}

// CreateServiceProvider converts echo context to params.
func (w *ServerInterfaceWrapper) CreateServiceProvider(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateServiceProvider(ctx)
	return err
}

// UpdateServiceProvider converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateServiceProvider(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateServiceProvider(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/web/auth", wrapper.CreateSession)
	router.GET(baseURL+"/web/customers", wrapper.GetCustomers)
	router.POST(baseURL+"/web/customers", wrapper.ConnectCustomer)
	router.GET(baseURL+"/web/customers/:id", wrapper.GetCustomer)
	router.PUT(baseURL+"/web/customers/:id", wrapper.UpdateCustomer)
	router.GET(baseURL+"/web/service-provider", wrapper.GetServiceProvider)
	router.POST(baseURL+"/web/service-provider", wrapper.CreateServiceProvider)
	router.PUT(baseURL+"/web/service-provider", wrapper.UpdateServiceProvider)

}

