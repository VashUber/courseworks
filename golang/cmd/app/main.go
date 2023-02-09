package main

import (
	"net/http"

	"github.com/VashUber/coursework/golang/internal/mux"
	"github.com/VashUber/coursework/golang/internal/server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(server.NewHttpServer, mux.NewServerMux),
		fx.Invoke(func(*http.Server) {}),
	)

	app.Run()
}
