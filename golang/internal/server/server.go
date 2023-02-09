package server

import (
	"context"
	"net/http"

	"go.uber.org/fx"
)

func NewHttpServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go srv.ListenAndServe()

			return nil
		},
	})

	return srv
}
