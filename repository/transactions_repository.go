package repository

import "bwa-campaign-app/model/domain"

type TransactionsRepository interface {
	FindByCampaignID(campaignID int) ([]domain.Transaction, error)
}
