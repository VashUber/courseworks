package httpserver

import (
	"net/http"

	"github.com/VashUber/coursework/golang/internal/config"
	"github.com/gin-gonic/gin"
)

func CreateNewHttpServer(cfg *config.Config, r *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: r,
	}

	return srv
}
