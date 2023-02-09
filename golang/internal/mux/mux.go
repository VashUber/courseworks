package mux

import "net/http"

func NewServerMux() *http.ServeMux {
	mx := http.NewServeMux()
	initRoutes(mx)

	return mx
}

func initRoutes(mx *http.ServeMux) {
	mx.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Works"))
	})
}
