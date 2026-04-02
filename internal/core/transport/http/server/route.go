package corehttpserver

import (
	"net/http"

	corehttpmiddleware "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/middleware"
)

type Route struct {
	Method string
	Path string
	Handler http.HandlerFunc
	Middleware []corehttpmiddleware.Middleware
}


func (r *Route) WithMiddleware() http.Handler {
	return corehttpmiddleware.ChainMiddleware(
		r.Handler,
		r.Middleware...,
	)
}