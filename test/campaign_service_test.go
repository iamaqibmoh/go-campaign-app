package test

import (
	"bwa-campaign-app/helper"
	"bwa-campaign-app/model/web"
	"bwa-campaign-app/service"
	"encoding/json"
	"fmt"
	"testing"
)

func newCampaignService() service.CampaignService {
	campaignRepository := newCampaignRepository()
	return service.NewCampaignService(campaignRepository)
}

func TestFindCampaigns(t *testing.T) {
	campaignService := newCampaignService()
	campaigns, err := campaignService.FindCampaigns(3)
	helper.PanicIfError(err)

	fmt.Println(len(campaigns))
}

func TestCreateCampaign(t *testing.T) {
	campaignService := newCampaignService()
	userService := newUserService()
	user, _ := userService.FindUserByID(12)
	input := web.CreateCampaignInput{}
	input.Name = "Penggalangan Dana Startup"
	input.ShortDescription = "Short Description"
	input.Description = "Long Description"
	input.GoalAmount = 100000000
	input.Perks = "gift 1, gift 2, gift 3"
	input.User = user

	campaign, err := campaignService.CreateCampaign(input)
	helper.PanicIfError(err)

	bytes, _ := json.Marshal(campaign)
	fmt.Println(string(bytes))
}
