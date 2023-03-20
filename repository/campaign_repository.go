package repository

import "bwa-campaign-app/model/domain"

type CampaignRepository interface {
	Save(campaign domain.Campaign) (domain.Campaign, error)
	FindAll() ([]domain.Campaign, error)
	FindByUserID(userID int) ([]domain.Campaign, error)
	FindByID(id int) (domain.Campaign, error)
	Update(campaign domain.Campaign) (domain.Campaign, error)
}
