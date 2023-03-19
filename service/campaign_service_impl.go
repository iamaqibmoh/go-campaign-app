package service

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/repository"
	"fmt"
	"github.com/gosimple/slug"
)

type CampaignServiceImpl struct {
	repository.CampaignRepository
}

func (s *CampaignServiceImpl) CreateCampaign(input web.CreateCampaignInput) (domain.Campaign, error) {
	campaign := domain.Campaign{}
	campaign.Name = input.Name
	campaign.Summary = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	//generate slug
	slugFormat := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugFormat)

	save, err := s.CampaignRepository.Save(campaign)
	if err != nil {
		return save, err
	}

	return save, nil
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
