package formatter

import (
	"bwa-campaign-app/model/domain"
	"bwa-campaign-app/model/web"
	"strings"
)

func CampaignDetailFormatter(campaign domain.Campaign) web.CampaignDetailResponse {
	campaignDetailResponse := web.CampaignDetailResponse{}
	campaignDetailResponse.ID = campaign.ID
	campaignDetailResponse.Name = campaign.Name
	campaignDetailResponse.UserID = campaign.UserID
	campaignDetailResponse.ShortDescription = campaign.Summary
	campaignDetailResponse.Description = campaign.Description

	//campaign images
	campaignDetailResponse.ImageURL = ""
	if len(campaign.CampaignImages) > 0 {
		var imageURL string
		for _, image := range campaign.CampaignImages {
			if image.IsPrimary == 1 {
				imageURL = image.FileName
			}
		}
		campaignDetailResponse.ImageURL = imageURL
	}
	campaignDetailResponse.BackerCount = campaign.BackerCount
	campaignDetailResponse.GoalAmount = campaign.GoalAmount
	campaignDetailResponse.CurrentAmount = campaign.CurrentAmount
	campaignDetailResponse.Slug = campaign.Slug

	//perks
	var perks []string
	splitPerks := strings.Split(campaign.Perks, ",")
	for _, perk := range splitPerks {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailResponse.Perks = perks

	//user of campaign
	campaignDetailResponse.User.Name = campaign.User.Name
	campaignDetailResponse.User.ImageURL = campaign.User.Avatar

	//campaign images of campaign
	images := []web.CampaignImagesOfCampaignDetail{}
	for _, image := range campaign.CampaignImages {
		img := web.CampaignImagesOfCampaignDetail{}
		img.ImageURL = image.FileName
		var isPrimary bool
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		img.IsPrimary = isPrimary
		images = append(images, img)
	}
	campaignDetailResponse.Images = images

	return campaignDetailResponse
}

func CampaignFormatter(campaign domain.Campaign) web.CampaignResponse {
	campaignResponse := web.CampaignResponse{}

	campaignResponse.ID = campaign.ID
	campaignResponse.UserID = campaign.UserID
	campaignResponse.Name = campaign.Name
	campaignResponse.ShortDescription = campaign.Summary

	campaignResponse.ImageURL = ""
	if len(campaign.CampaignImages) > 0 {
		campaignResponse.ImageURL = campaign.CampaignImages[0].FileName
	}

	campaignResponse.GoalAmount = campaign.GoalAmount
	campaignResponse.CurrentAmount = campaign.CurrentAmount
	campaignResponse.Slug = campaign.Slug

	return campaignResponse
}

func CampaignsFormatter(campaigns []domain.Campaign) []web.CampaignResponse {
	var campaignsResponse []web.CampaignResponse

	for _, campaign := range campaigns {
		campaignResponse := CampaignFormatter(campaign)
		campaignsResponse = append(campaignsResponse, campaignResponse)
	}

	return campaignsResponse
}
