package mux

import (
	"context"
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
	InternalErrorHandler    http.Handler
}

var DefaultRouterConfig = &RouterConfig{
	NotFoundHandler:         &defaultNotFoundHandler{},
	MethodNotAllowedHandler: &defaultMethodNotAllowedHandler{},
	InternalErrorHandler:    &defaultInternalErrorHandler{},
}

// Init route
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
		match, paramsMap, _ := UrlMatchesPattern(route.Path, r.URL.Path)
		if !match {
			continue
		}

		if !strings.EqualFold(route.Method, r.Method) {
			ro.routerConfig.MethodNotAllowedHandler.ServeHTTP(w, r)
			return
		}

		// Set params
		r = r.WithContext(context.WithValue(r.Context(), ParamsGetter{}, paramsMap))

		// We have a match! Let's call the handler =)
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
