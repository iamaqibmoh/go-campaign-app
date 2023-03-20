package controller

import "github.com/gin-gonic/gin"

type CampaignController interface {
	FindCampaigns(ctx *gin.Context)
	FindCampaignByID(ctx *gin.Context)
	CreateCampaign(ctx *gin.Context)
	UpdateCampaign(ctx *gin.Context)
}
