package controller

import "github.com/gin-gonic/gin"

type TransactionsController interface {
	GetByCampaignID(ctx *gin.Context)
	GetByUserID(ctx *gin.Context)
	CreateTransaction(ctx *gin.Context)
	GetMidtransNotification(ctx *gin.Context)
}
