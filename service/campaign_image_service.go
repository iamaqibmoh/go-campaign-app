package service

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
)

type CampaignImageService interface {
	CreateCampaignImage(input web.CreateCampaignImageInput, fileLocation string) (domain.CampaignImage, error)
}
