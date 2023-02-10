package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
}

func newUserController() userController {
	uc := userController{}

	return uc
}

func (uc *userController) initUserRoutes(e *gin.RouterGroup) {
	e.GET("/get", uc.get)
}

func (uc *userController) get(ctx *gin.Context) {
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, name)
}
