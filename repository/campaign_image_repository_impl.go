package repository

import (
	"bwa-campaign-app/model/domain"
	"gorm.io/gorm"
)

type CampaignImageRepositoryImpl struct {
	db *gorm.DB
}

func (c *CampaignImageRepositoryImpl) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	campaignImage := domain.CampaignImage{}
	err := c.db.Model(&campaignImage).
		Where("campaign_id=?", campaignID).
		Update("is_primary", false).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *CampaignImageRepositoryImpl) Save(campaignImage domain.CampaignImage) (domain.CampaignImage, error) {
	err := c.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func NewCampaignImageRepository(db *gorm.DB) CampaignImageRepository {
	return &CampaignImageRepositoryImpl{db: db}
}
