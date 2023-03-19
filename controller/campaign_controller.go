package controller

import "github.com/gin-gonic/gin"

type CampaignController interface {
	FindCampaigns(ctx *gin.Context)
	FindCampaignByID(ctx *gin.Context)
}