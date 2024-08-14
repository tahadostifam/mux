package main

import (
	"fmt"
	"net/http"

	"github.com/tahadostifam/go-mux"
)

func main() {
	router := mux.NewRouter(mux.DefaultRouterConfig)

	router.GET("/sample", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Sample"))
	})

	router.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	router.Route("GET", "/posts/{slug}", func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("Content-Type", "text/html")

		params := mux.Params(r)

		slug := params["slug"]

		w.Write([]byte("<h1>Post Slug:  " + slug + "</h>"))
	})

	fmt.Println("Server is listening on port 3000!")
	http.ListenAndServe(":3000", router)
}
