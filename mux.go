package mux

import "net/http"

type Route struct {
	Handler http.Handler
	Method  string
	Path    string
}

type Router struct {
	routes []Route
}

func (ro *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	route := &Route{
		Method:  method,
		Path:    path,
		Handler: handlerFunc,
	}

	ro.routes = append(ro.routes, *route)
}

func (re *Route) Match(r *http.Request) bool {
	if r.Method != re.Method {
		return false // Method mismatch
	}

	if r.URL.Path != re.Path {
		return false // Path mismatch
	}

	return true
}

func (re *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	re.Handler.ServeHTTP(w, r)
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range rtr.routes {
		match := route.Match(r)
		if !match {
			continue
		}

		// We have a match! Call the handler, and return
		route.ServeHTTP(w, r)
		return
	}

	// No matches, so it's a 404
	http.NotFound(w, r)
}

func NewRouter() *Router {
	return &Router{}
}

// func (r *Router) GetRoute(name string) *route {
// 	return r.namedRoutes[name]
// }

// func (r *Router) RegisterEmptyRoute(name string) *route {
// 	route := &route{}
// }
