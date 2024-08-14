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
	router := mux.NewRouter(*mux.DefaultRouterConfig)

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
```

Check `examples/` directory for more...

## Benchmark

```txt
goos: linux
goarch: amd64
pkg: github.com/tahadostifam/go-mux
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkRouter-8                 942849              1325 ns/op             623 B/op          2 allocs/op
BenchmarkUrlMatchesPattern-8    84036637                13.40 ns/op            0 B/op          0 allocs/op
```
