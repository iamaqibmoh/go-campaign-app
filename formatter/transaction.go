package formatter

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

func CampaignTransactionFormatter(transaction domain.Transaction) web.CampaignTransactionResponse {
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
		campaignTransaction := CampaignTransactionFormatter(tr)
		campaignTransactions = append(campaignTransactions, campaignTransaction)
	}

	return campaignTransactions
}
