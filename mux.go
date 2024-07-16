package mux

import (
	"net/http"
	"strings"
)

type Router struct {
	routes       []Route
	routerConfig *RouterConfig
}

type RouterConfig struct {
	NotFoundHandler         http.Handler
	MethodNotAllowedHandler http.Handler
}

var DefaultRouterConfig = &RouterConfig{
	NotFoundHandler:         &defaultNotFoundHandler{},
	MethodNotAllowedHandler: &defaultMethodNotAllowedHandler{},
}

func (ro *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	route := &Route{
		Method:  method,
		Path:    path,
		Handler: handlerFunc,
	}

	ro.routes = append(ro.routes, *route)
}

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range ro.routes {
		match := route.Match(r)
		if !match {
			continue
		}

		if !strings.EqualFold(route.Method, r.Method) {
			ro.routerConfig.MethodNotAllowedHandler.ServeHTTP(w, r)
			return
		}

		// We have a match! Call the handler, and return
		route.ServeHTTP(w, r)
		return
	}

	// No matches! Let's 404
	ro.routerConfig.NotFoundHandler.ServeHTTP(w, r)
}

func NewRouter(routerConfig RouterConfig) *Router {
	return &Router{
		routerConfig: &routerConfig,
	}
}
