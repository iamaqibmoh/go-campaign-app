package controller

import (
	"bwa-campaign-app/formatter"
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionControllerImpl struct {
	service.TransactionsService
}

func (c *TransactionControllerImpl) GetByUserID(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(domain.User)
	transactions, err := c.TransactionsService.GetByUserID(currentUser.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to user's transactions",
			http.StatusBadRequest,
			"BAD REQUEST",
			gin.H{"errors": err.Error()}))
		return
	}

	ctx.JSON(200, helper.APIResponse(
		"List of user's transactions",
		200, "success",
		formatter.UserTransactionsFormatter(transactions)))
}

func (c *TransactionControllerImpl) GetByCampaignID(ctx *gin.Context) {
	input := web.CampaignTransactionsInput{}
	err := ctx.ShouldBindUri(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to get campaign's transactions",
			http.StatusBadRequest,
			"BAD REQUEST",
			helper.ValidationErrorsFormatter(err)))
		return
	}

	currentUser := ctx.MustGet("currentUser").(domain.User)
	input.User = currentUser

	transactions, err := c.TransactionsService.GetByCampaignID(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to campaign's transactions",
			http.StatusBadRequest,
			"BAD REQUEST",
			gin.H{"errors": err.Error()}))
		return
	}

	ctx.JSON(200, helper.APIResponse(
		"List of campaign's transactions",
		200, "success",
		formatter.CampaignTransactionsFormatter(transactions)))
}

func NewTransactionController(transactionsService service.TransactionsService) TransactionsController {
	return &TransactionControllerImpl{TransactionsService: transactionsService}
}
