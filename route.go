package mux

import "net/http"

type Route struct {
	Handler http.Handler
	Method  string
	Path    string
}

func (re *Route) Match(req *http.Request) bool {
	return mathQueryString(re.Path, req.URL.Path)
}

func (re *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	re.Handler.ServeHTTP(w, r)
}
