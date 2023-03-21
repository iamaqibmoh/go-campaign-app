package formatter

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

func campaignTransactionFormatter(transaction domain.Transaction) web.CampaignTransactionResponse {
	return web.CampaignTransactionResponse{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
}

func CampaignTransactionsFormatter(transactions []domain.Transaction) []web.CampaignTransactionResponse {
	campaignTransactions := []web.CampaignTransactionResponse{}
	for _, tr := range transactions {
		campaignTransaction := campaignTransactionFormatter(tr)
		campaignTransactions = append(campaignTransactions, campaignTransaction)
	}

	return campaignTransactions
}

func userTransactionFormatter(transaction domain.Transaction) web.UserTransactionResponse {
	return web.UserTransactionResponse{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		Campaign: web.CampaignOfUserTransaction{
			Name:     transaction.Campaign.Name,
			ImageURL: transaction.Campaign.CampaignImages[0].FileName,
		},
	}
}

func UserTransactionsFormatter(transactions []domain.Transaction) []web.UserTransactionResponse {
	userTransactionResponses := []web.UserTransactionResponse{}

	for _, transaction := range transactions {
		userTransactionResponses = append(userTransactionResponses, userTransactionFormatter(transaction))
	}

	return userTransactionResponses
}
