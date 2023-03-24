package service

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

type MidtransService interface {
	GetPaymentURL(transaction domain.Transaction, user domain.User) string
	PaymentProcess(input web.MidtransNotificationInput) error
}
