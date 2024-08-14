package mux

import (
	"context"
	"net/http"
	"strings"
	"sync"
)

type Router struct {
	routes       []Route
	routerConfig *RouterConfig
	paramsPool   sync.Pool
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

// Route adds a new route to the router
func (ro *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	ro.routes = append(ro.routes, Route{
		Method:  method,
		Path:    path,
		Handler: handlerFunc,
	})
}

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paramsMap := ro.paramsPool.Get().(ParamsMap)
	defer ro.paramsPool.Put(paramsMap)

	for i := range ro.routes {
		match, params, err := urlMatchesPattern(ro.routes[i].Path, r.URL.Path, paramsMap)
		if err != nil {
			ro.routerConfig.InternalErrorHandler.ServeHTTP(w, r)
			return
		}
		if !match {
			continue
		}

		if !strings.EqualFold(ro.routes[i].Method, r.Method) {
			ro.routerConfig.MethodNotAllowedHandler.ServeHTTP(w, r)
			return
		}

		// Set params
		ctx := context.WithValue(r.Context(), paramsGetter{}, params)
		r = r.WithContext(ctx)

		// We have a match! Let's call the handler
		ro.routes[i].ServeHTTP(w, r)
		return
	}

	// No matches! Let's 404
	ro.routerConfig.NotFoundHandler.ServeHTTP(w, r)
}

func NewRouter(routerConfig *RouterConfig) *Router {
	if routerConfig == nil {
		routerConfig = DefaultRouterConfig
	}
	return &Router{
		routerConfig: routerConfig,
		paramsPool: sync.Pool{
			New: func() interface{} {
				return make(ParamsMap)
			},
		},
	}
}
