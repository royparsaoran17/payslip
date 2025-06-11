// Package router
package router

import (
	"net/http"

	"payroll-se/internal/appctx"
	"payroll-se/internal/ucase/contract"
	"payroll-se/pkg/routerkit"
)

// httpHandlerFunc is a contract http handler for router
type httpHandlerFunc func(request *http.Request, svc contract.UseCase, conf *appctx.Config) appctx.Response

// Router is a contract router and must implement this interface
type Router interface {
	Route() *routerkit.Router
}
