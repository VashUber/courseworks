package app

import (
	"net/http"

	"github.com/VashUber/coursework/golang/internal/fileserver"
	"github.com/VashUber/coursework/golang/internal/mux"
	"github.com/VashUber/coursework/golang/internal/server"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	app := fx.New(
		fx.Provide(server.NewHttpServer, mux.NewServerMux, fileserver.NewFileServer),
		fx.Invoke(func(*http.Server) {}),
	)

	return app
}
