package mux

import "net/http"

func NewServerMux(fs http.Handler) *http.ServeMux {
	mx := http.NewServeMux()

	mx.Handle("/", fs)
	initRoutes(mx)

	return mx
}

func initRoutes(mx *http.ServeMux) {
	mx.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Works"))
	})
}
