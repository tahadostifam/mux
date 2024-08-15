package main

import (
	"fmt"
	"net/http"

	"github.com/tahadostifam/go-mux"
)

func main() {
	router := mux.NewRouter(mux.DefaultRouterConfig)

	router.GET("/sample", func(c *mux.Context) {
		c.ResponseWriter.Write([]byte("Sample"))
	})

	router.Route("GET", "/", func(c *mux.Context) {
		c.ResponseWriter.Write([]byte("Hello world"))
	})

	router.Route("GET", "/posts/{slug}", func(c *mux.Context) {
		c.Request.Header.Add("Content-Type", "text/html")

		slug := c.Param("slug")

		c.ResponseWriter.Write([]byte("<h1>Post Slug:  " + slug + "</h>"))
	})

	fmt.Println("Server is listening on port 3000!")
	http.ListenAndServe(":3000", router)
}
