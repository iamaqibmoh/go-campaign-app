package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	CheckEmailAvailability(ctx *gin.Context)
	UploadAvatar(ctx *gin.Context)
}
