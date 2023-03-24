package test

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/service"
	"fmt"
	"testing"
)

func TestMidtrans(t *testing.T) {
	//serv := service.NewTransactionsService(
	//	repository.NewTransactionsRepository(app.DBConnection()),
	//	newCampaignRepository(),
	//	service.NewMidtransService())
	//
	midt := service.NewMidtransService()
	tr := domain.Transaction{
		ID:     7,
		Amount: 1234567,
	}

	user := domain.User{
		Name:  "Aqib",
		Email: "aqib@gmail.com",
	}
	paymentURL := midt.GetPaymentURL(tr, user)

	fmt.Println(paymentURL)
}
