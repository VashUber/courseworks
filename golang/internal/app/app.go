package app

import (
	"github.com/VashUber/coursework/golang/internal/components/controller"
	"github.com/VashUber/coursework/golang/internal/components/fileserver"
	"github.com/VashUber/coursework/golang/internal/components/httpserver"
	"github.com/VashUber/coursework/golang/internal/config"
)

func Run(cfg *config.Config) {
	r := controller.NewRouter()

	fileserver.RunFileServer(r.Engine, "./static/index.html")
	httpserver.RunHttpServer(cfg, r.Engine)
}
