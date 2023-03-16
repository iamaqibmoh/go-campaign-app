package controller

import (
	"bwa-campaign-app/app"
	"bwa-campaign-app/model/domain"
	"github.com/gin-gonic/gin"
)

func UserTableJson(ctx *gin.Context) {
	var users []domain.User

	db := app.DBConnection()
	db.Find(&users)

	ctx.JSON(200, &users)
}
