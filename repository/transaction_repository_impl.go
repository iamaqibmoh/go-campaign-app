package repository

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func (r *TransactionRepositoryImpl) FindAll() ([]domain.Transaction, error) {
	var trxs []domain.Transaction
	err := r.db.Preload("Campaign").Preload("User").Order("id desc").Find(&trxs).Error
	helper.PanicIfError(err)
	return trxs, nil
}

func (r *TransactionRepositoryImpl) FindByID(id int) (domain.Transaction, error) {
	tr := domain.Transaction{}
	err := r.db.Where("id=?", id).Find(&tr).Error
	helper.PanicIfError(err)
	return tr, nil
}

func (r *TransactionRepositoryImpl) Update(transaction domain.Transaction) (domain.Transaction, error) {
	err := r.db.Save(&transaction).Error
	helper.PanicIfError(err)

	return transaction, nil
}

func (r *TransactionRepositoryImpl) Save(transaction domain.Transaction) (domain.Transaction, error) {
	err := r.db.Create(&transaction).Error
	helper.PanicIfError(err)

	return transaction, nil
}

func (r *TransactionRepositoryImpl) FindByUserID(userID int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary=1").
		Where("user_id=?", userID).Order("id desc").
		Find(&transactions).Error

	helper.PanicIfError(err)

	return transactions, nil
}

func (r *TransactionRepositoryImpl) FindByCampaignID(campaignID int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Preload("User").Where("campaign_id=?", campaignID).Order("id desc").Find(&transactions).Error
	helper.PanicIfError(err)

	return transactions, nil
}

func NewTransactionsRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{db: db}
}
