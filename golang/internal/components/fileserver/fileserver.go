package fileserver

import "github.com/gin-gonic/gin"

func RunFileServer(e *gin.Engine, path, root string) {
	e.Static(path, root)
}
