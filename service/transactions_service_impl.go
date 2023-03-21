package service

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/repository"
	"errors"
)

type TransactionsServiceImpl struct {
	repository.TransactionsRepository
	repository.CampaignRepository
}

func (s *TransactionsServiceImpl) GetByCampaignID(input web.CampaignTransactionsInput) ([]domain.Transaction, error) {
	campaign, err2 := s.CampaignRepository.FindByID(input.ID)
	helper.PanicIfError(err2)

	if input.User.ID != campaign.UserID {
		return nil, errors.New("You're not an owner of this campaign")
	}

	transactions, err := s.TransactionsRepository.FindByCampaignID(input.ID)
	helper.PanicIfError(err)

	return transactions, nil
}

func NewTransactionsService(transactionRepository repository.TransactionsRepository, campaignRepository repository.CampaignRepository) TransactionsService {
	return &TransactionsServiceImpl{
		TransactionsRepository: transactionRepository,
		CampaignRepository:     campaignRepository,
	}
}
