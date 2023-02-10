package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func initSpaRoute(e *gin.Engine, root, spaRoot string) {
	e.NoRoute(func(ctx *gin.Context) {
		if strings.Contains(ctx.Request.URL.Path, spaRoot) {
			ctx.Status(http.StatusAccepted)
			ctx.File(root + ctx.Request.URL.Path)

			return
		}

		ctx.Status(http.StatusAccepted)
		ctx.File(fmt.Sprintf("%s%sindex.html", root, spaRoot))
	})
}
