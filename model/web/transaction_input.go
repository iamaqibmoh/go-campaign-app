package web

import "bwa-campaign-app/model/domain"

type CampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User domain.User
}
