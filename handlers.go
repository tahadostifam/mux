package mux

import "net/http"

func (ro *Router) GET(path string, handlerFunc http.HandlerFunc) {
	ro.Route("GET", path, handlerFunc)
}

func (ro *Router) POST(path string, handlerFunc http.HandlerFunc) {
	ro.Route("POST", path, handlerFunc)
}

func (ro *Router) DELETE(path string, handlerFunc http.HandlerFunc) {
	ro.Route("DELETE", path, handlerFunc)
}

func (ro *Router) PATCH(path string, handlerFunc http.HandlerFunc) {
	ro.Route("PATCH", path, handlerFunc)
}

// Default error handlers (NotFound, MethodNotAllowed)

type (
	defaultMethodNotAllowedHandler struct{}
	defaultNotFoundHandler         struct{}
)

func (*defaultMethodNotAllowedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"application/json"}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(errMethodNotAllowed)
}

func (*defaultNotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header()["Content-Type"] = []string{"application/json"}
	w.WriteHeader(http.StatusNotFound)
	w.Write(errNotFound)
}
