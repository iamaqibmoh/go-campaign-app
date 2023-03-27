package service

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/repository"
	"github.com/veritrans/go-midtrans"
	"strconv"
)

type MidtransServiceImpl struct {
	repository.TransactionRepository
	repository.CampaignRepository
}

func (m *MidtransServiceImpl) PaymentProcess(input web.MidtransNotificationInput) error {
	orderID, _ := strconv.Atoi(input.OrderID)
	transaction, err := m.TransactionRepository.FindByID(orderID)
	helper.PanicIfError(err)

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "canceled"
	}

	updatedTransaction, err := m.TransactionRepository.Update(transaction)
	helper.PanicIfError(err)

	campaign, err := m.CampaignRepository.FindByID(transaction.CampaignID)
	helper.PanicIfError(err)

	if updatedTransaction.Status == "paid" {
		campaign.CurrentAmount += updatedTransaction.Amount
		campaign.BackerCount += 1

		_, err := m.CampaignRepository.Update(campaign)
		helper.PanicIfError(err)
	}

	return nil
}

func (m *MidtransServiceImpl) GetPaymentURL(transaction domain.Transaction, user domain.User) string {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-45AdXWBv1X_mffLokUzpAX0q"
	midclient.ClientKey = "SB-Mid-client-_WBGGYDD9ytviwYB"
	midclient.APIEnvType = midtrans.Sandbox

	var snapGateway midtrans.SnapGateway
	snapGateway = midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	token, err := snapGateway.GetToken(snapReq)
	helper.PanicIfError(err)

	return token.RedirectURL
}

func NewMidtransService(transactionsRepository repository.TransactionRepository, campaignRepository repository.CampaignRepository) MidtransService {
	return &MidtransServiceImpl{transactionsRepository, campaignRepository}
}
