package helper

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

func SetTransactionStatus(input web.MidtransNotificationInput, transaction domain.Transaction) {
	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "canceled"
	}
}
