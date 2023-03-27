package controller

import (
	"bwa-campaign-app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionCMSController struct {
	service.TransactionsService
}

func NewTransactionCMSController(transactionsService service.TransactionsService) *TransactionCMSController {
	return &TransactionCMSController{TransactionsService: transactionsService}
}

func (c *TransactionCMSController) ShowAll(ctx *gin.Context) {
	transactions, err := c.TransactionsService.FindAllTransactions()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "error.gohtml", nil)
		return
	}

	ctx.HTML(200, "transactions_index.gohtml", gin.H{"transactions": transactions})
}
