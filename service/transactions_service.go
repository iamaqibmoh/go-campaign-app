package service

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

type TransactionsService interface {
	GetByCampaignID(input web.CampaignTransactionsInput) ([]domain.Transaction, error)
	GetByUserID(userID int) ([]domain.Transaction, error)
	CreateTransaction(input web.CreateTransactionInput) (domain.Transaction, error)
	FindAllTransactions() ([]domain.Transaction, error)
}
