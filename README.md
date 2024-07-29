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
		params := mux.Params(r)

		slug := params["slug"]

		w.Write([]byte("Post Slug => " + slug))
	})

	fmt.Println("Server is listening now!")
	http.ListenAndServe(":3000", router)
}
```

Check `examples/` directory for more...

## Benchmark

```txt
goos: linux
goarch: amd64
pkg: github.com/tahadostifam/go-mux
cpu: Intel(R) Core(TM) i7-6820HQ CPU @ 2.70GHz
BenchmarkUrlMatchesPattern
BenchmarkUrlMatchesPattern-8 [1093012] [1106 ns/op] [1345 B/op] [11 allocs/op]
```
