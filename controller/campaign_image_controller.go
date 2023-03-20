package controller

import "github.com/gin-gonic/gin"

type CampaignImageController interface {
	CreateCampaignImage(ctx *gin.Context)
}
