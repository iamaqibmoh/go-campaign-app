package repository

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"gorm.io/gorm"
)

type TransactionsRepositoryImpl struct {
	db *gorm.DB
}

func (r *TransactionsRepositoryImpl) FindByCampaignID(campaignID int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Preload("User").Where("campaign_id=?", campaignID).Order("id desc").Find(&transactions).Error
	helper.PanicIfError(err)

	return transactions, nil
}

func NewTransactionsRepository(db *gorm.DB) TransactionsRepository {
	return &TransactionsRepositoryImpl{db: db}
}
