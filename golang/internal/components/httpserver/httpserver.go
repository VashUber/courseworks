package httpserver

import (
	"net/http"

	"github.com/VashUber/coursework/golang/internal/config"
	"github.com/gin-gonic/gin"
)

func RunHttpServer(cfg config.HttpConfig, r *gin.Engine) {
	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: r,
	}

	srv.ListenAndServe()
}
