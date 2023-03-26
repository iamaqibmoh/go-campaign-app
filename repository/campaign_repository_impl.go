package repository

import (
	"bwa-campaign-app/model/domain"
	"gorm.io/gorm"
)

type CampaignRepositoryImpl struct {
	db *gorm.DB
}

func (r *CampaignRepositoryImpl) Update(campaign domain.Campaign) (domain.Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *CampaignRepositoryImpl) FindByID(id int) (domain.Campaign, error) {
	campaign := domain.Campaign{}
	err := r.db.Where("id=?", id).Preload("User").Preload("CampaignImages").Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *CampaignRepositoryImpl) Save(campaign domain.Campaign) (domain.Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *CampaignRepositoryImpl) FindAll() ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary=1").Order("id desc").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, err
}

func (r *CampaignRepositoryImpl) FindByUserID(userID int) ([]domain.Campaign, error) {
	var campaigns []domain.Campaign
	err := r.db.Where("user_id=?", userID).Preload("CampaignImages", "campaign_images.is_primary=1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, err
}

func NewCampaignRepository(db *gorm.DB) CampaignRepository {
	return &CampaignRepositoryImpl{db: db}
}
