package corehttpserver

import (
	"fmt"
	"net/http"

	corehttpmiddleware "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/middleware"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
	ApiVersion3 = ApiVersion("v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
	middleware []corehttpmiddleware.Middleware
}


func NewAPIVersionRouter(
	apiVersion ApiVersion,
	middleware ...corehttpmiddleware.Middleware,
) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux: http.NewServeMux(),
		apiVersion: apiVersion,
		middleware: middleware,
	}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.Handle(pattern, route.WithMiddleware())
	}
}

func (r *APIVersionRouter) WithMiddleware() http.Handler {
	return corehttpmiddleware.ChainMiddleware(
		r,
		r.middleware...,
	)
}





