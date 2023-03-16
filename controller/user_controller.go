package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	UserRegister(ctx *gin.Context)
}
