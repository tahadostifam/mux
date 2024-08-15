package mux

import "net/http"

type Context struct {
	Params         ParamsMap
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

type HandlerFunc func(*Context)

func (c *Context) Param(key string) string {
	return c.Params[key]
}
