package mux

import "net/http"

func (ro *Router) GET(path string, handlerFunc HandlerFunc) {
	ro.Route("GET", path, handlerFunc)
}

func (ro *Router) POST(path string, handlerFunc HandlerFunc) {
	ro.Route("POST", path, handlerFunc)
}

func (ro *Router) DELETE(path string, handlerFunc HandlerFunc) {
	ro.Route("DELETE", path, handlerFunc)
}

func (ro *Router) PATCH(path string, handlerFunc HandlerFunc) {
	ro.Route("PATCH", path, handlerFunc)
}

// Default error handlers (NotFound, MethodNotAllowed)

func defaultMethodNotAllowedHandler(c *Context) {
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.ResponseWriter.WriteHeader(http.StatusMethodNotAllowed)
	c.ResponseWriter.Write(errMethodNotAllowed)
}

func defaultNotFoundHandler(c *Context) {
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.ResponseWriter.WriteHeader(http.StatusNotFound)
	c.ResponseWriter.Write(errNotFound)
}

func defaultInternalErrorHandler(c *Context) {
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	c.ResponseWriter.Write(errInternalError)
}
