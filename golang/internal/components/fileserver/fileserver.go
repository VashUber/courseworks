package fileserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunFileServer(e *gin.Engine, root string) {
	e.NoRoute(func(ctx *gin.Context) {
		ctx.Status(http.StatusAccepted)
		ctx.File(root)
	})
}
