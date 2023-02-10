package app

import (
	"github.com/VashUber/coursework/golang/internal/components/fileserver"
	"github.com/VashUber/coursework/golang/internal/components/httpserver"
	"github.com/VashUber/coursework/golang/internal/config"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	r := gin.Default()

	srv := httpserver.CreateNewHttpServer(cfg, r)
	fileserver.RunFileServer(r, "/", "./static")

	srv.ListenAndServe()
}
