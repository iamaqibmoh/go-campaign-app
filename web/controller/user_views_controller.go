package controller

import "github.com/gin-gonic/gin"

type UserViewsController struct {
}

func NewUserViewsController() *UserViewsController {
	return &UserViewsController{}
}

func (c *UserViewsController) Index(ctx *gin.Context) {
	ctx.HTML(200, "user_index.gohtml", nil)
}
