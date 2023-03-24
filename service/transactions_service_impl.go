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
	MidtransService
}

func (s *TransactionsServiceImpl) CreateTransaction(input web.CreateTransactionInput) (domain.Transaction, error) {
	tr := domain.Transaction{}
	tr.Amount = input.Amount
	tr.CampaignID = input.CampaignID
	tr.UserID = input.User.ID
	tr.Status = "pending"

	transaction, err := s.TransactionsRepository.Save(tr)
	helper.PanicIfError(err)

	paymentURL := s.MidtransService.GetPaymentURL(transaction, input.User)
	transaction.PaymentURL = paymentURL

	update, err := s.TransactionsRepository.Update(transaction)
	helper.PanicIfError(err)

	return update, nil
}

func (s *TransactionsServiceImpl) GetByUserID(userID int) ([]domain.Transaction, error) {
	transactions, err := s.TransactionsRepository.FindByUserID(userID)
	helper.PanicIfError(err)

	return transactions, nil
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

func NewTransactionsService(transactionRepository repository.TransactionsRepository, campaignRepository repository.CampaignRepository, midtransService MidtransService) TransactionsService {
	return &TransactionsServiceImpl{
		TransactionsRepository: transactionRepository,
		CampaignRepository:     campaignRepository,
		MidtransService:        midtransService,
	}
}
