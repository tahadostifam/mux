package main

import (
	"fmt"
	"net/http"

	"github.com/tahadostifam/go-mux"
)

func main() {
	router := mux.NewRouter(*mux.DefaultRouterConfig)

	router.GET("/sample", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Sample"))
	})

	router.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	router.Route("GET", "/posts/{slug}", func(w http.ResponseWriter, r *http.Request) {
		params := r.Context().Value(mux.ParamsGetter{}).(mux.ParamsMap)

		slug := params["slug"]

		w.Write([]byte("Post Slug => " + slug))
	})

	fmt.Println("Server is listening now!")
	http.ListenAndServe(":3000", router)
}
