package mux

var (
	errMethodNotAllowed = []byte(`{"error": "method not allowed"}`)
	errNotFound         = []byte(`{"error": "not found"}`)
)
