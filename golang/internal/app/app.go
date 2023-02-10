package app

import (
	"github.com/VashUber/coursework/golang/internal/components/controller"
	"github.com/VashUber/coursework/golang/internal/components/httpserver"
	"github.com/VashUber/coursework/golang/internal/config"
)

func Run(cfg *config.Config) {
	r := controller.NewRouter(cfg.Controller)

	httpserver.RunHttpServer(cfg.Http, r.Engine)
}
