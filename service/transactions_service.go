package service

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

type TransactionsService interface {
	GetByCampaignID(input web.CampaignTransactionsInput) ([]domain.Transaction, error)
}
