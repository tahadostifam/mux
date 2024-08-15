package mux

type Route struct {
	Handler HandlerFunc
	Method  string
	Path    string
}
