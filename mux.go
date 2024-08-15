package mux

import (
	"net/http"
	"strings"
	"sync"
)

type Router struct {
	routes       []Route
	routerConfig *RouterConfig
	paramsPool   sync.Pool
	contextPool  sync.Pool
}

type RouterConfig struct {
	NotFoundHandler         HandlerFunc
	MethodNotAllowedHandler HandlerFunc
	InternalErrorHandler    HandlerFunc
}

var DefaultRouterConfig = &RouterConfig{
	NotFoundHandler:         defaultNotFoundHandler,
	MethodNotAllowedHandler: defaultMethodNotAllowedHandler,
	InternalErrorHandler:    defaultInternalErrorHandler,
}

func (ro *Router) Route(method, path string, handlerFunc HandlerFunc) {
	ro.routes = append(ro.routes, Route{
		Method:  method,
		Path:    path,
		Handler: handlerFunc,
	})
}

func (ro *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := ro.contextPool.Get().(*Context)
	ctx.Request = r
	ctx.ResponseWriter = w
	ctx.Params = ro.paramsPool.Get().(ParamsMap)

	defer func() {
		for k := range ctx.Params {
			delete(ctx.Params, k)
		}
		ro.paramsPool.Put(ctx.Params)
		ctx.Request = nil
		ctx.ResponseWriter = nil
		ctx.Params = nil
		ro.contextPool.Put(ctx)
	}()

	for i := range ro.routes {
		match, _, err := urlMatchesPattern(ro.routes[i].Path, r.URL.Path, ctx.Params)
		if err != nil {
			ro.routerConfig.InternalErrorHandler(ctx)
			return
		}
		if !match {
			continue
		}

		if !strings.EqualFold(ro.routes[i].Method, r.Method) {
			ro.routerConfig.MethodNotAllowedHandler(ctx)
			return
		}

		// We have a match! Let's call the handler
		ro.routes[i].Handler(ctx)
		return
	}

	// No matches! Let's 404
	ro.routerConfig.NotFoundHandler(ctx)
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
		contextPool: sync.Pool{
			New: func() interface{} {
				return &Context{}
			},
		},
	}
}
