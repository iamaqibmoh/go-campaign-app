package controller

import "github.com/gin-gonic/gin"

type TransactionsController interface {
	GetByCampaignID(ctx *gin.Context)
}
