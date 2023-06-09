package repository

import "bwa-campaign-app/model/domain"

type TransactionRepository interface {
	FindByCampaignID(campaignID int) ([]domain.Transaction, error)
	FindByUserID(userID int) ([]domain.Transaction, error)
	Save(transaction domain.Transaction) (domain.Transaction, error)
	Update(transaction domain.Transaction) (domain.Transaction, error)
	FindByID(id int) (domain.Transaction, error)
	FindAll() ([]domain.Transaction, error)
}
