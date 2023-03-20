package service

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/repository"
	"errors"
)

type CampaignImageServiceImpl struct {
	repository.CampaignImageRepository
	repository.CampaignRepository
}

func (c *CampaignImageServiceImpl) CreateCampaignImage(input web.CreateCampaignImageInput, fileLocation string) (domain.CampaignImage, error) {
	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := c.CampaignImageRepository.MarkAllImagesAsNonPrimary(input.CampaignID)
		helper.PanicIfError(err)
	}

	campaignImage := domain.CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.FileName = fileLocation
	campaignImage.IsPrimary = isPrimary

	campaign, err2 := c.CampaignRepository.FindByID(input.CampaignID)
	helper.PanicIfError(err2)

	if campaign.UserID != input.User.ID {
		return campaignImage, errors.New("You're not an owner of the campaign")
	}

	save, err := c.CampaignImageRepository.Save(campaignImage)
	helper.PanicIfError(err)

	return save, nil
}

func NewCampaignImageService(campaignImageRepository repository.CampaignImageRepository, campaignRepository repository.CampaignRepository) CampaignImageService {
	return &CampaignImageServiceImpl{
		CampaignImageRepository: campaignImageRepository,
		CampaignRepository:      campaignRepository,
	}
}
