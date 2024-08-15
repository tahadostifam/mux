package mux

import (
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"testing"
)

func BenchmarkRouter(b *testing.B) {
	f, err := os.Create("mux_allocs.pprof")
	if err != nil {
		log.Fatalln("Could not create file", err)
	}
	defer f.Close()

	pprof.Lookup("allocs").WriteTo(f, 0)

	router := NewRouter(DefaultRouterConfig)

	r, err := http.NewRequest("GET", "/posts/hello_world", nil)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < b.N; i++ {
		router.Route("GET", "/posts/{slug}", func(c *Context) {})
		router.ServeHTTP(nil, r)
	}
}
