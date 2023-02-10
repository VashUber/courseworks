package controller

import (
	"github.com/VashUber/coursework/golang/internal/config"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(cfg config.ControllerConfig) *Router {
	r := &Router{
		Engine: gin.Default(),
	}

	r.initRoutes(cfg)

	return r
}

func (r *Router) initRoutes(cfg config.ControllerConfig) {
	{
		initSpaRoute(r.Engine, cfg.StaticPath)
	}

	{
		uc := newUserController()
		user_group := r.Engine.Group("/api/user")

		uc.initUserRoutes(user_group)
	}
}
