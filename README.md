# 👾 MUX: A Simple and Lightweight Routing Library

Designed for fun and simplicity! =)

## Features:

-   **URL Path Matching:** Define routes with flexible path patterns to handle different URL structures.
-   **Handler Functions:** Associate handler functions with specific routes to determine how requests are processed.
-   **Method Support:** Handle various HTTP request methods like **GET**, **POST**, **PUT**, **DELETE**, and more.
-   **Parameter Extraction:** Extract dynamic values from URLs for flexible route handling.

## Install

```bash
go get -u github.com/tahadostifam/mux@latest
```

## Example

```go
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
```

Check `examples/` directory for more...

## Benchmark

```txt
goos: linux
goarch: amd64
pkg: github.com/tahadostifam/go-mux
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkRouter-8                2014885               567.0 ns/op           210 B/op          0 allocs/op
BenchmarkUrlMatchesPattern-8    85451278                13.62 ns/op            0 B/op          0 allocs/op
```
