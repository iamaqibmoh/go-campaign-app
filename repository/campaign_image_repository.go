package repository

import "bwa-campaign-app/model/domain"

type CampaignImageRepository interface {
	Save(campaignImage domain.CampaignImage) (domain.CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)
}
