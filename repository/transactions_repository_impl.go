package repository

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"gorm.io/gorm"
)

type TransactionsRepositoryImpl struct {
	db *gorm.DB
}

func (r *TransactionsRepositoryImpl) FindByID(id int) (domain.Transaction, error) {
	tr := domain.Transaction{}
	err := r.db.Where("id=?", id).Find(&tr).Error
	helper.PanicIfError(err)
	return tr, nil
}

func (r *TransactionsRepositoryImpl) Update(transaction domain.Transaction) (domain.Transaction, error) {
	err := r.db.Save(&transaction).Error
	helper.PanicIfError(err)

	return transaction, nil
}

func (r *TransactionsRepositoryImpl) Save(transaction domain.Transaction) (domain.Transaction, error) {
	err := r.db.Create(&transaction).Error
	helper.PanicIfError(err)

	return transaction, nil
}

func (r *TransactionsRepositoryImpl) FindByUserID(userID int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary=1").
		Where("user_id=?", userID).Order("id desc").
		Find(&transactions).Error

	helper.PanicIfError(err)

	return transactions, nil
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
