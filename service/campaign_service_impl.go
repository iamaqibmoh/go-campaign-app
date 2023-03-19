package service

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/repository"
)

type CampaignServiceImpl struct {
	repository.CampaignRepository
}

func (s *CampaignServiceImpl) FindDetailCampaignByID(id web.CampaignIDFromURI) (domain.Campaign, error) {
	campaign, err := s.CampaignRepository.FindByID(id.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *CampaignServiceImpl) FindCampaigns(userID int) ([]domain.Campaign, error) {
	if userID != 0 {
		campaigns, err := s.CampaignRepository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.CampaignRepository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func NewCampaignService(campaignRepository repository.CampaignRepository) CampaignService {
	return &CampaignServiceImpl{CampaignRepository: campaignRepository}
}
