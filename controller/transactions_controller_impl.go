package controller

import (
	"bwa-campaign-app/formatter"
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionsControllerImpl struct {
	service.TransactionsService
	service.MidtransService
}

func (c *TransactionsControllerImpl) GetMidtransNotification(ctx *gin.Context) {
	input := web.MidtransNotificationInput{}
	err := ctx.ShouldBindJSON(&input)

	bytes, _ := json.Marshal(input)
	fmt.Println(string(bytes))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to process notification transaction",
			http.StatusBadRequest,
			"BAD REQUEST",
			err.Error()))
		return
	}

	err = c.MidtransService.PaymentProcess(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Failed to process notification transaction",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			err.Error()))
		return
	}

	ctx.JSON(200, input)
}

func (c *TransactionsControllerImpl) CreateTransaction(ctx *gin.Context) {
	input := web.CreateTransactionInput{}
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to create transaction",
			http.StatusBadRequest,
			"BAD REQUEST",
			helper.ValidationErrorsFormatter(err)))
		return
	}

	input.User = ctx.MustGet("currentUser").(domain.User)

	transaction, err := c.TransactionsService.CreateTransaction(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Failed to create transactions",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			err.Error()))
		return
	}

	ctx.JSON(200, helper.APIResponse(
		"Success to create transactions",
		200, "success",
		formatter.MidtransTransactionFormatter(transaction)))
}

func (c *TransactionsControllerImpl) GetByUserID(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(domain.User)
	transactions, err := c.TransactionsService.GetByUserID(currentUser.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.APIResponse(
			"Failed to user's transactions",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			gin.H{"errors": err.Error()}))
		return
	}

	ctx.JSON(200, helper.APIResponse(
		"List of user's transactions",
		200, "success",
		formatter.UserTransactionsFormatter(transactions)))
}

func (c *TransactionsControllerImpl) GetByCampaignID(ctx *gin.Context) {
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
		ctx.JSON(http.StatusInternalServerError, helper.APIResponse(
			"Failed to campaign's transactions",
			http.StatusInternalServerError,
			"INTERNAL SERVER ERROR",
			gin.H{"errors": err.Error()}))
		return
	}

	ctx.JSON(200, helper.APIResponse(
		"List of campaign's transactions",
		200, "success",
		formatter.CampaignTransactionsFormatter(transactions)))
}

func NewTransactionsController(transactionsService service.TransactionsService, midtransService service.MidtransService) TransactionsController {
	return &TransactionsControllerImpl{transactionsService, midtransService}
}
