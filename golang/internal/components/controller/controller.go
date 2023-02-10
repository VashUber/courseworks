package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	r := &Router{
		Engine: gin.Default(),
	}

	api := r.Engine.Group("/api")

	{
		api.GET("/test", r.SomeApi)
	}

	return r
}

func (r *Router) SomeApi(ctx *gin.Context) {
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, name)
}
