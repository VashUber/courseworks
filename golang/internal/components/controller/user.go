package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
}

func NewUserController() userController {
	uc := userController{}

	return uc
}

func (uc *userController) InitUserRoutes(e *gin.RouterGroup) {
	e.GET("/get", uc.Get)
}

func (uc *userController) Get(ctx *gin.Context) {
	name := ctx.Query("name")

	ctx.JSON(http.StatusOK, name)
}
