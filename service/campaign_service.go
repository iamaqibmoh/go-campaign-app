package service

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

type CampaignService interface {
	FindCampaigns(userID int) ([]domain.Campaign, error)
	FindDetailCampaignByID(id web.CampaignIDFromURI) (domain.Campaign, error)
	CreateCampaign(input web.CreateCampaignInput) (domain.Campaign, error)
	UpdateCampaign(id int, input web.CreateCampaignInput) (domain.Campaign, error)
}
