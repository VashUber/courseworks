package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initSpaRoute(e *gin.Engine, root string) {
	e.NoRoute(func(ctx *gin.Context) {
		ctx.Status(http.StatusAccepted)
		ctx.File(root)
	})
}
